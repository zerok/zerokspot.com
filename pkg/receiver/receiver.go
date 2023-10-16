package receiver

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/gohugoio/hugo/parser/pageparser"
	"github.com/google/go-github/v52/github"
	"golang.org/x/oauth2"
	"gopkg.in/yaml.v2"
)

type Configuration struct {
	SkipPull     bool
	SkipCommit   bool
	RepoPath     string
	GitHubToken  string
	TimeLocation *time.Location
	ForceNow     time.Time
	AccessToken  string
}

type Receiver struct {
	cfg      Configuration
	router   chi.Router
	ghClient *github.Client
}

type Configurator func(cfg *Configuration)

func New(ctx context.Context, cfgs ...Configurator) *Receiver {
	router := chi.NewRouter()
	cfg := Configuration{}
	for _, c := range cfgs {
		c(&cfg)
	}
	recv := &Receiver{
		cfg:    cfg,
		router: router,
	}
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: cfg.GitHubToken},
	)
	tc := oauth2.NewClient(ctx, ts)
	recv.ghClient = github.NewClient(tc)
	router.Post("/", recv.handleRecieve)
	return recv
}

func (recv *Receiver) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	recv.router.ServeHTTP(w, r)
}

type postStatus struct {
	Path string `json:"path"`
}

func (recv *Receiver) handleRecieve(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	if recv.cfg.AccessToken != "" && token != recv.cfg.AccessToken {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}
	switch r.Header.Get("Content-Type") {
	case "text/markdown":
		if status, err := recv.handleReceiveMarkdown(w, r); err != nil {
			fmt.Println(err)
			http.Error(w, "Error processing markdown", http.StatusInternalServerError)
		} else {
			w.Header().Set("Content-Type", "text/json")
			json.NewEncoder(w).Encode(status)
		}
	default:
		http.Error(w, "Unsupported content type", http.StatusBadRequest)
	}
}

func (recv *Receiver) handleReceiveMarkdown(w http.ResponseWriter, r *http.Request) (*postStatus, error) {
	ctx := r.Context()
	defer r.Body.Close()
	data, err := pageparser.ParseFrontMatterAndContent(r.Body)
	if err != nil {
		return nil, err
	}
	slug := data.FrontMatter["slug"].(string)
	if slug == "" {
		return nil, fmt.Errorf("no slug specified")
	}
	delete(data.FrontMatter, "slug")
	now := recv.cfg.ForceNow
	if now.IsZero() {
		now = time.Now()
	}
	now = now.In(recv.cfg.TimeLocation)

	if !recv.cfg.SkipPull {
		slog.DebugContext(ctx, "Switching to main and pulling latest changes")
		if err := recv.callGit(ctx, "switch", "main"); err != nil {
			return nil, err
		}
		if err := recv.callGit(ctx, "pull", "origin", "main"); err != nil {
			return nil, err
		}
	}

	relDestDir := filepath.Join("content", "weblog", fmt.Sprintf("%d", now.Year()))
	relDestPath := filepath.Join(relDestDir, fmt.Sprintf("%s.md", slug))
	destDir := filepath.Join(recv.cfg.RepoPath, "content", "weblog", fmt.Sprintf("%d", now.Year()))
	if err := os.MkdirAll(destDir, 0700); err != nil {
		return nil, fmt.Errorf("failed to create  destination directory: %w", err)
	}
	destPath := filepath.Join(destDir, fmt.Sprintf("%s.md", slug))
	data.FrontMatter["date"] = now.Format(time.RFC3339)
	slog.InfoContext(ctx, fmt.Sprintf("Writing output to %s", destPath))
	fp, err := os.OpenFile(destPath, os.O_CREATE|os.O_RDWR|os.O_EXCL, 0600)
	if err != nil {
		return nil, fmt.Errorf("failed to open destination file `%s`: %w", destPath, err)
	}
	fp.WriteString("---\n")
	yaml.NewEncoder(fp).Encode(data.FrontMatter)
	fp.WriteString("---\n\n")
	fp.Write(data.Content)
	if err := fp.Close(); err != nil {
		return nil, fmt.Errorf("failed to close target file: %w", err)
	}

	if !recv.cfg.SkipCommit {
		branchname := fmt.Sprintf("articles/%s", slug)
		if err := recv.callGit(ctx, "switch", "-C", branchname); err != nil {
			return nil, err
		}
		if err := recv.callGit(ctx, "add", relDestPath); err != nil {
			recv.cleanupBranch(ctx, branchname)
			return nil, err
		}
		if err := recv.callGit(ctx, "commit", "-m", "Add "+slug); err != nil {
			recv.cleanupBranch(ctx, branchname)
			return nil, err
		}
		if err := recv.callGit(ctx, "push", "origin", branchname); err != nil {
			recv.cleanupBranch(ctx, branchname)
			return nil, err
		}
		if err := recv.cleanupBranch(ctx, branchname); err != nil {
			return nil, err
		}
		slog.DebugContext(ctx, "Creating pull-request")
		if err := recv.createPR(ctx, slug, branchname); err != nil {
			return nil, err
		}
	}
	return &postStatus{
		Path: relDestPath,
	}, nil
}

func (recv *Receiver) cleanupBranch(ctx context.Context, branchname string) error {
	if err := recv.callGit(ctx, "switch", "main"); err != nil {
		return err
	}
	if err := recv.callGit(ctx, "branch", "-D", branchname); err != nil {
		return err
	}
	return nil
}

func (recv *Receiver) callGit(ctx context.Context, args ...string) error {
	slog.DebugContext(ctx, fmt.Sprintf("Executing git-command: %s", args))
	cmd := exec.CommandContext(ctx, "git", args...)
	cmd.Dir = recv.cfg.RepoPath
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func (recv *Receiver) createPR(ctx context.Context, slug string, branchname string) error {
	title := fmt.Sprintf("Article: %s", slug)
	base := "main"
	pull := github.NewPullRequest{
		Title: &title,
		Head:  &branchname,
		Base:  &base,
	}
	pr, _, err := recv.ghClient.PullRequests.Create(ctx, "zerok", "zerokspot.com", &pull)
	if err != nil {
		return fmt.Errorf("failed to create pull-request: %w", err)
	}
	slog.InfoContext(ctx, fmt.Sprintf("Pull-request created: %s", pr.GetHTMLURL()))
	return nil
}

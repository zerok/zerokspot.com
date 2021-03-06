package textbundlereceiver

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/go-chi/chi"
	slugify "github.com/gosimple/slug"
	"github.com/rs/zerolog"
	"gitlab.com/zerok/zerokspot.com/pkg/textbundleimporter"
)

type Receiver struct {
	router      chi.Router
	AccessToken string
	RepoPath    string
	Importer    *textbundleimporter.Importer
	GitHubUser  string
	GitHubToken string
}

type Configurator func(*Receiver)

func New(configs ...Configurator) *Receiver {
	recv := &Receiver{}
	recv.router = chi.NewRouter()
	recv.router.Post("/", recv.handleReceive)
	for _, c := range configs {
		c(recv)
	}
	return recv
}

func (recv *Receiver) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	recv.router.ServeHTTP(w, r)
}

func (recv *Receiver) handleReceive(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := zerolog.Ctx(ctx)
	token := r.Header.Get("Authorization")
	if recv.AccessToken != "" && token != recv.AccessToken {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}
	if err := r.ParseMultipartForm(5_000_000); err != nil {
		logger.Debug().Err(err).Msg("Invalid form data")
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}
	slug := r.FormValue("slug")
	file, filehdr, err := r.FormFile("data")
	if err != nil {
		logger.Debug().Err(err).Msg("Invalid file data")
		http.Error(w, "Invalid file data", http.StatusBadRequest)
		return
	}
	packpath, err := recv.storePack(ctx, file, filehdr)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to store file")
		http.Error(w, "Failed to store file", http.StatusInternalServerError)
		return
	}
	logger.Info().Msgf("Stored pack under %s", packpath)
	defer os.RemoveAll(packpath)

	if slug == "" {
		elems := strings.SplitN(filehdr.Filename, ".", 2)
		if len(elems) < 1 {
			logger.Debug().Err(err).Msgf("Invalid file name: %s", filehdr.Filename)
			http.Error(w, "Invalid file data", http.StatusBadRequest)
			return
		}
		slug = elems[0]
	}
	slug = slugify.Make(slug)
	branchname := fmt.Sprintf("articles/%s", slug)
	if err := recv.callGit(ctx, "checkout", "-f", "main"); err != nil {
		logger.Error().Err(err).Msg("Git failed")
		http.Error(w, "Git failed", http.StatusInternalServerError)
		return
	}
	if err := recv.callGit(ctx, "clean", "-f"); err != nil {
		logger.Error().Err(err).Msg("Git failed")
		http.Error(w, "Git failed", http.StatusInternalServerError)
		return
	}
	if err := recv.callGit(ctx, "pull", "origin", "main"); err != nil {
		logger.Error().Err(err).Msg("Git failed")
		http.Error(w, "Git failed", http.StatusInternalServerError)
		return
	}
	if err := recv.callGit(ctx, "checkout", "-f", "-b", branchname); err != nil {
		logger.Error().Err(err).Msg("Git failed")
		http.Error(w, "Git failed", http.StatusInternalServerError)
		return
	}
	if err := recv.Importer.Import(ctx, packpath, slug); err != nil {
		logger.Error().Err(err).Msg("Import failed")
		http.Error(w, "Import failed", http.StatusInternalServerError)
		return
	}
	if err := recv.callGit(ctx, "add", "."); err != nil {
		logger.Error().Err(err).Msg("Git failed")
		http.Error(w, "Git failed", http.StatusInternalServerError)
		return
	}
	if err := recv.callGit(ctx, "commit", "-m", "Add "+slug); err != nil {
		logger.Error().Err(err).Msg("Git failed")
		http.Error(w, "Git failed", http.StatusInternalServerError)
		return
	}
	if err := recv.callGit(ctx, "push", "origin", branchname); err != nil {
		logger.Error().Err(err).Msg("Git failed")
		http.Error(w, "Git failed", http.StatusInternalServerError)
		return
	}
	if err := recv.callGit(ctx, "checkout", "-f", "main"); err != nil {
		logger.Error().Err(err).Msg("Git failed")
		http.Error(w, "Git failed", http.StatusInternalServerError)
		return
	}
	if err := recv.callGit(ctx, "branch", "-D", branchname); err != nil {
		logger.Error().Err(err).Msg("Git failed")
		http.Error(w, "Git failed", http.StatusInternalServerError)
		return
	}
	if err := recv.createPR(ctx, slug, branchname); err != nil {
		logger.Error().Err(err).Msg("Failed to create PR")
		http.Error(w, "PR creation failed", http.StatusInternalServerError)
		return
	}
}

func (recv *Receiver) createPR(ctx context.Context, slug string, branchname string) error {
	client := http.Client{}
	body := bytes.Buffer{}
	if err := json.NewEncoder(&body).Encode(&prBody{
		Title: "Article: " + slug,
		Head:  branchname,
		Base:  "main",
	}); err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPost, "https://api.github.com/repos/zerok/zerokspot.com/pulls", &body)
	if err != nil {
		return err
	}
	req.SetBasicAuth(recv.GitHubUser, recv.GitHubToken)
	resp, err := client.Do(req.WithContext(ctx))
	if err != nil {
		return err
	}
	if resp.StatusCode >= 300 {
		return fmt.Errorf("GitHub API returned an unexpected status code: %d", resp.StatusCode)
	}
	return nil
}

type prBody struct {
	Title string `json:"title"`
	Base  string `json:"base"`
	Head  string `json:"head"`
}

func (recv *Receiver) callGit(ctx context.Context, args ...string) error {
	cmd := exec.Command("git", args...)
	cmd.Dir = recv.RepoPath
	return cmd.Run()
}

func (recv *Receiver) storePack(ctx context.Context, file multipart.File, filehdr *multipart.FileHeader) (string, error) {
	out, err := ioutil.TempFile(os.TempDir(), "textpackrecv-")
	if err != nil {
		return "", err
	}
	if _, err := io.Copy(out, file); err != nil {
		out.Close()
		return "", err
	}
	out.Close()
	return out.Name(), nil
}

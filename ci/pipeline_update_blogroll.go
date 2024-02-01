package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"strings"
	"time"

	"dagger.io/dagger"
	"github.com/spf13/cobra"
)

var pipelineUpdateBlogrollCmd = &cobra.Command{
	Use: "update-blogroll",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()
		ctx, span := tracer.Start(ctx, "main")
		defer span.End()
		dc, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stderr))
		if err != nil {
			failSpan(ctx, span, "Failed to connect to Dagger", err)
			return err
		}
		defer dc.Close()
		return generateBlogroll(ctx, dc, versions)
	},
}

func generateBlogroll(ctx context.Context, dc *dagger.Client, versions *Versions) error {
	ctx, span := tracer.Start(ctx, "build")
	defer span.End()

	feedbinUsername := dc.SetSecret("feedbinUsername", os.Getenv("FEEDBIN_USER"))
	feedbinPassword := dc.SetSecret("feedbinPassword", os.Getenv("FEEDBIN_PASSWORD"))
	githubTokenSecret := dc.SetSecret("GITHUB_TOKEN", os.Getenv("GITHUB_TOKEN"))

	var blogBin *dagger.File
	binaryCacheHit := os.Getenv("CACHE_HIT_BLOG_BINARY") == "true"

	if binaryCacheHit {
		blogBin = dc.Host().File("./bin/blog")
	} else {
		goContainer := getGoContainer(dc)
		blogBin = getBlogBinary(dc, withOtelEnv(ctx, dc, goContainer))
		if _, err := blogBin.Export(ctx, "./bin/blog"); err != nil {
			return fmt.Errorf("failed to export the generated binary: %w", err)
		}
	}

	rootDirectory := dc.Host().Directory(".", dagger.HostDirectoryOpts{
		Exclude: []string{"output", "public"},
	})

	blogrollFile := dc.Container().From(versions.AlpineImage()).
		WithExec([]string{"apk", "add", "--no-cache", "tzdata"}).
		WithEnvVariable("CACHE_BUSTER", time.Now().Format(time.RFC3339Nano)).
		WithFile("/usr/local/bin/blog", blogBin).
		WithSecretVariable("FEEDBIN_USER", feedbinUsername).
		WithSecretVariable("FEEDBIN_PASSWORD", feedbinPassword).
		WithWorkdir("/data").
		WithExec([]string{"blog", "blogroll", "--output", "blogroll.json"}).
		File("/data/blogroll.json")

	// Now that we have a file, let's do a checkout, compare the content of
	// that file and if there is change, create a PR + automerge.
	container, err := dc.Container().From(versions.AlpineImage()).
		WithExec([]string{"apk", "add", "--no-cache", "git", "tzdata", "github-cli"}).
		WithEnvVariable("CACHE_BUSTER", time.Now().Format(time.RFC3339Nano)).
		WithSecretVariable("GITHUB_TOKEN", githubTokenSecret).
		WithMountedDirectory("/src", rootDirectory).
		WithWorkdir("/src").
		WithExec([]string{"/bin/sh", "-c", "git remote set-url origin https://oauth2:$GITHUB_TOKEN@github.com/zerok/zerokspot.com.git"}).
		WithExec([]string{"git", "config", "user.email", "bot@zerokspot.com"}).
		WithExec([]string{"git", "config", "user.name", "zerokspot-bot"}).
		WithFile("/input/data/blogroll.json", blogrollFile).
		WithExec([]string{"git", "checkout", "-f", "main"}).
		WithExec([]string{"git", "pull", "origin", "main"}).
		WithExec([]string{"git", "switch", "-C", "update-blogroll"}).
		WithExec([]string{"cp", "/input/data/blogroll.json", "/src/data/"}).
		WithExec([]string{"git", "status", "--porcelain", "data/blogroll.json"}).Sync(ctx)

	if err != nil {
		return err
	}
	output, err := container.Stdout(ctx)
	if err != nil {
		return err
	}
	if strings.TrimSpace(output) == "" {
		slog.InfoContext(ctx, "No changes found")
		return nil
	}
	_, err = container.
		WithExec([]string{"git", "add", "./data/blogroll.json"}).
		WithExec([]string{"git", "commit", "-m", "Update blogroll"}).
		WithExec([]string{"git", "push", "--set-upstream", "origin", "update-blogroll"}).
		WithExec([]string{"gh", "pr", "create", "--fill"}).
		WithExec([]string{"gh", "pr", "merge", "--auto", "--squash"}).
		Sync(ctx)
	return err
}

func init() {
	pipelinesCmd.AddCommand(pipelineUpdateBlogrollCmd)
}

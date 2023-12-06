package main

import (
	"fmt"
	"log/slog"
	"os"
	"strings"
	"time"

	"dagger.io/dagger"
	"github.com/spf13/cobra"
	"go.opentelemetry.io/otel/codes"
)

var importMastodonLinks = &cobra.Command{
	Use: "import-mastodon-links",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()
		ctx, span := tracer.Start(ctx, "import-mastodon-links")
		defer span.End()
		slog.InfoContext(ctx, fmt.Sprintf("TraceID: %s", span.SpanContext().TraceID().String()))

		dc, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stderr))
		if err != nil {
			span.SetStatus(codes.Error, "Failed to setup Dagger")
			slog.ErrorContext(ctx, "Failed to setup Dagger", slog.Any("err", err))
			span.End()
			return err
		}
		defer dc.Close()

		githubTokenSecret := dc.SetSecret("GITHUB_TOKEN", os.Getenv("GITHUB_TOKEN"))

		rootDirectory := dc.Host().Directory(".", dagger.HostDirectoryOpts{
			Exclude: []string{"output", "public"},
		})

		goContainer := getGoContainer(dc)
		blogBin := getBlogBinary(dc, withOtelEnv(ctx, dc, goContainer))

		container := dc.Container().From(versions.AlpineImage()).
			WithExec([]string{"apk", "add", "--no-cache", "git", "tzdata", "github-cli"}).
			WithDirectory("/src", rootDirectory).
			WithWorkdir("/src").
			WithFile("/src/bin/blog", blogBin).
			WithSecretVariable("GITHUB_TOKEN", githubTokenSecret).
			WithEnvVariable("CACHE_BUSTER", time.Now().Format(time.RFC3339Nano)).
			WithExec([]string{"git", "checkout", "-f", "main"}).
			WithExec([]string{"git", "config", "user.email", "bot@zerokspot.com"}).
			WithExec([]string{"git", "config", "user.name", "zerokspot-bot"}).
			WithExec([]string{"/bin/sh", "-c", "git remote set-url origin https://oauth2:$GITHUB_TOKEN@github.com/zerok/zerokspot.com.git"}).
			WithExec([]string{"git", "pull", "origin", "main"}).
			// Reset the import-mastodon-links branch if it already exists so
			// that there is only one in flight at the same time.
			WithExec([]string{"git", "switch", "-C", "import-mastodon-links"}).
			WithExec([]string{"/src/bin/blog", "import-mastodon-links"}).
			WithExec([]string{"git", "status", "--porcelain"})

		// Now check if we need to continue. If there are no changes, no need to
		// open a new PR:
		container, err = container.Sync(ctx)
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
			WithExec([]string{"/bin/sh", "-c", "(gh pr view import-mastodon-links && gh pr close import-mastodon-links) || true"}).
			WithExec([]string{"git", "add", "."}).
			WithExec([]string{"git", "commit", "-m", "Add mastodon link(s)"}).
			WithExec([]string{"git", "push", "origin", "import-mastodon-links"}).
			WithExec([]string{"gh", "pr", "create", "--fill"}).
			Sync(ctx)
		return err
	},
}

func init() {
	pipelinesCmd.AddCommand(importMastodonLinks)
}

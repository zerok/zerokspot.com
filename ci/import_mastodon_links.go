package main

import (
	"fmt"
	"log/slog"
	"os"

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

		rootDirectory := dc.Host().Directory(".", dagger.HostDirectoryOpts{
			Exclude: []string{"output", "public"},
		})

		goContainer := getGoContainer(dc)
		blogBin := getBlogBinary(dc, withOtelEnv(ctx, dc, goContainer))

		container := dc.Container().From(versions.AlpineImage()).
			WithExec([]string{"apk", "add", "--no-cache", "git", "tzdata"}).
			WithMountedDirectory("/src", rootDirectory).
			WithWorkdir("/src").
			WithFile("/src/bin/blog", blogBin).
			WithExec([]string{"/src/bin/blog", "import-mastodon-links"}).
			WithExec([]string{"git", "status"})
		_, err = container.Sync(ctx)
		return err
	},
}

func init() {
	pipelinesCmd.AddCommand(importMastodonLinks)
}

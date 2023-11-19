package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"strings"

	"dagger.io/dagger"
	"github.com/spf13/cobra"
	"gitlab.com/zerok/zerokspot.com/pkg/otelhandler"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
	"go.opentelemetry.io/otel/trace"
)

var tracer trace.Tracer

var rootCmd = &cobra.Command{}

var versions *Versions

func main() {
	logger := slog.New(otelhandler.OTELHandler{Handler: slog.NewTextHandler(os.Stderr, nil)})
	slog.SetDefault(logger)
	ctx := context.Background()

	var exporter sdktrace.SpanExporter
	var err error

	if os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT") != "" && os.Getenv("OTEL_EXPORTER_OTLP_HEADERS") != "" {
		otlpClient := otlptracehttp.NewClient()
		exporter, err = otlptrace.New(ctx, otlpClient)
	} else {
		exporter, err = stdouttrace.New()
	}
	if err != nil {
		slog.ErrorContext(ctx, "Failed to initialize OTLP output", slog.Any("err", err))
		os.Exit(1)
	}
	res, err := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("zerokspot-ci"),
		),
	)
	if err != nil {
		slog.ErrorContext(ctx, "Failed to generate OTLP resource", slog.Any("err", err))
		os.Exit(1)
	}
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSpanProcessor(sdktrace.NewBatchSpanProcessor(exporter)),
		sdktrace.WithResource(res),
	)
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			slog.ErrorContext(ctx, "Failed to shut tracer provider down", slog.Any("err", err))
		}
	}()
	otel.SetTracerProvider(tp)

	tracer = tp.Tracer("zerokspot-ci")
	versions, err = LoadVersions(ctx)
	if err != nil {
		if err := tp.Shutdown(context.Background()); err != nil {
			slog.ErrorContext(ctx, "Failed to shut tracer provider down", slog.Any("err", err))
		}
		os.Exit(1)
	}

	if err := rootCmd.ExecuteContext(ctx); err != nil {
		slog.Error("Command failed", slog.Any("error", err))
		if err := tp.Shutdown(context.Background()); err != nil {
			slog.ErrorContext(ctx, "Failed to shut tracer provider down", slog.Any("err", err))
		}
		os.Exit(1)
	}
}

func getPublicRev(ctx context.Context) (string, error) {
	ctx, span := tracer.Start(ctx, "getPublicRev")
	defer span.End()
	httpClient := http.Client{}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://zerokspot.com/.gitrev", nil)
	if err != nil {
		return "", err
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		return "", err
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(data)), nil

}

func getTraceParent(ctx context.Context) string {
	sctx := trace.SpanContextFromContext(ctx)
	return fmt.Sprintf("00-%s-%s-%s", sctx.TraceID().String(), sctx.SpanID().String(), sctx.TraceFlags()&trace.FlagsSampled)
}

func withOtelEnv(ctx context.Context, client *dagger.Client, container *dagger.Container) *dagger.Container {
	parentTrace := getTraceParent(ctx)
	slog.InfoContext(ctx, fmt.Sprintf("Container with parent trace: %s", parentTrace))
	c := container.
		WithEnvVariable("TRACEPARENT", parentTrace).
		WithEnvVariable("OVERRIDE_TRACEPARENT", parentTrace).
		WithoutEnvVariable("OTEL_TRACES_EXPORTER")

	for _, v := range []string{
		"OTEL_EXPORTER_OTLP_PROTOCOL",
		"OTEL_EXPORTER_OTLP_HEADERS",
		"OTEL_EXPORTER_OTLP_ENDPOINT",
	} {
		sec := client.SetSecret(v, os.Getenv(v))
		c = c.WithSecretVariable(v, sec)
	}
	return c
}

func build(ctx context.Context, client *dagger.Client, versions *Versions, publish bool) error {
	ctx, span := tracer.Start(ctx, "build")
	defer span.End()
	feedbinUsername := client.SetSecret("feedbinUsername", os.Getenv("FEEDBIN_USER"))
	feedbinPassword := client.SetSecret("feedbinPassword", os.Getenv("FEEDBIN_PASSWORD"))
	sshPrivateKey := os.Getenv("SSH_PRIVATE_KEY")

	publicRev, err := getPublicRev(ctx)
	if err != nil {
		failSpan(ctx, span, "Failed to retrieve public rev", err)
		return err
	}

	rootDirectory := client.Host().Directory(".", dagger.HostDirectoryOpts{
		Exclude: []string{"public", "bin"},
	})

	var gitRev string

	goContainer := getGoContainer(client)

	// Prime the Go cache
	if _, err := goContainer.
		WithMountedFile("/src/go.mod", rootDirectory.File("go.mod")).
		WithMountedFile("/src/go.sum", rootDirectory.File("go.sum")).
		WithWorkdir("/src").
		WithExec([]string{"go", "mod", "download"}).Sync(ctx); err != nil {
		return err
	}
	goContainer = withOtelEnv(ctx, client, goContainer)

	versionOutput, err := goContainer.
		WithWorkdir("/src").
		WithMountedDirectory("/src", rootDirectory).
		WithExec([]string{"bash", "-c", "go list -m github.com/gohugoio/hugo | cut -d ' ' -f 2"}).
		Stdout(ctx)
	if err != nil {
		failSpan(ctx, span, "Failed to detect Hugo version", err)
		return err
	}
	hugoVersion := strings.TrimSpace(strings.TrimPrefix(versionOutput, "v"))
	hugoContainer := withOtelEnv(ctx, client, getHugoContainer(client, hugoVersion))
	blogBin := getBlogBinary(client, withOtelEnv(ctx, client, goContainer))

	if err := func(ctx context.Context) error {
		ctx, span := tracer.Start(ctx, "buildWebsite")
		defer span.End()
		hugoContainer = hugoContainer.
			WithSecretVariable("FEEDBIN_USER", feedbinUsername).
			WithSecretVariable("FEEDBIN_PASSWORD", feedbinPassword).
			WithMountedDirectory("/src", rootDirectory).
			WithWorkdir("/src").
			WithMountedFile("/usr/local/bin/blog", blogBin).
			WithExec([]string{"blog", "build-archive"}).
			// This will *not* fail if hugo fails but that should be ok for now
			// as build-archive already does lots of content checking
			WithExec([]string{"/bin/sh", "-c", "unset OTEL_TRACES_EXPORTER && unset OTEL_EXPORTER_OTLP_TRACES_ENDPOINT && unset OTEL_EXPORTER_OTLP_TRACES_PROTOCOL && otel-cli exec --timeout 60s --verbose --name hugo --service zerokspot-cli hugo"}).
			WithExec([]string{"blog", "build-graph"}).
			WithExec([]string{"blog", "blogroll", "--output", "data/blogroll.json"}).
			WithExec([]string{"/bin/sh", "-c", "unset OTEL_TRACES_EXPORTER && unset OTEL_EXPORTER_OTLP_TRACES_ENDPOINT && unset OTEL_EXPORTER_OTLP_TRACES_PROTOCOL && otel-cli exec --timeout 60s --verbose --name hugo --service zerokspot-cli hugo"}).
			WithExec([]string{"blog", "books", "gen-opml"}).
			WithExec([]string{"blog", "build-mapping"}).
			WithExec([]string{"git", "rev-parse", "HEAD"})

		gitRev, err = hugoContainer.Stdout(ctx)
		if err != nil {
			return err
		}
		return nil
	}(ctx); err != nil {
		failSpan(ctx, span, "Failed to build website", err)
		return err
	}

	//
	// If we don't plan to publish anything, then we're done here
	if !publish {
		if _, err := hugoContainer.Directory("/src/public").Export(ctx, "./output"); err != nil {
			span.SetStatus(codes.Error, "Failed to build website")
			return err
		}
		return nil
	}

	hugoContainer = hugoContainer.
		WithExec([]string{"blog", "changes", "--since-rev", publicRev, "--url", "--output", "public/.changes.txt"}).
		WithNewFile("/src/public/.gitrev", dagger.ContainerWithNewFileOpts{
			Contents: gitRev,
		}).
		WithExec([]string{"chmod", "0755", "/src/public"})

	if err := os.MkdirAll("public", 0700); err != nil {
		failSpan(ctx, span, "Failed to create public folder", err)
		return err
	}
	if _, err := hugoContainer.File("/src/public/.changes.txt").Export(ctx, "./public/.changes.txt"); err != nil {
		failSpan(ctx, span, "Failed to export changes.txt", err)
		return err
	}

	if err := func(ctx context.Context) error {
		ctx, span := tracer.Start(ctx, "rsync")
		defer span.End()
		// Prepare an rsync container which we can then use to upload everything:
		rsyncContainer := withOtelEnv(ctx, client, getRsyncContainer(client)).
			WithExec([]string{"mkdir", "/root/.ssh"}).
			WithExec([]string{"chmod", "0700", "/root/.ssh"}).
			WithNewFile("/root/.ssh/id_rsa", dagger.ContainerWithNewFileOpts{
				Contents:    strings.TrimSpace(sshPrivateKey) + "\n",
				Permissions: 0600,
			}).
			WithMountedFile("/usr/local/bin/blogcli", blogBin).
			WithMountedDirectory("/src/public", hugoContainer.Directory("/src/public")).
			WithWorkdir("/src/public").
			WithExec([]string{"rsync", "-e", "ssh -o StrictHostKeyChecking=no", "-az", "/usr/local/bin/blogcli", "www-zerokspot@zs-web-001.nodes.h10n.me:/srv/www/zerokspot.com/www/bin/blog"}).
			WithExec([]string{"rsync", "-e", "ssh -o StrictHostKeyChecking=no", "-az", ".", ".mapping.json.xz", ".gitrev", "www-zerokspot@zs-web-001.nodes.h10n.me:/srv/www/zerokspot.com/www/htdocs/"}).
			WithExec([]string{"ssh", "www-zerokspot@zs-web-001.nodes.h10n.me", "touch /srv/www/zerokspot.com/www/deployed"})

		if _, err := rsyncContainer.Sync(ctx); err != nil {
			return err
		}
		return nil
	}(ctx); err != nil {
		failSpan(ctx, span, "Failed to rsync", err)
		return err
	}

	changes, err := getChanges(ctx, "./public/.changes.txt")
	if err != nil {
		failSpan(ctx, span, "Failed to get changes", err)
		return err
	}

	if err := func(ctx context.Context) error {
		if len(changes) == 0 {
			slog.InfoContext(ctx, "No changes found. Skipping webmentions.")
			return nil
		}

		ctx, span := tracer.Start(ctx, "sendWebmentions")
		defer span.End()
		slog.InfoContext(ctx, "Generating webmentions")
		mentionContainer := withOtelEnv(ctx, client, client.Container().From(versions.WebmentiondImage())).WithEntrypoint(nil)
		for _, change := range changes {
			logger := slog.With(slog.String("mentionFrom", change))
			logger.InfoContext(ctx, "Mentioning")
			mentionContainer = mentionContainer.WithExec([]string{"/usr/local/bin/webmentiond", "send", change})
			if _, err := mentionContainer.Sync(ctx); err != nil {
				return err
			}
		}

		return nil
	}(ctx); err != nil {
		failSpan(ctx, span, "Failed to execute webmentions", err)
		return err
	}

	if err := func(ctx context.Context) error {
		ntfyURL := os.Getenv("NTFY_URL")
		if ntfyURL == "" {
			return nil
		}
		ctx, span := tracer.Start(ctx, "sendNotification")
		defer span.End()
		slog.InfoContext(ctx, "Sending notification")
		ntfyURLSecret := client.SetSecret("NTFY_URL", ntfyURL)
		_, err := hugoContainer.
			WithSecretVariable("NTFY_URL", ntfyURLSecret).
			WithExec([]string{
				"/bin/sh", "-c",
				"curl -d 'Website updated' ${NTFY_URL}",
			}).Sync(ctx)
		return err
	}(ctx); err != nil {
		failSpan(ctx, span, "Failed to send notifications", err)
		return err
	}
	return nil
}

func getChanges(ctx context.Context, path string) ([]string, error) {
	_, span := tracer.Start(ctx, "getChanges")
	defer span.End()
	fp, err := os.Open(path)
	if err != nil {
		span.SetStatus(codes.Error, err.Error())
		return nil, err
	}
	defer fp.Close()
	result := make([]string, 0, 10)
	scanner := bufio.NewScanner(fp)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) > 0 {
			result = append(result, line)
		}

	}
	return result, nil
}

func failSpan(ctx context.Context, span trace.Span, msg string, err error) {
	span.SetStatus(codes.Error, msg)
	slog.ErrorContext(ctx, msg, slog.Any("err", err))
	span.End()
}

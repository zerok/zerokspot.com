package main

import (
	"bufio"
	"context"
	"flag"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"dagger.io/dagger"
	"github.com/rs/zerolog"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
	"go.opentelemetry.io/otel/trace"
)

var tracer trace.Tracer

func main() {
	var publish bool
	flag.BoolVar(&publish, "publish", false, "Also upload stuff")
	flag.Parse()

	logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).With().Timestamp().Logger()
	ctx := logger.WithContext(context.Background())

	otlpClient := otlptracehttp.NewClient()
	exporter, err := otlptrace.New(ctx, otlpClient)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to initialize oltp output")
	}
	res, err := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("zerokspot-ci"),
		),
	)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to generate OLTP resource")
	}
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSpanProcessor(sdktrace.NewBatchSpanProcessor(exporter)),
		sdktrace.WithResource(res),
	)
	defer tp.Shutdown(context.Background())
	otel.SetTracerProvider(tp)

	tracer = tp.Tracer("zerokspot-ci")
	ctx, span := tracer.Start(ctx, "main")
	defer span.End()
	span.SetAttributes(attribute.Bool("publish", publish))

	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stderr))
	if err != nil {
		span.SetStatus(codes.Error, "Failed to setup Dagger")
		logger.Fatal().Err(err).Msg("Failed to setup Dagger")
	}
	defer client.Close()

	if err := build(ctx, client, publish); err != nil {
		span.SetStatus(codes.Error, "Failed to setup build")
		logger.Fatal().Err(err).Msg("Failed to build")
	}
	span.SetStatus(codes.Ok, "")
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
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(data)), nil

}

func build(ctx context.Context, client *dagger.Client, publish bool) error {
	ctx, span := tracer.Start(ctx, "build")
	defer span.End()
	logger := zerolog.Ctx(ctx)
	feedbinUsername := client.Host().EnvVariable("FEEDBIN_USER").Secret()
	feedbinPassword := client.Host().EnvVariable("FEEDBIN_PASSWORD").Secret()
	sshPrivateKey, err := client.Host().EnvVariable("SSH_PRIVATE_KEY").Value(ctx)
	if err != nil {
		span.SetStatus(codes.Error, "Failed to SSH key")
		return err
	}

	publicRev, err := getPublicRev(ctx)
	if err != nil {
		span.SetStatus(codes.Error, "Failed to retrieve public rev")
		return err
	}

	rootDirectory := client.Host().Directory(".")

	goCacheVolume := client.CacheVolume("gocache")

	var buildContainer *dagger.Container
	var hugoContainer *dagger.Container
	var blogBin *dagger.File
	var gitRev string

	// Build a binary that can be used on a Ubuntu server
	if err := func(ctx context.Context) error {
		ctx, span := tracer.Start(ctx, "buildBlogBinary")
		defer span.End()
		buildContainer = client.Container().From("golang:1.19").
			WithEnvVariable("GOOS", "linux").
			WithEnvVariable("GOARCH", "amd64").
			WithEnvVariable("GOCACHE", "/go/pkg/cache").
			WithMountedCache("/go/pkg", goCacheVolume).
			WithMountedDirectory("/src/pkg", rootDirectory.Directory("pkg")).
			WithMountedDirectory("/src/cmd", rootDirectory.Directory("cmd")).
			WithMountedFile("/src/go.mod", rootDirectory.File("go.mod")).
			WithMountedFile("/src/go.sum", rootDirectory.File("go.sum")).
			WithExec([]string{"mkdir", "/src/bin"}).
			WithWorkdir("/src/cmd/blog").
			WithExec([]string{"go", "build", "-o", "../../bin/blog"})

		blogBin = buildContainer.File("/src/bin/blog")
		if _, err := buildContainer.ExitCode(ctx); err != nil {
			span.SetStatus(codes.Error, "Failed build binary")
			return err
		}
		return nil
	}(ctx); err != nil {
		span.SetStatus(codes.Error, "Failed build binary")
		return err
	}

	if err := func(ctx context.Context) error {
		ctx, span := tracer.Start(ctx, "buildWebsite")
		defer span.End()
		hugoContainer = client.Container().From("klakegg/hugo:0.107.0-ext-ubuntu").
			WithEntrypoint([]string{}).
			WithExec([]string{"apt-get", "update"}).
			WithExec([]string{"apt-get", "install", "-y", "git"}).
			WithSecretVariable("FEEDBIN_USER", feedbinUsername).
			WithSecretVariable("FEEDBIN_PASSWORD", feedbinPassword).
			WithMountedDirectory("/src", rootDirectory).
			WithWorkdir("/src").
			WithMountedFile("/usr/local/bin/blog", blogBin).
			WithExec([]string{"blog", "build-archive"}).
			WithExec([]string{"hugo"}).
			WithExec([]string{"blog", "build-graph"}).
			WithExec([]string{"blog", "blogroll", "--output", "data/blogroll.json"}).
			WithExec([]string{"hugo"}).
			WithExec([]string{"blog", "books", "gen-opml"}).
			WithExec([]string{"blog", "search", "build-mapping"}).
			WithExec([]string{"git", "rev-parse", "HEAD"})

		gitRev, err = hugoContainer.Stdout(ctx)
		if err != nil {
			span.SetStatus(codes.Error, "Failed build website")
			return err
		}
		return nil
	}(ctx); err != nil {
		span.SetStatus(codes.Error, "Failed build website")
		return err
	}

	hugoContainer = hugoContainer.
		WithExec([]string{"blog", "changes", "--since-rev", publicRev, "--url", "--output", "public/.changes.txt"}).
		WithNewFile("/src/public/.gitrev", dagger.ContainerWithNewFileOpts{
			Contents: gitRev,
		}).
		WithExec([]string{"chmod", "0755", "/src/public"})

	if err := os.MkdirAll("public", 0700); err != nil {
		span.SetStatus(codes.Error, "Failed to create public folder")
		return err
	}
	if _, err := hugoContainer.File("/src/public/.changes.txt").Export(ctx, "./public/.changes.txt"); err != nil {
		span.SetStatus(codes.Error, "Failed to export changes.txt")
		return err
	}

	// If we don't plan to publish anything, then we're done here
	if !publish {
		if _, err := hugoContainer.ExitCode(ctx); err != nil {
			span.SetStatus(codes.Error, "Failed to build website")
			return err
		}
		return nil
	}

	if err := func(ctx context.Context) error {
		ctx, span := tracer.Start(ctx, "rsync")
		defer span.End()
		// Prepare an rsync container which we can then use to upload everything:
		rsyncContainer := client.Container().From("alpine:3.17").
			WithExec([]string{"apk", "add", "rsync", "openssh-client-default"}).
			WithExec([]string{"mkdir", "/root/.ssh"}).
			WithExec([]string{"chmod", "0700", "/root/.ssh"}).
			WithNewFile("/root/.ssh/id_rsa", dagger.ContainerWithNewFileOpts{
				Contents:    strings.TrimSpace(sshPrivateKey) + "\n",
				Permissions: 0600,
			}).
			WithMountedFile("/usr/local/bin/blogcli", blogBin).
			WithMountedDirectory("/src/public", hugoContainer.Directory("/src/public")).
			WithWorkdir("/src/public").
			WithExec([]string{"rsync", "-e", "ssh -o StrictHostKeyChecking=no", "-az", "/usr/local/bin/blogcli", "www-zerokspot@zs-web-001.nodes.h10n.me:/srv/www/zerokspot.com/www/bin/"}).
			WithExec([]string{"rsync", "-e", "ssh -o StrictHostKeyChecking=no", "-az", ".", ".mapping.json.xz", ".gitrev", "www-zerokspot@zs-web-001.nodes.h10n.me:/srv/www/zerokspot.com/www/htdocs/"}).
			WithExec([]string{"ssh", "www-zerokspot@zs-web-001.nodes.h10n.me", "touch /srv/www/zerokspot.com/www/deployed"})

		if _, err := rsyncContainer.ExitCode(ctx); err != nil {
			span.SetStatus(codes.Error, "Failed to rsync")
			return err
		}
		return nil
	}(ctx); err != nil {
		span.SetStatus(codes.Error, "Failed to rsync")
		return err
	}

	changes, err := getChanges(ctx, "./public/.changes.txt")
	if err != nil {
		return err
	}

	if len(changes) == 0 {
		logger.Info().Msg("No changes found. Skipping webmentions.")
		return nil
	}

	if err := func(ctx context.Context) error {
		ctx, span := tracer.Start(ctx, "sendWebmentions")
		defer span.End()
		logger.Info().Msg("Generating webmentions")
		mentionContainer := client.Container().From("zerok/webmentiond:latest").WithEntrypoint([]string{})
		for _, change := range changes {
			logger.Info().Msgf("Mentioning from %s", change)
			mentionContainer = mentionContainer.WithExec([]string{"/usr/local/bin/webmentiond", "send", change})
			if _, err := mentionContainer.ExitCode(ctx); err != nil {
				span.SetStatus(codes.Error, "Failed to execute webmentions")
				return err
			}
		}

		return nil
	}(ctx); err != nil {
		span.SetStatus(codes.Error, "Failed to execute webmentions")
		return err
	}
	return nil
}

func getChanges(ctx context.Context, path string) ([]string, error) {
	ctx, span := tracer.Start(ctx, "getChanges")
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

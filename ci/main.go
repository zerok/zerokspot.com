package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"dagger.io/dagger"
	"github.com/rs/zerolog"
)

func main() {
	var publish bool
	flag.BoolVar(&publish, "publish", false, "Also upload stuff")
	flag.Parse()

	logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).With().Timestamp().Logger()
	ctx := context.Background()
	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stderr))
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to setup Dagger")
	}
	defer client.Close()

	if err := build(ctx, client, publish); err != nil {
		logger.Fatal().Err(err).Msg("Failed to build")
	}
}

func requireEnv(name string) (string, error) {
	v := os.Getenv(name)
	if v == "" {
		return "", fmt.Errorf("%s not set", name)
	}
	return v, nil
}

func getPublicRev(ctx context.Context) (string, error) {
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
	feedbinUsername := client.Host().EnvVariable("FEEDBIN_USER").Secret()
	feedbinPassword := client.Host().EnvVariable("FEEDBIN_PASSWORD").Secret()
	sshPrivateKey, err := client.Host().EnvVariable("SSH_PRIVATE_KEY").Value(ctx)
	if err != nil {
		return err
	}

	publicRev, err := getPublicRev(ctx)
	if err != nil {
		return err
	}

	rootDirectory := client.Host().Directory(".")

	goCacheVolume := client.CacheVolume("gocache")

	// Build a binary that can be used on a Ubuntu server
	buildContainer := client.Container().From("golang:1.19").
		WithEnvVariable("GOOS", "linux").
		WithEnvVariable("GOARCH", "amd64").
		WithEnvVariable("GOCACHE", "/go/pkg/cache").
		WithMountedCache("/go/pkg", goCacheVolume).
		WithExec([]string{"go", "env"}).
		WithMountedDirectory("/src/pkg", rootDirectory.Directory("pkg")).
		WithMountedDirectory("/src/cmd", rootDirectory.Directory("cmd")).
		WithMountedFile("/src/go.mod", rootDirectory.File("go.mod")).
		WithMountedFile("/src/go.sum", rootDirectory.File("go.sum")).
		WithExec([]string{"mkdir", "/src/bin"}).
		WithWorkdir("/src/cmd/blog").
		WithExec([]string{"go", "build", "-o", "../../bin/blog"})

	blogBin := buildContainer.File("/src/bin/blog")

	hugoContainer := client.Container().From("klakegg/hugo:0.107.0-ext-ubuntu").
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

	gitRev, err := hugoContainer.Stdout(ctx)
	if err != nil {
		return err
	}

	hugoContainer = hugoContainer.
		WithExec([]string{"blog", "changes", "--since-rev", publicRev, "--url", "--output", "public/.changes.txt"}).
		WithNewFile("/src/public/.gitrev", dagger.ContainerWithNewFileOpts{
			Contents: gitRev,
		}).
		WithExec([]string{"chmod", "0755", "/src/public"})

	if err := os.MkdirAll("public", 0700); err != nil {
		return err
	}
	if _, err := hugoContainer.File("/src/public/.changes.txt").Export(ctx, "./public/.changes.txt"); err != nil {
		return err
	}

	// If we don't plan to publish anything, then we're done here
	if !publish {
		if _, err := hugoContainer.ExitCode(ctx); err != nil {
			return err
		}
		return nil
	}

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
		WithExec([]string{"rsync", "-e", "ssh -o StrictHostKeyChecking=no", "-avz", "/usr/local/bin/blogcli", "www-zerokspot@zs-web-001.nodes.h10n.me:/srv/www/zerokspot.com/www/bin/"}).
		WithExec([]string{"rsync", "-e", "ssh -o StrictHostKeyChecking=no", "-avz", ".", ".mapping.json.xz", ".gitrev", "www-zerokspot@zs-web-001.nodes.h10n.me:/srv/www/zerokspot.com/www/htdocs/"}).
		WithExec([]string{"ssh", "www-zerokspot@zs-web-001.nodes.h10n.me", "touch /srv/www/zerokspot.com/www/deployed"})

	if _, err := rsyncContainer.ExitCode(ctx); err != nil {
		return err
	}
	return nil
}

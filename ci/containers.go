package main

import (
	"context"
	"fmt"
	"os"

	"dagger.io/dagger"
)

func getGoContainer(dc *dagger.Client) *dagger.Container {
	goCacheVolume := dc.CacheVolume("gocache")

	return dc.Container().From(versions.GoImage()).
		WithEnvVariable("GOOS", "linux").
		WithEnvVariable("GOARCH", "amd64").
		WithEnvVariable("GOCACHE", "/go/pkg/cache").
		WithEnvVariable("CGO_ENABLED", "0").
		WithMountedCache("/go/pkg", goCacheVolume)

}

func getHugoContainer(dc *dagger.Client, hugoVersion string) *dagger.Container {
	return dc.Container(dagger.ContainerOpts{
		Platform: "linux/amd64",
	}).From(versions.UbuntuImage()).
		WithEnvVariable("DEBIAN_FRONTEND", "noninteractive").
		WithExec([]string{"apt-get", "update"}).
		WithExec([]string{"apt-get", "install", "-y", "curl", "tzdata", "git"}).
		WithExec([]string{"curl", "-L", "-o", "/tmp/otel-cli.deb", "https://github.com/equinix-labs/otel-cli/releases/download/v0.4.1/otel-cli_0.4.1_linux_amd64.deb"}).
		WithExec([]string{"dpkg", "-i", "/tmp/otel-cli.deb"}).
		WithExec([]string{"curl", "-L", "-o", "/tmp/hugo.deb", fmt.Sprintf("https://github.com/gohugoio/hugo/releases/download/v%s/hugo_extended_%s_linux-amd64.deb", hugoVersion, hugoVersion)}).
		WithExec([]string{"dpkg", "-i", "/tmp/hugo.deb"})
}

func getRsyncContainer(dc *dagger.Client) *dagger.Container {
	return dc.Container().From(versions.AlpineImage()).
		WithExec([]string{"apk", "add", "rsync", "openssh-client-default"})
}

func getBlogBinary(dc *dagger.Client, buildContainer *dagger.Container) *dagger.File {
	rootDirectory := dc.Host().Directory(".", dagger.HostDirectoryOpts{
		Include: []string{"go.mod", "go.sum", "pkg", "cmd"},
	})
	return buildContainer.
		WithMountedDirectory("/src", rootDirectory).
		WithExec([]string{"mkdir", "/src/bin"}).
		WithWorkdir("/src/cmd/blog").
		WithExec([]string{"go", "build", "-o", "../../bin/blog"}).
		File("/src/bin/blog")
}

func getOrRestoreBlogBinary(ctx context.Context, dc *dagger.Client) (*dagger.File, error) {
	var blogBin *dagger.File
	binaryCacheHit := os.Getenv("CACHE_HIT_BLOG_BINARY") == "true"

	if binaryCacheHit {
		blogBin = dc.Host().File("./bin/blog")
	} else {
		goContainer := getGoContainer(dc)
		blogBin = getBlogBinary(dc, withOtelEnv(ctx, dc, goContainer))
		if _, err := blogBin.Export(ctx, "./bin/blog"); err != nil {
			return nil, fmt.Errorf("failed to export the generated binary: %w", err)
		}
	}
	return blogBin, nil
}

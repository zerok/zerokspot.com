---
title: Maintainable CI pipelines with Dagger
date: "2023-01-16T20:02:10+01:00"
tags:
- dagger
- cicd
- zerokspot
- golang
---

The CI pipeline for a blog such as zerokspot.com *should* be quite simple. Run Hugo and push the result onto a server via rsync. Over the years, though, the site has gained a lot of functionality that is (or at least was not possible at the time of implementing it) with a plain Hugo setup and so the pipeline became more complicated, too.

It also didn’t help that the whole setup started out on GitHub, then moved to GitLab, and then back again. I also wanted to add things like preview builds along the way but eventually abandoned that plan without properly cleaning the pipeline up again. One of the problems with all of that was also that I couldn’t simply test these things locally. I always had to modify the CI configuration, push it, and then see it either fail or succeed. And do that like 10 times in a row.

A couple of months ago I stumbled upon [Dagger](https://dagger.io/) thanks to [the ShipIt podcast](https://changelog.com/shipit/48), which might help with some of my issues. Months went by and I couldn’t find the time to properly play with it. Turns out, that was a good thing as it has matured quite a bit since then!

## What is Dagger?

But first, what is Dagger?

> Dagger is a programmable CI/CD engine that runs your pipelines in containers. — [dagger.io](https://dagger.io)

OK, but what does that mean? It means that you write your pipeline actions in (for instance) Go and use the Dagger SDK to interact with OCI containers, building them, and running operations inside of them.

The Dagger SDK itself just builds GraphQL requests and sends them to a special container running the “Dagger Engine”. This, in turn, translates those requests to instructions about building and coordinating other contains to run the operations you’ve requested. You can also re-use those contains, running multiple commands in them and also share data (like files and directories) with the container to get data in and out.

## Pipelines as code

Thanks to the Dagger SDK the whole pipeline (or at least those parts you want to handle through it) becomes just another piece of code inside your repository.

At this point, SDKs for the following languages are provided:

- Go
- JavaScript/TypeScript
- Python
- “GraphQL”

GraphQL is under quotes simply because the Dagger Engine itself offers the GraphQL endpoints and so this isn’t really an SDK but simply an API. The SDK then just talks to a container running the Dagger Engine which then just talks to yet another OCI runtime in order to handle the containers.

## Implications

This has a couple of implications but most of all it means that you can just run your pipeline locally if you have an OCI runtime like Docker installed. Things like secrets and artefacts that are handled by your current CI setup still need to be injected but that’s pretty much it!

As long as you don’t require the execution of jobs on multiple machines and orchestration of all of this, Dagger should do just fine!

## The tiniest example

Let’s go and look at a small example: The CI/CD pipeline for GoGraz.org. The pipeline needs to do the following things:

1. Checkout the source code
2. Build the website inside a Hugo container
3. Save the output as artefact
4. Deploy the artefact using GitHub Pages

Out of these, Dagger will handle step 2: The building of the website.

So where do we start? As far as I can tell there’s currently no clear recommendation or standard where to put your Dagger configuration so I’ll just use `./ci/main.go` for now.

The first thing that code needs to do is to connect to the Dagger Engine:

```go
ctx := context.Background()
client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stderr))
if err != nil {
	log.Fatal(err.Error())
}
defer client.Close()
```

That’s what  `dagger.Connect` is all about. To make the execution a bit more transparent I’ve also passed the `WithLogOutput(os.Stderr)` option which will cause any operations by the engine to be logged to Stderr.

Once that connection is established, we want a Hugo container with access to the local working directory in order to build the page:

```go
workdir := client.Host().Workdir()
targetURL := "https://gograz.org"

container := client.Container().
	From("klakegg/hugo:0.107.0-ext").
	WithMountedDirectory("/src", workdir).
	WithWorkdir("/src").
	WithEnvVariable("HUGO_ENVIRONMENT", "production").
	WithEnvVariable("HUGO_ENV", "production").
	WithExec([]string{
		"--minify",
		"--buildFuture",
		"--baseURL",
		targetURL,
	})

if _, err := container.ExitCode(ctx); err != nil {
	return err
}
```

The lines above do just that and also execute the container with some flags to generate output inside the `/src/public` folder.

One thing to keep in mind here: Dagger will not actually execute anything inside the container unless you query it for an output/exit-code. So the `container.ExitCode` call is *essential* for things to actually happen.

Finally, we need the content of that folder on the host operating system again so that we can save it as artefact and the publish it.

```go
if _, err := container.
	Directory("/src/public").
	Export(ctx, filepath.Join(pwd, "public")); err != nil {
    return err
}
```

Within the GitHub Actions workflow I now just need to install Go and then run the `ci/main.go` file:

```go
jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - name: Checkout
        uses: actions/checkout@v3
        with:
          submodules: recursive
      - name: Setup Pages
        id: pages
        uses: actions/configure-pages@v2
      - name: Build with Hugo
        env:
          TARGET_URL: "${{ steps.pages.outputs.base_url }}/"
        run: go run ci/main.go build
      - name: Upload artifact
        uses: actions/upload-pages-artifact@v1
        with:
          path: ./public

```

## A more advanced example

On zerokspot.com I have a [slightly more advanced pipeline](https://github.com/zerok/zerokspot.com/blob/6697f1d5a1eb7b5947b7918a44c50e29741f2a96/ci/main.go). I won’t post the whole source file here so please take a look instead It does a handful of command execution in the same calendar in then also share files and folders within that container with another one to upload the result to a server.

```go
feedbinUsername := client.Host().EnvVariable("FEEDBIN_USER").Secret()
feedbinPassword := client.Host().EnvVariable("FEEDBIN_PASSWORD").Secret()

hugoContainer := client.Container().From("klakegg/hugo:0.107.0-ext-ubuntu").
	WithEntrypoint([]string{}).
	WithExec([]string{"apt-get", "update"}).
	WithExec([]string{"apt-get", "install", "-y", "git"}).
	WithSecretVariable("FEEDBIN_USER", feedbinUsername).
	WithSecretVariable("FEEDBIN_PASSWORD", feedbinPassword).
	WithMountedDirectory("/src", rootDirectory).
	WithWorkdir("/src").
    // That blogBin is actually a file producted in another
    // container and just mounted in:
	WithMountedFile("/usr/local/bin/blog", blogBin).
	WithExec([]string{"blog", "build-archive"}).
	WithExec([]string{"hugo"}).
	WithExec([]string{"blog", "build-graph"}).
	WithExec([]string{"blog", "blogroll", "--output", "data/blogroll.json"}).
	WithExec([]string{"hugo"}).
	WithExec([]string{"blog", "books", "gen-opml"}).
	WithExec([]string{"blog", "search", "build-mapping"}).
	WithExec([]string{"git", "rev-parse", "HEAD"})
```

This also uses “secrets” that are just environment variables on the outside. When Dagger sees those, though, they receive some special treatment in the logs and the output of the pipeline so that it’s harder to leak them.

Finally, it also generates a list of URLs that should be called via webmentions. If that list is empty, though, we can skip this step:

```go
changes, err := getChanges(ctx, "./public/.changes.txt")
if err != nil {
	return err
}

if len(changes) == 0 {
	logger.Info().Msg("No changes found. Skipping webmentions.")
	return nil
}
logger.Info().Msg("Generating webmentions")
mentionContainer := client.Container().From("zerok/webmentiond:latest").WithEntrypoint([]string{})
for _, change := range changes {
	logger.Info().Msgf("Mentioning from %s", change)
	mentionContainer = mentionContainer.WithExec([]string{"/usr/local/bin/webmentiond", "send", change})
	if _, err := mentionContainer.ExitCode(ctx); err != nil {
		return err
	}
}
```


## Some things to try out

With Dagger now in place for these two pipelines, I want to try a handful of things eventually:

- Stop using GitHub Secrets for secrets and use a central HashiCorp Vault server for that. The secrets can then be loaded directly from there and GH Secrets just need to provide me with an access token for that Vault.
- Open Telemetry support: I want to generate traces for all the pipeline executions in order to look for issues in a central place (e.g. Grafana).
- Perhaps play around with using a remote Docker instance to host the engine (just as experiment!)
- Build Docker images with Dagger instead of using Dockerfiles

## Conclusion

If you haven’t noticed it yet: I’m having a lot of fun with Dagger right now! Being able to have my pipeline inside a Go file but still use Docker for isolation is just nice!

I’ve also found a lot of useful resources including demos and presentations inside [Dagger’s Discord community](https://discord.gg/dagger-io)!

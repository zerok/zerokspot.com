---
title: Building a CLI using urfave/cli
date: "2021-01-25T19:35:10+01:00"
tags:
- golang
- cli
- development
- 100daystooffload
incoming:
- url: https://chaos.social/@zerok/105617847327860451
---

I’ve been using [Cobra](https://cobra.dev/) for a really long time. Whenever I start a project that needs to handle more than one command add Cobra. For a small toolbox for datasphere I wanted to try something different: [urfave/cli](https://github.com/urfave/cli).

Compared to Cobra, urfave/cli is much more lightweight by not integrating any kind of configuration management, fewer hooks, and also handling flags differently, but for that small project I don’t need that anyway 😅

Right when you start with this library you run into the first API difference: While in Cobra you always work with `cobra.Command` objects, in urfave/cli you start off with a `cli.App` and only work with `cli.Command` instances if you’re dealing with sub-commands:

	package main
	
	import (
	  "fmt"
	
	  "github.com/urfave/cli"
	)
	
	func main() {
	  app := cli.App{
	    Action: function(ctx *cli.Context) error {
	      fmt.Println("hello world")
	    },
	  }
	
	  app.Run(os.Args)
	}

## Flags

If you need any flags, you add them to an array that is attached to the App instance:

	cli.App{
	  ...
	
	  Flags: []cli.Flag{
	    &cli.BoolFlag{
	      Name: "verbose",
	      Aliases: []string{"v"},
	    },
	  },
	
	  ...
	}

Flag-names longer than one character are rendered with two dashes in the usage message, those with just one character receive only a single dash:

	   --verbose, -v           Verbose logging (default: false)

As you saw in the example above, the action of an app receives only a single argument: a `cli.Context` instance. This is also how you gain access to the parsed flags using functions like `GetBool(name)`. 

Don’t let that name confuse you, though: `cli.Context` doesn’t implement the `context.Context` interface but you can still work with those thanks to `cli.Context.Context` 😬

## Sub-commands

Sub-commands work pretty much like Apps but can be nested:

	app := cli.App{}
	app.Commands = []*cli.Command{
		&cli.Command{
			Name: "subcommand",
			Action: func(ctx *cli.Context) error {
				return nil
			},
			Subcommands: []*cli.Command{
				&cli.Command{
					Name: "subsubcommand",
					Action: func(ctx *cli.Context) error {
						return nil
					},
				},
			},
		},
	}

One thing to keep in mind with sub-commands, though, is that you have to pass the flags at the same level as the one where it has been declared. If a flag is defined for the application, you cannot specify it after the sub-command.

That’s a bit counter-intuitive coming from Cobra, but it’s no show-stopper for me. And that’s as far as I’ve come so far. It seems to get the job done while not getting in my way (similar to Cobra) 👍

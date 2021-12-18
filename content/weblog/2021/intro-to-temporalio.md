---
title: Intro to Temporal.io
date: "2021-12-18T18:07:45+01:00"
tags: []
---

About a month ago I listened to [episode #467 of The Changelog](https://changelog.com/podcast/467) where Adam and Jerod interviewed Shawn Wang. During that interview he mentioned a newish project called [Temporal.io](https://temporal.io/) which aims at providing a platform for building and running reliable microservices build around workflows.

Temporal consists of two main components: (1) a **server** that does the orchestration of workflows and (2) **worker processes** that implement the workflows.

To get started with a simple developer setup for the server, you can use for instance the [temporal/docker-compose](github.com/temporalio/docker-compose/) project which provides various docker-compose files including one that just launched the Temporal server including a PostgreSQL server for state handling:

	git clone git@github.com:temporalio/docker-compose.git
	cd docker-compose
	docker-compose -f docker-compose-postgres.yml up

After a couple of seconds you should then have a working server which also exposes a nice web-interface on http://localhost:8088. But all of that doesn’t yet give you any out-of-the-box workflows. So let’s go through a simple example step by step by handling each core-concept of Temporal one after another starting with *workflows*:

## Workflows

A [workflow](https://docs.temporal.io/docs/concepts/workflows) is a sequence of activities that should be performed based on a trigger or input. It is implemented as a function in whatever language you’re using and communicates with the Temporal server whenever something about its progress or state changes. That function is then registered as workflow within a worker process (e.g. a simple Go executable).

Let’s us a little example here: A workflow that implements all the activities around a user subscribing to a service:

- If a user cancels the subscription the workflow should end.
- Every month we want to remind the user of the subscription.
- Initially, there should be some kind of trial period. Once that’s over the user should receive a notification about the trial having ended.

In order to do all of that, a workflow runs various activities and can even trigger other workflows to be run right away or on a schedule.

From the outside you can even send a workflow signals (e.g. for cancelling the subscription) or querying its state. Temporal also keep track of the current state of of the workflow. This means you don’t have to create your own little state manager inside a database etc. if your workflow is complicated. The Temporal Server will do all that for you so that state becomes just another variable in your code.

Implementing a workflow is implementing a function in Go with a specific signature:

	func Subscription(ctx workflow.Context, username string) error {
		// ...
	}

Within that function you then add all the steps that should be performed as long as a user is subscribed. Once the function returns, the workflow run ends and it is marked as completed, failed, or something similar.

If you now restart the worker process that contains this function, Temporal will do its best to restore the state of the workflow and continue from there. This also means that you can do changes on the function (fixes, new feature, …), redeploy the process, and it *should* work. There are even ways to do explicit workflow versioning which are [documented here](https://docs.temporal.io/docs/go/versioning).

On thing to keep in mind here, though, is that state has to be changed in a reproducible manner. For this reason, you are not allowed to do things like calling an external API or I/O within a workflow directly. This is where *activities* come in.

## Activities

An [activity](https://docs.temporal.io/docs/concepts/activities) is a piece of business logic that is part of a workflow. Here you can do pretty much everything as long as it handles retries. If an activity fails for some reason (i.e. returning an error) then it will be retried by the workflow. Once completed, an activity can return data that can then be added to the state of the workflow and subsequent activities can use that result as input.

Once again, activities are just normal Go functions:

	func SendTrialInfo(ctx context.Context, name string) error {
		logger := activity.GetLogger(ctx)
		logger.Info("Trial starting", "name", name)
		return nil
	}

The workflow can execute an activity and wait for it’s completion like this:

	err := workflow.ExecuteActivity(ctx, SendTrialInfo, name).Get(ctx, nil)
	if err != nil {
		return err
	}

## Signals & Queries

Since workflows are potentially very long-running, you might need to interact with them or ask them about their current state. In our example workflow, we need a way to cancel the subscription (and so also the workflow) but also a way to find out if the user is still in their trial phase.

These things are handled using [signals](https://docs.temporal.io/docs/concepts/signals) and [queries](https://docs.temporal.io/docs/concepts/queries). With signals you can send messages into the workflow and the workflow can then react to them. Inside the workflow you have a selector bound to a channel through which signals are received:

	sigChan := workflow.GetSignalChannel(ctx, sigCancel)
	selector := workflow.NewSelector(ctx)
	selector.AddReceive(sigChan, func(c workflow.ReceiveChannel, more bool) {
		logger.Info("Subscription should be cancelled")
		cancelled = true
	})

	// ...
	
	for {
		if cancelled {
			break
		}
		// Select processes all the pending actions for the
		// selector and calls the respective handlers
		selector.Select(ctx)
	}

From the outside you can send signals using this:

	if err := c.SignalWorkflow(ctx, workflowID, runID, sigCancel, nil); err != nil {
		logger.Error().Err(err).Msg("Failed to signal workflow")
	}

On the other hand, workflows can also provide query handlers so that they can answer questions from the outside:

	cancelled := false
	inTrial := true
	
	workflow.SetQueryHandler(ctx, "current_state", func() (map[string]bool, error) {
		return map[string]bool{
			"cancelled": cancelled,
			"inTrial":   inTrial,
		}, nil
	})
	

You can then, for instance, go to the workflow run in the web-ui and execute the `current_state` query to get information about the cancelled state and if the workflow is still in its trial phase.

## The whole implementation

Let’s put everything together now. The whole implementation of the workflow and activities just takes about 140 lines:

	import (
		"context"
		"flag"
		"fmt"
		"os"
		"time"
	
		"github.com/rs/zerolog"
		"go.temporal.io/sdk/activity"
		"go.temporal.io/sdk/client"
		"go.temporal.io/sdk/worker"
		"go.temporal.io/sdk/workflow"
	)
	
	const queue = "subscriptions"
	const sigCancel = "cancelSubscription"
	
	func SendSubscriptionReminder(ctx workflow.Context) error {
		logger := workflow.GetLogger(ctx)
		logger.Info("Sending reminder")
		return nil
	}
	
	// Cancelling the manager will also cancel the subruns
	func SendSubscriptionReminderManager(ctx workflow.Context) error {
		cancelled := false
		sigChan := workflow.GetSignalChannel(ctx, sigCancel)
		selector := workflow.NewSelector(ctx)
		selector.AddReceive(sigChan, func(c workflow.ReceiveChannel, more bool) {
			cancelled = true
		})
	
		cctx := workflow.WithChildOptions(ctx, workflow.ChildWorkflowOptions{
			CronSchedule: "* * * * *",
		})
		workflow.ExecuteChildWorkflow(cctx, SendSubscriptionReminder)
		for {
			if cancelled {
				break
			}
			selector.Select(ctx)
		}
		return nil
	}
	
	func Subscription(ctx workflow.Context, name string) error {
		logger := workflow.GetLogger(ctx)
		logger.Info("User has subscribed", "name", name)
		trialDuration := time.Second * 30
	
		var reminderWorkflow workflow.ChildWorkflowFuture
	
		// All activities need to be run within a certain context that configures
		// their retry policy and various timeouts. For this one we just say that
		// each activity must not take longer than a minute.
		ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
			StartToCloseTimeout: time.Minute,
		})
	
		cancelled := false
		inTrial := true
	
		workflow.SetQueryHandler(ctx, "current_state", func() (map[string]bool, error) {
			return map[string]bool{
				"cancelled": cancelled,
				"inTrial":   inTrial,
			}, nil
		})
	
		// In order to learn if a user wants to cancel or not, we have to be able
		// to process signals:
		sigChan := workflow.GetSignalChannel(ctx, sigCancel)
		selector := workflow.NewSelector(ctx)
		selector.AddReceive(sigChan, func(c workflow.ReceiveChannel, more bool) {
			logger.Info("Subscription should be cancelled")
			cancelled = true
	
			// Also cancel the reminder workflow
			if reminderWorkflow != nil {
				reminderWorkflow.SignalChildWorkflow(ctx, sigCancel, nil)
			}
		})
	
		if cancelled {
			return nil
		}
	
		err := workflow.ExecuteActivity(ctx, SendTrialInfo, name).Get(ctx, nil)
		if err != nil {
			return err
		}
	
		// Now we can wait for the trial period to end. If the subscription is
		// cancelled during this time, we can skip the period and proceed right
		// away.
		workflow.AwaitWithTimeout(ctx, trialDuration, func() bool {
			return cancelled == true
		})
	
		if !cancelled {
			inTrial = false
			if err := workflow.ExecuteActivity(ctx, SendTrialEndInfo, name).Get(ctx, nil); err != nil {
				return err
			}
		}
	
		// Now we start a workflow that will send out a reminder every minute:
		if !cancelled {
			cctx := workflow.WithChildOptions(ctx, workflow.ChildWorkflowOptions{})
			reminderWorkflow = workflow.ExecuteChildWorkflow(cctx, SendSubscriptionReminderManager)
		}
	
		for {
			if cancelled {
				break
			}
			selector.Select(ctx)
		}
	
		if reminderWorkflow != nil {
			logger.Info("Waiting for reminder workflow")
			reminderWorkflow.Get(ctx, nil)
			logger.Info("Reminder workflow complete")
		}
	
		logger.Info("Done")
		return nil
	}
	
	func SendTrialInfo(ctx context.Context, name string) error {
		logger := activity.GetLogger(ctx)
		logger.Info("Trial starting", "name", name)
		return nil
	}
	
	func SendTrialEndInfo(ctx context.Context, name string) error {
		logger := activity.GetLogger(ctx)
		logger.Info("Trial ended", "name", name)
		return nil
	}
	
	func main() {
		ctx := context.Background()
		logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout})
		ctx = logger.WithContext(ctx)
		c, err := client.NewClient(client.Options{})
		if err != nil {
			logger.Fatal().Err(err).Msg("Failed to create client")
		}
		defer c.Close()
	
		logger.Info().Msg("Registering worker")
		w := worker.New(c, queue, worker.Options{})
		w.RegisterWorkflow(Subscription)
		w.RegisterWorkflow(SendSubscriptionReminder)
		w.RegisterWorkflow(SendSubscriptionReminderManager)
		w.RegisterActivity(SendTrialInfo)
		w.RegisterActivity(SendTrialEndInfo)
		if err := w.Run(worker.InterruptCh()); err != nil {
			logger.Fatal().Err(err).Msg("Failed to run worker")
		}
	}

This is a complete worker implementation. If you now run that file it registers itself on the server and the server can then use it for workflow runs:

	go run main.go

Another part of your application that would now trigger such a workflow can do so using this piece of code:

	run, err := c.ExecuteWorkflow(ctx, client.StartWorkflowOptions{
		TaskQueue: queue,
	}, Subscription, "Horst Gutmann")
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to execute workflow")
	}
	wfID := run.GetID()
	runID := run.GetRunID()
	fmt.Printf("Workflow ID: %s\nRun ID: %s\n", wfID, runID)

In order to query the workflow and send signals you’ll have to store the workflow and run ID somewhere but that’s about it.

## Final thoughts

I haven’t yet used Temporal for anything productive. In fact, I’ve only played around with it for a couple of hours but my final initial thoughts of it are very positive. This looks like a very convenient system for implementing workflows that should also scale quite will. It comes with some concepts that are quite new to me so if I’d pick it for some actual project, I’d probably need a bit of time  for getting used to how stuff should work in this system compared to “classic” approaches. Sure, there have been other systems with a similar approach (e.g. [Uber’s Cadence](https://github.com/uber/cadence)) but since I haven’t worked with these either, I’m not yet used to this approach.

Given the good documentation around the project and lots of talks and presentations, learning shouldn’t be the issue here, though. Let’s see. Perhaps I’ll have some ideas around the GoGraz user-group where something like Temporal might come in handy!
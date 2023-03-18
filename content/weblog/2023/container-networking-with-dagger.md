---
title: Container networking with Dagger
date: "2023-03-18T21:18:21+01:00"
tags:
- dagger
- cicd
- golang
- containers
---

One of the biggest limitations I saw with Dagger up until now was that you couldn’t connect containers on a network level. This prevented use-cases like testing against database containers (e.g. for end-to-end testing scenarios). 

[Dagger 0.4](https://dagger.io/blog/dagger-engine-0-4-0) now supports what they call “services”. You can create a container, define exposed ports, and then add it as a dependency to another container so that it can connect via TCP to it. I’ve used that so far for verifying that a binary I’m building for [gograz-meetup](https://github.com/gograz/gograz-meetup/blob/master/magefiles/main.go) actually starts a working HTTP server:

	// Create a backend container that exposes its service on
	// port 8080:
	backendContainer := goContainer.WithExposedPort(8080).
	    WithExec([]string{"./gograz-meetup", "--addr", "0.0.0.0:8080"})
	
	// Bind the backendContainer as service with the hostname
	// "backend" to a new container:
	code, err := goContainer.
		WithServiceBinding("backend", backendContainer).
		WithExec([]string{"curl", "--fail", "http://backend:8080/alive"}).
		ExitCode(ctx)
	
	if err != nil {
		return err
	}

This is just the simplest application of services. You can learn more (especially about the life-cycle of containers and best-practices on persisting data) in the [official documentation](https://docs.dagger.io/757394/use-service-containers/)!

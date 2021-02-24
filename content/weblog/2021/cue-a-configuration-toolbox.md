---
title: 'CUE: A configuration toolbox'
date: "2021-02-24T21:07:34+01:00"
tags:
- 100daystooffload
- configuration
- opensource
- validation
---

While making my way through my podcast backlog I listened to [Go Time #163](https://changelog.com/gotime/163) today which was about a project I hadn’t heard of before: [CUE](https://cuelang.org/).

CUE is a language for basically defining and validating various kinds of data. That sounds extremely generic, so what are some of the use-cases with CUE?

## Configuration validation

For me, the first thing I wanted to try out is if I could write a simple schema for the configuration files I use for the [podcast-archiver](https://github.com/zerok/podcast-archiver) project. A sample configuration file (in YAML) looks like this:

	sink:
	  dropbox_oauth_access_token: "12345"
	feeds:
	  - url: https://changelog.com/gotime/feed
	    name: "gotime"

The `sink` is quite complex, though, as it might have different properties depending on what target system should be used. This example uses Dropbox as sink but the project also supports Google Cloud Storage and S3.

With CUE I’ve now defined the following schema and saved it inside a `configuration.cue` file:

	sink: #Sink
	feeds: [...#Feed]
	
	#Sink: #GoogleSink | #DropboxSink | #S3Sink
	
	#GoogleSink: {
	  google_project_id: string,
	}
	
	#DropboxSink: {
	  dropbox_oauth_access_token: string,
	}
	
	#S3Sink: {
	  bucket: string,
	  region: string,
	  access_key_id: string,
	  access_key_secret: string,
	}
	
	#Feed: {
	  url: string,
	  folder: string,
	}

Let’s go through what this all means:
- The file should contain two fields: `sink` and `feeds`.
- A sink is either a GoogleSink, DropboxSink, or S3Sink.
- All the sinks have different required fields

If the sample configuration I’ve posted above is stored inside a `config.yaml` file, then I can validate it against the schema with the following command:

	$ cue vet config.yaml configuration.cue
	feeds.0: field `name` not allowed:
	    ./config.yaml:5:6
	    ./configuration.cue:2:12
	    ./configuration.cue:21:8

Turns out, I had a typo inside the configuration above. Feeds don’t have a `name` property! What should have been in there instead is a `folder`. After changing that, vet returns without any error.

## File format

The CUE file format itself is basically a superset of JSON. It comes with a couple of additions (like comments) and shortcuts esp. regarding objects that only contain a single field. The example above converted from YAML to CUE would look like this with the sink now fitting into a single line:

	sink: dropbox_oauth_access_token: "12345"
	feeds: [{
		url:    "https://changelog.com/gotime/feed"
		folder: "gotime"
	}]

CUE is a lot more powerful than something like YAML as you can also re-use other definitions through [packages](https://cuelang.org/docs/tutorials/tour/packages/imports/)! Combined with the template-based approach for re-use in configurations this should be immensely useful for large configurations.

One part of that is that CUE will merge data as long as it doesn’t stay in conflict. Let’s say that I want to have multiple feeds and every single one should have the same `folder` value set:

	sink: dropbox_oauth_access_token: "12345"
	feeds: [...{folder: "gotime"}]
	feeds: [{
		url:    "https://changelog.com/gotime/feed1"
	},{
		url:    "https://changelog.com/gotime/feed2"
	}]

Here I defined that every item inside `feeds` should have the folder `gotime` but with both items the URL should be different. This works because the two definitions don’t contradict each other.  The following configuration would not work as the folder on the second item would contradict the definition in line 2:

	sink: dropbox_oauth_access_token: "12345"
	feeds: [...{folder: "gotime"}]
	feeds: [{
		url:    "https://changelog.com/gotime/feed1"
	},{
	    folder: "gotime2",
		url:    "https://changelog.com/gotime/feed2"
	}]

## Interoperability

OK, so there is yet another configuration file format … that’s not supported by the apps I’m using. Luckily,  CUE comes with two commands that allow you to convert YAML/JSON files into CUE files and the other way around:

	# This will create a something.cue file:
	$ cue import something.yaml
	
	# And back (just to STDOUT):
	$ cue export something.cue --out yaml

All of a sudden you can use the power of CUE for pretty much every system that support JSON or YAML configuration files! 

In general, I really like what I’ve seen of CUE so far. I’ve only played around with it for a couple of hours but I can absolutely see myself using it more and also spending some time looking into the packaging features here in order to simplify, for instance, some of my Kubernetes configurations.

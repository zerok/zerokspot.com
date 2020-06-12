---
title: Automating notes on iOS
date: "2020-06-12T18:57:56+02:00"
tags:
- automation
- getdrafts
- ios
- workingcopy
- 100daystooffload
---

I’ve finally started using one of the big features of [Drafts](https://getdrafts.com/) properly: Actions. For the longest time I just did some basic reformatting: adding a Hugo-compatible frontmatter to drafts before moving them manually into WorkingCopy. Since I’m now using notes on here more and more to collect links, comment on other people’s posts etc. I grew tired of creating all these folders in the Git repository manually.

Turns out, [WorkingCopy](https://workingcopyapp.com/) (the app I use for any interaction with Git) has quite extensive support for automating such things either through Siri Shortcuts or through x-callback-urls. Drafts, at the other end of the workflow pipeline, supports calling x-callback-urls as part of an action and so I now have the following steps automated:

1. Pull the latest upstream changes for my website’s repository
2. Create all the necessary folders within it based on the current date
3. Create a new file and put the content of the current draft in it
4. Commit and push using the name of the draft

I originally wanted to show you just the action consisting of various “Callback URL” steps but I thought it might be more useful to present the JavaScript version 😅 An hour later and here you go:

	function guessTitle(content) {
		const titleRe = new RegExp('^title: "?(.*)"?$');
		for (let line of content.split('\n')) {
			let mo = titleRe.exec(line);
			if (mo) {
				var raw = mo[1];
				raw = raw.toLowerCase().replaceAll(/[^a-z0-9]/g, '-').replaceAll(/--/g, '-');
				return raw;
			}
		}
		return null;
	}
	
	function prepareCallback(cmd, apiKey) {
	  const cb = new CallbackURL();
	  cb.waitForResponse = true;
	  cb.baseURL = `working-copy://x-callback-url/${cmd}/`;
	  cb.addParameter('key', apiKey);
	  return cb;
	}
	
	(function() {
	  // Ask the user for their WorkingCopy API key and store it
	  // for later calls within Draft:
	  const credential = Credential.create("WorkingCopy", "Working Copy");
	  credential.addTextField("apiKey", "API key");
	  credential.authorize();
	
	  const repo = 'zerokspot.com';
	  const wcBaseURL = 'working-copy://x-callback-url';
	  const wcKey = credential.getValue('apiKey');
	  const dateFolder = strftime(new Date(), '%Y/%m/%d');
	
	  // Try to get a good slug/filename based on the title of
	  // the post but let the user override it:
	  var slug = guessTitle(draft.content);
	  const slugPrompt = new Prompt();
	  slugPrompt.addTextField('slug', 'Post slug', slug);
	  slugPrompt.isCancellable = true;
	  slugPrompt.addButton('OK');
	  if (!slugPrompt.show()) {
	    context.cancel();
	    return
	  }
	  slug = slugPrompt.fieldValues['slug'];
	
	  // Now let's fetch the latest changes from upstream:
	  var cb = prepareCallback('pull', wcKey);
	  cb.addParameter('repo', repo);
	  if (!cb.open()) {
	    context.fail('Failed to pull from upstream');
	    return;
	  }
	
	  cb = prepareCallback('write', wcKey);
	  cb.addParameter('repo', repo);
	  cb.addParameter('text', draft.content);
	  cb.addParameter('path', `content/notes/${dateFolder}/${slug}.md`);
	  if (!cb.open()) {
	    context.fail('Failed to write note');
	    return;
	  }
	
	  cb = prepareCallback('commit', wcKey);
	  cb.addParameter('repo', repo);
	  cb.addParameter('text', draft.content);
	  cb.addParameter('path', `content/notes/${dateFolder}/${slug}.md`);
	  cb.addParameter('message', slug);
	  if (!cb.open()) {
	    context.fail('Failed to commit note');
	    return;
	  }
	
	  cb = prepareCallback('push', wcKey);
	  cb.addParameter('repo', repo);
	  if (!cb.open()) {
	    context.fail('Failed to push note');
	    return;
	  }
	})();

This was also the first time that I did anything with [x-callback-urls](http://x-callback-url.com/), a mechanism I didn’t know even existed before wanting to implement this automation and now … I know that it exists and also to some degree how to use it 😅

---
date: '2009-09-09T12:00:00-00:00'
language: en
tags:
- ckeditor
- javascript
title: Custom filebrowser callbacks in CKEditor 3.0
---


Last weekend I started playing around a little bit with the recently released <a href="http://ckeditor.com">CKEditor</a> 3.0. I had wanted to have a WYSIWYG editor for my blog for quite a while now but was always either too lazy in general or just too lazy to write a custom uploader.&nbsp; Last weekend was perfect for something like that, though, since I didn&#39;t receive a book and a DVD-box I had originally planed to spend the whole weekend on ;-)

--------------

The first part was naturally easy: Downloading it, placing all the right JS and CSS files on the server and writing some small JavaScript-embed code so that I can dynamically switch from Markdown and ReST (which I use normally for writing posts here) to HTML which automatically should then load the editor.

Then came the part I never found the motivation for before: Writing about 50 lines of code and about 200 lines of HTML and CSS to get the custom file uploader and browser to work. This part was also quite simple for CKEditor at first, but then reached a hurdle: What are the callback functions that actually tell the original upload dialog the URL of the file? The only <a href="http://docs.fckeditor.net/FCKeditor_2.x/Developers_Guide/Customization/Custom_File_Browser">official documentation</a> I could find was about FCKEditor which used something like <code>window.opener.SetUrl(yourUrl)</code> when you select a file from the browser and probably a similar function when confirming the upload. Well, this doesn&#39;t work anymore with CKEditor 3.0.

The new version has a real callback system where the caller actually can determine the callback code. So the link that opens for instance the filebrowser or the upload-form&#39;s action includes an additional numeric parameter called &quot;CKEditorFuncNum&quot; that tells the callee a parameter to be used for <code>window.opener.CKEDITOR.tools.callFunction</code>.

So all you have to do from the <strong>file browser</strong> when you have a file selected is to call this code with the right callback number (normally 1) and the URL of the selected file:

<pre class="code">window.opener.CKEDITOR.tools.callFunction(CKEditorFuncNum, url);</pre>

For the quick-uploader the process is quite similar. At first I thought that the editor might be listening for a 200 HTTP return code and perhaps look into some header field or something like that to determine the location of the uploaded file, but then - through some <a href="http://getfirebug.com/">Firebug</a> monitoring - I noticed that all that happens after an upload is the following code:

<pre class="code">&lt;script type=&quot;text/javascript&quot;&gt;
  window.parent.CKEDITOR.tools.callFunction(CKEditorFuncNum, 
    url, errorMessage);
&lt;/script&gt;
</pre>

If the upload failed, set the errorMessage to some non-zero-length string and empty the url, and vice versa on success.

Something like this could perhaps be a little bit better documented :-)</p>

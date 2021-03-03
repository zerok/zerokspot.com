---
date: '2008-02-04T12:00:00-00:00'
language: en
tags:
- java
title: 'TinyTip: Where are my cookies in commons-httpclient?'
---


[commons-httpclient](http://hc.apache.org/httpclient-3.x/) is for me perhaps one of the most useful libraries in Java when it comes to doing HTTP requests. It not only abstracts HTTP logically but at the same time keeps you at a low enough level. That said, it still abstracts some stuff perhaps a little bit too far away for some use cases. 

-------------------------------

When you want to see what cookies the server responded with on a request, I personally would look first directly in the response headers:
    
    HttpClient client = ...;
    HttpMethod method = ...;
    client.executeMethod(method);
    for(Header h : method.getResponseHeaders()){
        ...
    }
    
But httpclient comes with a feature that I hadn't anticipated there: It automatically manages the HttpClient's cookies unless you tell not to. That said, I'm currently working on a program where I want to store the cookies over multiple sessions so I wanted to access them after receiving a response and then also attach them to a different HttpClient instance. After some browsing through the libs' javadocs I finally also found where the cookies are stored: In the state of the client itself:
    
    for(Cookie c : client.getState().getCookies){
        ...
    }

So if all you want are the cookies, you can get them here. But HttpState doesn't only hold the cookies, but also HTTP credentials, so if you're interested in those, you might as well just move/dump&restore the whole HttpState :-)

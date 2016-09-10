---
date: '2010-02-09T12:00:00-00:00'
language: en
link: http://gowalla.com/api/
tags:
- api
- geolocation
- gowalla
title: Gowalla finally got an API
url_title: Gowalla Developers on Gowalla
---


It's been <a href="http://getsatisfaction.com/gowalla/topics/gowalla_api">quite a while in the making</a>&nbsp;but Gowalla today finally made an API for their geo-location service <a href="http://gowalla.com/blog/2010/02/announcing-the-gowalla-api/">publicly available</a> :-) Compared to what Foursquare has to offer, this one is rather limited because it&#39;s&nbsp;<em>read-only</em>. But this is IMO actually a good thing since then it&#39;s not that simple to easily cheat the system from the outside. On the other hand, the <a href="http://gowalla.com/api/docs">documentation</a> mentions POST-requests often enough to suggest that check-in support or some other operation that requires write-access is planned for the future.

<meta charset="utf-8" /><meta charset="utf-8" id="webkit-interchange-charset" /></p>

Another absent feature is OAuth support, which isn&#39;t really an issue right now, since you barely need to do authenticated requests anyway, right now. I guess, we will see it when the time comes for write-access. Applications, on the other side, require an application key in order to access any of the provided data.

Given that the API is read-only for the time being, there are some obvious limitations of what you can do with it. Nevertheless, this would have been more than enough to add some more information onto the Twitterwall for the recent <a href="http://www.barcamp.at/BarCamp_Klagenfurt_2010">Barcamp Klagenfurt 2010</a> and I&#39;m pretty sure, many people will create some useful and some just plain funny tools with it :-)

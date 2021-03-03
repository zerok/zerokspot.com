---
date: '2008-05-07T12:00:00-00:00'
language: en
tags:
- photography
- online-services
title: Some Flickr alternatives
---


Every once in a while it's good to take a look what alternatives to a service you use and like very much. So given the current Microsoft/Yahoo-dicussions, I thought it was once again time to look at some of the "alternative" photo communities out there: [Smugmug](http://smugmug.com), [Zooomr](http://zooomr.com), [23hq](http://23hq.com),  [Fotocommunity](http://fotocommunity.com) and [ipernity](http://ipernity.com). I didn't really try most of these services but just looked at their feature-set to see if they offered what I need. Also don't expect a complete comparison off these sites here, since this would probably easily fill a 20 pages article. I just want to point out a few things I've noticed.

So let's get started.

-------------------------------

## Smugmug

[Smugmug](http://smugmug.com) is the only one of the services I looked at that doesn't offer a free account nor do they have the usual ~20 USD account. Their pricing starts at 39.95 USD and will probably keep many people out of the community who just want to get started with that whole photography 2.0 thing (although, given the weak dollar right now, 40 USD for a whole year shouldn't really stop anyone). 

On the feature-side, Smugmug has it all: Photosets (called galleries there), tagging, geotagging, photogroups/communities, unlimited storage and traffic, OpenID support and much much more. For me personally important is a way to easily get all of my photos off the system again if I want to move on. There are generally 2 ways I would solve this: (1) By offering an API or (2) by offering generated archives for photosets (or even the whole account). While Smugmug first of all offers an API, they also have multiple ways to generate backups of your whole account. You can either order DVDs or CDs with your photos on them, or you can use something they call [SmugDAV](http://wiki.smugmug.com/display/SmugMug/SmugDAV), which I guess boils down to a read-only WebDAV-access to your whole account. If true, this is absolutely stunning.

One feature that Smugmug is missing, though, is that communities lack a forum like section to discuss topics relevant to this community. This way the interaction within a group is limited to comments on photos and I can only assume that it will be hard to create some kind of contest-centric communities without setting up some kind of external board.

For me personally this is a real alternative to Flickr thanks to its rich feature-set. Compared to what Flickr and to some degree Zooomr and 23hq offer, the UI isn't as polished, although the JavaScript loading of photos within a group ([example](http://themes.smugmug.com/gallery/1519826_ftW6B/1#73017541_vPBya)) is a good idea. But this is a non-issue if you consider a "Power User"-Account (for 59.95 USD) which allows you to completely customize your theme by applying your own CSS rules. At this account-level you can also get your own [DNS handler](http://www.smugmug.com/help/power-custom-domain) to make your galleries appear even more integrated with your site than a mere layout could do.


## Zooomr

Back when I first saw [Zooomr](http://zooomr.com), it appeared like a shameless Flickr ripoff. While some aspects of the design have changed, there are still too many similarities in the layout for me, but OK, since I guess you're actually looking for something exactly like Flickr but simply not Flickr itself let's ignore this for now.

First of all: This is the only photo community site in this comparison, where the free account is actually worth something.

<blockquote>
    <p>Zooomr is free.  There is no upload limit, bandwidth limit, download limit, file size limit, set limit, or number of photo limits. </p>
    <cite><a href="http://zooomrlearnmore.wiki.zoho.com/Zooomr-FAQ-June-15-2007.html">Zooomr's help section</a></cite>
</blockquote>

That said, this site has its share of problems:

*   When I first signed up for it, I couldn't even change my password (not that this was I guess more than a year after the site was launched).
*   According to [this blog post](http://blog.zooomr.com/2006/03/27/attention-developers/) from more than 2 years ago, there is an API that tries to emulate Flickr's in order to make it easier for app-developers to support it, yet it's way to complicated (impossible?) to get your hands on an API key as this email by Kristopher Tate shows:
    
    <blockquote><p>Okay guys, to help me out, I need your Zooomr Account URLs so that I
    can link your account to the API Key.</p>

    <p>I also need a formal description and url where users can click. Thank
    much.</p>

    <p>kristopher</p><cite><a href="http://groups.google.com/group/zooomr-dev/msg/b0d15cda2b05a188">Mail on the Google Groups list</a></cite></blockquote>
    
    ... and in a mail from this March:
    
    <blockquote><p>Cmoewes & Friends, the API has been launched with Z2k8! :)</p>

    <p>I want to get you guys in before we open-up the flood gates, so please
    send me an email at kristopher [at] bluebridge.jp with [ZAPI]
    <APP_NAME_HERE> as the title, and I'll personally register your api
    key.</p>

    <p>Please also send a URL and description of your app, along with the URL
    of your Zooomr Account (example: zooomr.com/kristopher )</p>

    <p>Thanks for waiting patiently - I can't wait to see what we can come up
    with together!</p>

    <p>kristopher</p><cite><a href="http://groups.google.com/group/zooomr-dev/msg/6730bc7190b589c4">Mail on the Google Groups list</a></cite></blockquote>
    
    With that I'm not sure how I should count the API. Yes, it exists. No, you won't get easy access to it for now. I guess there is from now on a "Yes", a "No" and a "Meh" a scores for rankings in this category, with Zooomr getting a "Meh".
    
*   Sometimes the service looks a little bit unreliable. Back when I first tried it, the missing password-recovery turned me down, so I ignored Zooomr for nearly a year. Then I wanted to check it out again just to see the OpenID login being totally broken. Sure, probably just bad timing, but it still stuck in my head.


<div class="figure">
    <img src="http://img.skitch.com/20080503-k39f43k4awp6uuawp2y7ink31i.png" alt="" />
    <p class="caption">Choose what's in your SmartSet on Zooomr</p>
</div>


One feature, that Zooomr has that no one else has, though, are the so-called "SmartSets". From Flickr you're probably used to a set being more or less just a container where you put photos into to easier find them or to tell a story with them. SmartSet are something like stored search-queries, where a SmartSet for instance can contain all of your photos tagged with "pizza" and "restaurant" and having been viewed by more than 100 people. While a great idea, it has one drawback, though: You can't manually sort the resultset, which is interesting one some occasions. But for most situations, this feature is just great (although probably quite intense on the backend).

When it comes to tagging, Zooomr (together with ipernity) also lets you link a photo with accounts of other Zooomr users to show, that they were in the picture. Definitely a welcome improvement over the situation over at Flickr.

As with Smugmug, Zooomr offers you a way to customize your site (even on the free account) but limits you (similar to how [Pownce](http://pownce.com) does it) to choosing from a set of themes or changing a couple of colors and setting a background image. So integrating Zooomr completely with your site (layout-wise) won't be possible unless you have a really basic design there.

A feature that I was really missing in Flickr, is a way to find photos takes with a specific lense or a specific camera. I was actually quite and very positively surprised by Zooomr actually having this and basically a filter for each EXIF-tag associated with a photo. This could well be my personally killer feature here (hence I added it to the summary).

<div class="figure">
    <img src="http://img.skitch.com/20080509-b5k1paxpd4ksd5cq658227xsnh.png" alt="" />
    <p class="caption">Filters per EXIF-tag</p>
</div>

Zooomr is definitely a nice service and I see why many people like it, it's just probably nothing for me. There are just a couple of details like the whole "Getting an API-key" business that would annoy me way too much.

## 23hq

I've already written about [23hq](http://23hq.com) [quite](http://zerokspot.com/weblog/920/) [a lot](http://zerokspot.com/weblog/470/) in the last years, so I'll try to keep this short here: Good small site with a very simple yet nice interface, an API, discussions in photo communities. A serious downside for me is the lack of geotagging. The free account has its limit at 30 photos per month, so it's basically a trial yet still usable.

## Fotocommunity

[Fotocommunity](http://fotocommunity.com) seems to be quite popular esp. in the German scene and seems to be targeted at (semi-) professionals. It features a completely different structure than sites like Flickr or Zooomr with photos not being organized by tags or sets or something like that but by a number of categories and subcategories with "Subjects", "Specials", "People", "Nude", "Nature", "Digiart", "Youth" and "World" being at the very top. So no tagging here. No organizing your *own* photos. You can create folders but IMO the whole site looks kind of "old" in this department. For example there is also only one big forum or seems there to be anything comparable to the groups of Flickr.

<div class="figure">
    <img width="450px" src="http://img.skitch.com/20080503-tb76pba4sme7tn61t5yb76ct.png" alt="" />
    <p class="caption">Paid accounts on <a href="http://fotocommunity.com">fotocommunity</a></p>
</div>

For those of us who have enough photos to need a pro-account on Flickr, you will also need one on Fotocommunity. But honestly, the pricing for the paid accounts isn't even remotely comparable to what Flickr offers. I currently have more than 2500 photos on my Flickr account and I'd have to go with a "World" account for 8 EUR per month. Sure, I'm probably not in the demography for this site since I tend to also upload photos that I myself don't consider good, but this is really too much.

If you can live with being able to just upload a very small amount of photos per day, take a look, otherwise there are in my opinion many better options out there.

## ipernity

If Flickr and Virb could have a child together, I guess [ipernity](http://ipernity.com) would come pretty close to what would it look like. The whole UI looks very inspired by Flickr (even more so than Zooomr, yet more polished), but on the feature-side ipernity offers not only sharing of photos but also videos and audio-files -- called "documents" here. Still, the site seems to be very photo-centric.

But wait a second. I said, the UI seemed to be "inspired" by Flickr. <s>inspired</s>. In most sections it honestly looks more like a straight copy without actually bringing the whole functionality over from the original. For example: You can put <s>photos</s> documents into albums, but you cannot have albums of albums.

On the other hand, ipernity also improves on Flickr's UI in some places. For example -- as mentioned in the section about Zooomr -- if you're logged in and view a photo, you can add a "member tag" to a photos in order to show, that this person is actually in this photo. This is a really nice idea esp. since on Flickr I often see tagging photos manually with their name in order to get or less the same result, although without the distinction between a normal tag and a tag that's supposed to represent a person.

<div class="figure">
    <img src="http://img.skitch.com/20080509-nrn9f4swm19rtmity3rg3ctxib.png" alt=""/>
    <p class="caption">"I am in this photo"</p>
</div>

ipernity offers 2 types of accounts: A free one, that let's you upload 200 MB per month and only shows you your latest 1000 documents; and a pro account for about 24 EUR/year which removes those restrictions, ads and lets you download photos in their original format and size.

As with Zooomr, you can also customize your home page there more or less to the same extend as on Kristopher Tate's service. So no 100% integration with your site, but at least to some degree.

From what I can tell, there is currently not really anything like an "open" API in the sense that anyone can make their applications easily compatible with the site. There are some mostly 1st party apps that connect to the site (uploader, etc.) but I couldn't find any documentation on this obviously existing connectivity nor did I find any hints on when such an API might become open to the public.

A little bit weird is, that I couldn't find any information about who is actually behind this site. whois doesn't return anything useful, there is no imprint, nothing on the about page, and only an address in the [Privacy Policy](http://www.ipernity.com/about/you). No idea, why there isn't any info on this on the about page :-/

Anyway, the site looks really nice despite being quite a copycat in some areas. The community seems to be quite big, too (around 29,000 users judging from [this account](http://www.ipernity.com/user/lea.ipernity) who automatically adds anyone who registers as a contact). For me personally the mixture of photos, audio and video is a little bit too much. I wasn't really a fan of Flickr adding video functionality (although they at least did it in a way, that was not completely ortogonal to what they did before), so ipernity for me personally tries to many things. From a photography standpoint, the EXIF tags in Zooomr are just more interesting than videos and audio-files.

I guess, by now you have noticed, that I'm not really sure how I should feel about this site. One the one hand it's nice with a fast and slick interface, but on the other hand, it doesn't offer anything truly unique for what I would use it for -- namely sharing and finding photos.

## Summary

Not that I don't even consider posting about limits here, because as always, the free accounts are always (except for Zooomr as mentioned above) basically trials and nothing more. 

<table>
    <thead>
        <tr>
            <th></th>
            <th>Flickr</th>
            <th>Smugmug</th>
            <th>Zooomr</th>
            <th>23hq</th>
            <th>Fotocommunity</th>
            <th>ipernity</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td class="feature">Free account?</td>
            <td class="flickr">Yes</td>
            <td class="smugmug">No</td>
            <td class="zooomr">Yes</td>
            <td class="23hq">Yes</td>
            <td class="fotocommunity">Yes</td>
            <td class="ipernity">Yes</td>
        </tr>
        <tr>
            <td class="feature">Price(s) for pro account</td>
            <td class="flickr">25 USD/year</td>
            <td class="smugmug">39.95 USD/year<br />59.95 USD/year<br />149.95 USD/year</td>
            <td class="zooomr">19.95 USD/year</td>
            <td class="23hq">20 EUR/year</td>
            <td class="fotocommunity">48 EUR/year<br />72 EUR/year<br />96 EUR/year</td>
            <td class="ipernity">23.88 EUR</td>
        </tr>
        <tr>
            <td class="feature">API</td>
            <td class="flickr">Yes</td>
            <td class="smugmug">Yes</td>
            <td class="zooomr">Meh</td>
            <td class="23hq">Yes</td>
            <td class="fotocommunity">No</td>
            <td class="ipernity">No</td>
        </tr>
        <tr>
            <td class="feature">Photogroups</td>
            <td class="flickr">Yes</td>
            <td class="smugmug">Yes</td>
            <td class="zooomr">Yes</td>
            <td class="23hq">Yes</td>
            <td class="fotocommunity">No</td>
            <td class="ipernity">Yes</td>
        </tr>
        <tr>
            <td class="feature">Discussions</td>
            <td class="flickr">Yes</td>
            <td class="smugmug">No</td>
            <td class="zooomr">Yes</td>
            <td class="23hq">Yes</td>
            <td class="fotocommunity">No</td>
            <td class="ipernity">Yes</td>
        </tr>
        <tr>
            <td class="feature">Organization</td>
            <td class="flickr">Tags, Sets</td>
            <td class="smugmug">Tags, Galleries</td>
            <td class="zooomr">Tags, <a href="#smartsets">SmartSets</a></td>
            <td class="23hq">Tags, Albums</td>
            <td class="fotocommunity">Global(!) categories</td>
            <td class="ipernity">Albums, Tags</td>
        </tr>
        <tr>
            <td class="feature">OpenID</td>
            <td class="flickr">?</td>
            <td class="smugmug">Yes</td>
            <td class="zooomr">Yes</td>
            <td class="23hq">No</td>
            <td class="fotocommunity">No</td>
            <td class="ipernity">Yes</td>
        </tr>
        <tr>
            <td class="feature">Geotagging</td>
            <td class="flickr">Yes</td>
            <td class="smugmug">Yes</td>
            <td class="zooomr">Yes</td>
            <td class="23hq">No</td>
            <td class="fotocommunity">No</td>
            <td class="ipernity">Yes</td>
        </tr>
        <tr>
            <td class="feature">EXIF-Filter</td>
            <td class="flickr">No</td>
            <td class="smugmug">No</td>
            <td class="zooomr">Yes</td>
            <td class="23hq">No</td>
            <td class="fotocommunity">No</td>
            <td class="ipernity">No</td>
        </tr>
    </tbody>
</table>

## Updates

*   2008.05.09:
    *   Added ipernity (thanks, [Patrick](http://www.schreiblogade.de/))
    *   Noticed the EXIF-tags in Zooomr
    

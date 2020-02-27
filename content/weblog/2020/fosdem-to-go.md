---
title: "FosdemToGo: A new FOSDEM companion app for iOS"
date: "2020-02-27T17:25:24+0100"
---

Over the course of the last two months I've been playing around with Swift. As with any kind of learning I want to combine it with an actual practical need so I started looking around for things I wanted to have on iOS but didn't. Then I remembered that @ulope and I had a discussion at [FOSDEM 2019](https://zerokspot.com/weblog/2019/02/08/fosdem2019/) about the [fosdem-org app](https://github.com/johanhenselmans/fosdem/) by Johan Henselmans which has served us well for many, many years now. In January I started the app again in order to take a look at all those sessions I'd get into a waiting line for in just a couple of weeks and noticed that it didn't yet support newer iPhones with a notch (basically a fourth third of the screen was left empty).

So I checked out the project and tried to get it to run. I immediately remembered [CocoaPods](https://cocoapods.org/) thanks to its dependency on MobileVLCKit ... and couldn't get the app to build because of this dependency. So I removed it and tried to just fix my little problem. Turns out, I didn't know a thing about frameworks on iOS and Objective-C wasn't helping there either. After a couple of hours I simply quit and set out to write a FOSDEM companion app from scratch in order not only to learn Swift but also how to do iOS development and also play around with things like SwiftUI and CoreData.

The result of this effort is [FosdemToGo](https://github.com/zerok/fosdem-to-go) (I'm still bad at naming things 😅). 

<figure>
<div style="display:flex;justify-content:space-between">
<img src="/media/2020/fosdemtogo-day.png" style="margin:0 1px 0 1px" alt="Listing of tracks for a single day">
<img src="/media/2020/fosdemtogo-myschedule.png" style="margin: 0 1px 0 1px" alt="Listing of bookmarks">
</div>
</figure>

There are a couple of things I wanted to achieve with this app:

- The app should be offline by default. Once a schedule has been downloaded from [fosdem.org](https://fosdem.org/) no online requests should be made unless the user explicitly initiates them. The WIFI at FOSDEM (and any larger event for that matter) is historically flakey. If the app tries to refresh a schedule in such an environment on its own and blocks the user from taking a look at the schedule because of it, I've done something wrong.
- I want to have a quick overview what tracks are available on what day. I usually go to FOSDEM not wanting to stick to a single track for a day (or two where possible) but instead like to take a look on the evening before what tracks/talks I want to see the next day.
- Bookmarking of talks should be possible and they should be listed grouped, again, by day.

That's all the app does right now. There are some little details like that if you tap on the location of talk then [nav.fosdem.org](https://nav.fosdem.org/) will open showing you the room, but that's pretty much it for now! More will definitely come, but it will always be an offline-first app! If you are on iOS then you can find the app in the **[AppStore](https://apps.apple.com/at/app/fosdemtogo/id1500192366?l=en)**. Please let me know what you think!

P.S.: The app does not work on anything below iOS 13 but I might add support for iOS 12 at least some time in the future. It's just not a priority right now.

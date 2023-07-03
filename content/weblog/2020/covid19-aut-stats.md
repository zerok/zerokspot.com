---
title: "Collecting COVID-19 stats for Austria"
date: "2020-03-28T16:04:12+0100"
tags:
- covid19
- datasette
- opendata
---


**Note:** This project has been archived on [2023-07-03](https://zerokspot.com/weblog/2023/07/03/sunsetting-covid19-aut-stats/).

For 10 days now I've been scraping information from the [Austrian "Sozialministerium"][a] (and since yesterday also from [info.gesundheitsministerium.at][f]) to get a historical view on the course COVID-19 is taking in Austria. That page doesn't offer any kind of history to its values and so I thought it might be useful to some to have that data available as CSV or SQLite database. You can find the result on [GitHub][b] where you can also see exactly how I'm fetching the data as well as the resulting [CSV file][c] with values going back to March 13 thanks to [Archive.org's Wayback machine][g].

The CI job in that repository also bundles that CSV with [datasette][e] and publishes it onto [covid19-aut-stats.h10n.me][d]. Thanks to the datasette-vega plugin you can even generate some nice graphs out of the data:

<figure>
<img src="/media/2020/covid19-graph.svg" style="width:400px">
<figcaption>Curve of confirmed cases in Austria (<a href="https://covid19-aut-stats.h10n.me/covid19-aut?sql=select+date%28date%29%2C+max%28confirmed%29+from+%5Bcovid19-aut%5D+group+by+date%28date%29+order+by+date+asc&_hide_sql=1#g.mark=bar&g.x_column=date(date)&g.x_type=temporal&g.y_column=max(confirmed)&g.y_type=quantitative">source</a>)
</figure>

[a]: https://www.sozialministerium.at/Informationen-zum-Coronavirus/Neuartiges-Coronavirus-(2019-nCov).html
[b]: https://github.com/zerok/covid19-aut-stats
[c]: https://github.com/zerok/covid19-aut-stats/blob/master/covid19-aut.csv
[d]: https://covid19-aut-stats.h10n.me/covid19-aut/covid19-aut
[e]: https://datasette.io
[f]: https://info.gesundheitsministerium.at/
[g]: https://archive.org/web/

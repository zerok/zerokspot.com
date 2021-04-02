---
title: Ledger as value tracker
date: "2021-04-02T14:08:10+02:00"
tags:
- 100daystooffload
- finances
---

I just realised that [ledger](https://www.ledger-cli.org/), the tool that I’m using to track my income and expenses, can also be used to keep track of the value of your assets over time! For me this is mostly interesting because of some of the crypto coins I have. Internally, this feature relies on a separate price-db which can be just a simple text file with a structure like this:

	P 2021/04/02 12:06:17 BTC €50594.905
	P 2021/04/02 12:06:17 ETH €1692.985
	P 2021/04/02 12:06:17 XLM €0.383824

Let’s say, I then have such a transaction for when I purchase some ETH:

	2021/02/01 * Bought some ETH
	    ; :cryptoinvest:
	    Assets:Crypto:ETH  0.00611607 ETH
	    Liabilities:CreditCard  -9.01 €

Thanks to the price-db I can now look at that value over time:

	# Current value:
	$ ledger -f index.ledger --price-db prices.db balance \
		Assets:Crypto:ETH -X €
	€ 10.35  Assets:Crypto:ETH
	
	# Value from last week:
	$ ledger -f index.ledger --price-db prices.db balance \
		Assets:Crypto:ETH -X € --end 2021-03-27
	€ 8.63  Assets:Crypto:ETH

To make working with those prices a bit easier, I’ve written myself a little exchange-rate importer from Coinbase which you can find [here](https://gitlab.com/zerok/ledger-tools). Basically, whenever I want to update the prices.db file, I just run the following command:

	$ lt update-prices --overwrite --coinbase-currency ETH

… and it downloads the latest exchange rate for ETH and appends it to the prices.db file.

In the future I might even add some diagrams but since I know how much I’ve spent so far (in EUR) in ETH and other currencies and I can extract the current value of these assets thanks to ledger, I’m already quite pleased with the setup 😅

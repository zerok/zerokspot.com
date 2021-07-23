---
title: csvgen
date: "2021-07-23T20:37:07+02:00"
tags:
- golang
- opensource
- 100daystooffload
- releases
---

I’m working quite a bit with CSV files, from imports I’m writing at work to imports I need for my personal [ledger](https://www.ledger-cli.org/) setup where I convert those exports I get from my bank about my accounts and convert them to ledger entries.

So, I need to create new parsers quite often and dealing with the same setup and parsing logic has now annoyed me enough that I’ve created a little code-generator for Go that does that for me. The idea is, that I just put a simple configuration file (i.e. `csvgen.yml`) into a Go package, add a `go:generate` line and I’d get all the scaffolding. The output of all that is [gitlab.com/zerok/csvgen](https://gitlab.com/zerok/csvgen/) 🥳

So, let’s assume that I have a CSV file that looks like this:

	15.12.2003,Name 1
	18.03.2004,Name 2

The first column is a date formatted as is common in German-speaking countries while the other column is just some plain text. The configuration file for this would look like that:

	---
	processors:
	- type: RecordProcessor
	  recordType: Record
	  comma: ","
	  fields:
	    Date:
	      column: 0
	      decoder:
	        name: parseDate
	    Description:
	      column: 1

And finally, the code I’d have to write myself is just this one:

	package main
	
	import (
		"context"
		"encoding/csv"
		"fmt"
		"os"
		"time"
	)
	
	type Record struct {
		Date        time.Time
		Description string
	}
	
	func parseDate(ctx context.Context, value string) (time.Time, error) {
		return time.Parse("02.01.2006", value)
	}
	
	//go:generate csvgen --output gen_recordprocessor.go
	type RecordProcessor struct {
		reader *csv.Reader
	}
	
	func main() {
		fp, _ := os.Open("test.csv")
		defer fp.Close()
		processor := NewRecordProcessor(fp)
		records, err := processor.ReadAll(context.Background())
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(records)
	}

When I now run `go generate .` within this package, it will generate a constructor for the `RecordProcessor` called `NewRecordProcessor` which accepts an `io.Reader` and also produce reader methods which convert the records from the CSV file into properly typed values for the `Record` struct.

Right now, that’s all that csvgen does. As a next step, I want it to also offer some value parsers that I’ve need again and again. In the future there might also be some CSV-writing functionality alongside the reader which was the main reason why I called the main type “Processor” and not just “Reader” or “Parser”. Let’s see 🙂

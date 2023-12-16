.PHONY: all
all: bin/blog

.PHONY: test
test:
	go test ./... -v

.PHONY: alldata
alldata:
	go run ./ci pipelines pr

.PHONY: clean
clean:
	rm -rf bin data/blogroll.json content/archive

.PHONY: run
run: alldata
	hugo serve -D


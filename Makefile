all: bin/blog

.PHONY: test
test:
	go test ./... -v

bin:
	mkdir -p bin

.PHONY: alldata
alldata:
	./bin/blog build-archive
	hugo
	./bin/blog build-graph
	./bin/blog blogroll --output ./data/blogroll.json
	hugo
	./bin/blog books gen-opml
	./bin/blog build-mapping

clean:
	rm -rf bin data/blogroll.json

bin/blog: $(shell find . -name '*.go') bin go.mod
	cd cmd/blog && go build -o ../../bin/blog


.PHONY: run
run: alldata
	hugo serve -D

.PHONY: clean all

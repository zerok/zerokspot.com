main_js=static/js/archive.js

all: bin/blog $(main_js)

.PHONY: test
test:
	go test ./... -v

prepare:
	yarn

frontend: $(main_js)

static/js/archive.js: $(shell find ./static/app -name '*.js')
	yarn run webpack

bin:
	mkdir -p bin

clean:
	rm -rf bin $(main_js)

bin/blog: $(shell find . -name '*.go') bin
	cd cmd/blog && go build -o ../../bin/blog

.PHONY: clean all prepare frontend

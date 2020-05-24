main_js=static/js/archive.js

all: bin/blogsearch bin/blog $(main_js)

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

bin/blogsearch: bin $(shell find . -name '*.go')
	cd cmd/blogsearch && CGO_ENABLED=0 go build -o ../../bin/blogsearch

update-blogsearch-image:
	docker build -t registry.gitlab.com/zerok/zerokspot.com/blogsearch:latest .
	docker push registry.gitlab.com/zerok/zerokspot.com/blogsearch:latest

.PHONY: clean all prepare update-blogsearch-image frontend

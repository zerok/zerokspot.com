all: bin/blogsearch

bin:
	mkdir -p bin

clean:
	rm -rf bin

bin/blogsearch: $(shell find . -name '*.go')
	cd cmd/blogsearch && go build -o ../../bin/blogsearch

update-blogsearch-image:
	docker build -t registry.gitlab.com/zerok/zerokspot.com/blogsearch:latest .
	docker push registry.gitlab.com/zerok/zerokspot.com/blogsearch:latest

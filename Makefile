all: bin/blogsearch bin/blog

bin:
	mkdir -p bin

clean:
	rm -rf bin

bin/blog: $(shell find . -name '*.go') bin
	cd cmd/blog && go build -o ../../bin/blog

bin/blogsearch: bin $(shell find . -name '*.go')
	cd cmd/blogsearch && go build -o ../../bin/blogsearch

update-blogsearch-image:
	docker build -t registry.gitlab.com/zerok/zerokspot.com/blogsearch:latest .
	docker push registry.gitlab.com/zerok/zerokspot.com/blogsearch:latest

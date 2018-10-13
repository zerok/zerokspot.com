all: bin/blogsearch

bin:
	mkdir -p bin

clean:
	rm -rf bin

bin/blogsearch: $(shell find . -name '*.go')
	cd cmd/blogsearch && go build -o ../../bin/blogsearch

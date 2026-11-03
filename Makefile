all: build

-include config.sh
export

bin:
	mkdir -p bin

.PHONY: build
build: bin embedded
	if [ ! -d bin/repo ]; then \
		mkdir -p bin/repo; \
		git clone "$(REPO)" bin/repo; \
	fi
	cp -v -p -r -f bin/repo/webfonts bin/repo/css embedded/
	go generate ./...
	go fmt ./...

.PHONY: clean
clean:
	rm -rf bin

embedded:
	mkdir -p embedded

.PHONY: test
test:
	go test ./...

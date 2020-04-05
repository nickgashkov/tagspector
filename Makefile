.PHONY: build fmt test docker-run docker-shell docker-build

build:
	./scripts/build.sh

fmt:
	go fmt ./...

test:
	go test ./...

docker-run: docker-build
	docker run --rm go/tagspector

docker-shell: docker-build
	docker run --rm -it -v $(pwd):/app go/tagspector sh

docker-build:
	docker build -t go/tagspector .

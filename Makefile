.PHONY: run build docker-run docker-shell docker-build

build:
	go build -o ./bin/tagspector ./cmd/tagspector

fmt:
	go fmt ./...

docker-run: docker-build
	docker run --rm go/tagspector

docker-shell: docker-build
	docker run --rm -it -v $(pwd):/app go/tagspector sh

docker-build:
	docker build -t go/tagspector .

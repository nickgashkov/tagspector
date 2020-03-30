.PHONY: run build

run: build
	docker run --rm go/tagspector

shell: build
	docker run --rm -it -v $(pwd):/app go/tagspector sh

build:
	docker build -t go/tagspector .

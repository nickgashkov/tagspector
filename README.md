# tagspector [![Build Status](https://github.com/nickgashkov/tagspector/workflows/build/badge.svg)](https://github.com/nickgashkov/tagspector/workflows/build) [![codecov](https://codecov.io/gh/nickgashkov/tagspector/branch/master/graph/badge.svg)](https://codecov.io/gh/nickgashkov/tagspector)

**tagspector** introspects your source code for the codetags.

> :warning: This project makes no sense and overengineered by design. It's whole
> purpose is to flex Golang muscles and try various stuff like `net/http` and
> `github.com/lib/pq`.

## Prerequisites

To run **tagspector** in container, [Docker](https://docs.docker.com/install/)
must be installed. To run **tagspector** on the host machine, Golang 1.14 must
be installed.

## Building

**tagspector** can be built using `make`:

```
$ make build         # using Go 1.14
$ make docker-build  # using Docker
```

## Running

**tagspector** binaries can be found in `./bin` directory and run directly
accordingly to platform. Run `./bin/tagspector -help` to get usage help.

## Features backlog

Here goes the ideas to make this monstrosity even scarier:

- `tagspectord` — simple HTTP server with an ability to introspect source code
itself or path on the host machine.
- `.tagspector.conf` — configure the Tagspector from the configuration file
- `TAGSPECTOR_*` — configure the Tagspector from environment
- `--url` — Tagspector gets the code from the URL
- `cat code.py | tagspector -` — Tagspector gets the code from the `stdin`
- Language-specific codetags introspection — `x = "TODO"` in Python isn't a
codetag, but `x = ""  # TODO` is
- Language guessing
- Add filepath and line numbers for codetags output
- Support multiline codetags
- Add PostgreSQL integration

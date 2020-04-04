# Tagspector

Tagspector introspects your source code for the codetags.

[![codecov](https://codecov.io/gh/nickgashkov/tagspector/branch/master/graph/badge.svg)](https://codecov.io/gh/nickgashkov/tagspector) 
[![Build Status](https://github.com/nickgashkov/tagspector/workflows/build/badge.svg)](https://github.com/nickgashkov/tagspector/workflows/build)


## Features backlog

- `tagspectord` — simple HTTP server with an ability to introspect source code itself or path on the host machine.
- `.tagspector.conf` — configure the Tagspector from the configuration file
- `TAGSPECTOR_*` — configure the Tagspector from environment
- `--url` — Tagspector gets the code from the URL
- `cat code.py | tagspector -` — Tagspector gets the code from the `stdin`
- Language-specific codetags introspection — `x = "TODO"` in Python isn't a code tag, but `x = ""  # TODO` is
- Language guessing
- Add filepath and line numbers for codetags output
- Support multiline codetags

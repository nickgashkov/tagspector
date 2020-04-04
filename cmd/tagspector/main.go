package main

import (
	"flag"
	"github.com/nickgashkov/tagspector/internal/cli"
	"os"
)

func main() {
	codetags := flag.String(
		"codetags", "TODO,FIXME",
		"comma-separated codetags to search for",
	)

	flag.Parse()
	os.Exit(cli.Cli(os.Stdout, os.Stdin, flag.Args(), *codetags))
}

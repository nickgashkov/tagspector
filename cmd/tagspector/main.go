package main

import (
	"flag"
	"os"
)

func main() {
	codetags := flag.String(
		"codetags", "TODO,FIXME",
		"comma-separated codetags to search for",
	)

	flag.Parse()
	os.Exit(cli(os.Stdout, os.Stdin, flag.Args(), *codetags))
}

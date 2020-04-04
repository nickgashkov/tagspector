package main

import (
	"flag"
	"fmt"
	"github.com/nickgashkov/tagspector/internal/codetags"
	"github.com/nickgashkov/tagspector/internal/slices"
	"io"
	"os"
	"strings"
)

func main() {
	var readers []io.Reader

	targetCodetags := flag.String(
		"codetags",
		"TODO,FIXME",
		"comma-separated codetags to search for",
	)

	flag.Parse()

	if flag.NArg() != 1 {
		fmt.Println("expected 1 argument but got", flag.NArg())
		os.Exit(2)
	}

	targetCodetagsSlice := strings.Split(*targetCodetags, ",")

	if slices.StringsEqual(flag.Args(), []string{"-"}) {
		readers = append(readers, os.Stdin)
	}

	parsed := codetags.Parse(readers, targetCodetagsSlice)
	parsedString := strings.Join(parsed, "\n")

	fmt.Println(parsedString)
}

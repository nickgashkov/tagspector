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

	flag.Parse()

	if flag.NArg() != 1 {
		fmt.Println("expected 1 argument but got", flag.NArg())
		os.Exit(2)
	}

	if slices.StringsEqual(flag.Args(), []string{"-"}) {
		readers = append(readers, os.Stdin)
	}

	parsed := codetags.Parse(readers, []string{"TODO"})
	parsedString := strings.Join(parsed, "\n")

	fmt.Println(parsedString)
}

package main

import (
	"fmt"
	"github.com/nickgashkov/tagspector/internal/codetags"
	"github.com/nickgashkov/tagspector/internal/slices"
	"io"
	"strings"
)

func cli(output io.Writer, input io.Reader, args []string, targetCodetags string) int {
	if len(args) != 1 {
		_, _ = fmt.Fprintln(output, "expected 1 argument but got", len(args))
		return 2
	}

	targetCodetagsSlice := strings.Split(targetCodetags, ",")

	readers := getReaders(input, args)
	parsed := codetags.Parse(readers, targetCodetagsSlice)
	parsedString := strings.Join(parsed, "\n")

	_, _ = fmt.Fprintln(output, parsedString)
	return 0
}

func getReaders(input io.Reader, args []string) []io.Reader {
	if slices.StringsEqual(args, []string{"-"}) {
		return []io.Reader{input}
	}

	return []io.Reader{}
}

package cli

import (
	"fmt"
	"github.com/nickgashkov/tagspector/internal/codetags"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func Cli(output io.Writer, input io.Reader, args []string, codetagsFlag string) (resultCode int) {
	pathname, pathnameMessage := getPathname(args)
	codetagsSlice, codetagsSliceMessage := getCodetagsSlice(codetagsFlag)

	if maybeComplain(output, pathnameMessage, codetagsSliceMessage) {
		return 2
	}

	readers := getReaders(input, pathname)
	parsed := codetags.Parse(readers, codetagsSlice)
	parsedString := strings.Join(parsed, "\n")

	_, _ = fmt.Fprintln(output, parsedString)
	return 0
}

func getPathname(args []string) (string, string) {
	if len(args) != 1 {
		return "", fmt.Sprint("expected 1 argument but got ", len(args))
	}

	return args[0], ""
}

func getCodetagsSlice(codetagsFlag string) ([]string, string) {
	var codetagsSlice []string

	for _, codetag := range strings.Split(codetagsFlag, ",") {
		if codetag := strings.TrimSpace(codetag); codetag != "" {
			codetagsSlice = append(codetagsSlice, codetag)
		}
	}

	if len(codetagsSlice) == 0 {
		return []string{}, "expected codetags but got none"
	}

	return codetagsSlice, ""
}

func getReaders(input io.Reader, pathname string) []io.Reader {
	if pathname == "-" {
		return []io.Reader{input}
	}

	var readers []io.Reader
	matches, err := filepath.Glob(pathname)

	if err != nil {
		panic(err)
	}

	for _, match := range matches {
		f, err := os.Open(match)

		if err != nil {
			panic(err)
		}

		readers = append(readers, f)
	}

	return readers
}

func maybeComplain(output io.Writer, messageCandidates ...string) (complained bool) {
	var complainAbout []string

	for _, messageCandidate := range messageCandidates {
		if messageCandidate != "" {
			complainAbout = append(complainAbout, messageCandidate)
		}
	}

	if len(complainAbout) == 0 {
		return false
	}

	for _, message := range complainAbout {
		_, _ = fmt.Fprintln(output, message)
	}

	return true
}

package codetags

import (
	"bufio"
	"io"
	"strings"
)

func Parse(readers []io.Reader, codetags []string) []string {
	var parsed = make([]string, 0)

	for _, reader := range readers {
		parsed = append(parsed, parseIo(reader, codetags)...)
	}

	return parsed
}

func parseIo(reader io.Reader, codetags []string) []string {
	var parsed = make([]string, 0)
	var scanner = bufio.NewScanner(reader)

	for scanner.Scan() {
		line := scanner.Text()

		if parsedLn := parseLn(line, codetags); parsedLn != "" {
			parsed = append(parsed, parsedLn)
		}
	}

	return parsed
}

func parseLn(s string, codetags []string) string {
	for _, codetag := range codetags {
		if strings.Contains(s, codetag) {
			return s
		}
	}

	return ""
}

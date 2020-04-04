package codetags

import (
	"bufio"
	"io"
	"strings"
)

func Parse(reader io.Reader, codetags []string) []string {
	var parsed = make([]string, 0)
	var scanner = bufio.NewScanner(reader)

	for scanner.Scan() {
		line := scanner.Text()

		if parsedLn := parseln(line, codetags); parsedLn != "" {
			parsed = append(parsed, parsedLn)
		}
	}

	return parsed
}

func parseln(s string, codetags []string) string {
	for _, codetag := range codetags {
		if strings.Contains(s, codetag) {
			return s
		}
	}

	return ""
}

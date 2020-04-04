package codetags

import (
	"github.com/nickgashkov/tagspector/testing/assert"
	"io"
	"strings"
	"testing"
)

func TestParseReturnsEmptyCodetagsForEmptyBuffer(t *testing.T) {
	// Given
	reader := strings.NewReader("")
	readers := []io.Reader{reader}
	codetags := []string{"TODO"}

	// When
	parsed := Parse(readers, codetags)

	// Then
	assert.Equal(t, parsed, []string{})
}

func TestParseReturnsMultipleCodetagsForNonEmptyBuffer(t *testing.T) {
	// Given
	reader := strings.NewReader("TODO: Do one thing\nTODO: Do another thing")
	readers := []io.Reader{reader}
	codetags := []string{"TODO"}

	// When
	parsed := Parse(readers, codetags)

	// Then
	assert.Equal(t, parsed, []string{
		"TODO: Do one thing",
		"TODO: Do another thing",
	})
}

func TestParseReturnsEmptyCodetagsForNonEmptyBufferIfTheresNoDesiredCodetags(t *testing.T) {
	// Given
	reader := strings.NewReader("TODO: Do one thing\nTODO: Do another thing")
	readers := []io.Reader{reader}
	codetags := []string{"FIXME"}

	// When
	parsed := Parse(readers, codetags)

	// Then
	assert.Equal(t, parsed, []string{})
}

func TestParseReturnsOnlyDesiredCodetagsForNonEmptyBuffer(t *testing.T) {
	// Given
	reader := strings.NewReader("FIXME: Do one thing\nWTF: Do another thing")
	readers := []io.Reader{reader}
	codetags := []string{"FIXME"}

	// When
	parsed := Parse(readers, codetags)

	// Then
	assert.Equal(t, parsed, []string{"FIXME: Do one thing"})
}

func TestParseReturnsMultipleDesiredCodetagsForNonEmptyBuffer(t *testing.T) {
	// Given
	reader := strings.NewReader("FIXME: Do one thing\nWTF: Do another thing")
	readers := []io.Reader{reader}
	codetags := []string{"FIXME", "WTF"}

	// When
	parsed := Parse(readers, codetags)

	// Then
	assert.Equal(t, parsed, []string{
		"FIXME: Do one thing",
		"WTF: Do another thing",
	})
}

func TestParseReturnsMultipleDesiredCodetagsForMultipleBuffers(t *testing.T) {
	// Given
	readers := []io.Reader{
		strings.NewReader("FIXME: From first\nWTF: From first"),
		strings.NewReader("FIXME: From second\nWTF: From second"),
	}
	codetags := []string{"FIXME"}

	// When
	parsed := Parse(readers, codetags)

	// Then
	assert.Equal(t, parsed, []string{
		"FIXME: From first",
		"FIXME: From second",
	})
}

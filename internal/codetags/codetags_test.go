package codetags

import (
	"github.com/nickgashkov/tagspector/testing/assert"
	"strings"
	"testing"
)

func TestParseReturnsEmptyCodetagsForEmptyBuffer(t *testing.T) {
	// Given
	reader := strings.NewReader("")
	codetags := []string{"TODO"}

	// When
	parsed := Parse(reader, codetags)

	// Then
	assert.Equal(t, parsed, []string{})
}

func TestParseReturnsMultipleCodetagsForNonEmptyBuffer(t *testing.T) {
	// Given
	reader := strings.NewReader("TODO: Do one thing\nTODO: Do another thing")
	codetags := []string{"TODO"}

	// When
	parsed := Parse(reader, codetags)

	// Then
	assert.Equal(t, parsed, []string{
		"TODO: Do one thing",
		"TODO: Do another thing",
	})
}

func TestParseReturnsEmptyCodetagsForNonEmptyBufferIfTheresNoDesiredCodetags(t *testing.T) {
	// Given
	reader := strings.NewReader("TODO: Do one thing\nTODO: Do another thing")
	codetags := []string{"FIXME"}

	// When
	parsed := Parse(reader, codetags)

	// Then
	assert.Equal(t, parsed, []string{})
}

func TestParseReturnsOnlyDesiredCodetagsForNonEmptyBuffer(t *testing.T) {
	// Given
	reader := strings.NewReader("FIXME: Do one thing\nWTF: Do another thing")
	codetags := []string{"FIXME"}

	// When
	parsed := Parse(reader, codetags)

	// Then
	assert.Equal(t, parsed, []string{"FIXME: Do one thing"})
}

func TestParseReturnsMultipleDesiredCodetagsForNonEmptyBuffer(t *testing.T) {
	// Given
	reader := strings.NewReader("FIXME: Do one thing\nWTF: Do another thing")
	codetags := []string{"FIXME", "WTF"}

	// When
	parsed := Parse(reader, codetags)

	// Then
	assert.Equal(t, parsed, []string{
		"FIXME: Do one thing",
		"WTF: Do another thing",
	})
}

package cli

import (
	"bytes"
	"github.com/nickgashkov/tagspector/testing/assert"
	"strings"
	"testing"
)

func TestCliExitsWithResultCode2IfMoreThan1ArgumentSupplied(t *testing.T) {
	// Given
	output := bytes.NewBuffer([]byte{})
	input := strings.NewReader("")
	args := []string{"multiple", "args"}
	codetags := "TODO,FIXME"

	// When
	resultCode := Cli(output, input, args, codetags)

	// Then
	assert.Equal(t, resultCode, 2)
	assert.Equal(t, output.String(), "expected 1 argument but got 2\n")
}

func TestCliExitsWithResultCode2IfLessThan1ArgumentSupplied(t *testing.T) {
	// Given
	output := bytes.NewBuffer([]byte{})
	input := strings.NewReader("")
	args := []string{}
	codetags := "TODO,FIXME"

	// When
	resultCode := Cli(output, input, args, codetags)

	// Then
	assert.Equal(t, resultCode, 2)
	assert.Equal(t, output.String(), "expected 1 argument but got 0\n")
}

func TestCliExitsWithResultCode0If1ArgumentSupplied(t *testing.T) {
	// Given
	output := bytes.NewBuffer([]byte{})
	input := strings.NewReader("")
	args := []string{"single"}
	codetags := "TODO,FIXME"

	// When
	resultCode := Cli(output, input, args, codetags)

	// Then
	assert.Equal(t, resultCode, 0)
}

func TestCliExitsWithResultCode2IfCodetagsIsEmptyString(t *testing.T) {
	// Given
	output := bytes.NewBuffer([]byte{})
	input := strings.NewReader("")
	args := []string{"single"}
	codetags := ""

	// When
	resultCode := Cli(output, input, args, codetags)

	// Then
	assert.Equal(t, resultCode, 2)
	assert.Equal(t, output.String(), "expected codetags but got none\n")
}

func TestCliExitsWithResultCode2IfCodetagsIsMultipleEmptyStrings(t *testing.T) {
	// Given
	output := bytes.NewBuffer([]byte{})
	input := strings.NewReader("")
	args := []string{"single"}
	codetags := ",,,"

	// When
	resultCode := Cli(output, input, args, codetags)

	// Then
	assert.Equal(t, resultCode, 2)
	assert.Equal(t, output.String(), "expected codetags but got none\n")
}

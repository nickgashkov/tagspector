package main

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
	codetags := ""

	// When
	resultCode := cli(output, input, args, codetags)

	// Then
	assert.Equal(t, resultCode, 2)
	assert.Equal(t, output.String(), "expected 1 argument but got 2\n")
}

func TestCliExitsWithResultCode2IfLessThan1ArgumentSupplied(t *testing.T) {
	// Given
	output := bytes.NewBuffer([]byte{})
	input := strings.NewReader("")
	args := []string{}
	codetags := ""

	// When
	resultCode := cli(output, input, args, codetags)

	// Then
	assert.Equal(t, resultCode, 2)
	assert.Equal(t, output.String(), "expected 1 argument but got 0\n")
}

func TestCliExitsWithResultCode0If1ArgumentSupplied(t *testing.T) {
	// Given
	output := bytes.NewBuffer([]byte{})
	input := strings.NewReader("")
	args := []string{"single"}
	codetags := ""

	// When
	resultCode := cli(output, input, args, codetags)

	// Then
	assert.Equal(t, resultCode, 0)
}

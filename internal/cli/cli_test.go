package cli

import (
	"bytes"
	"github.com/nickgashkov/tagspector/testing/assert"
	"github.com/nickgashkov/tagspector/testing/fixtures"
	"io"
	"io/ioutil"
	"os"
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

func TestGetReadersPassesInputBackIfPathnameIsDash(t *testing.T) {
	// Given
	input := strings.NewReader("")
	pathname := "-"

	// When
	readers := getReaders(input, pathname)

	// Then
	assert.Equal(t, readers, []io.Reader{input})
}

func TestGetReadersRetrievesAllGlobbedFilesWhenPathnameIsGlob(t *testing.T) {
	// Given
	input := strings.NewReader("")
	pathname := fixtures.AbsFilepath("TestGetReadersRetrievesAllGlobbedFilesWhenPathnameIsGlob*.dat")

	// When
	readers := getReaders(input, pathname)

	// Then
	var readerTexts []string

	for _, reader := range readers {
		readerText, err := ioutil.ReadAll(reader)
		readerTexts = append(readerTexts, string(readerText))
		assert.Ok(t, err)
	}

	assert.Equal(t, readerTexts, []string{
		"TestGetReadersRetrievesAllGlobbedFilesWhenPathnameIsGlob0 data",
		"TestGetReadersRetrievesAllGlobbedFilesWhenPathnameIsGlob1 data",
	})
}

func TestGetReadersPanicsUponInvalidGlob(t *testing.T) {
	// Given
	input := strings.NewReader("")
	pathname := "\\"

	// Then
	defer assert.Panic(t)

	// When
	getReaders(input, pathname)
}

func TestGetReadersPanicsUponInvalidFiles(t *testing.T) {
	// Given
	input := strings.NewReader("")
	f, err := ioutil.TempFile("", "TestGetReadersSkipsInvalidFiles.dat")
	assert.Ok(t, err)
	pathname := f.Name()
	err = os.Chmod(pathname, 000)
	assert.Ok(t, err)

	// Then
	defer assert.Panic(t)

	// When
	getReaders(input, pathname)
}

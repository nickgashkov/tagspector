package config

import (
	"github.com/nickgashkov/tagspector/testing/assert"
	"testing"
)

func TestNewWithDefaults(t *testing.T) {
	var (
		args    []string
		environ []string
	)

	config := New(args, environ)
	assert.Equal(t, *config, Config{pattern: ""})
}

func TestNewWithArgsInMultipleItems(t *testing.T) {
	var args = []string{"--pattern", "*/**.py"}
	var environ []string

	config := New(args, environ)
	assert.Equal(t, *config, Config{pattern: "*/**.py"})
}

func TestNewWithEnviron(t *testing.T) {
	var args []string
	var environ = []string{"TAGSPECTOR_PATTERN=*/**.py"}

	config := New(args, environ)
	assert.Equal(t, *config, Config{pattern: "*/**.py"})
}

func TestNewWithArgsHasPrecedenceOverEnviron(t *testing.T) {
	var args = []string{"--pattern", "higher_priority.py"}
	var environ = []string{"TAGSPECTOR_PATTERN=lower_priority.py"}

	config := New(args, environ)
	assert.Equal(t, *config, Config{pattern: "higher_priority.py"})
}

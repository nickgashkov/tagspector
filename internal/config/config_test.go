package config

import (
	"github.com/nickgashkov/tagspector/testing/assert"
	"testing"
)

func TestNewWithoutArgsAndEnviron(t *testing.T) {
	var (
		args    []string
		environ []string
	)

	config := New(args, environ)
	assert.Equal(t, *config, Config{pattern: ""})
}

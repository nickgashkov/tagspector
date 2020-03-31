package main

import (
	"github.com/nickgashkov/tagspector/internal/config"
	"os"
)

func main() {
	config.New(os.Args, os.Environ())
}

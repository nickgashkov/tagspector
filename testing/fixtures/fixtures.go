package fixtures

import (
	"path/filepath"
	"runtime"
)

func AbsFilepath(name string) string {
	_, filename, _, _ := runtime.Caller(0)
	return filepath.Join(filename, "..", "..", "testdata", name)
}

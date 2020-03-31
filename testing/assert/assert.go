package assert

import (
	"fmt"
	"github.com/nickgashkov/tagspector/internal/ansi"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

func True(t testing.TB, condition bool, msg string) {
	if !condition {
		_, file, line, _ := runtime.Caller(1)
		printFailDetails()
		printErrorSummary(file, line, msg)
		t.FailNow()
	}
}

func Ok(t testing.TB, err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		printFailDetails()
		printErrorSummary(file, line, "unexpected error")
		printErrorDetails("err.Error()", err.Error())
		t.FailNow()
	}
}

func Equal(t testing.TB, exp, act interface{}) {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		printFailDetails()
		printErrorSummary(file, line, "unexpected value")
		printErrorDetails("exp", exp)
		printErrorDetails("act", act)
		t.FailNow()
	}
}

const spaces = "    "
const newline = "\n"

func printFailDetails() {
	fmt.Println(newline + newline + "--- FAIL DETAILS:")
}

func printErrorSummary(file string, line int, msg string) {
	fmt.Printf(ansi.Error(spaces+"%s:%d: %s"+newline), filepath.Base(file), line, msg)
}

func printErrorDetails(key string, value interface{}) {
	fmt.Printf(ansi.Error(spaces+spaces+"%s: %#v"+newline), key, value)
}

package log

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRedirectStderrToFile(t *testing.T) {
	*stderrToFile = true
	RedirectStderrToFile()

	fmt.Println("fmt to stdout.") // stdout is not redirected.
	line := "fmt to stderr. panic works too."
	fmt.Fprint(os.Stderr, line) // stderr is redirect to file.

	// panic can also be captured to file, but will cause this test "fail",
	// so comment out here.
	// panic("panic to stderr")

	b, err := ioutil.ReadFile(stderrFile)
	if nil != err {
		t.Error(err)
	}
	str := string(b)
	Convey("When reading stderr file content", t, func() {
		So(str, ShouldEqual, line)
	})
}

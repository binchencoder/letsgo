package log

import (
	"bytes"
	"flag"
	"os"

	"github.com/golang/glog"
)

// PrintSysVars outputs all system variables to log file, including all command
// line arguments (flags), all environment variables.
func PrintSysVars() {
	var buf bytes.Buffer

	// All flags.
	buf.WriteString("Command line arguments: ")
	flag.VisitAll(func(f *flag.Flag) {
		buf.WriteString(f.Name)
		buf.WriteString("=")
		buf.WriteString(f.Value.String())
		buf.WriteString(", ")
	})
	glog.Infoln(buf.String())

	// All env vars.
	buf.Reset()
	buf.WriteString("Environment variables: ")
	for _, v := range os.Environ() {
		buf.WriteString(v)
		buf.WriteString(", ")
	}
	glog.Infoln(buf.String())
}

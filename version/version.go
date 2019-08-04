package version

import (
	"fmt"
	"io"
	"runtime"
	"strconv"
	"time"

	"github.com/golang/glog"
)

var (
	// Use ALL_CAP here because bazel outputs key/value pairs in this format.

	BUILD_EMBED_LABEL string
	BUILD_HOST        string
	BUILD_TIMESTAMP   string
	BUILD_USER        string
)

// Print outputs version info to w.
func Print(w io.Writer) {
	io.WriteString(w, "===== Build Info =====\n")
	io.WriteString(w, fmt.Sprintf("BUILD_EMBED_LABEL: %s\n", BUILD_EMBED_LABEL))
	io.WriteString(w, fmt.Sprintf("BUILD_HOST: %s\n", BUILD_HOST))

	io.WriteString(w, fmt.Sprintf("BUILD_TIMESTAMP: %s\n", BUILD_TIMESTAMP))
	if BUILD_TIMESTAMP != "" {
		ms, err := strconv.ParseInt(BUILD_TIMESTAMP, 10, 64)
		if err == nil {
			tm := time.Unix(ms/1000, 0)
			io.WriteString(w, fmt.Sprintf("BUILD_TIMESTAMP (LOCAL): %s, (UTC): %s\n", tm, tm.UTC()))
		} else {
			io.WriteString(w, fmt.Sprintf("Invalid timestamp: %s\n", BUILD_TIMESTAMP))
		}
	}

	io.WriteString(w, fmt.Sprintf("BUILD_USER: %s\n", BUILD_USER))
	io.WriteString(w, fmt.Sprintf("BUILD_GO_VERSION: %s\n", runtime.Version()))
}

// PrintGlog outputs version info to glog.
func PrintGlog() {
	glog.Info("===== Build Info =====")
	glog.Infof("BUILD_EMBED_LABEL: %s", BUILD_EMBED_LABEL)
	glog.Infof("BUILD_HOST: %s", BUILD_HOST)

	glog.Infof("BUILD_TIMESTAMP: %s", BUILD_TIMESTAMP)
	if BUILD_TIMESTAMP != "" {
		ms, err := strconv.ParseInt(BUILD_TIMESTAMP, 10, 64)
		if err == nil {
			tm := time.Unix(ms/1000, 0)
			glog.Infof("BUILD_TIMESTAMP (LOCAL): %s, (UTC): %s", tm, tm.UTC())
		} else {
			glog.Infof("Invalid timestamp: %s", BUILD_TIMESTAMP)
		}
	}

	glog.Infof("BUILD_USER: %s", BUILD_USER)
	glog.Infof("BUILD_GO_VERSION: %s", runtime.Version())
}

package log

import (
	"flag"
	"fmt"
	"os"
	"syscall"
	"time"
)

var (
	stderrToFile = flag.Bool("stderr-to-file", false, "Redirect stderr to file")

	stderrFile string
)

func RedirectStderrToFile() {
	if !*stderrToFile {
		return
	}

	stderrFile = fmt.Sprintf("/tmp/stderr-%s.log", time.Now().Format("20060102-150405"))
	logFile, err := os.OpenFile(stderrFile, os.O_WRONLY|os.O_CREATE, 0644)
	if nil != err {
		fmt.Printf("os.OpenFile: %v\n", err)
		return
	}
	defer logFile.Close()

	if err = syscall.Dup2(int(logFile.Fd()), int(os.Stderr.Fd())); nil != err {
		fmt.Printf("syscall.Dup2: %v\n", err)
		return
	}
	fmt.Printf("Redirected stderr to %s\n", stderrFile)
}

package letsgo

import (
	"flag"
	"os"

	"jingoal.com/letsgo/log"
	"jingoal.com/letsgo/runtime"
	"jingoal.com/letsgo/version"
)

var (
	flagVersion = flag.Bool("version", false, "If true, print the build version info")
)

// InitOption defines the initialization option.
type InitOption func()

// FlagUsage returns an InitOption which sets the flag usage function.
func FlagUsage(usage func()) InitOption {
	return func() {
		flag.Usage = usage
	}
}

// Init initializes Golang application. It must be called in every
// main function as the first line.
func Init(opts ...InitOption) {
	for _, opt := range opts {
		opt()
	}
	flag.Parse()
	if *flagVersion {
		version.Print(os.Stdout)
		os.Exit(0)
	}
	runtime.EnableGoroutineDump()
	version.PrintGlog()

	log.PrintSysVars()
	log.RedirectStderrToFile()
}

// Cleanup performs cleanup at panic.
// Should be defer-called in every main function next to letsgo.Init().
func Cleanup() {
	log.Flush()
}

// Exit performs cleanup before exits program with os.Exit().
// os.Exit: program terminates immediately; deferred functions are not run.
func Exit(errorCode int) {
	Cleanup()
	os.Exit(errorCode)
}

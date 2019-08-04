package log

import (
	"fmt"

	"github.com/golang/glog"
	"golang.org/x/net/context"

	"jingoal.com/letsgo/grpc"
	"jingoal.com/letsgo/trace"
)

const (
	traceFormat          = "(trace: %s) "
	callerFormat         = "(caller: %s) "
	traceAndCallerFormat = "(trace: %s, caller: %s) "
)

// Switch is a boolean type with a method Context.
// See the documentation of V for more information.
type Switch bool

// Logger is a struct contains MDC for log methods. It provides functions Info,
// Warning, Error, Fatal, plus formatting variants such as Infof.
// See the documentation of github.com/golang/glog for more information.
type Logger struct {
	mdc string
}

// Verbose is a variant type of Logger. It provides less log methods.
// If the value is nil, it doesn't record log.
type Verbose Logger

// One may write either
//	if log.V(2) { log.Context(ctx).Info("log this") }
// or
//	glog.V(2).Context(ctx).Info("log this")
// The second form is shorter but the first is cheaper if logging is off because it does
// not evaluate its arguments.
//
// See the documentation of github.com/golang/glog.V for more information.
func V(level glog.Level) Switch {
	return Switch(glog.VDepth(1, level))
}

// Context return a Logger with MDC from ctx.
func Context(ctx context.Context) *Logger {
	// caller is the service name of the caller.
	ctx, caller := grpc.FromMetadataIncoming(ctx)
	tid := trace.GetTraceIdOrEmpty(ctx)

	logger := &Logger{}

	switch {
	case tid != "" && caller != "":
		logger.mdc = fmt.Sprintf(traceAndCallerFormat, tid, caller)
	case tid != "":
		logger.mdc = fmt.Sprintf(traceFormat, tid)
	case caller != "":
		logger.mdc = fmt.Sprintf(callerFormat, caller)
	}

	return logger
}

// Flush flushes log to disk.
func Flush() {
	glog.Flush()
}

// Context return a Verbose with MDC from ctx.
func (s Switch) Context(ctx context.Context) *Verbose {
	if s {
		return (*Verbose)(Context(ctx))
	}

	return nil
}

func (l *Logger) args(args ...interface{}) []interface{} {
	if l.mdc == "" {
		return args
	}

	return append([]interface{}{l.mdc}, args...)
}

// Info logs to the INFO log.
// Arguments are handled in the manner of fmt.Print; a newline is appended if missing.
func (l *Logger) Info(args ...interface{}) {
	glog.InfoDepth(1, l.args(args)...)
}

// InfoDepth acts as Info but uses depth to determine which call frame to log.
// InfoDepth(0, "msg") is the same as Info("msg").
func (l *Logger) InfoDepth(depth int, args ...interface{}) {
	glog.InfoDepth(depth+1, l.args(args)...)
}

// Infoln logs to the INFO log.
// Arguments are handled in the manner of fmt.Println; a newline is appended if missing.
func (l *Logger) Infoln(args ...interface{}) {
	args = append(args, "\n")

	glog.InfoDepth(1, l.args(args)...)
}

// Infof logs to the INFO log.
// Arguments are handled in the manner of fmt.Printf; a newline is appended if missing.
func (l *Logger) Infof(format string, args ...interface{}) {
	glog.InfoDepth(1, l.mdc, fmt.Sprintf(format, args...))
}

// Warning logs to the WARNING and INFO logs.
// Arguments are handled in the manner of fmt.Print; a newline is appended if missing.
func (l *Logger) Warning(args ...interface{}) {
	glog.WarningDepth(1, l.args(args)...)
}

// WarningDepth acts as Warning but uses depth to determine which call frame to log.
// WarningDepth(0, "msg") is the same as Warning("msg").
func (l *Logger) WarningDepth(depth int, args ...interface{}) {
	glog.WarningDepth(depth+1, l.args(args)...)
}

// Warningln logs to the WARNING and INFO logs.
// Arguments are handled in the manner of fmt.Println; a newline is appended if missing.
func (l *Logger) Warningln(args ...interface{}) {
	args = append(args, "\n")

	glog.WarningDepth(1, l.args(args)...)
}

// Warningf logs to the WARNING and INFO logs.
// Arguments are handled in the manner of fmt.Printf; a newline is appended if missing.
func (l *Logger) Warningf(format string, args ...interface{}) {
	glog.WarningDepth(1, l.mdc, fmt.Sprintf(format, args...))
}

// Error logs to the ERROR, WARNING, and INFO logs.
// Arguments are handled in the manner of fmt.Print; a newline is appended if missing.
func (l *Logger) Error(args ...interface{}) {
	glog.ErrorDepth(1, l.args(args)...)
}

// ErrorDepth acts as Error but uses depth to determine which call frame to log.
// ErrorDepth(0, "msg") is the same as Error("msg").
func (l *Logger) ErrorDepth(depth int, args ...interface{}) {
	glog.ErrorDepth(depth+1, l.args(args)...)
}

// Errorln logs to the ERROR, WARNING, and INFO logs.
// Arguments are handled in the manner of fmt.Println; a newline is appended if missing.
func (l *Logger) Errorln(args ...interface{}) {
	args = append(args, "\n")

	glog.ErrorDepth(1, l.args(args)...)
}

// Errorf logs to the ERROR, WARNING, and INFO logs.
// Arguments are handled in the manner of fmt.Printf; a newline is appended if missing.
func (l *Logger) Errorf(format string, args ...interface{}) {
	glog.ErrorDepth(1, l.mdc, fmt.Sprintf(format, args...))
}

// Fatal logs to the FATAL, ERROR, WARNING, and INFO logs,
// including a stack trace of all running goroutines, then calls os.Exit(255).
// Arguments are handled in the manner of fmt.Print; a newline is appended if missing.
func (l *Logger) Fatal(args ...interface{}) {
	glog.FatalDepth(1, l.args(args)...)
}

// FatalDepth acts as Fatal but uses depth to determine which call frame to log.
// FatalDepth(0, "msg") is the same as Fatal("msg").
func (l *Logger) FatalDepth(depth int, args ...interface{}) {
	glog.FatalDepth(depth+1, l.args(args)...)
}

// Fatalln logs to the FATAL, ERROR, WARNING, and INFO logs,
// including a stack trace of all running goroutines, then calls os.Exit(255).
// Arguments are handled in the manner of fmt.Println; a newline is appended if missing.
func (l *Logger) Fatalln(args ...interface{}) {
	args = append(args, "\n")

	glog.FatalDepth(1, l.args(args)...)
}

// Fatalf logs to the FATAL, ERROR, WARNING, and INFO logs,
// including a stack trace of all running goroutines, then calls os.Exit(255).
// Arguments are handled in the manner of fmt.Printf; a newline is appended if missing.
func (l *Logger) Fatalf(format string, args ...interface{}) {
	glog.FatalDepth(1, l.mdc, fmt.Sprintf(format, args...))
}

// Exit logs to the FATAL, ERROR, WARNING, and INFO logs, then calls os.Exit(1).
// Arguments are handled in the manner of fmt.Print; a newline is appended if missing.
func (l *Logger) Exit(args ...interface{}) {
	glog.ExitDepth(1, l.args(args)...)
}

// ExitDepth acts as Exit but uses depth to determine which call frame to log.
// ExitDepth(0, "msg") is the same as Exit("msg").
func (l *Logger) ExitDepth(depth int, args ...interface{}) {
	glog.ExitDepth(depth+1, l.args(args)...)
}

// Exitln logs to the FATAL, ERROR, WARNING, and INFO logs, then calls os.Exit(1).
func (l *Logger) Exitln(args ...interface{}) {
	args = append(args, "\n")

	glog.ExitDepth(1, l.args(args)...)
}

// Exitf logs to the FATAL, ERROR, WARNING, and INFO logs, then calls os.Exit(1).
// Arguments are handled in the manner of fmt.Printf; a newline is appended if missing.
func (l *Logger) Exitf(format string, args ...interface{}) {
	glog.ExitDepth(1, l.mdc, fmt.Sprintf(format, args...))
}

// Info is equivalent to the Logger.Info function, guarded by the value of v.
// See the documentation of V for usage.
func (v *Verbose) Info(args ...interface{}) {
	if v != nil {
		(*Logger)(v).InfoDepth(1, args...)
	}
}

// Infoln is equivalent to the Logger.Infoln function, guarded by the value of v.
// See the documentation of V for usage.
func (v *Verbose) Infoln(args ...interface{}) {
	if v != nil {
		args = append(args, "\n")
		(*Logger)(v).InfoDepth(1, args)
	}
}

// Infof is equivalent to the Logger.Infof function, guarded by the value of v.
// See the documentation of V for usage.
func (v *Verbose) Infof(format string, args ...interface{}) {
	if v != nil {
		(*Logger)(v).InfoDepth(1, fmt.Sprintf(format, args...))
	}
}

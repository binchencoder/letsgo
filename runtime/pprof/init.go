package pprof

import (
	"net/http"
	"sync"

	"github.com/golang/glog"
)

const (
	pprofPrefix = "/_/debug/pprof"
)

var (
	once sync.Once
)

// EnablePprof registers the HTTP handler to serve the pprof
// on the same port as the main service (such as gRPC service).
func EnablePprof(mux *http.ServeMux) {
	once.Do(func() {
		glog.Infof("Serving pprof debug at %s", pprofPrefix)
		mux.Handle(pprofPrefix+"/", http.HandlerFunc(Index))
		mux.Handle(pprofPrefix+"/cmdline", http.HandlerFunc(Cmdline))
		mux.Handle(pprofPrefix+"/profile", http.HandlerFunc(Profile))
		mux.Handle(pprofPrefix+"/symbol", http.HandlerFunc(Symbol))
		mux.Handle(pprofPrefix+"/trace", http.HandlerFunc(Trace))
	})
}

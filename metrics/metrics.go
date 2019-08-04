package metrics

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/golang/glog"
	prom "github.com/prometheus/client_golang/prometheus"
)

const (
	metricsPath = "/_/metrics"
)

var (
	// TODO(zhwang): remove mux once all apps stop using it.
	mux  *http.ServeMux
	once sync.Once
)

func init() {
	mux = http.NewServeMux()
}

// StartPrometheusServer starts prometheus server.
// Deprecated. Prefer to use one port to serve prometheus metrics. Use
// EnablePrometheus() instead, and programs should call ListenAndServe()
// by themselves.
func StartPrometheusServer(scrapePort int) {
	once.Do(func() {
		mux.Handle(metricsPath, prom.Handler())

		hostAddr := fmt.Sprintf(":%d", scrapePort)
		glog.Infof("Starting prometheus server at %s.", hostAddr)
		if err := http.ListenAndServe(hostAddr, mux); err != nil {
			glog.Errorf("Failed to start prometheus server at %s: err: %v", hostAddr, err)
			glog.Flush()
			panic(err)
		}
	})
}

// RegisterHandleFunc registers http handler.
// Deprecated.
func RegisterHandleFunc(path string, handler func(http.ResponseWriter, *http.Request)) {
	mux.HandleFunc(path, handler)
}

// EnablePrometheus registers the HTTP handler to serve the prometheus metrics
// on the same port as the main service (such as gRPC service).
func EnablePrometheus(mux *http.ServeMux) {
	once.Do(func() {
		glog.Infof("Serving prometheus metrics at %s", metricsPath)
		mux.Handle(metricsPath, prom.UninstrumentedHandler())
	})
}

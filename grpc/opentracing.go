package grpc

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"strings"
	"time"

	"github.com/golang/glog"
	opentracing "github.com/opentracing/opentracing-go"
	jaeger "github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"github.com/uber/jaeger-client-go/rpcmetrics"
	xkit "github.com/uber/jaeger-lib/metrics/go-kit"
	"github.com/uber/jaeger-lib/metrics/go-kit/expvar"
)

const (
	sampleRateTotal = 10000
)

var (
	// Disable by default, typically set to "localhost:5775".
	localAgentHostPort = flag.String("jaeger-agent-host-port", "", "Jaeger local agent host port; empty means disable tracing")

	// For debug, use "const" sampler, 1 param.
	samplerType  = flag.String("jaeger-sampler-type", "probabilistic", "Jaeger sampler type")
	samplerParam = flag.Float64("jaeger-sampler-param", 0.001, "Jaeger sampler param")

	bufFlushInterval = flag.Duration("jaeger-buf-flush-interval", 10*time.Second, "Jaeger buffer flush interval")
)

// InitOpenTracing initializes opentracing and returns the tracer.
func InitOpenTracing(serviceName string) opentracing.Tracer {
	glog.V(3).Infof(
		"InitOpenTracing: jaeger-agent-host-port = %v, jaeger-sampler-type = %v, jaeger-sampler-param = %v, jaeger-buf-flush-interval = %v ",
		*localAgentHostPort,
		*samplerType,
		*samplerParam,
		*bufFlushInterval)

	cfg := config.Configuration{
		Sampler: &config.SamplerConfig{
			Type:  *samplerType,
			Param: *samplerParam,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:            false,
			BufferFlushInterval: *bufFlushInterval,
			LocalAgentHostPort:  *localAgentHostPort,
		},
	}

	if *localAgentHostPort == "" {
		cfg.Disabled = true
	}

	metricsFactory := xkit.Wrap("", expvar.NewFactory(10)) // 10 buckets for histograms

	tracer, _, err := cfg.New(
		serviceName,
		config.Logger(jaeger.StdLogger),
		config.Observer(rpcmetrics.NewObserver(metricsFactory, rpcmetrics.DefaultNameNormalizer)),
	)
	if err != nil {
		msg := fmt.Sprintf("InitOpenTracing %s %v", serviceName, err)
		panic(msg)
	}

	return tracer
}

// GetOutboundIP gets the outbound IP of the local machine.
func GetOutboundIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().String()
	idx := strings.LastIndex(localAddr, ":")

	return localAddr[0:idx]
}

// NewContextFromParentSpan creates a new context with a new span, which
// inherits from parent span in the existing context. This is typically for
// tracing an async call.
func NewContextFromParentSpan(ctx context.Context) context.Context {
	if parent := opentracing.SpanFromContext(ctx); parent != nil {
		return opentracing.ContextWithSpan(context.Background(), parent)
	}
	return context.Background()
}

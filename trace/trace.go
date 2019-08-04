package trace

import (
	"encoding/hex"
	"net/http"

	"github.com/golang/glog"
	"github.com/pborman/uuid"
	"golang.org/x/net/context"
)

type traceIdKey struct{}

// Struct traceInfo defines the trace info.
type traceInfo struct {
	Tid string // The trace ID.
}

func GenerateTraceId() string {
	id := uuid.NewRandom()

	var buf [32]byte
	hex.Encode(buf[:], id[:])

	tid := string(buf[:])
	glog.V(2).Infof("Generated traceId:%s", tid)
	return tid
}

// NewTraceId returns a copy of parent context in which
// a new trace ID is created and attached.
func NewTraceId(parent context.Context) context.Context {
	ti := traceInfo{
		Tid: GenerateTraceId(),
	}
	return context.WithValue(parent, traceIdKey{}, &ti)
}

// WithTraceId returns a copy of parent context in which
// the given trace ID is attached.
func WithTraceId(parent context.Context, tid string) context.Context {
	ti := traceInfo{
		Tid: tid,
	}
	return context.WithValue(parent, traceIdKey{}, &ti)
}

// GetTraceId returns the trace ID from the given context.
func GetTraceId(ctx context.Context) (tid string, ok bool) {
	if ti, ok := ctx.Value(traceIdKey{}).(*traceInfo); ok {
		return ti.Tid, ok
	}
	return "", false
}

// GetTraceIdOrEmpty returns the trace ID from the given context or
// an empty string if the trace ID is not found.
func GetTraceIdOrEmpty(ctx context.Context) string {
	if ti, ok := ctx.Value(traceIdKey{}).(*traceInfo); ok {
		return ti.Tid
	}
	return ""
}

// SetTraceIdHeader add trace id to HTTTP Header X-Request-Id.
func SetTraceIdHeader(ctx context.Context, w http.ResponseWriter) {
	if ti, ok := ctx.Value(traceIdKey{}).(*traceInfo); ok {
		w.Header().Set("X-Request-Id", ti.Tid)
	}
}

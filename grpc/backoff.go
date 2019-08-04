package grpc

import (
	"time"

	"github.com/cenkalti/backoff"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type expBackOffKey struct{}

// ExpBackoffUnaryClientInterceptor is a gRPC client-side interceptor that
// provides retry with backoff for unary RPCs.
func ExpBackoffUnaryClientInterceptor(ctx context.Context, method string, req, reply interface{},
	cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {

	if ebo := GetExpBackOff(ctx); ebo != nil {
		op := func() error {
			err := invoker(ctx, method, req, reply, cc, opts...)
			// Either nil or PermanentError will stop the retry.
			return err
		}

		return backoff.Retry(op, ebo)
	}

	return invoker(ctx, method, req, reply, cc, opts...)
}

// WithExpBackOff returns a copy of parent context with the given
// ExponentialBackOff.
func WithExpBackOff(parent context.Context, ebo *backoff.ExponentialBackOff) context.Context {
	return context.WithValue(parent, expBackOffKey{}, ebo)
}

// GetExpBackOff returns ExponentialBackOff from context.
func GetExpBackOff(ctx context.Context) *backoff.ExponentialBackOff {
	ebo, ok := ctx.Value(expBackOffKey{}).(*backoff.ExponentialBackOff)
	if !ok {
		return nil
	}

	return ebo
}

// GetCustomizedExpBackOff returns a customized ExponentialBackOff.
func GetCustomizedExpBackOff(initialInterval, maxInterval, maxElapsedTime time.Duration,
	randFactor, multiplier float64) *backoff.ExponentialBackOff {

	b := &backoff.ExponentialBackOff{
		InitialInterval:     initialInterval,
		RandomizationFactor: randFactor,
		Multiplier:          multiplier,
		MaxInterval:         maxInterval,
		MaxElapsedTime:      maxElapsedTime,
		Clock:               backoff.SystemClock,
	}
	b.Reset()
	return b
}

// GetFastExpBackOff returns an ExponentialBackOff for fast requests.
// For latency of 30ms request, this will do max 3~4 retries, with max latency
// of 250~380ms.
func GetFastExpBackOff() *backoff.ExponentialBackOff {
	b := &backoff.ExponentialBackOff{
		InitialInterval:     50 * time.Millisecond,
		RandomizationFactor: 0.5,
		Multiplier:          2,
		MaxInterval:         150 * time.Millisecond,
		MaxElapsedTime:      250 * time.Millisecond,
		Clock:               backoff.SystemClock,
	}
	b.Reset()
	return b
}

// GetMediumExpBackOff returns an ExponentialBackOff for medium speed requests.
// For latency of 100ms request, this will do max 5 retries, with max latency
// of 1~1.5s.
func GetMediumExpBackOff() *backoff.ExponentialBackOff {
	b := &backoff.ExponentialBackOff{
		InitialInterval:     100 * time.Millisecond,
		RandomizationFactor: 0.5,
		Multiplier:          1.5,
		MaxInterval:         500 * time.Millisecond,
		MaxElapsedTime:      1 * time.Second,
		Clock:               backoff.SystemClock,
	}
	b.Reset()
	return b
}

// GetSlowExpBackOff returns an ExponentialBackOff for slow requests.
// For latency of 300ms request, this will do max 8-9 retries, with max latency
// of 10-14s.
func GetSlowExpBackOff() *backoff.ExponentialBackOff {
	b := &backoff.ExponentialBackOff{
		InitialInterval:     100 * time.Millisecond,
		RandomizationFactor: 0.5,
		Multiplier:          2,
		MaxInterval:         3 * time.Second,
		MaxElapsedTime:      10 * time.Second,
		Clock:               backoff.SystemClock,
	}
	b.Reset()
	return b
}

package grpc

import (
	"fmt"
	"testing"
	"time"

	"github.com/cenkalti/backoff"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type invokerFunc func(ctx context.Context, method string, req, reply interface{},
	cc *grpc.ClientConn, retryCount *int, opDur int, opts ...grpc.CallOption) error

func TestExpBackoffUnaryClientInterceptor(t *testing.T) {
	testExpBackoffUnaryClientInterceptorImpl(t, testInvokerNoErr)
	testExpBackoffUnaryClientInterceptorImpl(t, testInvokerPermanentErr)
}

func testExpBackoffUnaryClientInterceptorImpl(t *testing.T, f invokerFunc) {
	ebo := GetMediumExpBackOff()
	ctx := WithExpBackOff(context.Background(), ebo)
	retryCount := 0
	opDur := 100

	invoker := func(ctx context.Context, method string, req, reply interface{},
		cc *grpc.ClientConn, opts ...grpc.CallOption) error {

		return f(ctx, method, req, reply, cc, &retryCount, opDur, opts...)
	}

	opts := []grpc.CallOption{}
	err := ExpBackoffUnaryClientInterceptor(ctx, "testMethod", nil, nil, nil, invoker, opts...)
	if err != nil {
		t.Errorf("ExpBackoffUnaryClientInterceptor should succeed, %v", err)
	}

	if retryCount < 2 {
		t.Errorf("ExpBackoffUnaryClientInterceptor expect at least 1 retry")
	}
}

func TestExpBackOffWithContext(t *testing.T) {
	ebo := GetExpBackOff(context.Background())
	if ebo != nil {
		t.Errorf("Context should NOT enable retry with exponential backoff")
	}

	ctx := WithExpBackOff(context.Background(), GetFastExpBackOff())
	ebo = GetExpBackOff(ctx)
	if ebo == nil {
		t.Errorf("Context should enable retry with exponential backoff, %+v", ctx)
	}
}

func TestGetSlowExpBackOff(t *testing.T) {
	opDur := 30
	retryCount := 0
	op := func() error {
		a := 1
		b := 2
		c := "abc"
		d := "efg"
		return funcWithRetryErr(opDur, a, b, c, d, &retryCount)
	}
	ebo := GetSlowExpBackOff()
	backoff.Retry(op, ebo)

	if retryCount < 3 {
		t.Errorf("Expected at least %d retries, actually %d", 2, retryCount)
	}
}

func funcWithRetryErr(opDur, a, b int, c, d string, retryCount *int) error {
	*retryCount++
	time.Sleep(time.Duration(opDur) * time.Millisecond)

	return fmt.Errorf("Some retryable error")
}

func testInvokerPermanentErr(ctx context.Context, method string, req, reply interface{},
	cc *grpc.ClientConn, retryCount *int, opDur int, opts ...grpc.CallOption) error {

	*retryCount++
	time.Sleep(time.Duration(opDur) * time.Millisecond)

	// Return PermanentError after 2 retries.
	if *retryCount > 3 {
		return &backoff.PermanentError{}
	}

	return fmt.Errorf("Some retryable error")
}

func testInvokerNoErr(ctx context.Context, method string, req, reply interface{},
	cc *grpc.ClientConn, retryCount *int, opDur int, opts ...grpc.CallOption) error {

	*retryCount++
	time.Sleep(time.Duration(opDur) * time.Millisecond)

	// Succeeds after 2 retries.
	if *retryCount > 3 {
		return nil
	}

	return fmt.Errorf("Some retryable error")
}

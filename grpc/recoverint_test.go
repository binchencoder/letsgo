package grpc

import (
	"testing"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func TestUnaryRecoverServerInterceptor(t *testing.T) {
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		panic("panic from a unary invoke.")
		return nil, nil
	}

	_, err := UnaryRecoverServerInterceptor(context.Background(), nil, nil, handler)
	if err == nil || err.Error() != "panic from a unary invoke." {
		t.Errorf("UnaryRecoverServerInterceptor error %v", err)
	}
}

func TestStreamRecoverServerInterceptor(t *testing.T) {
	handler := func(srv interface{}, stream grpc.ServerStream) error {
		panic("panic from a stream invoke.")
		return nil
	}

	err := StreamRecoverServerInterceptor(nil, nil, nil, handler)
	if err == nil || err.Error() != "panic from a stream invoke." {
		t.Errorf("StreamRecoverServerInterceptor error %v", err)
	}
}

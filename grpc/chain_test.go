package grpc

import (
	"testing"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var (
	someService      = "SomeService.StreamMethod"
	parentContext    = context.WithValue(context.TODO(), "parent", 1)
	parentUnaryInfo  = &grpc.UnaryServerInfo{FullMethod: someService}
	parentStreamInfo = &grpc.StreamServerInfo{
		FullMethod:     someService,
		IsServerStream: true,
	}
)

func TestChainUnaryServer(t *testing.T) {
	oneInt := func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		ctx = context.WithValue(ctx, "first", 1)
		return handler(ctx, req)
	}

	twoInt := func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		if !checkValueInContext(ctx, "first", 1) {
			t.Errorf("TestChainUnary, context value error.")
		}
		ctx = context.WithValue(ctx, "second", 2)

		return handler(ctx, req)
	}

	threeInt := func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		if !checkValueInContext(ctx, "first", 1) {
			t.Errorf("TestChainUnary, context value error.")
		}
		if !checkValueInContext(ctx, "second", 2) {
			t.Errorf("TestChainUnary, context value error.")
		}
		ctx = context.WithValue(ctx, "third", 3)
		return handler(ctx, req)
	}

	chain := ChainUnaryServer(oneInt, twoInt, threeInt, oneInt)

	_, err := chain(parentContext, "input", parentUnaryInfo, func(ctx context.Context, req interface{}) (interface{}, error) {
		if !checkValueInContext(ctx, "first", 1) {
			t.Errorf("TestChainUnary, context value error.")
		}
		if !checkValueInContext(ctx, "second", 2) {
			t.Errorf("TestChainUnary, context value error.")
		}
		if !checkValueInContext(ctx, "third", 3) {
			t.Errorf("TestChainUnary, context value error.")
		}
		return nil, nil
	})

	if err != nil {
		t.Errorf("TestChainUnaryServer %v", err)
	}
}

func TestChainStreamServer(t *testing.T) {
	someService := 1

	oneInt := func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		srv = add(srv, 2)
		return handler(srv, stream)
	}

	twoInt := func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		srv = add(srv, 4)
		return handler(srv, stream)
	}

	threeInt := func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		srv = add(srv, 8)
		return handler(srv, stream)
	}

	chain := ChainStreamServer(oneInt, twoInt, threeInt, oneInt)

	err := chain(someService, nil, parentStreamInfo, func(srv interface{}, stream grpc.ServerStream) error {
		srv = add(srv, 16)
		if srv != 33 {
			t.Errorf("TestChainStreamServer error, result %v", srv)
		}
		return nil
	})

	if err != nil {
		t.Errorf("TestChainStreamServer %v", err)
	}

	chain0 := ChainStreamServer()

	err = chain0(someService, nil, parentStreamInfo, func(srv interface{}, stream grpc.ServerStream) error {
		if srv != 1 {
			t.Errorf("TestChainStreamServer, chain contains zero interceptor error.")
		}
		return nil
	})

	if err != nil {
		t.Errorf("TestChainStreamServer %v", err)
	}

	chain1 := ChainStreamServer(oneInt)

	err = chain1(someService, nil, parentStreamInfo, func(srv interface{}, stream grpc.ServerStream) error {
		if srv != 3 {
			t.Errorf("TestChainStreamServer, chain contains one interceptor error.")
		}
		return nil
	})

	if err != nil {
		t.Errorf("TestChainStreamServer %v", err)
	}
}

func checkValueInContext(ctx context.Context, key interface{}, expectedValue interface{}) bool {
	value := ctx.Value(key)
	if value == nil || value != expectedValue {
		return false
	}

	return true
}

func add(o interface{}, j int) interface{} {
	if i, ok := o.(int); ok {
		return i + j
	}

	return 0
}

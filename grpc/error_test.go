package grpc

import (
	"errors"
	"fmt"
	"testing"

	"google.golang.org/grpc/codes"

	pb "github.com/binchencoder/ease-gateway/proto/frontend"
)

const (
	expected = `rpc error: code = InvalidArgument desc = {"code":500000,"params":["a","b","c"]}`
)

func TestErrorConversion(t *testing.T) {
	e := pb.Error{
		Code:   pb.ErrorCode_UNDEFINED,
		Params: []string{"a", "b", "c"},
	}

	err := ToGrpcError(codes.InvalidArgument, &e)
	if err.Error() != expected {
		t.Errorf("expect frontend error '%s' but got '%s'", expected, err.Error())
	}

	code, ee := FromGrpcError(err)
	if code != codes.InvalidArgument {
		t.Errorf("expect gRPC code '%d' but got '%d'", codes.InvalidArgument, code)
	}
	if ee.Code != pb.ErrorCode_UNDEFINED {
		t.Errorf("expect overall code '%d' but got '%d'", pb.ErrorCode_UNDEFINED, ee.Code)
	}
	if fmt.Sprintf("%s", ee.Params) != fmt.Sprintf("%s", e.Params) {
		t.Errorf("expect params '%s' but got '%s'", e.Params, ee.Params)
	}

	err = errors.New("err")
	code, ee = FromGrpcError(err)
	if code != codes.Unknown && ee != nil {
		t.Error("unexpected result when err is not grpc error and FromGrpcError(err) is called.")
	}
}

// TestGrpcInternalServerError tests the GrpcInternalServerError function.
func TestGrpcInternalServerError(t *testing.T) {
	msg := "Fake Internal Server error!"
	grpcErr := ToGrpcInternalError(errors.New(msg))

	code, pbErr := FromGrpcError(grpcErr)
	if code != codes.Internal {
		t.Errorf("TestGrpcInternalServerError(): expect code as %d. While it's: %d ", codes.Internal, code)
	}
	if len(pbErr.Details) == 0 || pbErr.Details[0].Code != pb.ErrorCode_SERVER_ERROR ||
		len(pbErr.Details[0].Params) == 0 || pbErr.Details[0].Params[0] != msg {
		t.Error("TestGrpcInternalServerError(): invalid error details. Detail: ", pbErr.Details)
	}

	grpcErr2 := ToGrpcInternalError(grpcErr)
	if grpcErr.Error() != grpcErr2.Error() {
		t.Error("TestGrpcInternalServerError(): ToGrpcInternalError should not wrap exiting grpc internal error again.")
	}

	e := &pb.Error{
		Code:   pb.ErrorCode_UNDEFINED,
		Params: []string{"a", "b", "c"},
	}
	grpcErr3 := ToGrpcError(codes.NotFound, e)
	grpcErr4 := ToGrpcInternalError(grpcErr3)
	expteced := ToGrpcError(codes.Internal, e)
	if grpcErr4.Error() != expteced.Error() {
		t.Error("TestGrpcInternalServerError(): ToGrpcInternalError should convert Code from codes.NotFound to codes.Internal.")
	}
}

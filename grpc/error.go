package grpc

import (
	"github.com/golang/protobuf/jsonpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "binchencoder.com/gateway-proto/frontend"
)

var marshaler = jsonpb.Marshaler{
	EnumsAsInts: true,
}

const (
	// 自定义的grpc codes:

	// LimitExceeded indecates 达到限流.
	LimitExceeded = codes.Code(150)
	// StatusFound indicates 302跳转.
	StatusFound = codes.Code(151)
)

// ToGrpcError converts the given error code and frontend error into
// a gRPC error. The error message in the gRPC error is a JSON
// representation of the frontend error. In gRPC gateway, the error
// code will be mapped to proper HTTP status code.
func ToGrpcError(code codes.Code, e *pb.Error) error {
	s, err := marshaler.MarshalToString(e)
	if err != nil {
		return err
	}

	return status.Errorf(code, s)
}

// FromGrpcError extracts the gRPC code and frontend error from
// the given gRPC error. The gRPC error's error message is
// treated as a JSON pb and converted to a frontend error. If the
// conversion fails, the returned code will be codes.Unknown.
func FromGrpcError(err error) (codes.Code, *pb.Error) {
	s, ok := status.FromError(err)
	if !ok {
		return codes.Unknown, nil
	}

	pbe := pb.Error{}
	if e := jsonpb.UnmarshalString(s.Message(), &pbe); e != nil {
		return codes.Unknown, nil
	}
	return s.Code(), &pbe
}

// ToGrpcInternalError creates a grpc internal server error based on normal
// golang error. Any service hit system error, e.g. can not connect to DB, etc.
// It should return an grpc error with SERVER_ERROR code.
// ToGrpcInternalError converts a normal golang error to a grpc error with
// SERVER_ERROR code. grpc gateway will return 500 error code to frontend for
// this grpc error.
// If the error is a grpc error (code != codes.Unknown):
//   If is a grpc internal error then return the error.
//   Else replace the error's Code with codes.Internal.
// Else wrap the error into a grpc internal error.
func ToGrpcInternalError(err error) error {
	code, pbErr := FromGrpcError(err)
	if code != codes.Unknown {
		if code == codes.Internal {
			return err
		}
		return ToGrpcError(codes.Internal, pbErr)
	}

	e := &pb.Error{
		Code: pb.ErrorCode_SERVER_ERROR,
		Details: []*pb.ErrorMessage{
			{
				Code:   pb.ErrorCode_SERVER_ERROR,
				Params: []string{err.Error()},
			},
		},
	}

	return ToGrpcError(codes.Internal, e)
}

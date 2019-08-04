package grpc

import (
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/stretchr/testify/mock"
)

// TracerMock mocks grpc opentracing.Tracer.
type TracerMock struct {
	mock.Mock
}

func (tm *TracerMock) StartSpan(operationName string, opts ...opentracing.StartSpanOption) opentracing.Span {
	args := tm.Called(operationName, opts)
	if res, ok := args.Get(0).(opentracing.Span); ok {
		return res
	}
	return nil
}

func (tm *TracerMock) Inject(sm opentracing.SpanContext, format interface{}, carrier interface{}) error {
	args := tm.Called(sm, format, carrier)
	return args.Error(0)
}

func (tm *TracerMock) Extract(format interface{}, carrier interface{}) (opentracing.SpanContext, error) {
	args := tm.Called(format, carrier)
	if res, ok := args.Get(0).(opentracing.SpanContext); ok {
		return res, args.Error(1)
	}
	return nil, args.Error(1)
}

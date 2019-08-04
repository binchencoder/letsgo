// Package trace defines context functions to work with trace ID.
//
// Trace ID is a string which uniquely identifies an organic request. It should
// be generated at the very beginning of receiving the organic request, put in
// context and passed around to everywhere in the call path. At proper location
// in the call path, the trace ID should be used when output logging. Thus in
// future the same request can be traced through services by matching the trace
// ID.
//
// For gRPC, metadata API should be used to attach the trace ID. Since there
// are other values to be put with metadata API, we'll create an rpc package
// to hold related functions.
//
// Example to work with trace ID:
//
// ctx := trace.NewTraceId(context.Background())
//
// // Pass around the ctx ...
//
// tid := trace.GetTraceIdOrEmpty(ctx) or
// tid, err := trace.GetTraceId(ctx)
//
// // Output tid to log file ...
//
// In case the trace ID could not be passed around through context (for example,
// when it reaches a call with JSON on HTTP), function WithTraceId can be used
// on the other end of the call:
//
// ctx := trace.WithTraceId(json.TraceId)
//
package trace

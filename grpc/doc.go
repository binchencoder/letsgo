// Package grpc defines the context function to work with grpc metadata.
//
// In gRPC, the RPC request is a protocol buffer representing the RPC call's
// business information, and the protocol buffer is not intended to pass
// values other than business information. But in reality the client might
// have information useful for the gRPC server. For example, gRPC server
// might be interested in the client names so that it can monitor different
// clients' QPS or latency.
//
// gRPC provides a metadata API for this very purpose. see
// https://github.com/grpc/grpc-go/blob/master/Documentation/grpc-metadata.md
// for more info about metadata.
//
// This package utilizes the metadata API to pass useful information from
// client to gRPC server:
//     - client name
//     - user identity
//     - custom identity
//     - trace ID
// Since the last three values are already saved in context, this package
// provides functions to convert from/to context of gRPC metadata.
//
// On client side, convert values in context to metadata before calling a gRPC
// method, from the enterprise-circle app:
//
// rpcCtx := ToMetadataOutgoing(ctx, "ecircle")
// resp, err := cli.RpcMethod(rpcCtx, ...)
//
// On server side, convert values in metadata to context:
//
// ctx, clientName := FromMetadataIncoming(rpcCtx)
// requestCount.Label(clientName).Inc()
// traceId, ok := trace.GetTraceIdOrEmpty(ctx)
//
// ToMetadataOutgoing() and FromMetadataIncoming() are embeded in SkyLB as
// interceptors so that client side automatically convert values of interests
// in context to metadata to pass down via grpc call, and server side
// automatically converts metadata as context values.
//
// In client-side interceptors, metadata in context is considered to be outgoing.
// In server-side interceptors, metadata in context is considered to be incoming.
//
// To add a new metadata XXX to pass down via context,
// 1. Implement WithXXX(ctx) to set value in context
// 2. Implement GetXXX(ctx) to get value from context
// 3. Update FromMetadataIncoming() and ToMetadataOutgoing() accordingly
//
// TODO(yuekui): Improve adding a new metadata logic to be single point change.
package grpc

package grpc

import (
	"github.com/golang/glog"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	vexpb "github.com/binchencoder/gateway-proto/data"
	"github.com/binchencoder/letsgo/hashring"
	"github.com/binchencoder/letsgo/ident"
	"github.com/binchencoder/letsgo/trace"
)

const (
	mdKeyAid        = "mdkey_aid"
	mdKeyClientName = "mdkey_client_name"
	mdKeyCid        = "mdkey_cid"
	mdKeyUid        = "mdkey_uid"
	mdKeyCustomid   = "mdkey_customid"
	mdKeyHashKey    = "mdkey_hashkey"
	mdKeyTraceid    = "mdkey_traceid"
)

// FromMetadataIncoming returns a copy of incoming parent context in which the
// known context values (trace ID, user entities, etc.) will be copied from
// a gRPC metadata to context. Also the client's name will be return.
// If the metadata doesn't exist, the incoming parent context will be returned
// with an empty string for client's name.
//
// If the metadata in the incoming parent context is unknown, it will not be
// passed down.
func FromMetadataIncoming(incoming context.Context) (context.Context, string) {
	md, ok := metadata.FromIncomingContext(incoming)
	if !ok {
		return incoming, ""
	}
	glog.V(4).Infof("FromMetadataIncoming: %#v", md)

	var userDetails *vexpb.UserDetails
	ctx := incoming
	if cid, ok := md[mdKeyCid]; ok {
		if uid, ok := md[mdKeyUid]; ok {
			userDetails = &vexpb.UserDetails{
				CompanyId: getLastEle(cid),
				UserId:    getLastEle(uid),
			}
		}
	}
	if aid, ok := md[mdKeyAid]; ok {
		if userDetails == nil {
			userDetails = &vexpb.UserDetails{
				AccountId: getLastEle(aid),
			}
		} else {
			userDetails.AccountId = getLastEle(aid)
		}
	}
	if userDetails != nil {
		ctx = ident.WithUserDetails(ctx, userDetails)
	}

	if custId, ok := md[mdKeyCustomid]; ok {
		ctx = ident.WithCustomIdent(ctx, getLastEle(custId))
	}
	if tid, ok := md[mdKeyTraceid]; ok {
		ctx = trace.WithTraceId(ctx, getLastEle(tid))
	} else {
		ctx = trace.NewTraceId(ctx)
	}
	if hashKey, ok := md[mdKeyHashKey]; ok {
		ctx = hashring.WithHashKey(ctx, getLastEle(hashKey))
	}

	var clientName string
	if cname, ok := md[mdKeyClientName]; ok {
		clientName = getLastEle(cname)
	}

	return ctx, clientName
}

// ToMetadataOutgoing returns a copy of outgoing parent context in which the known
// context values (trace ID, user entities, etc.) will be copied to
// a gRPC metadata, being able to transmitted to server side.
//
// If the outgoing parent context already has metadata, the metadata will be
// joined.
func ToMetadataOutgoing(outgoing context.Context, clientName string) context.Context {
	values := make(map[string]string)
	if clientName != "" {
		values[mdKeyClientName] = clientName
	}

	// User identity.
	if userDetails, ok := ident.GetUserDetails(outgoing); ok {
		if userDetails.CompanyId != "" {
			values[mdKeyCid] = userDetails.CompanyId
		}
		if userDetails.UserId != "" {
			values[mdKeyUid] = userDetails.UserId
		}
		if userDetails.AccountId != "" {
			values[mdKeyAid] = userDetails.AccountId
		}
	}

	// Custom identity.
	if custId, ok := ident.GetCustomIdent(outgoing); ok {
		values[mdKeyCustomid] = custId
	}

	// Trace ID.
	if tid, ok := trace.GetTraceId(outgoing); ok {
		values[mdKeyTraceid] = tid
	} else {
		values[mdKeyTraceid] = trace.GenerateTraceId()
	}

	// Hash Key.
	if hkey, ok := hashring.GetHashKey(outgoing); ok {
		values[mdKeyHashKey] = hkey
	}

	md := metadata.New(values)

	// If outgoing parent has already metadata, pass it down.
	existingMd, ok := metadata.FromOutgoingContext(outgoing)
	if ok {
		md = metadata.Join(existingMd, md)
	}

	glog.V(4).Infof("ToMetadataOutgoing: %#v", md)
	return metadata.NewOutgoingContext(outgoing, md)
}

// ClientToMetadataInterceptor is a gRPC client-side interceptor that put
// context value of interest into metadata for Unary RPCs.
func ClientToMetadataInterceptor(svcName string, ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	ctx = ToMetadataOutgoing(ctx, svcName)
	return invoker(ctx, method, req, reply, cc, opts...)
}

// ServerFromMetadataInterceptor is a gRPC server-side interceptor that put
// metadata as context values for Unary RPCs.
func ServerFromMetadataInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	ctx, _ = FromMetadataIncoming(ctx)
	if glog.V(2) {
		if tid, ok := trace.GetTraceId(ctx); ok {
			glog.Infof("traceId: %s", tid)
		}
	}
	return handler(ctx, req)
}

// ToIncomingCtx creates a new incoming context with the metadata from an
// outgoing context.
func ToIncomingCtx(outgoing context.Context) context.Context {
	if md, ok := metadata.FromOutgoingContext(outgoing); ok {
		return metadata.NewIncomingContext(context.Background(), md)
	}

	return metadata.NewIncomingContext(context.Background(), nil)
}

func getLastEle(slice []string) string {
	count := len(slice)
	if count == 0 {
		return ""
	}
	return slice[count-1]
}

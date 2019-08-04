package grpc

import (
	"testing"

	"golang.org/x/net/context"

	vexpb "github.com/binchencoder/ease-gateway/proto/data"
	"github.com/binchencoder/letsgo/hashring"
	"github.com/binchencoder/letsgo/ident"
	"github.com/binchencoder/letsgo/trace"
)

const (
	companyId  = "1001"
	userId     = "501001"
	customId   = "f0001"
	clientName = "ecircle"
	hashKey    = "abcdefg"
)

func TestMetadata(t *testing.T) {
	ctxs := []context.Context{trace.NewTraceId(context.Background()), context.Background()}

	for _, ctx := range ctxs {
		userDetails := &vexpb.UserDetails{
			CompanyId: companyId,
			UserId:    userId,
		}
		ctx = ident.WithUserDetails(ctx, userDetails)
		ctx = ident.WithCustomIdent(ctx, customId)
		ctx = hashring.WithHashKey(ctx, hashKey)

		rpcCtx := ToMetadataOutgoing(ctx, clientName)
		rpcCtx = ToIncomingCtx(rpcCtx)
		ctx, cname := FromMetadataIncoming(rpcCtx)

		if cname != clientName {
			t.Errorf("Expect gRPC client name %q but got %q.", clientName, cname)
		}

		tid := trace.GetTraceIdOrEmpty(ctx)
		if len(tid) != 32 {
			t.Errorf("expect %d length of trace id, but got %d", 32, len(tid))
		}

		newUserDetails, ok := ident.GetUserDetails(ctx)
		if !ok {
			t.Errorf("Expect getting back user identity but failed.")
		}
		if userDetails.CompanyId != newUserDetails.CompanyId {
			t.Errorf("Expect company ID %s but got %s.", userDetails.CompanyId, newUserDetails.CompanyId)
		}
		if userDetails.UserId != newUserDetails.UserId {
			t.Errorf("Expect user ID %s but got %s.", userDetails.UserId, newUserDetails.UserId)
		}

		custId, ok := ident.GetCustomIdent(ctx)
		if !ok {
			t.Errorf("Expect getting back custom identity but failed.")
		}
		if custId != customId {
			t.Errorf("Expect custom ID %s but got %s.", customId, custId)
		}

		hkey, ok := hashring.GetHashKey(ctx)
		if !ok {
			t.Errorf("Expect getting back hash key but failed")
		}
		if hkey != hashKey {
			t.Errorf("Expect custom ID %s but got %s.", hashKey, hkey)
		}
	}
}

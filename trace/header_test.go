package trace

import (
	"testing"

	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
)

func TestHeaderFromContext(t *testing.T) {
	context := context.Background()
	if header := HeaderFromContext(context); header == nil {
		t.Log("header is nil.")
	} else {
		t.Error("HeaderFromContext return not correct.")
	}

	m := make(map[string]string)
	md := metadata.New(m)
	ctx := metadata.NewOutgoingContext(context, md)
	ctx = toIncomingCtx(ctx)
	if header := HeaderFromContext(ctx); header == nil {
		t.Error("HeaderFromContext return not correct.")
	} else {
		t.Log(header)
	}

	m[XUid] = "uid"
	m[XCid] = "cid"
	m[XAid] = "aid"
	m[XRequestId] = "requestid"
	md = metadata.New(m)
	ctx = metadata.NewOutgoingContext(ctx, md)
	ctx = toIncomingCtx(ctx)
	if header := HeaderFromContext(ctx); header == nil {
		t.Error("HeaderFromContext return not correct.")
	} else {
		t.Log(header)
	}

	m[XSource] = "source"
	m[XClient] = "client"
	m[XAppVersion] = "appVersion"
	m[XDid] = "did"
	md = metadata.New(m)
	ctx = metadata.NewOutgoingContext(ctx, md)
	ctx = toIncomingCtx(ctx)
	if header := HeaderFromContext(ctx); header == nil {
		t.Error("HeaderFromContext return not correct.")
	} else {
		t.Log(header)
	}
}

func TestTokenFromContext(t *testing.T) {
	context := context.Background()
	if token := TokenFromContext(context); token == "" {
		t.Log("token is null.")
	} else {
		t.Error("TokenFromContext return not correct.")
	}

	m := make(map[string]string)
	md := metadata.New(m)
	ctx := metadata.NewOutgoingContext(context, md)
	ctx = toIncomingCtx(ctx)
	if token := TokenFromContext(ctx); token == "" {
		t.Log("token is null.")
	} else {
		t.Error("TokenFromContext return not correct.")
	}

	m[Cookie] = "JINSESSIONID=c19f3976-ce6b-4746-b36a-f8b6774a8b39; aid=cHEIZrMbR2Ku2SfwbT9VOg; cid=2; uid=17198277"
	md = metadata.New(m)
	ctx = metadata.NewOutgoingContext(ctx, md)
	ctx = toIncomingCtx(ctx)
	if token := TokenFromContext(ctx); token != "c19f3976-ce6b-4746-b36a-f8b6774a8b39" {
		t.Errorf("TokenFromContext return not correct.")
	} else {
		t.Log(token)
	}
}

func toIncomingCtx(outgoing context.Context) context.Context {
	if md, ok := metadata.FromOutgoingContext(outgoing); ok {
		return metadata.NewIncomingContext(context.Background(), md)
	}

	return metadata.NewIncomingContext(context.Background(), nil)
}

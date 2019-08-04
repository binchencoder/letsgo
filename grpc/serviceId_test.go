package grpc

import (
	"strings"
	"testing"

	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"

	"github.com/binchencoder/letsgo/ident"
	vexpb "github.com/binchencoder/ease-gateway/proto/data"
)

func TestServiceId(t *testing.T) {
	// Test to add and get service ID to context without existing metadata.
	serviceId := int(123)
	ctx := WithServiceId(context.Background(), serviceId)
	id, err := GetServiceId(ToIncomingCtx(ctx))
	if err != nil {
		t.Errorf("GetServiceId() failed, %v", err)
	}
	if id != serviceId {
		t.Errorf("Get wrong service ID, expected = %d, actual = %d", serviceId, id)
	}

	// Test to have more than one service IDs due to MD join.
	serviceId = int(789)
	ctx = WithServiceId(ctx, serviceId)
	id, err = GetServiceId(ToIncomingCtx(ctx))
	if err != nil {
		t.Errorf("GetServiceId() failed, %v", err)
	}
	if id != serviceId {
		t.Errorf("Get wrong service ID, expected = %d, actual = %d", serviceId, id)
	}

	// Test to add and get service ID to context with existing metadata.

	testMdKey := "testMdKey"
	testMdVal := "testMdVal"
	md := metadata.Pairs(testMdKey, testMdVal)
	ctx = metadata.NewOutgoingContext(context.Background(), md)

	serviceId = int(456)
	ctx = WithServiceId(ctx, serviceId)
	id, err = GetServiceId(ToIncomingCtx(ctx))
	if err != nil {
		t.Errorf("GetServiceId() failed, %v", err)
	}
	if id != serviceId {
		t.Errorf("Got wrong service ID, expected = %d, actual = %d", serviceId, id)
	}

	// Keys in metadata are always lower-case.
	val, err := getLastVal(ToIncomingCtx(ctx), strings.ToLower(testMdKey))
	if err != nil {
		t.Errorf("Failed to get %q from metadata, %v", testMdKey, err)
	}
	if val != testMdVal {
		t.Errorf("Got wrong value for %s, expected = %s, actual = %s", testMdKey, testMdVal, val)
	}

	// Test to get service ID from context without metadata.
	_, err = GetServiceId(ToIncomingCtx(context.Background()))
	if err == nil {
		t.Errorf("Expect error when context does NOT have metadata")
	}

	// Test to get service ID from metadata does not have service ID.
	userDetails := &vexpb.UserDetails{
		CompanyId: "companyId",
		UserId:    "userId",
	}
	ctx = ident.WithUserDetails(ctx, userDetails)
	_, err = GetServiceId(ToIncomingCtx(context.Background()))
	if err == nil {
		t.Errorf("Expect error when context metadata does NOT have service ID")
	}
}

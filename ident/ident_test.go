package ident

import (
	"testing"

	"golang.org/x/net/context"

	vexpb "github.com/binchencoder/gateway-proto/proto/data"
)

const (
	companyId = "1001"
	userId    = "501001"
	customId  = "f0001"
)

func TestUserIdent(t *testing.T) {
	userDetails := &vexpb.UserDetails{
		CompanyId: companyId,
		UserId:    userId,
	}
	ctx := WithUserDetails(context.Background(), userDetails)
	newUserDetails, ok := GetUserDetails(ctx)
	if !ok {
		t.Errorf("Expect getting back userDetails but failed.")
	}
	if newUserDetails.CompanyId != userDetails.CompanyId {
		t.Errorf("Expect company ID %s but got %s.", userDetails.CompanyId, newUserDetails.CompanyId)
	}
	if newUserDetails.UserId != userDetails.UserId {
		t.Errorf("Expect user ID %s but got %s.", userDetails.UserId, newUserDetails.UserId)
	}

	ctx = context.Background()
	if newUserDetails, ok = GetUserDetails(ctx); ok {
		t.Errorf("Expect user ID not found but got %+v.", *newUserDetails)
	}
}

func TestCustomIdent(t *testing.T) {
	ctx := WithCustomIdent(context.Background(), customId)
	custId, ok := GetCustomIdent(ctx)
	if !ok {
		t.Errorf("Expect getting back custom identity but failed.")
	}
	if custId != customId {
		t.Errorf("Expect custom ID %s but got %s.", customId, custId)
	}

	ctx = context.Background()
	if custId, ok = GetCustomIdent(ctx); ok {
		t.Errorf("Expect custom ID not found but got %s.", custId)
	}
}

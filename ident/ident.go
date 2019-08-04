package ident

import (
	"golang.org/x/net/context"

	data "github.com/binchencoder/ease-gateway/proto/data"
)

type userDetailsKey struct{}
type customIdKey struct{}

type customIdent struct {
	customId string // The custom ID.
}

// WithUserDetails returns a copy of parent context in which
// the userDetails is attached.
func WithUserDetails(parent context.Context, userDetails *data.UserDetails) context.Context {
	return context.WithValue(parent, userDetailsKey{}, userDetails)
}

// GetUserDetails returns the userDetails from the given context.
func GetUserDetails(ctx context.Context) (*data.UserDetails, bool) {
	if userDetais, ok := ctx.Value(userDetailsKey{}).(*data.UserDetails); ok {
		return userDetais, true
	}
	return nil, false
}

// WithCustomIdent returns a copy of parent context in which
// the custom identity is attached.
func WithCustomIdent(parent context.Context, customId string) context.Context {
	ident := customIdent{
		customId: customId,
	}
	return context.WithValue(parent, customIdKey{}, &ident)
}

// GetCustomIdent returns the custom identity from the given context.
func GetCustomIdent(ctx context.Context) (customId string, ok bool) {
	if ident, ok := ctx.Value(customIdKey{}).(*customIdent); ok {
		return ident.customId, ok
	}
	return "", false
}

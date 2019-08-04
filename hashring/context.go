package hashring

import (
	"encoding/hex"

	"github.com/pborman/uuid"
	"golang.org/x/net/context"
)

type consistentHashKey struct{}

// Struct hashKeyInfo defines the hash key info.
type hashKeyInfo struct {
	hashKey string // The hash key.
}

func GenerateHashKey() string {
	id := uuid.NewRandom()

	var buf [32]byte
	hex.Encode(buf[:], id[:])

	return string(buf[:])
}

// NewHashKey returns a copy of parent context in which
// a new consistent has key is created and attached.
func NewHashKey(parent context.Context) context.Context {
	hki := hashKeyInfo{
		hashKey: GenerateHashKey(),
	}
	return context.WithValue(parent, consistentHashKey{}, &hki)
}

// WithHashKey returns a copy of parent context in which
// the given hash key is attached.
func WithHashKey(parent context.Context, hashKey string) context.Context {
	hki := hashKeyInfo{
		hashKey: hashKey,
	}
	return context.WithValue(parent, consistentHashKey{}, &hki)
}

// GetHashKey returns the hash key from the given context.
func GetHashKey(ctx context.Context) (string, bool) {
	if hki, ok := ctx.Value(consistentHashKey{}).(*hashKeyInfo); ok {
		return hki.hashKey, ok
	}
	return "", false
}

// GetHashKeyOrEmpty returns the hash key ID from the given context or
// an empty string if the hash key is not found.
func GetHashKeyOrEmpty(ctx context.Context) string {
	if hki, ok := ctx.Value(consistentHashKey{}).(*hashKeyInfo); ok {
		return hki.hashKey
	}
	return ""
}

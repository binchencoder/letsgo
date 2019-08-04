package hashring

import (
	"testing"

	"golang.org/x/net/context"
)

func TestGenerateHashKey(t *testing.T) {
	key1 := GenerateHashKey()
	key2 := GenerateHashKey()

	if len(key1) != 32 {
		t.Errorf("expect %d length of key1, but got %d", 32, len(key1))
	}

	if len(key2) != 32 {
		t.Errorf("expect %d length of key2, but got %d", 32, len(key2))
	}

	if key1 == key2 {
		t.Errorf("expect key1 is different from key2 but they are the same: %s", key1)
	}
}

func TestHashKey(t *testing.T) {
	ctx := NewHashKey(context.Background())

	id := GetHashKeyOrEmpty(ctx)
	if len(id) != 32 {
		t.Errorf("expect %d length of id, but got %d", 32, len(id))
	}
}

func TestEmptyHashKey(t *testing.T) {
	ctx := context.Background()

	_, ok := GetHashKey(ctx)
	if ok {
		t.Errorf("expect no hash key but got it.")
	}

	id := GetHashKeyOrEmpty(ctx)
	if id != "" {
		t.Errorf("expect empty string but got non-empty string.")
	}
}

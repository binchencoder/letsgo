package trace

import (
	"testing"

	"golang.org/x/net/context"
)

func TestGenerateTraceId(t *testing.T) {
	id1 := GenerateTraceId()
	id2 := GenerateTraceId()

	if len(id1) != 32 {
		t.Errorf("expect %d length of id1, but got %d", 32, len(id1))
	}

	if len(id2) != 32 {
		t.Errorf("expect %d length of id2, but got %d", 32, len(id2))
	}

	if id1 == id2 {
		t.Errorf("expect id1 is different from id2 but they are the same: %s", id1)
	}
}

func TestTraceId(t *testing.T) {
	ctx := NewTraceId(context.Background())

	id := GetTraceIdOrEmpty(ctx)
	if len(id) != 32 {
		t.Errorf("expect %d length of id, but got %d", 32, len(id))
	}
}

func TestEmptyTraceId(t *testing.T) {
	ctx := context.Background()

	_, ok := GetTraceId(ctx)
	if ok {
		t.Errorf("expect no trace id but got it.")
	}

	id := GetTraceIdOrEmpty(ctx)
	if id != "" {
		t.Errorf("expect empty string but got non-empty string.")
	}
}

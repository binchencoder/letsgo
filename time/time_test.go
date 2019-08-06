package time

import (
	"strconv"
	"testing"
	"time"
)

func TestMillisecondSince(t *testing.T) {
	start := time.Now()
	time.Sleep(time.Millisecond)
	if MillisecondSince(start) < 1 {
		t.Error("Expect the time since start greater than one millisecond but was not.")
	}
}

func TestMillisecondNow(t *testing.T) {
	value := MillisecondNow()
	t.Log("value:", value)
	s := strconv.FormatInt(value, 10)
	t.Log("s:", s)
	if len(s) != 13 {
		t.Error("Except millisecond.")
	}
}

func TestMillisecond(t *testing.T) {
	value := Millisecond(time.Now())
	t.Log("value:", value)
	s := strconv.FormatInt(value, 10)
	t.Log("s:", s)
	if len(s) != 13 {
		t.Error("Except millisecond.")
	}
}

package time

import (
	"time"
)

// MillisecondSince returns the time duration since the start time
// in milliseconds.
func MillisecondSince(start time.Time) float64 {
	return float64(time.Since(start)) / float64(time.Millisecond)
}

// MillisecondNow returns the current time in milliseconds.
func MillisecondNow() int64 {
	return time.Now().UnixNano() / (int64(time.Millisecond) / int64(time.Nanosecond))
}

// Millisecond returns t in milliseconds.
func Millisecond(t time.Time) int64 {
	return t.UnixNano() / (int64(time.Millisecond) / int64(time.Nanosecond))
}

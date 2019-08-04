package sync

import (
	"sync"
)

// RWLock encapsulates an RWMutex with handy lock operations.
// The typical use of a RWLock is through embedding:
//
// // Define a custom struct.
// type MyStrcut struct {
//     sync.RWLock
//     protected int
// }
//
// // Define an instance.
// mystruct := &MyStrcut{}
//
// // Apply the read lock.
// mystruct.WithRLock(func() error {
//     // Read "protected".
// })
//
// // Apply the write lock.
// mystruct.WithWLock(func() error {
//     // Write "protected".
// })
type RWLock struct {
	lock sync.RWMutex
}

// WithRLock runs the given function with the read lock grabbed.
func (rwm *RWLock) WithRLock(f func() error) error {
	(&rwm.lock).RLock()
	defer (&rwm.lock).RUnlock()
	return f()
}

// WithRLock runs the given function with the write lock grabbed.
func (rwm *RWLock) WithWLock(f func() error) error {
	(&rwm.lock).Lock()
	defer (&rwm.lock).Unlock()
	return f()
}

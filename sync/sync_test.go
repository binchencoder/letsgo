package sync

import (
	"math/rand"
	"sync"
	"testing"
	"time"
)

const (
	concurrency = 10
	jobs        = 100
)

type myLock struct {
	RWLock
	count int
}

func TestRWLock(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	l := myLock{}

	var wg sync.WaitGroup
	wg.Add(3 * concurrency)

	for i := 0; i < concurrency; i++ {
		go func() {
			for j := 0; j < jobs; j++ {
				l.WithWLock(func() error {
					l.count++
					return nil
				})
				time.Sleep(time.Duration(rand.Int63n(100)) * time.Microsecond)
			}
			wg.Done()
		}()
		go func() {
			for j := 0; j < jobs; j++ {
				l.WithRLock(func() error {
					if l.count > concurrency*jobs {
						t.Errorf("Wrong count number, it implies a read lock bug.")
					}
					return nil
				})
				time.Sleep(time.Duration(rand.Int63n(100)) * time.Microsecond)
			}
			wg.Done()
		}()
		go func() {
			for j := 0; j < jobs; j++ {
				l.WithWLock(func() error {
					l.count--
					return nil
				})
				time.Sleep(time.Duration(rand.Int63n(100)) * time.Microsecond)
			}
			wg.Done()
		}()
	}

	wg.Wait()

	if l.count != 0 {
		t.Errorf("Expect final count 0 but got %d, it implies a write lock bug.")
	}
}

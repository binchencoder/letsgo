package grpc

import (
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc/naming"
)

// WatcherMock mocks grpc naming.Watcher.
type WatcherMock struct {
	mock.Mock
}

// Next blocks until an update or error happens.
func (wm *WatcherMock) Next() ([]*naming.Update, error) {
	args := wm.Called()
	if res, ok := args.Get(0).([]*naming.Update); ok {
		return res, nil
	}
	return nil, args.Error(1)
}

// Close closes the Watcher.
func (wm *WatcherMock) Close() {
	wm.Called()
}

// ResolverMock mocks grpc naming.Resolver.
type ResolverMock struct {
	mock.Mock
}

// Resolve creates a Watcher for target.
func (bm *ResolverMock) Resolve(target string) (naming.Watcher, error) {
	args := bm.Called(target)
	if res, ok := args.Get(0).(naming.Watcher); ok {
		return res, nil
	}
	return nil, args.Error(1)
}

package etcd

import (
	etcd "github.com/coreos/etcd/client"
	"github.com/stretchr/testify/mock"
	"golang.org/x/net/context"
)

// KeysAPIMock mocks etcd.KeysAPI.
type KeysAPIMock struct {
	mock.Mock
}

func (kapim *KeysAPIMock) Get(ctx context.Context, key string, opts *etcd.GetOptions) (*etcd.Response, error) {
	args := kapim.Called(ctx, key, opts)
	if res, ok := args.Get(0).(*etcd.Response); ok {
		return res, args.Error(1)
	}
	return nil, args.Error(1)
}

func (kapim *KeysAPIMock) Set(ctx context.Context, key, value string, opts *etcd.SetOptions) (*etcd.Response, error) {
	args := kapim.Called(ctx, key, value, opts)
	if res, ok := args.Get(0).(*etcd.Response); ok {
		return res, args.Error(1)
	}
	return nil, args.Error(1)
}

func (kapim *KeysAPIMock) Delete(ctx context.Context, key string, opts *etcd.DeleteOptions) (*etcd.Response, error) {
	args := kapim.Called(ctx, key, opts)
	if res, ok := args.Get(0).(*etcd.Response); ok {
		return res, args.Error(1)
	}
	return nil, args.Error(1)
}

func (kapim *KeysAPIMock) Create(ctx context.Context, key, value string) (*etcd.Response, error) {
	args := kapim.Called(ctx, key, value)
	if res, ok := args.Get(0).(*etcd.Response); ok {
		return res, args.Error(1)
	}
	return nil, args.Error(1)
}

func (kapim *KeysAPIMock) CreateInOrder(ctx context.Context, dir, value string, opts *etcd.CreateInOrderOptions) (*etcd.Response, error) {
	args := kapim.Called(ctx, dir, value, opts)
	if res, ok := args.Get(0).(*etcd.Response); ok {
		return res, args.Error(1)
	}
	return nil, args.Error(1)
}

func (kapim *KeysAPIMock) Update(ctx context.Context, key, value string) (*etcd.Response, error) {
	args := kapim.Called(ctx, key, value)
	if res, ok := args.Get(0).(*etcd.Response); ok {
		return res, args.Error(1)
	}
	return nil, args.Error(1)
}

func (kapim *KeysAPIMock) Watcher(key string, opts *etcd.WatcherOptions) etcd.Watcher {
	args := kapim.Called(key, opts)
	if res, ok := args.Get(0).(etcd.Watcher); ok {
		return res
	}
	return nil
}

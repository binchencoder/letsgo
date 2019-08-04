package grpc

import (
	"github.com/stretchr/testify/mock"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// BalancerMock mocks grpc.Balancer.
type BalancerMock struct {
	mock.Mock
}

func (bm *BalancerMock) Start(target string, config grpc.BalancerConfig) error {
	args := bm.Called(target, config)
	return args.Error(0)
}

func (bm *BalancerMock) Up(addr grpc.Address) (down func(error)) {
	args := bm.Called(addr)
	if res, ok := args.Get(0).(func(error)); ok {
		return res
	}
	return nil
}

func (bm *BalancerMock) Get(ctx context.Context, opts grpc.BalancerGetOptions) (addr grpc.Address, put func(), err error) {
	args := bm.Called(ctx, opts)
	addr = args.Get(0).(grpc.Address)
	put = args.Get(1).(func())
	err = args.Get(2).(error)
	return
}

func (bm *BalancerMock) Notify() <-chan []grpc.Address {
	args := bm.Called()
	if addrs, ok := args.Get(0).(<-chan []grpc.Address); ok {
		return addrs
	}
	return nil
}

func (bm *BalancerMock) Close() error {
	args := bm.Called()
	if err, ok := args.Get(0).(error); ok {
		return err
	}
	return nil
}

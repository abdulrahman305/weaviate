//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2025 Weaviate B.V. All rights reserved.
//
//  CONTACT: hello@weaviate.io
//

// Code generated by mockery v2.53.2. DO NOT EDIT.

package lsmkv

import (
	context "context"

	logrus "github.com/sirupsen/logrus"
	cyclemanager "github.com/weaviate/weaviate/entities/cyclemanager"

	mock "github.com/stretchr/testify/mock"
)

// MockBucketCreator is an autogenerated mock type for the BucketCreator type
type MockBucketCreator struct {
	mock.Mock
}

type MockBucketCreator_Expecter struct {
	mock *mock.Mock
}

func (_m *MockBucketCreator) EXPECT() *MockBucketCreator_Expecter {
	return &MockBucketCreator_Expecter{mock: &_m.Mock}
}

// NewBucket provides a mock function with given fields: ctx, dir, rootDir, logger, metrics, compactionCallbacks, flushCallbacks, opts
func (_m *MockBucketCreator) NewBucket(ctx context.Context, dir string, rootDir string, logger logrus.FieldLogger, metrics *Metrics, compactionCallbacks cyclemanager.CycleCallbackGroup, flushCallbacks cyclemanager.CycleCallbackGroup, opts ...BucketOption) (*Bucket, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, dir, rootDir, logger, metrics, compactionCallbacks, flushCallbacks)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for NewBucket")
	}

	var r0 *Bucket
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, logrus.FieldLogger, *Metrics, cyclemanager.CycleCallbackGroup, cyclemanager.CycleCallbackGroup, ...BucketOption) (*Bucket, error)); ok {
		return rf(ctx, dir, rootDir, logger, metrics, compactionCallbacks, flushCallbacks, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, logrus.FieldLogger, *Metrics, cyclemanager.CycleCallbackGroup, cyclemanager.CycleCallbackGroup, ...BucketOption) *Bucket); ok {
		r0 = rf(ctx, dir, rootDir, logger, metrics, compactionCallbacks, flushCallbacks, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Bucket)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, logrus.FieldLogger, *Metrics, cyclemanager.CycleCallbackGroup, cyclemanager.CycleCallbackGroup, ...BucketOption) error); ok {
		r1 = rf(ctx, dir, rootDir, logger, metrics, compactionCallbacks, flushCallbacks, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockBucketCreator_NewBucket_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'NewBucket'
type MockBucketCreator_NewBucket_Call struct {
	*mock.Call
}

// NewBucket is a helper method to define mock.On call
//   - ctx context.Context
//   - dir string
//   - rootDir string
//   - logger logrus.FieldLogger
//   - metrics *Metrics
//   - compactionCallbacks cyclemanager.CycleCallbackGroup
//   - flushCallbacks cyclemanager.CycleCallbackGroup
//   - opts ...BucketOption
func (_e *MockBucketCreator_Expecter) NewBucket(ctx interface{}, dir interface{}, rootDir interface{}, logger interface{}, metrics interface{}, compactionCallbacks interface{}, flushCallbacks interface{}, opts ...interface{}) *MockBucketCreator_NewBucket_Call {
	return &MockBucketCreator_NewBucket_Call{Call: _e.mock.On("NewBucket",
		append([]interface{}{ctx, dir, rootDir, logger, metrics, compactionCallbacks, flushCallbacks}, opts...)...)}
}

func (_c *MockBucketCreator_NewBucket_Call) Run(run func(ctx context.Context, dir string, rootDir string, logger logrus.FieldLogger, metrics *Metrics, compactionCallbacks cyclemanager.CycleCallbackGroup, flushCallbacks cyclemanager.CycleCallbackGroup, opts ...BucketOption)) *MockBucketCreator_NewBucket_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]BucketOption, len(args)-7)
		for i, a := range args[7:] {
			if a != nil {
				variadicArgs[i] = a.(BucketOption)
			}
		}
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(logrus.FieldLogger), args[4].(*Metrics), args[5].(cyclemanager.CycleCallbackGroup), args[6].(cyclemanager.CycleCallbackGroup), variadicArgs...)
	})
	return _c
}

func (_c *MockBucketCreator_NewBucket_Call) Return(_a0 *Bucket, _a1 error) *MockBucketCreator_NewBucket_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockBucketCreator_NewBucket_Call) RunAndReturn(run func(context.Context, string, string, logrus.FieldLogger, *Metrics, cyclemanager.CycleCallbackGroup, cyclemanager.CycleCallbackGroup, ...BucketOption) (*Bucket, error)) *MockBucketCreator_NewBucket_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockBucketCreator creates a new instance of MockBucketCreator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockBucketCreator(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockBucketCreator {
	mock := &MockBucketCreator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

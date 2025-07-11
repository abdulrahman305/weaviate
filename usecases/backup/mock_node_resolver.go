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

package backup

import mock "github.com/stretchr/testify/mock"

// MockNodeResolver is an autogenerated mock type for the NodeResolver type
type MockNodeResolver struct {
	mock.Mock
}

type MockNodeResolver_Expecter struct {
	mock *mock.Mock
}

func (_m *MockNodeResolver) EXPECT() *MockNodeResolver_Expecter {
	return &MockNodeResolver_Expecter{mock: &_m.Mock}
}

// AllNames provides a mock function with no fields
func (_m *MockNodeResolver) AllNames() []string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for AllNames")
	}

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// MockNodeResolver_AllNames_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AllNames'
type MockNodeResolver_AllNames_Call struct {
	*mock.Call
}

// AllNames is a helper method to define mock.On call
func (_e *MockNodeResolver_Expecter) AllNames() *MockNodeResolver_AllNames_Call {
	return &MockNodeResolver_AllNames_Call{Call: _e.mock.On("AllNames")}
}

func (_c *MockNodeResolver_AllNames_Call) Run(run func()) *MockNodeResolver_AllNames_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockNodeResolver_AllNames_Call) Return(_a0 []string) *MockNodeResolver_AllNames_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockNodeResolver_AllNames_Call) RunAndReturn(run func() []string) *MockNodeResolver_AllNames_Call {
	_c.Call.Return(run)
	return _c
}

// LeaderID provides a mock function with no fields
func (_m *MockNodeResolver) LeaderID() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for LeaderID")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockNodeResolver_LeaderID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'LeaderID'
type MockNodeResolver_LeaderID_Call struct {
	*mock.Call
}

// LeaderID is a helper method to define mock.On call
func (_e *MockNodeResolver_Expecter) LeaderID() *MockNodeResolver_LeaderID_Call {
	return &MockNodeResolver_LeaderID_Call{Call: _e.mock.On("LeaderID")}
}

func (_c *MockNodeResolver_LeaderID_Call) Run(run func()) *MockNodeResolver_LeaderID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockNodeResolver_LeaderID_Call) Return(_a0 string) *MockNodeResolver_LeaderID_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockNodeResolver_LeaderID_Call) RunAndReturn(run func() string) *MockNodeResolver_LeaderID_Call {
	_c.Call.Return(run)
	return _c
}

// NodeCount provides a mock function with no fields
func (_m *MockNodeResolver) NodeCount() int {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for NodeCount")
	}

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// MockNodeResolver_NodeCount_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'NodeCount'
type MockNodeResolver_NodeCount_Call struct {
	*mock.Call
}

// NodeCount is a helper method to define mock.On call
func (_e *MockNodeResolver_Expecter) NodeCount() *MockNodeResolver_NodeCount_Call {
	return &MockNodeResolver_NodeCount_Call{Call: _e.mock.On("NodeCount")}
}

func (_c *MockNodeResolver_NodeCount_Call) Run(run func()) *MockNodeResolver_NodeCount_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockNodeResolver_NodeCount_Call) Return(_a0 int) *MockNodeResolver_NodeCount_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockNodeResolver_NodeCount_Call) RunAndReturn(run func() int) *MockNodeResolver_NodeCount_Call {
	_c.Call.Return(run)
	return _c
}

// NodeHostname provides a mock function with given fields: nodeName
func (_m *MockNodeResolver) NodeHostname(nodeName string) (string, bool) {
	ret := _m.Called(nodeName)

	if len(ret) == 0 {
		panic("no return value specified for NodeHostname")
	}

	var r0 string
	var r1 bool
	if rf, ok := ret.Get(0).(func(string) (string, bool)); ok {
		return rf(nodeName)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(nodeName)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) bool); ok {
		r1 = rf(nodeName)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// MockNodeResolver_NodeHostname_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'NodeHostname'
type MockNodeResolver_NodeHostname_Call struct {
	*mock.Call
}

// NodeHostname is a helper method to define mock.On call
//   - nodeName string
func (_e *MockNodeResolver_Expecter) NodeHostname(nodeName interface{}) *MockNodeResolver_NodeHostname_Call {
	return &MockNodeResolver_NodeHostname_Call{Call: _e.mock.On("NodeHostname", nodeName)}
}

func (_c *MockNodeResolver_NodeHostname_Call) Run(run func(nodeName string)) *MockNodeResolver_NodeHostname_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockNodeResolver_NodeHostname_Call) Return(_a0 string, _a1 bool) *MockNodeResolver_NodeHostname_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockNodeResolver_NodeHostname_Call) RunAndReturn(run func(string) (string, bool)) *MockNodeResolver_NodeHostname_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockNodeResolver creates a new instance of MockNodeResolver. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockNodeResolver(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockNodeResolver {
	mock := &MockNodeResolver{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

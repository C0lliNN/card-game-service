// Code generated by mockery v2.20.0. DO NOT EDIT.

package game

import mock "github.com/stretchr/testify/mock"

// IDGeneratorMock is an autogenerated mock type for the IDGenerator type
type IDGeneratorMock struct {
	mock.Mock
}

// NewID provides a mock function with given fields:
func (_m *IDGeneratorMock) NewID() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

type mockConstructorTestingTNewIDGeneratorMock interface {
	mock.TestingT
	Cleanup(func())
}

// NewIDGeneratorMock creates a new instance of IDGeneratorMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIDGeneratorMock(t mockConstructorTestingTNewIDGeneratorMock) *IDGeneratorMock {
	mock := &IDGeneratorMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	core "github.com/LambdaTest/test-at-scale/pkg/core"
	mock "github.com/stretchr/testify/mock"
)

// Driver is an autogenerated mock type for the Driver type
type Driver struct {
	mock.Mock
}

// RunDiscovery provides a mock function with given fields: ctx, payload, taskPayload, oauth, coverageDir, secretMap
func (_m *Driver) RunDiscovery(ctx context.Context, payload *core.Payload, taskPayload *core.TaskPayload, oauth *core.Oauth, coverageDir string, secretMap map[string]string) error {
	ret := _m.Called(ctx, payload, taskPayload, oauth, coverageDir, secretMap)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *core.Payload, *core.TaskPayload, *core.Oauth, string, map[string]string) error); ok {
		r0 = rf(ctx, payload, taskPayload, oauth, coverageDir, secretMap)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RunExecution provides a mock function with given fields: ctx, payload, taskPayload, oauth, coverageDir, secretMap
func (_m *Driver) RunExecution(ctx context.Context, payload *core.Payload, taskPayload *core.TaskPayload, oauth *core.Oauth, coverageDir string, secretMap map[string]string) error {
	ret := _m.Called(ctx, payload, taskPayload, oauth, coverageDir, secretMap)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *core.Payload, *core.TaskPayload, *core.Oauth, string, map[string]string) error); ok {
		r0 = rf(ctx, payload, taskPayload, oauth, coverageDir, secretMap)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewDriver interface {
	mock.TestingT
	Cleanup(func())
}

// NewDriver creates a new instance of Driver. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewDriver(t mockConstructorTestingTNewDriver) *Driver {
	mock := &Driver{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
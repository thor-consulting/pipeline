// +build !ignore_autogenerated

// Code generated by mga tool. DO NOT EDIT.

package workflow

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/stretchr/testify/mock"
)

// MockAWSFactory is an autogenerated mock for the AWSFactory type.
type MockAWSFactory struct {
	mock.Mock
}

// New provides a mock function.
func (_m *MockAWSFactory) New(organizationID uint, secretID string, region string) (_result_0 *session.Session, _result_1 error) {
	ret := _m.Called(organizationID, secretID, region)

	var r0 *session.Session
	if rf, ok := ret.Get(0).(func(uint, string, string) *session.Session); ok {
		r0 = rf(organizationID, secretID, region)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*session.Session)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint, string, string) error); ok {
		r1 = rf(organizationID, secretID, region)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

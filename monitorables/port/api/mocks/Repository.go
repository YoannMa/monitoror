// Code generated by mockery v2.7.5. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// OpenSocket provides a mock function with given fields: hostname, port
func (_m *Repository) OpenSocket(hostname string, port int) error {
	ret := _m.Called(hostname, port)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, int) error); ok {
		r0 = rf(hostname, port)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

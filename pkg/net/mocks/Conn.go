// Code generated by mockery v2.7.5. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	net2 "net"

	time "time"
)

// Conn is an autogenerated mock type for the Conn type
type Conn struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *Conn) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LocalAddr provides a mock function with given fields:
func (_m *Conn) LocalAddr() net2.Addr {
	ret := _m.Called()

	var r0 net2.Addr
	if rf, ok := ret.Get(0).(func() net2.Addr); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(net2.Addr)
		}
	}

	return r0
}

// Read provides a mock function with given fields: b
func (_m *Conn) Read(b []byte) (int, error) {
	ret := _m.Called(b)

	var r0 int
	if rf, ok := ret.Get(0).(func([]byte) int); ok {
		r0 = rf(b)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(b)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoteAddr provides a mock function with given fields:
func (_m *Conn) RemoteAddr() net2.Addr {
	ret := _m.Called()

	var r0 net2.Addr
	if rf, ok := ret.Get(0).(func() net2.Addr); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(net2.Addr)
		}
	}

	return r0
}

// SetDeadline provides a mock function with given fields: t
func (_m *Conn) SetDeadline(t time.Time) error {
	ret := _m.Called(t)

	var r0 error
	if rf, ok := ret.Get(0).(func(time.Time) error); ok {
		r0 = rf(t)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetReadDeadline provides a mock function with given fields: t
func (_m *Conn) SetReadDeadline(t time.Time) error {
	ret := _m.Called(t)

	var r0 error
	if rf, ok := ret.Get(0).(func(time.Time) error); ok {
		r0 = rf(t)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetWriteDeadline provides a mock function with given fields: t
func (_m *Conn) SetWriteDeadline(t time.Time) error {
	ret := _m.Called(t)

	var r0 error
	if rf, ok := ret.Get(0).(func(time.Time) error); ok {
		r0 = rf(t)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Write provides a mock function with given fields: b
func (_m *Conn) Write(b []byte) (int, error) {
	ret := _m.Called(b)

	var r0 int
	if rf, ok := ret.Get(0).(func([]byte) int); ok {
		r0 = rf(b)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(b)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

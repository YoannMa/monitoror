// Code generated by mockery v2.7.5. DO NOT EDIT.

package mocks

import (
	gojenkins "github.com/jsdidierlaurent/golang-jenkins"
	mock "github.com/stretchr/testify/mock"
)

// Jenkins is an autogenerated mock type for the Jenkins type
type Jenkins struct {
	mock.Mock
}

// GetBuildByJobId provides a mock function with given fields: jobID, number
func (_m *Jenkins) GetBuildByJobId(jobID string, number int) (gojenkins.Build, error) {
	ret := _m.Called(jobID, number)

	var r0 gojenkins.Build
	if rf, ok := ret.Get(0).(func(string, int) gojenkins.Build); ok {
		r0 = rf(jobID, number)
	} else {
		r0 = ret.Get(0).(gojenkins.Build)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, int) error); ok {
		r1 = rf(jobID, number)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetJob provides a mock function with given fields: jobName
func (_m *Jenkins) GetJob(jobName string) (gojenkins.Job, error) {
	ret := _m.Called(jobName)

	var r0 gojenkins.Job
	if rf, ok := ret.Get(0).(func(string) gojenkins.Job); ok {
		r0 = rf(jobName)
	} else {
		r0 = ret.Get(0).(gojenkins.Job)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(jobName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLastBuildByJobId provides a mock function with given fields: jobID
func (_m *Jenkins) GetLastBuildByJobId(jobID string) (gojenkins.Build, error) {
	ret := _m.Called(jobID)

	var r0 gojenkins.Build
	if rf, ok := ret.Get(0).(func(string) gojenkins.Build); ok {
		r0 = rf(jobID)
	} else {
		r0 = ret.Get(0).(gojenkins.Build)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(jobID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

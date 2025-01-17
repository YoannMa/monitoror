// Code generated by mockery v2.7.5. DO NOT EDIT.

package mocks

import (
	models "github.com/monitoror/monitoror/monitorables/github/api/models"
	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// GetChecks provides a mock function with given fields: owner, repository, ref
func (_m *Repository) GetChecks(owner string, repository string, ref string) (*models.Checks, error) {
	ret := _m.Called(owner, repository, ref)

	var r0 *models.Checks
	if rf, ok := ret.Get(0).(func(string, string, string) *models.Checks); ok {
		r0 = rf(owner, repository, ref)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Checks)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string) error); ok {
		r1 = rf(owner, repository, ref)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCommit provides a mock function with given fields: owner, repository, sha
func (_m *Repository) GetCommit(owner string, repository string, sha string) (*models.Commit, error) {
	ret := _m.Called(owner, repository, sha)

	var r0 *models.Commit
	if rf, ok := ret.Get(0).(func(string, string, string) *models.Commit); ok {
		r0 = rf(owner, repository, sha)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Commit)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string) error); ok {
		r1 = rf(owner, repository, sha)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCount provides a mock function with given fields: query
func (_m *Repository) GetCount(query string) (int, error) {
	ret := _m.Called(query)

	var r0 int
	if rf, ok := ret.Get(0).(func(string) int); ok {
		r0 = rf(query)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPullRequest provides a mock function with given fields: owner, repository, id
func (_m *Repository) GetPullRequest(owner string, repository string, id int) (*models.PullRequest, error) {
	ret := _m.Called(owner, repository, id)

	var r0 *models.PullRequest
	if rf, ok := ret.Get(0).(func(string, string, int) *models.PullRequest); ok {
		r0 = rf(owner, repository, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.PullRequest)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, int) error); ok {
		r1 = rf(owner, repository, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPullRequests provides a mock function with given fields: owner, repository
func (_m *Repository) GetPullRequests(owner string, repository string) ([]models.PullRequest, error) {
	ret := _m.Called(owner, repository)

	var r0 []models.PullRequest
	if rf, ok := ret.Get(0).(func(string, string) []models.PullRequest); ok {
		r0 = rf(owner, repository)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.PullRequest)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(owner, repository)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

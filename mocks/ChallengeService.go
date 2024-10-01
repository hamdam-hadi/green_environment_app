// Code generated by mockery v2.46.0. DO NOT EDIT.

package mocks

import (
	models "green_environment_app/models"

	mock "github.com/stretchr/testify/mock"
)

// ChallengeService is an autogenerated mock type for the ChallengeService type
type ChallengeService struct {
	mock.Mock
}

// CreateChallenge provides a mock function with given fields: challenge
func (_m *ChallengeService) CreateChallenge(challenge *models.Challenge) error {
	ret := _m.Called(challenge)

	if len(ret) == 0 {
		panic("no return value specified for CreateChallenge")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Challenge) error); ok {
		r0 = rf(challenge)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteChallenge provides a mock function with given fields: id
func (_m *ChallengeService) DeleteChallenge(id uint) error {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteChallenge")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllChallenges provides a mock function with given fields:
func (_m *ChallengeService) GetAllChallenges() ([]models.Challenge, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetAllChallenges")
	}

	var r0 []models.Challenge
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]models.Challenge, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []models.Challenge); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Challenge)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetChallengeByID provides a mock function with given fields: id
func (_m *ChallengeService) GetChallengeByID(id uint) (*models.Challenge, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for GetChallengeByID")
	}

	var r0 *models.Challenge
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (*models.Challenge, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(uint) *models.Challenge); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Challenge)
		}
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateChallenge provides a mock function with given fields: challenge
func (_m *ChallengeService) UpdateChallenge(challenge *models.Challenge) error {
	ret := _m.Called(challenge)

	if len(ret) == 0 {
		panic("no return value specified for UpdateChallenge")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Challenge) error); ok {
		r0 = rf(challenge)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewChallengeService creates a new instance of ChallengeService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewChallengeService(t interface {
	mock.TestingT
	Cleanup(func())
}) *ChallengeService {
	mock := &ChallengeService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
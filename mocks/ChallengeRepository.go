package mocks

import (
	models "green_environment_app/models"

	mock "github.com/stretchr/testify/mock"
)

type ChallengeRepository struct {
	mock.Mock
}

func (_m *ChallengeRepository) Delete(id uint) error {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (_m *ChallengeRepository) FindAll() ([]models.Challenge, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for FindAll")
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

func (_m *ChallengeRepository) FindByID(id uint) (*models.Challenge, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for FindByID")
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

func (_m *ChallengeRepository) Save(challenge *models.Challenge) error {
	ret := _m.Called(challenge)

	if len(ret) == 0 {
		panic("no return value specified for Save")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Challenge) error); ok {
		r0 = rf(challenge)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (_m *ChallengeRepository) Update(challenge *models.Challenge) error {
	ret := _m.Called(challenge)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Challenge) error); ok {
		r0 = rf(challenge)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func NewChallengeRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *ChallengeRepository {
	mock := &ChallengeRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
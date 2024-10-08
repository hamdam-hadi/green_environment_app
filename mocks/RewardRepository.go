package mocks

import (
	models "green_environment_app/models"

	mock "github.com/stretchr/testify/mock"
)

type RewardRepository struct {
	mock.Mock
}

func (_m *RewardRepository) Delete(id uint) error {
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

func (_m *RewardRepository) FindAll() ([]models.Reward, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for FindAll")
	}

	var r0 []models.Reward
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]models.Reward, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []models.Reward); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Reward)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *RewardRepository) FindByID(id uint) (*models.Reward, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for FindByID")
	}

	var r0 *models.Reward
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (*models.Reward, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(uint) *models.Reward); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Reward)
		}
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *RewardRepository) Save(reward *models.Reward) error {
	ret := _m.Called(reward)

	if len(ret) == 0 {
		panic("no return value specified for Save")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Reward) error); ok {
		r0 = rf(reward)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (_m *RewardRepository) Update(reward *models.Reward) error {
	ret := _m.Called(reward)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Reward) error); ok {
		r0 = rf(reward)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func NewRewardRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *RewardRepository {
	mock := &RewardRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

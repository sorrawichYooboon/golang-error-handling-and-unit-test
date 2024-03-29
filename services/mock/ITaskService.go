// Code generated by mockery v2.20.0. DO NOT EDIT.

package mockService

import (
	models "github.com/sorrawichYooboon/golang-error-handling-and-unit-test/models"
	mock "github.com/stretchr/testify/mock"
)

// ITaskService is an autogenerated mock type for the ITaskService type
type ITaskService struct {
	mock.Mock
}

// CreateTask provides a mock function with given fields: task
func (_m *ITaskService) CreateTask(task models.Task) (models.Task, error) {
	ret := _m.Called(task)

	var r0 models.Task
	var r1 error
	if rf, ok := ret.Get(0).(func(models.Task) (models.Task, error)); ok {
		return rf(task)
	}
	if rf, ok := ret.Get(0).(func(models.Task) models.Task); ok {
		r0 = rf(task)
	} else {
		r0 = ret.Get(0).(models.Task)
	}

	if rf, ok := ret.Get(1).(func(models.Task) error); ok {
		r1 = rf(task)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteTaskById provides a mock function with given fields: id
func (_m *ITaskService) DeleteTaskById(id string) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetTaskById provides a mock function with given fields: id
func (_m *ITaskService) GetTaskById(id string) (models.Task, error) {
	ret := _m.Called(id)

	var r0 models.Task
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (models.Task, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) models.Task); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(models.Task)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTasks provides a mock function with given fields:
func (_m *ITaskService) GetTasks() ([]models.Task, error) {
	ret := _m.Called()

	var r0 []models.Task
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]models.Task, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []models.Task); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Task)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateTaskById provides a mock function with given fields: id, task
func (_m *ITaskService) UpdateTaskById(id string, task models.Task) (models.Task, error) {
	ret := _m.Called(id, task)

	var r0 models.Task
	var r1 error
	if rf, ok := ret.Get(0).(func(string, models.Task) (models.Task, error)); ok {
		return rf(id, task)
	}
	if rf, ok := ret.Get(0).(func(string, models.Task) models.Task); ok {
		r0 = rf(id, task)
	} else {
		r0 = ret.Get(0).(models.Task)
	}

	if rf, ok := ret.Get(1).(func(string, models.Task) error); ok {
		r1 = rf(id, task)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewITaskService interface {
	mock.TestingT
	Cleanup(func())
}

// NewITaskService creates a new instance of ITaskService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewITaskService(t mockConstructorTestingTNewITaskService) *ITaskService {
	mock := &ITaskService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

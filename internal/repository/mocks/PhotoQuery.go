package mocks

import (
	"github.com/salamalfis/projectfinal/internal/model"
	"github.com/stretchr/testify/mock"
)

type PhotoQuery struct {
	mock.Mock
}

func (m *PhotoQuery) FindAll() ([]*model.Photo, error) {
	ret := m.Called()

	var r0 []*model.Photo
	if rf, ok := ret.Get(0).(func() []*model.Photo); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Photo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (m *PhotoQuery) FindByID(id int) (*model.Photo, error) {
	ret := m.Called(id)

	var r0 *model.Photo
	if rf, ok := ret.Get(0).(func(int) *model.Photo); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Photo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

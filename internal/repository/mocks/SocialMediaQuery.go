package mocks

import (
	"context"

	model "github.com/salamalfis/projectfinal/internal/model"
	mock "github.com/stretchr/testify/mock"
)

type SocialMediaQuery struct {
	mock.Mock
}

func (_m *SocialMediaQuery) CreateSocialMedia(ctx context.Context, socialMedia model.SocialMedia) (model.SocialMedia, error) {
	ret := _m.Called(ctx, socialMedia)

	if len(ret) == 0 {
		panic("no return value specified for CreateSocialMedia")
	}

	var r0 model.SocialMedia
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, model.SocialMedia) (model.SocialMedia, error)); ok {
		return rf(ctx, socialMedia)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.SocialMedia) model.SocialMedia); ok {
		r0 = rf(ctx, socialMedia)
	} else {
		r0 = ret.Get(0).(model.SocialMedia)
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.SocialMedia) error); ok {
		r1 = rf(ctx, socialMedia)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *SocialMediaQuery) DeleteSocialMediaByID(ctx context.Context, id uint64) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteSocialMediaByID")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

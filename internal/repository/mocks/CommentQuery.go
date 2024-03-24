package mocks

import (
	"context"

	"github.com/salamalfis/projectfinal/internal/model"
	mock "github.com/stretchr/testify/mock"
	//"github.com/stretchr/testify/mock"
)

// CommentQuery is an autogenerated mock type for the CommentQuery type
type CommentQuery struct {
	mock.Mock
}

// CreateComment provides a mock function with given fields: ctx, comment
func (_m *CommentQuery) CreateComment(ctx context.Context, comment model.Comment) (model.Comment, error) {
	ret := _m.Called(ctx, comment)

	var r0 model.Comment
	if rf, ok := ret.Get(0).(func(context.Context, model.Comment) model.Comment); ok {
		r0 = rf(ctx, comment)
	} else {
		r0 = ret.Get(0).(model.Comment)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.Comment) error); ok {
		r1 = rf(ctx, comment)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

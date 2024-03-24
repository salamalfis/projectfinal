package service

import (
	"context"

	"github.com/salamalfis/projectfinal/internal/model"
	"github.com/salamalfis/projectfinal/internal/repository"
)

// CommentsService ...
type CommentsService interface {
	// comments
	GetComments(ctx context.Context) ([]model.Comment, error)
	DeleteCommentsById(ctx context.Context, id uint64) (model.Comment, error)
	UpdateCommentsById(ctx context.Context, id uint64) (model.Comment, error)
	GetCommentsById(ctx context.Context, id uint64) (model.Comment, error)

	// activity
	CreateComments(ctx context.Context) (model.Comment, error)
}

type CommentsServiceImpl struct {
	repo repository.CommentQuery
}

func NewCommentService(repo repository.CommentQuery) CommentsService {
	return &CommentsServiceImpl{repo: repo}
}

func (c *CommentsServiceImpl) GetComments(ctx context.Context) ([]model.Comment, error) {
	comment, err := c.repo.GetComments(ctx)
	if err != nil {
		return nil, err
	}
	return comment, err
}

func (c *CommentsServiceImpl) CreateComments(ctx context.Context) (model.Comment, error) {
	comment, err := c.repo.CreateComments(ctx)
	if err != nil {
		return model.Comment{}, err
	}
	return comment, nil
}

func (c *CommentsServiceImpl) DeleteCommentsById(ctx context.Context, id uint64) (model.Comment, error) {
	comment, err := c.repo.GetCommentsByID(ctx, id)
	if err != nil {
		return model.Comment{}, err
	}
	return comment, nil
}

func (c *CommentsServiceImpl) UpdateCommentsById(ctx context.Context, id uint64) (model.Comment, error) {
	comment, err := c.repo.GetCommentsByID(ctx, id)
	if err != nil {
		return model.Comment{}, err
	}
	return comment, nil
}

func (c *CommentsServiceImpl) GetCommentsById(ctx context.Context, id uint64) (model.Comment, error) {
	comment, err := c.repo.GetCommentsByID(ctx, id)
	if err != nil {
		return model.Comment{}, err
	}
	return comment, nil
}


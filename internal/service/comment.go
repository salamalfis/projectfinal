package service

import (
	"context"

	"github.com/salamalfis/projectfinal/internal/model"
	"github.com/salamalfis/projectfinal/internal/repository"
)

// CommentsService ...
type CommentsService interface {
	// comments
	GetComments(ctx context.Context, photoId uint64) ([]model.CommentView, error)
	DeleteCommentsById(ctx context.Context, id uint64) error
	UpdateCommentsById(ctx context.Context, comment model.Comment) (*model.UpdateCommentRes, error)
	GetCommentsById(ctx context.Context, commentId uint64) (*model.Comment, error)

	// activity
	CreateComments(ctx context.Context, userid uint64, comment model.CreateComment) (*model.CreateCommentRespon, error)
}

type CommentsServiceImpl struct {
	repo repository.CommentQuery
}

func NewCommentService(repo repository.CommentQuery) CommentsService {
	return &CommentsServiceImpl{repo: repo}
}

func (c *CommentsServiceImpl) CreateComments(ctx context.Context, User_id uint64, newComment model.CreateComment) (*model.CreateCommentRespon, error) {
	comment := model.Comment{}
	comment.User_id = User_id
	comment.Message = newComment.Message
	comment.Photo_id = newComment.Photo_id

	err := c.repo.CreateComment(ctx, &comment)
	if err != nil {
		return nil, err
	}
	commentRes := model.CreateCommentRespon{}
	commentRes.ID = comment.ID
	commentRes.Message = comment.Message
	commentRes.User_id = comment.User_id
	commentRes.Photo_id = comment.Photo_id
	commentRes.CreatedAt = comment.CreatedAt

	return &commentRes, nil
}

func (c *CommentsServiceImpl) GetComments(ctx context.Context, photoId uint64) ([]model.CommentView, error) {
	comments, err := c.repo.GetComments(ctx, photoId)
	if err != nil {
		return nil, err
	}

	return comments, nil
}

func (c *CommentsServiceImpl) DeleteCommentsById(ctx context.Context, id uint64) error {
	err := c.repo.DeleteCommentsByID(ctx, id)

	return err
}

func (c *CommentsServiceImpl) UpdateCommentsById(ctx context.Context, comment model.Comment) (*model.UpdateCommentRes, error) {
	err := c.repo.UpdateCommentsByID(ctx, &comment)

	if err != nil {
		return nil, err
	}

	commentRes := model.UpdateCommentRes{}
	commentRes.ID = comment.ID
	commentRes.Message = comment.Message
	commentRes.PhotoId = comment.Photo_id
	commentRes.UserId = comment.User_id
	commentRes.UpdatedAt = comment.UpdatedAt

	return &commentRes, nil
}
func (c *CommentsServiceImpl) GetCommentsById(ctx context.Context, id uint64) (*model.Comment, error) {
	comment, err := c.repo.GetCommentsByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return comment, nil
}

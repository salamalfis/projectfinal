package repository

import (
	"context"

	"github.com/salamalfis/projectfinal/internal/infrastructure"
	"github.com/salamalfis/projectfinal/internal/model"
)

type CommentQuery interface {
	CreateComments(ctx context.Context) (model.Comment, error)
	GetComments(ctx context.Context) ([]model.Comment, error)
	GetCommentsByID(ctx context.Context, id uint64) (model.Comment, error)
	UpdateCommentsByID(ctx context.Context, id uint64) (model.Comment, error)
	DeleteCommentsByID(ctx context.Context, id uint64) error
}

type CommentComand interface {
	CreateComment(ctx context.Context) (model.Comment, error)
}

type commentQueryImpl struct {
	db infrastructure.GormPostgres
}

func NewCommentQuery(db infrastructure.GormPostgres) CommentQuery {
	return &commentQueryImpl{db: db}
}

func (u *commentQueryImpl) CreateComments(ctx context.Context) (model.Comment, error) {
	db := u.db.GetConnection()
	comment := model.Comment{}
	if err := db.
		WithContext(ctx).
		Table("comments").
		Create(&comment).Error; err != nil {
		return model.Comment{}, err
	}
	return comment, nil
}

func (u *commentQueryImpl) GetComments(ctx context.Context) ([]model.Comment, error) {
	db := u.db.GetConnection()
	comments := []model.Comment{}
	if err := db.
		WithContext(ctx).
		Table("comments").
		Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}

func (u *commentQueryImpl) GetCommentsByID(ctx context.Context, id uint64) (model.Comment, error) {

	db := u.db.GetConnection()
	comment := model.Comment{}
	if err := db.
		WithContext(ctx).
		Table("comments").
		Where("id = ?", id).
		Find(&comment).Error; err != nil {
		return model.Comment{}, err
	}
	return comment, nil
}

func (u *commentQueryImpl) UpdateCommentsByID(ctx context.Context, id uint64) (model.Comment, error) {

	db := u.db.GetConnection()
	comment := model.Comment{}
	if err := db.
		WithContext(ctx).
		Table("comments").
		Where("id = ?", id).
		Updates(&comment).Error; err != nil {
		return model.Comment{}, err
	}
	return comment, nil
}

func (u *commentQueryImpl) DeleteCommentsByID(ctx context.Context, id uint64) error {
	db := u.db.GetConnection()
	if err := db.
		WithContext(ctx).
		Table("comments").
		Delete(&model.Comment{ID: uint(id)}).
		Error; err != nil {
		return err
	}
	return nil
}

package repository

import (
	"context"

	"github.com/salamalfis/projectfinal/internal/infrastructure"
	"github.com/salamalfis/projectfinal/internal/model"
	"gorm.io/gorm"
)

type CommentQuery interface {
	CreateComment(ctx context.Context, comment *model.Comment) error
	GetComments(ctx context.Context, photoid uint64) ([]model.CommentView, error)
	GetCommentsByID(ctx context.Context, id uint64) (*model.Comment, error)
	UpdateCommentsByID(ctx context.Context, comment *model.Comment) error
	DeleteCommentsByID(ctx context.Context, id uint64) error
}

type commentQueryImpl struct {
	db infrastructure.GormPostgres
}

func NewCommentQuery(db infrastructure.GormPostgres) CommentQuery {
	return &commentQueryImpl{db: db}
}

func (c *commentQueryImpl) CreateComment(ctx context.Context, comment *model.Comment) error {
	db := c.db.GetConnection()

	err := db.
		WithContext(ctx).
		Table("comments").
		Create(&comment).
		Error

	return err
}

func (u *commentQueryImpl) GetComments(ctx context.Context, photoid uint64) ([]model.CommentView, error) {
	db := u.db.GetConnection()
	comments := []model.CommentView{}

	err := db.
		WithContext(ctx).
		Table("comments").
		Where("photo_id = ?", photoid).
		Where("deleted_at IS NULL").
		Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, email, username").Table("users").Where("deleted_at is null")
		}).
		Preload("Photo", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, title, caption, photo_url, user_id").Table("photos").Where("deleted_at is null")
		}).
		Find(&comments).
		Error

	if err != nil {
		return nil, err
	}

	return comments, nil
}

func (u *commentQueryImpl) GetCommentsByID(ctx context.Context, id uint64) (*model.Comment, error) {

	db := u.db.GetConnection()
	comment := model.Comment{}

	err := db.
		WithContext(ctx).
		Table("comments").
		Where("id = ?", id).
		Find(&comment).
		Error

	if err != nil {
		return nil, err
	}

	return &comment, nil
}

func (u *commentQueryImpl) UpdateCommentsByID(ctx context.Context, comment *model.Comment) error {

	db := u.db.GetConnection()
	err := db.
		WithContext(ctx).
		Updates(&comment).
		Error

	return err
}

func (u *commentQueryImpl) DeleteCommentsByID(ctx context.Context, id uint64) error {
	db := u.db.GetConnection()
	if err := db.
		WithContext(ctx).
		Table("comments").
		Delete(&model.Comment{ID: uint64(id)}).
		Error; err != nil {
		return err
	}
	return nil
}

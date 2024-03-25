package service

import (
	"context"

	"github.com/salamalfis/projectfinal/internal/model"
	"github.com/salamalfis/projectfinal/internal/repository"
)

type SocialMediaService interface {
	AddSocialMedia(ctx context.Context, userId uint64, social model.NewSocialMedia) (*model.CreateSocialMediaRes, error)
	GetSocialMedia(ctx context.Context, userId uint64) ([]model.SocialMediaView, error)
	GetSocialMediaById(ctx context.Context, socialId uint64) (*model.SocialMedia, error)
	UpdateSocialMediaById(ctx context.Context, social model.SocialMedia) (*model.UpdateSocialMediaRes, error)
	DeleteSocialMediaById(ctx context.Context, socialId uint64) error
}

type socialMediaServiceImpl struct {
	repo repository.SocialMediaQuery
}

func NewSocialMediaService(repo repository.SocialMediaQuery) SocialMediaService {
	return &socialMediaServiceImpl{repo: repo}
}

func (s *socialMediaServiceImpl) AddSocialMedia(ctx context.Context, userId uint64, newSocial model.NewSocialMedia) (*model.CreateSocialMediaRes, error) {
	social := model.SocialMedia{}
	social.Name = newSocial.Name
	social.SocialMediaUrl = newSocial.SocialMediaUrl
	social.UserId = userId

	err := s.repo.AddSocialMedia(ctx, &social)
	if err != nil {
		return nil, err
	}

	socialMediaRes := model.CreateSocialMediaRes{}
	socialMediaRes.ID = social.ID
	socialMediaRes.Name = social.Name
	socialMediaRes.UserId = social.UserId
	socialMediaRes.SocialMediaUrl = social.SocialMediaUrl
	socialMediaRes.CreatedAt = social.CreatedAt

	return &socialMediaRes, nil
}

func (s *socialMediaServiceImpl) GetSocialMedia(ctx context.Context, userId uint64) ([]model.SocialMediaView, error) {
	socials, err := s.repo.GetSocialMedia(ctx, userId)
	if err != nil {
		return nil, err
	}

	return socials, nil
}

func (s *socialMediaServiceImpl) GetSocialMediaById(ctx context.Context, socialId uint64) (*model.SocialMedia, error) {
	social, err := s.repo.GetSocialMediaById(ctx, socialId)
	if err != nil {
		return nil, err
	}

	return social, nil
}

func (s *socialMediaServiceImpl) UpdateSocialMediaById(ctx context.Context, social model.SocialMedia) (*model.UpdateSocialMediaRes, error) {
	err := s.repo.UpdateSocialMediaById(ctx, &social)

	if err != nil {
		return nil, err
	}

	socialRes := model.UpdateSocialMediaRes{}
	socialRes.ID = social.ID
	socialRes.UserId = social.UserId
	socialRes.Name = social.Name
	socialRes.SocialMediaUrl = social.SocialMediaUrl
	socialRes.UpdatedAt = social.UpdatedAt

	return &socialRes, nil
}

func (s *socialMediaServiceImpl) DeleteSocialMediaById(ctx context.Context, socialId uint64) error {
	return s.repo.DeleteSocialMediaById(ctx, socialId)
}

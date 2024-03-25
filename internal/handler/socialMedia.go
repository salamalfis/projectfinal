package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/salamalfis/projectfinal/internal/model"
	"github.com/salamalfis/projectfinal/internal/service"
	"github.com/salamalfis/projectfinal/pkg"
	"github.com/salamalfis/projectfinal/pkg/helper"
)

// SocialMediaHandler is a struct that defines the SocialMediaHandler
type SocialMediaHandler interface {
	AddSocialMedia(ctx *gin.Context)
	GetSocialMedia(ctx *gin.Context)
	GetSocialMediaById(ctx *gin.Context)
	UpdateSocialMediaById(ctx *gin.Context)
	DeleteSocialMediaById(ctx *gin.Context)
}

type socialMediaHandlerImpl struct {
	svc service.SocialMediaService
}

func NewSocialMediaHandler(svc service.SocialMediaService) SocialMediaHandler {
	return &socialMediaHandlerImpl{
		svc: svc,
	}
}

// AddSocialMedia is a function to add social media
func (s *socialMediaHandlerImpl) AddSocialMedia(ctx *gin.Context) {
	userId, err := helper.GetUserID(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}

	newSocial := model.NewSocialMedia{}
	err = ctx.Bind(&newSocial)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}

	validate := validator.New()
	err = validate.Struct(newSocial)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: err.Error()})
		return
	}

	socialMediaRes, err := s.svc.AddSocialMedia(ctx, userId, newSocial)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, socialMediaRes)
}

// GetSocialMedia is a function to get social media
func (s *socialMediaHandlerImpl) GetSocialMedia(ctx *gin.Context) {
	userIdStr := ctx.Request.URL.Query().Get("userId")
	if userIdStr == "" {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: "Missing User id in query"})
		return
	}
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: err.Error()})
		return
	}

	socials, err := s.svc.GetSocialMedia(ctx, uint64(userId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, socials)
}

// Update is a function to get social media by id
func (s *socialMediaHandlerImpl) UpdateSocialMediaById(ctx *gin.Context) {
	socialId, err := strconv.Atoi(ctx.Param("id"))
	if socialId == 0 || err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: err.Error()})
		return
	}

	social, err := s.svc.GetSocialMediaById(ctx, uint64(socialId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}

	if social.ID == 0 {
		ctx.JSON(http.StatusNotFound, pkg.ErrorResponse{Message: "User Social Media data did not exist"})
		return
	}

	userId, err := helper.GetUserID(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}

	if userId != uint64(social.UserId) {
		ctx.JSON(http.StatusUnauthorized, pkg.ErrorResponse{Message: "unauthorized to do this request"})
		return
	}

	socialUpdateData := model.NewSocialMedia{}
	err = ctx.ShouldBindJSON(&socialUpdateData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}

	validate := validator.New()
	err = validate.Struct(socialUpdateData)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: err.Error()})
		return
	}

	social.Name = socialUpdateData.Name
	social.SocialMediaUrl = socialUpdateData.SocialMediaUrl

	socialMediaRes, err := s.svc.UpdateSocialMediaById(ctx, *social)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, socialMediaRes)
}

// DeleteSocialMedia is a function to delete social media
func (s *socialMediaHandlerImpl) DeleteSocialMediaById(ctx *gin.Context) {
	socialId, err := strconv.Atoi(ctx.Param("id"))
	if socialId == 0 || err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: err.Error()})
		return
	}

	social, err := s.svc.GetSocialMediaById(ctx, uint64(socialId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}

	if social.ID == 0 {
		ctx.JSON(http.StatusNotFound, pkg.ErrorResponse{Message: "User Social Media data did not exist"})
		return
	}

	userId, err := helper.GetUserID(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}

	if userId != uint64(social.UserId) {
		ctx.JSON(http.StatusUnauthorized, pkg.ErrorResponse{Message: "unauthorized to do this request"})
		return
	}

	err = s.svc.DeleteSocialMediaById(ctx, uint64(socialId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, pkg.SuccessResponse{Message: "Your social media has been successfully deleted"})
}

// GetSocialMediaById is a function to get social media by id
func (s *socialMediaHandlerImpl) GetSocialMediaById(ctx *gin.Context) {
	socialId, err := strconv.Atoi(ctx.Param("id"))
	if socialId == 0 || err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: err.Error()})
		return
	}

	social, err := s.svc.GetSocialMediaById(ctx, uint64(socialId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}

	if social.ID == 0 {
		ctx.JSON(http.StatusNotFound, pkg.ErrorResponse{Message: "User Social Media data did not exist"})
		return
	}

	ctx.JSON(http.StatusOK, social)
}

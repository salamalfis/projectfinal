package handler

import (
	"github.com/gin-gonic/gin"
)

// SocialMediaHandler is a struct that defines the SocialMediaHandler
type SocialMediaHandler interface {
	GetSocialMedia(c *gin.Context)
	AddSocialMedia(c *gin.Context)
	DeleteSocialMediaById(c *gin.Context)
	UpdateSocialMediaById(c *gin.Context)
	GetSocialMediaById(c *gin.Context)
}

// SocialMediaHandlerImpl is a struct that defines the SocialMediaHandlerImpl
type SocialMediaHandlerImpl struct {
}


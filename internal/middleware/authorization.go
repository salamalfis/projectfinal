package middleware

import (
	
	"net/http"
	"strings"

	"github.com/salamalfis/projectfinal/pkg"
	"github.com/salamalfis/projectfinal/pkg/helper"
	"github.com/gin-gonic/gin"
)

const (
	STATIC_USERNAME = "golang006awesome"
	STATIC_PASSWORD = "mysecretpassword"

	CLAIM_USER_ID  = "claim_user_id"
	CLAIM_USERNAME = "claim_username"
)



func CheckAuthBearer(ctx *gin.Context) {
	auth := ctx.GetHeader("Authorization")

	authArr := strings.Split(auth, " ")
	if len(authArr) < 2 {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, pkg.ErrorResponse{
			Message: "unauthorized",
			Errors:  []string{"invalid token"},
		})
		return
	}
	if authArr[0] != "Bearer" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, pkg.ErrorResponse{
			Message: "unauthorized",
			Errors:  []string{"invalid authorization method"},
		})
		return
	}

	token := authArr[1]
	claims, err := helper.ValidateToken(token)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, pkg.ErrorResponse{
			Message: "unauthorized",
			Errors:  []string{"invalid token", "failed to decode"},
		})
		return
	}
	ctx.Set(CLAIM_USER_ID, claims["user_id"])
	ctx.Set(CLAIM_USERNAME, claims["username"])
	ctx.Next()
}



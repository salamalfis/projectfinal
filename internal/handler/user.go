package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/salamalfis/projectfinal/internal/middleware"
	"github.com/salamalfis/projectfinal/internal/model"
	"github.com/salamalfis/projectfinal/internal/service"
	"github.com/salamalfis/projectfinal/pkg"
)

type UserHandler interface {
	// users
	GetUsers(ctx *gin.Context)
	GetUsersById(ctx *gin.Context)
	DeleteUsersById(ctx *gin.Context)
	UpdateUsersById(ctx *gin.Context)

	// activity
	UserSignUp(ctx *gin.Context)
	UserLogin(ctx *gin.Context)
}

type userHandlerImpl struct {
	svc service.UserService
}

func NewUserHandler(svc service.UserService) UserHandler {
	return &userHandlerImpl{
		svc: svc,
	}
}

// ShowUsers godoc
//
//	@Summary		Show users list
//	@Description	will fetch 3rd party server to get users data
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	[]model.User
//	@Failure		400	{object}	pkg.ErrorResponse
//	@Failure		404	{object}	pkg.ErrorResponse
//	@Failure		500	{object}	pkg.ErrorResponse
//	@Router			/users [get]
func (u *userHandlerImpl) GetUsers(ctx *gin.Context) {
	users, err := u.svc.GetUsers(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

// ShowUsersById godoc
//
//	@Summary		Show users detail
//	@Description	will fetch 3rd party server to get users data to get detail user
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"User ID"
//	@Success		200	{object}	model.User
//	@Failure		400	{object}	pkg.ErrorResponse
//	@Failure		404	{object}	pkg.ErrorResponse
//	@Failure		500	{object}	pkg.ErrorResponse
//	@Router			/users/{id} [get]
func (u *userHandlerImpl) GetUsersById(ctx *gin.Context) {
	// get id user
	id, err := strconv.Atoi(ctx.Param("id"))
	if id == 0 || err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: "invalid required param"})
		return
	}
	user, err := u.svc.GetUsersById(ctx, uint64(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}
	if user.ID == 0 {
		ctx.JSON(http.StatusNotFound, pkg.ErrorResponse{Message: "user not found"})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (u *userHandlerImpl) UserSignUp(ctx *gin.Context) {
	// binding sign-up body
	userSignUp := model.UserSignUp{}
	if err := ctx.Bind(&userSignUp); err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: err.Error()})
		return
	}

	if err := userSignUp.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: err.Error()})
		return
	}

	user, err := u.svc.SignUp(ctx, userSignUp)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}

	token, err := u.svc.GenerateUserAccessToken(ctx, user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
	}

	ctx.JSON(http.StatusOK, map[string]any{
		"token": token,
	})
}

// DeleteUsersById godoc
//
//		@Summary		Delete user by selected id
//		@Description	will delete user with given id from param
//		@Tags			users
//		@Accept			json
//		@Produce		json
//	 	@Param 			Authorization header string true "bearer token"
//		@Param			id	path		int	true	"User ID"
//		@Success		200	{object}	model.User
//		@Failure		400	{object}	pkg.ErrorResponse
//		@Failure		404	{object}	pkg.ErrorResponse
//		@Failure		500	{object}	pkg.ErrorResponse
//		@Router			/users/{id} [delete]
func (u *userHandlerImpl) DeleteUsersById(ctx *gin.Context) {
	// get id user
	id, err := strconv.Atoi(ctx.Param("id"))
	if id == 0 || err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: "invalid required param"})
		return
	}

	// check user id session from context
	userId, ok := ctx.Get(middleware.CLAIM_USER_ID)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, pkg.ErrorResponse{Message: "invalid user session"})
		return
	}
	userIdInt, ok := userId.(float64)
	if !ok {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: "invalid user id session"})
		return
	}
	if id != int(userIdInt) {
		ctx.JSON(http.StatusUnauthorized, pkg.ErrorResponse{Message: "invalid user request"})
		return
	}

	user, err := u.svc.DeleteUsersById(ctx, uint64(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}
	if user.ID == 0 {
		ctx.JSON(http.StatusNotFound, pkg.ErrorResponse{Message: "user not found"})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

// UpdateUsersById godoc
// @Summary Update user by selected id
// @Description will update user with given id from param
// @Tags users
// @Accept json
// @Produce json
// @Param Authorization header string true "
// @Param id path int true "User ID"
// @Param username body string true "Username"
// @Param email body string true "Email"
// @Success 200 {object} model.User
// @Failure 400 {object} pkg.ErrorResponse
// @Failure 404 {object} pkg.ErrorResponse
// @Failure 500 {object} pkg.ErrorResponse
// @Router /users/{id} [put]

func (u *userHandlerImpl) UpdateUsersById(ctx *gin.Context) {
	// get id user
	id, err := strconv.Atoi(ctx.Param("id"))
	if id == 0 || err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: "invalid required param"})
		return
	}

	// check user id session from context
	userId, ok := ctx.Get(middleware.CLAIM_USER_ID)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, pkg.ErrorResponse{Message: "invalid user session"})
		return
	}
	userIdInt, ok := userId.(float64)
	if !ok {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: "invalid user id session"})
		return
	}
	if id != int(userIdInt) {
		ctx.JSON(http.StatusUnauthorized, pkg.ErrorResponse{Message: "invalid user request"})
		return
	}

	// binding update body
	userUpdate := model.UserUpdate{}
	if err := ctx.Bind(&userUpdate); err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: err.Error()})
		return
	}

	validate := validator.New()
	err = validate.Struct(userUpdate)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: err.Error()})
		return
	}

	user := model.User{ID: uint64(id), Username: userUpdate.Username, Email: userUpdate.Email}

	userRes, err := u.svc.UpdateUsersById(ctx, user.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, userRes)
}

// UserLogin godoc
// @Summary User login
// @Description will login user with given username and password
// @Tags users
// @Accept json
// @Produce json
// @Param username body string true "Username"
// @Param password body string true "Password"
// @Success 200 {object} model.User
// @Failure 400 {object} pkg.ErrorResponse
// @Failure 404 {object} pkg.ErrorResponse
// @Failure 500 {object} pkg.ErrorResponse
// @Router /users/login [post]
func (u *userHandlerImpl) UserLogin(ctx *gin.Context) {
	user := model.UserLogin{}
	if err := ctx.Bind(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: err.Error()})
		return
	}

	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: err.Error()})
		return
	}

	userRes, err := u.svc.UserLogin(ctx, user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}
	token, err := u.svc.GenerateUserAccessToken(ctx, userRes)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
	}

	ctx.JSON(http.StatusOK, pkg.TokenResponse{Token: token})

}

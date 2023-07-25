package handlers

import (
	"github.com/rierarizzo/cafelatte/internal/core/constants"
	"github.com/rierarizzo/cafelatte/internal/core/errors"
	errorHandler "github.com/rierarizzo/cafelatte/internal/infrastructure/api/error"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/api/mappers"
	"github.com/rierarizzo/cafelatte/internal/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rierarizzo/cafelatte/internal/core/ports"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/api/dto"
)

type UserHandler struct {
	userService ports.IUserService
}

func (uc *UserHandler) SignUp(c *gin.Context) {
	var signUpRequest dto.SignUpRequest
	err := c.BindJSON(&signUpRequest)
	if errorHandler.Error(c, err) {
		return
	}

	authorizedUser, err := uc.userService.SignUp(*mappers.FromSignUpReqToUser(signUpRequest))
	if errorHandler.Error(c, err) {
		return
	}

	c.JSON(http.StatusCreated, mappers.FromAuthorizedUserToAuthorizationRes(*authorizedUser))
}

func (uc *UserHandler) SignIn(c *gin.Context) {
	var signInRequest dto.SignInRequest
	err := c.BindJSON(&signInRequest)
	if errorHandler.Error(c, err) {
		return
	}

	authorizedUser, err := uc.userService.SignIn(signInRequest.Email, signInRequest.Password)
	if errorHandler.Error(c, err) {
		return
	}

	c.JSON(http.StatusOK, mappers.FromAuthorizedUserToAuthorizationRes(*authorizedUser))
}

func (uc *UserHandler) GetAllUsers(c *gin.Context) {
	users, err := uc.userService.GetAllUsers()
	if errorHandler.Error(c, err) {
		return
	}

	userResponse := make([]dto.UserResponse, 0)
	for _, k := range users {
		userResponse = append(userResponse, *mappers.FromUserToUserRes(k))
	}

	c.JSON(http.StatusOK, userResponse)
}

func (uc *UserHandler) FindUserByID(c *gin.Context) {
	userIDParam := c.Param("userID")
	userID, err := strconv.Atoi(userIDParam)
	if errorHandler.Error(c, err) {
		return
	}

	user, err := uc.userService.FindUserByID(userID)
	if errorHandler.Error(c, err) {
		return
	}

	c.JSON(http.StatusOK, *mappers.FromUserToUserRes(*user))
}

func getUserIDFromClaims(c *gin.Context) (int, error) {
	userClaims, exists := c.Get(constants.UserClaimsKey)
	if !exists {
		return 0, errors.WrapError(errors.ErrUnauthorizedUser, "claims not present in args")
	}

	return userClaims.(*utils.UserClaims).ID, nil
}

func NewUserHandler(userService ports.IUserService) *UserHandler {
	return &UserHandler{userService}
}

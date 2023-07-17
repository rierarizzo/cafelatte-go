package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rierarizzo/cafelatte/internal/core"
	"github.com/rierarizzo/cafelatte/internal/core/ports"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/api/dto"
	"github.com/rierarizzo/cafelatte/internal/utils"
)

type UserHandler struct {
	userService ports.IUserService
}

func (uc *UserHandler) SignUp(c *gin.Context) {
	var signUpRequest dto.SignUpRequest
	if err := c.BindJSON(&signUpRequest); err != nil {
		utils.HTTPError(core.BadRequest, c)
		return
	}

	user, err := uc.userService.SignUp(*signUpRequest.ToUserCore())
	if err != nil {
		utils.HTTPError(err, c)
		return
	}

	var authResponse dto.AuthResponse
	authResponse.LoadFromAuthorizedUserCore(*user)
	c.JSON(http.StatusOK, authResponse)
}

func (uc *UserHandler) SignIn(c *gin.Context) {
	var signInRequest dto.SignInRequest
	if err := c.BindJSON(&signInRequest); err != nil {
		utils.HTTPError(core.BadRequest, c)
		return
	}

	user, err := uc.userService.SignIn(signInRequest.Email, signInRequest.Password)
	if err != nil {
		utils.HTTPError(err, c)
		return
	}

	var authResponse dto.AuthResponse
	authResponse.LoadFromAuthorizedUserCore(*user)
	c.JSON(http.StatusOK, authResponse)
}

func (uc *UserHandler) GetAllUsers(c *gin.Context) {
	users, err := uc.userService.GetAllUsers()
	if err != nil {
		utils.HTTPError(err, c)
		return
	}

	var userResponse []dto.UserResponse
	for _, k := range users {
		var res dto.UserResponse
		res.LoadFromUserCore(k)
		userResponse = append(userResponse, res)
	}

	c.JSON(http.StatusOK, userResponse)
}

func (uc *UserHandler) FindUser(c *gin.Context) {
	userIDParam := c.Param("userID")
	userID, err := strconv.Atoi(userIDParam)
	if err != nil {
		utils.HTTPError(core.BadRequest, c)
		return
	}

	user, err := uc.userService.FindUserById(userID)
	if err != nil {
		utils.HTTPError(err, c)
		return
	}

	var userResponse dto.UserResponse
	userResponse.LoadFromUserCore(*user)

	authResponse := dto.AuthResponse{
		User:        userResponse,
		AccessToken: "",
	}

	c.JSON(http.StatusCreated, authResponse)
}

func NewUserHandler(userService ports.IUserService) *UserHandler {
	return &UserHandler{userService}
}

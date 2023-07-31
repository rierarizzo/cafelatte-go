package handlers

import (
	"github.com/rierarizzo/cafelatte/internal/infra/api/mappers"
	"github.com/rierarizzo/cafelatte/internal/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rierarizzo/cafelatte/internal/domain/ports"
	"github.com/rierarizzo/cafelatte/internal/infra/api/dto"
)

type UserHandler struct {
	userService ports.IUserService
}

func (h *UserHandler) SignUp(c *gin.Context) {
	var signUpRequest dto.SignUpRequest
	err := c.BindJSON(&signUpRequest)
	if err != nil {
		utils.AbortWithError(c, err)
		return
	}

	authorizedUser, err := h.userService.SignUp(*mappers.FromSignUpReqToUser(signUpRequest))
	if err != nil {
		utils.AbortWithError(c, err)
		return
	}

	c.JSON(http.StatusCreated,
		mappers.FromAuthorizedUserToAuthorizationRes(*authorizedUser))
}

func (h *UserHandler) SignIn(c *gin.Context) {
	var signInRequest dto.SignInRequest
	err := c.BindJSON(&signInRequest)
	if err != nil {
		utils.AbortWithError(c, err)
		return
	}

	authorizedUser, err := h.userService.SignIn(signInRequest.Email,
		signInRequest.Password)
	if err != nil {
		utils.AbortWithError(c, err)
		return
	}

	c.JSON(http.StatusOK,
		mappers.FromAuthorizedUserToAuthorizationRes(*authorizedUser))
}

func (h *UserHandler) GetAllUsers(c *gin.Context) {
	users, err := h.userService.GetUsers()
	if err != nil {
		utils.AbortWithError(c, err)
		return
	}

	userResponse := make([]dto.UserResponse, 0)
	for _, k := range users {
		userResponse = append(userResponse, *mappers.FromUserToUserRes(k))
	}

	c.JSON(http.StatusOK, userResponse)
}

func (h *UserHandler) FindUserByID(c *gin.Context) {
	userIDParam := c.Param("userID")
	userID, err := strconv.Atoi(userIDParam)
	if err != nil {
		utils.AbortWithError(c, err)
		return
	}

	user, err := h.userService.FindUserByID(userID)
	if err != nil {
		utils.AbortWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, *mappers.FromUserToUserRes(*user))
}

func NewUserHandler(userService ports.IUserService) *UserHandler {
	return &UserHandler{userService}
}

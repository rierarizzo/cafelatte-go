package handlers

import (
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
	if err := c.BindJSON(&signUpRequest); err != nil {
		c.Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := uc.userService.SignUp(*signUpRequest.ToUserCore())
	if err != nil {
		c.Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var authResponse dto.AuthResponse
	authResponse.LoadFromAuthorizedUserCore(*user)
	c.JSON(http.StatusCreated, authResponse)
}

func (uc *UserHandler) SignIn(c *gin.Context) {
	var signInRequest dto.SignInRequest
	if err := c.BindJSON(&signInRequest); err != nil {
		c.Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := uc.userService.SignIn(signInRequest.Email, signInRequest.Password)
	if err != nil {
		c.Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var authResponse dto.AuthResponse
	authResponse.LoadFromAuthorizedUserCore(*user)
	c.JSON(http.StatusCreated, authResponse)
}

func (uc *UserHandler) FindUser(c *gin.Context) {
	userIDParam := c.Param("userID")
	userID, err := strconv.Atoi(userIDParam)
	if err != nil {
		c.Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := uc.userService.FindUserById(userID)
	if err != nil {
		c.Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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

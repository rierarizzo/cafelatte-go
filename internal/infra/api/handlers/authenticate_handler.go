package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/rierarizzo/cafelatte/internal/domain/ports"
	"github.com/rierarizzo/cafelatte/internal/infra/api/dto"
	"github.com/rierarizzo/cafelatte/internal/infra/api/mappers"
	"github.com/rierarizzo/cafelatte/internal/utils"
	"net/http"
)

type AuthenticateHandler struct {
	authUsecase ports.IAuthenticateUsecase
}

func (h *AuthenticateHandler) SignUp(c *gin.Context) {
	var signUpRequest dto.SignUpRequest
	err := c.BindJSON(&signUpRequest)
	if err != nil {
		utils.AbortWithError(c, err)
		return
	}

	authorizedUser, err := h.authUsecase.SignUp(*mappers.FromSignUpReqToUser(signUpRequest))
	if err != nil {
		utils.AbortWithError(c, err)
		return
	}

	c.JSON(http.StatusCreated,
		mappers.FromAuthorizedUserToAuthorizationRes(*authorizedUser))
}

func (h *AuthenticateHandler) SignIn(c *gin.Context) {
	var signInRequest dto.SignInRequest
	err := c.BindJSON(&signInRequest)
	if err != nil {
		utils.AbortWithError(c, err)
		return
	}

	authorizedUser, err := h.authUsecase.SignIn(signInRequest.Email,
		signInRequest.Password)
	if err != nil {
		utils.AbortWithError(c, err)
		return
	}

	c.JSON(http.StatusOK,
		mappers.FromAuthorizedUserToAuthorizationRes(*authorizedUser))
}

func NewAuthHandler(authUsecase ports.IAuthenticateUsecase) *AuthenticateHandler {
	return &AuthenticateHandler{authUsecase}
}

package handlers

import (
	"github.com/gin-gonic/gin"
	domain "github.com/rierarizzo/cafelatte/internal/domain/errors"
	"github.com/rierarizzo/cafelatte/internal/domain/ports"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/api/dto"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/api/mappers"
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
		utils.AbortWithError(c, domain.NewAppError(err, domain.BadRequestError))
		return
	}

	authorizedUser, appErr := h.authUsecase.SignUp(mappers.FromSignUpReqToUser(signUpRequest))
	if appErr != nil {
		utils.AbortWithError(c, appErr)
		return
	}

	response := mappers.FromAuthorizedUserToAuthorizationRes(*authorizedUser)
	utils.RespondWithJSON(c, http.StatusCreated, response)
}

func (h *AuthenticateHandler) SignIn(c *gin.Context) {
	var signInRequest dto.SignInRequest
	err := c.BindJSON(&signInRequest)
	if err != nil {
		utils.AbortWithError(c, domain.NewAppError(err, domain.BadRequestError))
		return
	}

	authorizedUser, appErr := h.authUsecase.SignIn(signInRequest.Email,
		signInRequest.Password)
	if appErr != nil {
		utils.AbortWithError(c, appErr)
		return
	}

	response := mappers.FromAuthorizedUserToAuthorizationRes(*authorizedUser)
	utils.RespondWithJSON(c, http.StatusOK, response)
}

func NewAuthHandler(authUsecase ports.IAuthenticateUsecase) *AuthenticateHandler {
	return &AuthenticateHandler{authUsecase}
}

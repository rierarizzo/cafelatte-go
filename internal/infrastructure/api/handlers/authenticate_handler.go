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
	if err := c.BindJSON(&signUpRequest); err != nil {
		appErr := domain.NewAppError(err, domain.BadRequestError)
		utils.AbortWithError(c, appErr)
		return
	}

	authorized, appErr := h.authUsecase.SignUp(mappers.FromSignUpReqToUser(signUpRequest))
	if appErr != nil {
		utils.AbortWithError(c, appErr)
		return
	}

	utils.RespondWithJSON(c, http.StatusCreated,
		mappers.FromAuthorizedUserToAuthorizationRes(*authorized))
}

func (h *AuthenticateHandler) SignIn(c *gin.Context) {
	var signInRequest dto.SignInRequest
	if err := c.BindJSON(&signInRequest); err != nil {
		appErr := domain.NewAppError(err, domain.BadRequestError)
		utils.AbortWithError(c, appErr)
		return
	}

	authorized, appErr := h.authUsecase.SignIn(signInRequest.Email,
		signInRequest.Password)
	if appErr != nil {
		utils.AbortWithError(c, appErr)
		return
	}

	utils.RespondWithJSON(c, http.StatusOK,
		mappers.FromAuthorizedUserToAuthorizationRes(*authorized))
}

func NewAuthHandler(authUsecase ports.IAuthenticateUsecase) *AuthenticateHandler {
	return &AuthenticateHandler{authUsecase}
}

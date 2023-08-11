package authenticate

import (
	"github.com/gin-gonic/gin"
	"github.com/rierarizzo/cafelatte/internal/domain/authenticate"
	domain "github.com/rierarizzo/cafelatte/internal/domain/errors"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/api/http/user"
	http2 "github.com/rierarizzo/cafelatte/pkg/utils/http"
	"net/http"
)

type Handler struct {
	authUsecase authenticate.IAuthenticateUsecase
}

func (h *Handler) SignUp(c *gin.Context) {
	var signUpRequest user.SignUpRequest
	if err := c.BindJSON(&signUpRequest); err != nil {
		appErr := domain.NewAppError(err, domain.BadRequestError)
		http2.AbortWithError(c, appErr)
		return
	}

	authorized, appErr := h.authUsecase.SignUp(user.FromSignUpReqToUser(signUpRequest))
	if appErr != nil {
		http2.AbortWithError(c, appErr)
		return
	}

	http2.RespondWithJSON(c, http.StatusCreated,
		FromAuthorizedUserToAuthorizationRes(*authorized))
}

func (h *Handler) SignIn(c *gin.Context) {
	var signInRequest user.SignInRequest
	if err := c.BindJSON(&signInRequest); err != nil {
		appErr := domain.NewAppError(err, domain.BadRequestError)
		http2.AbortWithError(c, appErr)
		return
	}

	authorized, appErr := h.authUsecase.SignIn(signInRequest.Email,
		signInRequest.Password)
	if appErr != nil {
		http2.AbortWithError(c, appErr)
		return
	}

	http2.RespondWithJSON(c, http.StatusOK,
		FromAuthorizedUserToAuthorizationRes(*authorized))
}

func NewAuthHandler(authUsecase authenticate.IAuthenticateUsecase) *Handler {
	return &Handler{authUsecase}
}

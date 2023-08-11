package authenticate

import (
	"github.com/gin-gonic/gin"
	"github.com/rierarizzo/cafelatte/internal/domain/authenticate"
	domain "github.com/rierarizzo/cafelatte/internal/domain/errors"
	httpUtil "github.com/rierarizzo/cafelatte/pkg/utils/http"
	"net/http"
)

type Handler struct {
	authUsecase authenticate.IAuthenticateUsecase
}

func (h *Handler) SignUp(c *gin.Context) {
	var signUpRequest SignUpRequest
	if err := c.BindJSON(&signUpRequest); err != nil {
		appErr := domain.NewAppError(err, domain.BadRequestError)
		httpUtil.AbortWithError(c, appErr)
		return
	}

	authorized, appErr := h.authUsecase.SignUp(fromSignUpRequestToUser(signUpRequest))
	if appErr != nil {
		httpUtil.AbortWithError(c, appErr)
		return
	}

	httpUtil.RespondWithJSON(c, http.StatusCreated,
		fromAuthUserToResponse(*authorized))
}

func (h *Handler) SignIn(c *gin.Context) {
	var signInRequest SignInRequest
	if err := c.BindJSON(&signInRequest); err != nil {
		appErr := domain.NewAppError(err, domain.BadRequestError)
		httpUtil.AbortWithError(c, appErr)
		return
	}

	authorized, appErr := h.authUsecase.SignIn(signInRequest.Email,
		signInRequest.Password)
	if appErr != nil {
		httpUtil.AbortWithError(c, appErr)
		return
	}

	httpUtil.RespondWithJSON(c, http.StatusOK,
		fromAuthUserToResponse(*authorized))
}

func NewAuthHandler(authUsecase authenticate.IAuthenticateUsecase) *Handler {
	return &Handler{authUsecase}
}

package authenticator

import (
	"github.com/gin-gonic/gin"
	"github.com/rierarizzo/cafelatte/internal/domain"
	"github.com/rierarizzo/cafelatte/internal/domain/authenticator"
	httpUtil "github.com/rierarizzo/cafelatte/pkg/utils/http"
	"net/http"
)

type Handler struct {
	authenticator authenticator.Authenticator
}

func (h *Handler) SignUp(c *gin.Context) {
	var signUpRequest SignUpRequest
	if err := c.BindJSON(&signUpRequest); err != nil {
		appErr := domain.NewAppError(err, domain.BadRequestError)
		httpUtil.AbortWithError(c, appErr)
		return
	}

	authorized, appErr := h.authenticator.SignUp(fromRequestToUser(signUpRequest))
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

	authorized, appErr := h.authenticator.SignIn(signInRequest.Email,
		signInRequest.Password)
	if appErr != nil {
		httpUtil.AbortWithError(c, appErr)
		return
	}

	httpUtil.RespondWithJSON(c, http.StatusOK,
		fromAuthUserToResponse(*authorized))
}

func NewAuthHandler(authenticator authenticator.Authenticator) *Handler {
	return &Handler{authenticator}
}

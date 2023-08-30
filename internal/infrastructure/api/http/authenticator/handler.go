package authenticator

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rierarizzo/cafelatte/internal/domain"
	"github.com/rierarizzo/cafelatte/internal/domain/authenticator"
)

type Handler struct {
	authenticator authenticator.Authenticator
}

func Router(group *echo.Group) func(authenticatorHandler *Handler) {
	return func(h *Handler) {
		group.POST("/signup", h.SignUp)
		group.POST("/signin", h.SignIn)
	}
}

func (h *Handler) SignUp(c echo.Context) error {
	var req UserSignup
	if err := c.Bind(&req); err != nil {
		return domain.NewAppError(err, domain.BadRequestError)
	}

	authorized, appErr := h.authenticator.SignUp(fromRequestToUser(req))
	if appErr != nil {
		return appErr
	}

	return c.JSON(http.StatusCreated, fromAuthUserToResponse(*authorized))
}

func (h *Handler) SignIn(c echo.Context) error {
	var req UserSignin
	if err := c.Bind(&req); err != nil {
		return domain.NewAppError(err, domain.BadRequestError)
	}

	authorized, appErr := h.authenticator.SignIn(req.Email, req.Password)
	if appErr != nil {
		return appErr
	}

	return c.JSON(http.StatusOK, fromAuthUserToResponse(*authorized))
}

func New(authenticator authenticator.Authenticator) *Handler {
	return &Handler{authenticator}
}

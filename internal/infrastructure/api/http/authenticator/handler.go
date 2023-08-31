package authenticator

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rierarizzo/cafelatte/internal/domain"
	"github.com/rierarizzo/cafelatte/internal/domain/authenticator"
)

func ConfigureRouting(group *echo.Group) func(a authenticator.Authenticator) {
	return func(a authenticator.Authenticator) {
		group.POST("/signup", signUp(a))
		group.POST("/signin", signIn(a))
	}
}

func signUp(a authenticator.Authenticator) echo.HandlerFunc {
	return func(c echo.Context) error {
		var req UserSignup
		if err := c.Bind(&req); err != nil {
			appErr := domain.NewAppError(err, domain.BadRequestError)
			return appErr
		}

		authorized, appErr := a.SignUp(fromRequestToUser(req))
		if appErr != nil {
			return appErr
		}

		response := fromAuthUserToResponse(authorized)
		return c.JSON(http.StatusCreated, response)
	}
}

func signIn(a authenticator.Authenticator) echo.HandlerFunc {
	return func(c echo.Context) error {
		var req UserSignin
		if err := c.Bind(&req); err != nil {
			appErr := domain.NewAppError(err, domain.BadRequestError)
			return appErr
		}

		authorized, appErr := a.SignIn(req.Email, req.Password)
		if appErr != nil {
			return appErr
		}

		response := fromAuthUserToResponse(authorized)
		return c.JSON(http.StatusOK, response)
	}
}

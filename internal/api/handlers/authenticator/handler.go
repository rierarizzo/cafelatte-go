package authenticator

import (
	"github.com/rierarizzo/cafelatte/internal/usecases/authenticator"
	httpUtil "github.com/rierarizzo/cafelatte/pkg/utils/http"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rierarizzo/cafelatte/internal/domain"
)

func ConfigureRouting(group *echo.Group) func(a authenticator.Authenticator) {
	return func(a authenticator.Authenticator) {
		group.POST("/signup", signUp(a))
		group.POST("/signin", signIn(a))
	}
}

func signUp(a authenticator.Authenticator) echo.HandlerFunc {
	return func(c echo.Context) error {
		var request UserSignup
		err := c.Bind(&request)
		if err != nil {
			appErr := domain.NewAppError(err, domain.BadRequestError)
			return appErr
		}

		err = request.Validate()
		if err != nil {
			appErr := domain.NewAppError(err, domain.BadRequestError)
			return appErr
		}

		authorized, appErr := a.SignUp(fromRequestToUser(request))
		if appErr != nil {
			return appErr
		}

		response := fromAuthUserToResponse(authorized)
		return httpUtil.RespondWithJSON(c, http.StatusCreated, response)
	}
}

func signIn(a authenticator.Authenticator) echo.HandlerFunc {
	return func(c echo.Context) error {
		var request UserSignin
		err := c.Bind(&request)
		if err != nil {
			appErr := domain.NewAppError(err, domain.BadRequestError)
			return appErr
		}

		err = request.Validate()
		if err != nil {
			appErr := domain.NewAppError(err, domain.BadRequestError)
			return appErr
		}

		authorized, appErr := a.SignIn(request.Email, request.Password)
		if appErr != nil {
			return appErr
		}

		response := fromAuthUserToResponse(authorized)
		return httpUtil.RespondWithJSON(c, http.StatusOK, response)
	}
}

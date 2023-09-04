package usermanager

import (
	"fmt"
	"github.com/rierarizzo/cafelatte/internal/usecases/usermanager"
	httpUtil "github.com/rierarizzo/cafelatte/pkg/utils/http"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/rierarizzo/cafelatte/internal/domain"
)

func getAllUsers(m usermanager.Manager) echo.HandlerFunc {
	return func(c echo.Context) error {
		users, appErr := m.GetUsers()
		if appErr != nil {
			return appErr
		}

		response := fromUsersToResponse(users)
		return httpUtil.RespondWithJSON(c, http.StatusOK, response)
	}
}

func findUserById(m usermanager.Manager) echo.HandlerFunc {
	return func(c echo.Context) error {
		userId, err := strconv.Atoi(c.Param("userId"))
		if err != nil {
			return domain.NewAppError(err, domain.BadRequestError)
		}

		user, appErr := m.FindUserById(userId)
		if appErr != nil {
			return appErr
		}

		response := fromUserToResponse(user)
		return httpUtil.RespondWithJSON(c, http.StatusOK, response)
	}
}

func updateUserById(m usermanager.Manager) echo.HandlerFunc {
	return func(c echo.Context) error {
		userId, err := strconv.Atoi(c.Param("userId"))
		if err != nil {
			return domain.NewAppError(err, domain.BadRequestError)
		}

		var request UserUpdate
		if err = c.Bind(&request); err != nil {
			return domain.NewAppError(err, domain.BadRequestError)
		}

		appErr := m.UpdateUserById(userId, fromRequestToUser(request))
		if appErr != nil {
			return appErr
		}

		response := fmt.Sprintf("User %v updated", userId)
		return httpUtil.RespondWithJSON(c, http.StatusOK, response)
	}
}

func deleteUserById(m usermanager.Manager) echo.HandlerFunc {
	return func(c echo.Context) error {
		userId, err := strconv.Atoi(c.Param("userId"))
		if err != nil {
			return domain.NewAppError(err, domain.BadRequestError)
		}

		if appErr := m.DeleteUserById(userId); appErr != nil {
			return appErr
		}

		response := fmt.Sprintf("User %v deleted", userId)
		return httpUtil.RespondWithJSON(c, http.StatusOK, response)
	}
}

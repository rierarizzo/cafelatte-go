package usermanager

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/rierarizzo/cafelatte/internal/domain"
	"github.com/rierarizzo/cafelatte/internal/domain/usermanager"
)

func ConfigureRouting(g *echo.Group) func(m usermanager.Manager) {
	return func(m usermanager.Manager) {
		g.GET("/find", getAllUsers(m))
		g.GET("/find/:userId", findUserById(m))
		g.PUT("/update/:userId", updateUserById(m))
		g.DELETE("/delete/:userId", deleteUserById(m))
	}
}

func getAllUsers(m usermanager.Manager) echo.HandlerFunc {
	return func(c echo.Context) error {
		users, appErr := m.GetUsers()
		if appErr != nil {
			return appErr
		}

		response := fromUsersToResponse(users)
		return c.JSON(http.StatusOK, response)
	}
}

func findUserById(m usermanager.Manager) echo.HandlerFunc {
	return func(c echo.Context) error {
		userId, err := strconv.Atoi(c.Param("userId"))
		if err != nil {
			appErr := domain.NewAppError(err, domain.BadRequestError)
			return appErr
		}

		user, appErr := m.FindUserById(userId)
		if appErr != nil {
			return appErr
		}

		response := fromUserToResponse(user)
		return c.JSON(http.StatusOK, response)
	}
}

func updateUserById(m usermanager.Manager) echo.HandlerFunc {
	return func(c echo.Context) error {
		userId, err := strconv.Atoi(c.Param("userId"))
		if err != nil {
			appErr := domain.NewAppError(err, domain.BadRequestError)
			return appErr
		}

		var request UserUpdate
		err = c.Bind(&request)
		if err != nil {
			appErr := domain.NewAppError(err, domain.BadRequestError)
			return appErr
		}

		appErr := m.UpdateUserById(userId, fromRequestToUser(request))
		if appErr != nil {
			return appErr
		}

		response := fmt.Sprintf("User %v updated", userId)
		return c.JSON(http.StatusOK, response)
	}
}

func deleteUserById(m usermanager.Manager) echo.HandlerFunc {
	return func(c echo.Context) error {
		userId, err := strconv.Atoi(c.Param("userId"))
		if err != nil {
			appErr := domain.NewAppError(err, domain.BadRequestError)
			return appErr
		}

		appErr := m.DeleteUserById(userId)
		if appErr != nil {
			return appErr
		}

		response := fmt.Sprintf("User %v deleted", userId)
		return c.JSON(http.StatusOK, response)
	}
}

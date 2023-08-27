package usermanager

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/rierarizzo/cafelatte/internal/domain"
	"github.com/rierarizzo/cafelatte/internal/domain/usermanager"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/api/http/authenticator"
	"net/http"
	"strconv"
)

type Handler struct {
	userManager usermanager.Manager
}

func Router(group *echo.Group) func(userManagerHandler *Handler) {
	return func(handler *Handler) {
		group.Use(authenticator.Middleware)

		group.GET("/find", handler.GetUsers)
		group.GET("/find/:userId", handler.FindUserById)
		group.PUT("/update/:userId", handler.UpdateUserById)
		group.DELETE("/delete/:userId", handler.DeleteUserById)
	}
}

func (handler *Handler) GetUsers(c echo.Context) error {
	users, appErr := handler.userManager.GetUsers()
	if appErr != nil {
		return appErr
	}

	return c.JSON(http.StatusOK, fromUsersToResponse(users))
}

func (handler *Handler) FindUserById(c echo.Context) error {
	userId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		return domain.NewAppError(err, domain.BadRequestError)
	}

	user, appErr := handler.userManager.FindUserByID(userId)
	if appErr != nil {
		return appErr
	}

	return c.JSON(http.StatusOK, fromUserToResponse(*user))
}

func (handler *Handler) UpdateUserById(c echo.Context) error {
	var req UpdateUserRequest
	userId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		return domain.NewAppError(err, domain.BadRequestError)
	}

	if err = c.Bind(&req); err != nil {
		return domain.NewAppError(err, domain.BadRequestError)
	}

	if appErr := handler.userManager.UpdateUser(userId,
		fromRequestToUser(req)); appErr != nil {
		return appErr
	}

	return c.JSON(http.StatusOK, fmt.Sprintf("User %v updated", userId))
}

func (handler *Handler) DeleteUserById(c echo.Context) error {
	userId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		return domain.NewAppError(err, domain.BadRequestError)
	}

	if appErr := handler.userManager.DeleteUser(userId); appErr != nil {
		return appErr
	}

	return c.JSON(http.StatusOK, fmt.Sprintf("User %v deleted", userId))
}

func New(userManager usermanager.Manager) *Handler {
	return &Handler{userManager}
}

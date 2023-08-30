package usermanager

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/rierarizzo/cafelatte/internal/domain"
	"github.com/rierarizzo/cafelatte/internal/domain/usermanager"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/api/http/authenticator"
)

type Handler struct {
	userManager usermanager.Manager
}

func Router(group *echo.Group) func(userManagerHandler *Handler) {
	return func(h *Handler) {
		group.Use(authenticator.Middleware)

		group.GET("/find", h.GetUsers)
		group.GET("/find/:userId", h.FindUserById)
		group.PUT("/update/:userId", h.UpdateUserById)
		group.DELETE("/delete/:userId", h.DeleteUserById)
	}
}

func (h *Handler) GetUsers(c echo.Context) error {
	users, appErr := h.userManager.GetUsers()
	if appErr != nil {
		return appErr
	}

	return c.JSON(http.StatusOK, fromUsersToResponse(users))
}

func (h *Handler) FindUserById(c echo.Context) error {
	userId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		return domain.NewAppError(err, domain.BadRequestError)
	}

	user, appErr := h.userManager.FindUserById(userId)
	if appErr != nil {
		return appErr
	}

	return c.JSON(http.StatusOK, fromUserToResponse(*user))
}

func (h *Handler) UpdateUserById(c echo.Context) error {
	var req UserUpdate
	userId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		return domain.NewAppError(err, domain.BadRequestError)
	}

	if err = c.Bind(&req); err != nil {
		return domain.NewAppError(err, domain.BadRequestError)
	}

	if appErr := h.userManager.UpdateUserById(userId,
		fromRequestToUser(req)); appErr != nil {
		return appErr
	}

	return c.JSON(http.StatusOK, fmt.Sprintf("User %v updated", userId))
}

func (h *Handler) DeleteUserById(c echo.Context) error {
	userId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		return domain.NewAppError(err, domain.BadRequestError)
	}

	if appErr := h.userManager.DeleteUserById(userId); appErr != nil {
		return appErr
	}

	return c.JSON(http.StatusOK, fmt.Sprintf("User %v deleted", userId))
}

func New(userManager usermanager.Manager) *Handler {
	return &Handler{userManager}
}

package usermanager

import (
	"github.com/rierarizzo/cafelatte/internal/domain"
	"github.com/rierarizzo/cafelatte/internal/domain/usermanager"
	http2 "github.com/rierarizzo/cafelatte/pkg/utils/http"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	userManager usermanager.Manager
}

func (h *Handler) GetUsers(c *gin.Context) {
	users, appErr := h.userManager.GetUsers()
	if appErr != nil {
		http2.AbortWithError(c, appErr)
		return
	}

	http2.RespondWithJSON(c, http.StatusOK,
		fromUsersToResponse(users))
}

func (h *Handler) FindUserByID(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		appErr := domain.NewAppError(err, domain.BadRequestError)
		http2.AbortWithError(c, appErr)
		return
	}

	user, appErr := h.userManager.FindUserByID(userID)
	if appErr != nil {
		http2.AbortWithError(c, appErr)
		return
	}

	http2.RespondWithJSON(c, http.StatusOK, fromUserToResponse(*user))
}

func (h *Handler) UpdateUser(c *gin.Context) {
	var req UpdateUserRequest
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		appErr := domain.NewAppError(err, domain.BadRequestError)
		http2.AbortWithError(c, appErr)
		return
	}

	if err := c.BindJSON(&req); err != nil {
		appErr := domain.NewAppError(err, domain.BadRequestError)
		http2.AbortWithError(c, appErr)
		return
	}

	appErr := h.userManager.UpdateUser(userID,
		fromRequestToUser(req))
	if appErr != nil {
		http2.AbortWithError(c, appErr)
		return
	}

	http2.RespondWithJSON(c, http.StatusOK, "usermanager updated")
}

func (h *Handler) DeleteUser(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		appErr := domain.NewAppError(err, domain.BadRequestError)
		http2.AbortWithError(c, appErr)
		return
	}

	appErr := h.userManager.DeleteUser(userID)
	if appErr != nil {
		http2.AbortWithError(c, appErr)
		return
	}

	http2.RespondWithJSON(c, http.StatusOK, "usermanager deleted")
}

func NewUserHandler(userManager usermanager.Manager) *Handler {
	return &Handler{userManager}
}

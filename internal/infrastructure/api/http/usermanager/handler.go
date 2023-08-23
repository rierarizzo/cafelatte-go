package usermanager

import (
	"github.com/rierarizzo/cafelatte/internal/domain"
	"github.com/rierarizzo/cafelatte/internal/domain/usermanager"
	httpUtil "github.com/rierarizzo/cafelatte/pkg/utils/http"

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
		httpUtil.AbortWithError(c, appErr)
		return
	}

	httpUtil.RespondWithJSON(c, http.StatusOK, fromUsersToResponse(users))
}

func (h *Handler) FindUserByID(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		httpUtil.AbortWithError(c, domain.NewAppError(err, domain.BadRequestError))
		return
	}

	user, appErr := h.userManager.FindUserByID(userID)
	if appErr != nil {
		httpUtil.AbortWithError(c, appErr)
		return
	}

	httpUtil.RespondWithJSON(c, http.StatusOK, fromUserToResponse(*user))
}

func (h *Handler) UpdateUser(c *gin.Context) {
	var req UpdateUserRequest
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		httpUtil.AbortWithError(c, domain.NewAppError(err, domain.BadRequestError))
		return
	}

	if err = c.BindJSON(&req); err != nil {
		httpUtil.AbortWithError(c, domain.NewAppError(err, domain.BadRequestError))
		return
	}

	appErr := h.userManager.UpdateUser(userID, fromRequestToUser(req))
	if appErr != nil {
		httpUtil.AbortWithError(c, appErr)
		return
	}

	httpUtil.RespondWithJSON(c, http.StatusOK, "usermanager updated")
}

func (h *Handler) DeleteUser(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		httpUtil.AbortWithError(c, domain.NewAppError(err, domain.BadRequestError))
		return
	}

	appErr := h.userManager.DeleteUser(userID)
	if appErr != nil {
		httpUtil.AbortWithError(c, appErr)
		return
	}

	httpUtil.RespondWithJSON(c, http.StatusOK, "usermanager deleted")
}

func NewUserHandler(userManager usermanager.Manager) *Handler {
	return &Handler{userManager}
}

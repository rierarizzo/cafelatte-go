package user

import (
	domain "github.com/rierarizzo/cafelatte/internal/domain/errors"
	userDomain "github.com/rierarizzo/cafelatte/internal/domain/user"
	http2 "github.com/rierarizzo/cafelatte/pkg/utils/http"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	userService userDomain.IUserService
}

func (h *Handler) GetUsers(c *gin.Context) {
	users, appErr := h.userService.GetUsers()
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

	user, appErr := h.userService.FindUserByID(userID)
	if appErr != nil {
		http2.AbortWithError(c, appErr)
		return
	}

	http2.RespondWithJSON(c, http.StatusOK, fromUserToResponse(*user))
}

func (h *Handler) UpdateUser(c *gin.Context) {
	var req UpdateRequest
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

	appErr := h.userService.UpdateUser(userID,
		fromUpdateRequestToUser(req))
	if appErr != nil {
		http2.AbortWithError(c, appErr)
		return
	}

	http2.RespondWithJSON(c, http.StatusOK, "user updated")
}

func (h *Handler) DeleteUser(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		appErr := domain.NewAppError(err, domain.BadRequestError)
		http2.AbortWithError(c, appErr)
		return
	}

	appErr := h.userService.DeleteUser(userID)
	if appErr != nil {
		http2.AbortWithError(c, appErr)
		return
	}

	http2.RespondWithJSON(c, http.StatusOK, "user deleted")
}

func NewUserHandler(userService userDomain.IUserService) *Handler {
	return &Handler{userService}
}

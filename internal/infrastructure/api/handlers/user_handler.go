package handlers

import (
	domain "github.com/rierarizzo/cafelatte/internal/domain/errors"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/api/mappers"
	"github.com/rierarizzo/cafelatte/internal/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rierarizzo/cafelatte/internal/domain/ports"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/api/dto"
)

type UserHandler struct {
	userService ports.IUserService
}

func (h *UserHandler) GetAllUsers(c *gin.Context) {
	users, appErr := h.userService.GetUsers()
	if appErr != nil {
		utils.AbortWithError(c, appErr)
		return
	}

	userResponse := make([]dto.UserResponse, 0)
	for _, k := range users {
		userResponse = append(userResponse, mappers.FromUserToUserRes(k))
	}

	c.JSON(http.StatusOK, userResponse)
}

func (h *UserHandler) FindUserByID(c *gin.Context) {
	userIDParam := c.Param("userID")
	userID, err := strconv.Atoi(userIDParam)
	if err != nil {
		utils.AbortWithError(c, domain.NewAppError(err, domain.BadRequestError))
		return
	}

	user, appErr := h.userService.FindUserByID(userID)
	if appErr != nil {
		utils.AbortWithError(c, appErr)
		return
	}

	c.JSON(http.StatusOK, mappers.FromUserToUserRes(*user))
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	userIDParam := c.Param("userID")
	userID, err := strconv.Atoi(userIDParam)
	if err != nil {
		utils.AbortWithError(c, domain.NewAppError(err, domain.BadRequestError))
		return
	}

	var updUserReq dto.UpdateUserRequest
	err = c.BindJSON(&updUserReq)
	if err != nil {
		utils.AbortWithError(c, domain.NewAppError(err, domain.BadRequestError))
		return
	}

	appErr := h.userService.UpdateUser(userID,
		mappers.FromUpdateUserReqToUser(updUserReq))
	if appErr != nil {
		utils.AbortWithError(c, appErr)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"msg":    "user updated",
	})
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	userIDParam := c.Param("userID")
	userID, err := strconv.Atoi(userIDParam)
	if err != nil {
		utils.AbortWithError(c, domain.NewAppError(err, domain.BadRequestError))
		return
	}

	appErr := h.userService.DeleteUser(userID)
	if appErr != nil {
		utils.AbortWithError(c, appErr)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"msg":    "user deleter",
	})
}

func NewUserHandler(userService ports.IUserService) *UserHandler {
	return &UserHandler{userService}
}

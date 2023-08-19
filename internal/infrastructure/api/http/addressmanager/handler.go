package addressmanager

import (
	"github.com/gin-gonic/gin"
	"github.com/rierarizzo/cafelatte/internal/domain"
	"github.com/rierarizzo/cafelatte/internal/domain/addressmanager"
	httpUtil "github.com/rierarizzo/cafelatte/pkg/utils/http"
	"net/http"
	"strconv"
)

type Handler struct {
	addressService addressmanager.Manager
}

func (h *Handler) GetAddressesByUserID(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		appErr := domain.NewAppError(err, domain.BadRequestError)
		httpUtil.AbortWithError(c, appErr)
		return
	}

	addresses, appErr := h.addressService.GetAddressesByUserID(userID)
	if appErr != nil {
		httpUtil.AbortWithError(c, appErr)
		return
	}

	httpUtil.RespondWithJSON(c, http.StatusOK,
		fromAddressesToResponse(addresses))
}

func (h *Handler) AddUserAddresses(c *gin.Context) {
	var req []RegisterAddressRequest
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		appErr := domain.NewAppError(err, domain.BadRequestError)
		httpUtil.AbortWithError(c, appErr)
		return
	}

	if err := c.BindJSON(&req); err != nil {
		appErr := domain.NewAppError(err, domain.BadRequestError)
		httpUtil.AbortWithError(c, appErr)
		return
	}

	addresses, appErr := h.addressService.AddUserAddresses(userID,
		fromRequestToAddresses(req))
	if appErr != nil {
		httpUtil.AbortWithError(c, appErr)
		return
	}

	httpUtil.RespondWithJSON(c, http.StatusCreated,
		fromAddressesToResponse(addresses))
}

func NewAddressHandler(addressService addressmanager.Manager) *Handler {
	return &Handler{addressService}
}

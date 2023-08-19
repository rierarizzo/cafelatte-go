package productmanager

import (
	"github.com/gin-gonic/gin"
	"github.com/rierarizzo/cafelatte/internal/domain/productmanager"
	"github.com/rierarizzo/cafelatte/pkg/constants/misc"
	"github.com/rierarizzo/cafelatte/pkg/params/request"
	httpUtil "github.com/rierarizzo/cafelatte/pkg/utils/http"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Handler struct {
	productManager productmanager.Manager
}

func (h Handler) GetProducts(c *gin.Context) {
	log := logrus.WithField(misc.RequestIDKey, request.ID())

	var category = c.Query("category")

	if category == "" {
		products, appErr := h.productManager.GetProducts()
		if appErr != nil {
			httpUtil.AbortWithError(c, appErr)
			return
		}

		httpUtil.RespondWithJSON(c, http.StatusOK,
			fromProductsToResponse(products))
	} else {
		log.Debugf("Executing with query: %s", category)

		products, appErr := h.productManager.GetProductsByCategory(category)
		if appErr != nil {
			httpUtil.AbortWithError(c, appErr)
			return
		}

		httpUtil.RespondWithJSON(c, http.StatusOK,
			fromProductsToResponse(products))
	}
}

func (h Handler) GetProductCategories(c *gin.Context) {
	categories, appErr := h.productManager.GetProductCategories()
	if appErr != nil {
		httpUtil.AbortWithError(c, appErr)
		return
	}

	httpUtil.RespondWithJSON(c, http.StatusOK,
		fromProductCategoriesToResponse(categories))
}

func NewProductHandler(productManager productmanager.Manager) *Handler {
	return &Handler{productManager}
}

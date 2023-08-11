package product

import (
	"github.com/gin-gonic/gin"
	"github.com/rierarizzo/cafelatte/internal/domain/product"
	"github.com/rierarizzo/cafelatte/pkg/constants/misc"
	"github.com/rierarizzo/cafelatte/pkg/params/request"
	http2 "github.com/rierarizzo/cafelatte/pkg/utils/http"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Handler struct {
	productService product.IProductService
}

func (h Handler) GetProducts(c *gin.Context) {
	log := logrus.WithField(misc.RequestIDKey, request.ID())

	var category = c.Query("category")

	if category == "" {
		products, appErr := h.productService.GetProducts()
		if appErr != nil {
			http2.AbortWithError(c, appErr)
			return
		}

		http2.RespondWithJSON(c, http.StatusOK,
			fromProductsToResponse(products))
	} else {
		log.Debugf("Executing with query: %s", category)

		products, appErr := h.productService.GetProductsByCategory(category)
		if appErr != nil {
			http2.AbortWithError(c, appErr)
			return
		}

		http2.RespondWithJSON(c, http.StatusOK,
			fromProductsToResponse(products))
	}
}

func (h Handler) GetProductCategories(c *gin.Context) {
	categories, appErr := h.productService.GetProductCategories()
	if appErr != nil {
		http2.AbortWithError(c, appErr)
		return
	}

	http2.RespondWithJSON(c, http.StatusOK,
		fromProductCategoriesToResponse(categories))
}

func NewProductHandler(productService product.IProductService) *Handler {
	return &Handler{productService}
}

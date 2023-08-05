package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/rierarizzo/cafelatte/internal/domain/ports"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/api/mappers"
	"github.com/rierarizzo/cafelatte/internal/utils"
	"net/http"
)

type ProductHandler struct {
	productService ports.IProductService
}

func (h ProductHandler) GetProducts(c *gin.Context) {
	var category = c.Query("category")

	if category == "" {
		products, appErr := h.productService.GetProducts()
		if appErr != nil {
			utils.AbortWithError(c, appErr)
			return
		}

		utils.RespondWithJSON(c, http.StatusOK,
			mappers.FromProductSliceToProductResSlice(products))
	} else {
		products, appErr := h.productService.GetProductsByCategory(category)
		if appErr != nil {
			utils.AbortWithError(c, appErr)
			return
		}

		utils.RespondWithJSON(c, http.StatusOK,
			mappers.FromProductSliceToProductResSlice(products))
	}
}

func (h ProductHandler) GetProductCategories(c *gin.Context) {
	categories, appErr := h.productService.GetProductCategories()
	if appErr != nil {
		utils.AbortWithError(c, appErr)
		return
	}

	utils.RespondWithJSON(c, http.StatusOK,
		mappers.FromProductCategorySliceToProductCategoryResSlice(categories))
}

func NewProductHandler(productService ports.IProductService) *ProductHandler {
	return &ProductHandler{productService}
}

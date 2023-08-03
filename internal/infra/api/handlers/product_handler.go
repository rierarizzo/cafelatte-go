package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/rierarizzo/cafelatte/internal/domain/entities"
	"github.com/rierarizzo/cafelatte/internal/domain/ports"
	"github.com/rierarizzo/cafelatte/internal/infra/api/dto"
	"github.com/rierarizzo/cafelatte/internal/infra/api/mappers"
	"github.com/rierarizzo/cafelatte/internal/utils"
	"net/http"
)

type ProductHandler struct {
	productService ports.IProductService
}

func (h ProductHandler) GetProducts(c *gin.Context) {
	category := c.Query("category")

	var products []entities.Product

	if category == "" {
		retrvProducts, appErr := h.productService.GetProducts()
		if appErr != nil {
			utils.AbortWithError(c, appErr)
			return
		}

		products = retrvProducts
	} else {
		retrvProducts, appErr := h.productService.GetProductsByCategory(category)
		if appErr != nil {
			utils.AbortWithError(c, appErr)
			return
		}

		products = retrvProducts
	}

	var response []dto.ProductResponse
	for _, v := range products {
		response = append(response, mappers.FromProductToProductResponse(v))
	}

	c.JSON(http.StatusOK, response)
}

func (h ProductHandler) GetProductCategories(c *gin.Context) {
	categories, appErr := h.productService.GetProductCategories()
	if appErr != nil {
		utils.AbortWithError(c, appErr)
		return
	}

	var response []dto.ProductCategoryResponse
	for _, v := range categories {
		response = append(response,
			mappers.FromProductCategoryToProductCategoryResponse(v))
	}

	c.JSON(http.StatusOK, response)
}

func NewProductHandler(productService ports.IProductService) *ProductHandler {
	return &ProductHandler{productService}
}

package productmanager

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rierarizzo/cafelatte/internal/domain/productmanager"
	"github.com/rierarizzo/cafelatte/pkg/constants/misc"
	"github.com/rierarizzo/cafelatte/pkg/params/request"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	productManager productmanager.Manager
}

func Router(group *echo.Group) func(productManagerHandler *Handler) {
	return func(h *Handler) {
		group.GET("/find", h.GetProducts)
		group.GET("/find/categories", h.GetProductCategories)
	}
}

func (h Handler) GetProducts(c echo.Context) error {
	log := logrus.WithField(misc.RequestIdKey, request.Id())

	var category = c.QueryParam("category")

	if category == "" {
		log.Debug("Getting all products")

		products, appErr := h.productManager.GetProducts()
		if appErr != nil {
			return appErr
		}

		return c.JSON(http.StatusOK, fromProductsToResponse(products))
	} else {
		log.Debugf("Getting all products with query: %s", category)

		products, appErr := h.productManager.GetProductsByCategory(category)
		if appErr != nil {
			return appErr
		}

		return c.JSON(http.StatusOK, fromProductsToResponse(products))
	}
}

func (h Handler) GetProductCategories(c echo.Context) error {
	categories, appErr := h.productManager.GetProductCategories()
	if appErr != nil {
		return appErr
	}

	return c.JSON(http.StatusOK, fromProductCategoriesToResponse(categories))
}

func New(productManager productmanager.Manager) *Handler {
	return &Handler{productManager}
}

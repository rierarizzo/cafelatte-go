package productmanager

import (
	"github.com/labstack/echo/v4"
	"github.com/rierarizzo/cafelatte/internal/domain/productmanager"
	"github.com/rierarizzo/cafelatte/pkg/constants/misc"
	"github.com/rierarizzo/cafelatte/pkg/params/request"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Handler struct {
	productManager productmanager.Manager
}

func Router(group *echo.Group) func(productManagerHandler *Handler) {
	return func(handler *Handler) {
		group.GET("/find", handler.GetProducts)
		group.GET("/find/categories", handler.GetProductCategories)
	}
}

func (handler Handler) GetProducts(c echo.Context) error {
	log := logrus.WithField(misc.RequestIDKey, request.ID())

	var category = c.QueryParam("category")

	if category == "" {
		products, appErr := handler.productManager.GetProducts()
		if appErr != nil {
			return appErr
		}

		return c.JSON(http.StatusOK, fromProductsToResponse(products))
	} else {
		log.Debugf("Executing with query: %s", category)

		products, appErr := handler.productManager.GetProductsByCategory(category)
		if appErr != nil {
			return appErr
		}

		return c.JSON(http.StatusOK, fromProductsToResponse(products))
	}
}

func (handler Handler) GetProductCategories(c echo.Context) error {
	categories, appErr := handler.productManager.GetProductCategories()
	if appErr != nil {
		return appErr
	}

	return c.JSON(http.StatusOK, fromProductCategoriesToResponse(categories))
}

func New(productManager productmanager.Manager) *Handler {
	return &Handler{productManager}
}

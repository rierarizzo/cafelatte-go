package productmanager

import (
	"github.com/rierarizzo/cafelatte/internal/domain"
	"github.com/rierarizzo/cafelatte/internal/usecases/productmanager"
	httpUtil "github.com/rierarizzo/cafelatte/pkg/utils/http"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func getAllProducts(m productmanager.Manager) echo.HandlerFunc {
	return func(c echo.Context) error {
		var products []domain.Product
		var appErr *domain.AppError

		var category = c.QueryParam("category")
		if category != "" {
			logrus.Debugf("Getting all products with query: %s", category)

			products, appErr = m.GetProductsByCategory(category)
			if appErr != nil {
				return appErr
			}
		} else {
			logrus.Debug("Getting all products")

			products, appErr = m.GetProducts()
			if appErr != nil {
				return appErr
			}
		}

		response := fromProductsToResponse(products)
		return httpUtil.RespondWithJSON(c, http.StatusOK, response)
	}
}

func getProductCategories(m productmanager.Manager) echo.HandlerFunc {
	return func(c echo.Context) error {
		categories, appErr := m.GetProductCategories()
		if appErr != nil {
			return appErr
		}

		response := fromProductCategoriesToResponse(categories)
		return httpUtil.RespondWithJSON(c, http.StatusOK, response)
	}
}

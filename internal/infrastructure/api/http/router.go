package http

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	addressmanagerDomain "github.com/rierarizzo/cafelatte/internal/domain/addressmanager"
	authenticatorDomain "github.com/rierarizzo/cafelatte/internal/domain/authenticator"
	cardmanagerDomain "github.com/rierarizzo/cafelatte/internal/domain/cardmanager"
	productmanagerDomain "github.com/rierarizzo/cafelatte/internal/domain/productmanager"
	productpurchaserDomain "github.com/rierarizzo/cafelatte/internal/domain/productpurchaser"
	usermanagerDomain "github.com/rierarizzo/cafelatte/internal/domain/usermanager"
	addressmanagerHTTP "github.com/rierarizzo/cafelatte/internal/infrastructure/api/http/addressmanager"
	authenticatorHTTP "github.com/rierarizzo/cafelatte/internal/infrastructure/api/http/authenticator"
	cardmanagerHTTP "github.com/rierarizzo/cafelatte/internal/infrastructure/api/http/cardmanager"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/api/http/error"
	productmanagerHTTP "github.com/rierarizzo/cafelatte/internal/infrastructure/api/http/productmanager"

	productpurchaserHTTP "github.com/rierarizzo/cafelatte/internal/infrastructure/api/http/productpurchaser"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/api/http/saverequestid"
	usermanagerHTTP "github.com/rierarizzo/cafelatte/internal/infrastructure/api/http/usermanager"
)

func Router(userManager usermanagerDomain.Manager,
	authenticator authenticatorDomain.Authenticator,
	addressManager addressmanagerDomain.Manager,
	cardManager cardmanagerDomain.Manager,
	productManager productmanagerDomain.Manager,
	purchaser productpurchaserDomain.Purchaser) *echo.Echo {
	e := echo.New()

	/* Middlewares */
	e.HTTPErrorHandler = error.CustomHTTPErrorHandler
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	e.Use(middleware.RequestID())
	e.Use(saverequestid.Middleware)

	/* Groups */
	auth := e.Group("/auth")
	users := e.Group("/users")
	products := e.Group("/products")
	purchase := e.Group("/purchase")

	/* Routers */
	authenticatorHTTP.Router(auth)(authenticatorHTTP.New(authenticator))
	usermanagerHTTP.Router(users)(usermanagerHTTP.New(userManager))
	cardmanagerHTTP.Router(users)(cardmanagerHTTP.New(cardManager))
	addressmanagerHTTP.Router(users)(addressmanagerHTTP.New(addressManager))
	productmanagerHTTP.Router(products)(productmanagerHTTP.New(productManager))
	productpurchaserHTTP.Router(purchase)(productpurchaserHTTP.New(purchaser))

	return e
}

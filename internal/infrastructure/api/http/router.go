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
	errorHTTP "github.com/rierarizzo/cafelatte/internal/infrastructure/api/http/error"
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
	e.HTTPErrorHandler = errorHTTP.CustomHTTPErrorHandler

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}, id=${id}, latency=${latency}\n",
	}))
	e.Use(middleware.CORS())
	e.Use(middleware.RequestID())
	e.Use(saverequestid.Middleware)

	/* Groups */
	auth := e.Group("/auth")
	products := e.Group("/products")
	users := e.Group("/users", authenticatorHTTP.Middleware)
	purchase := e.Group("/purchase", authenticatorHTTP.Middleware)

	/* Routers */
	authenticatorHTTP.Router(auth)(authenticatorHTTP.New(authenticator))
	usermanagerHTTP.Router(users)(usermanagerHTTP.New(userManager))
	cardmanagerHTTP.Router(users)(cardmanagerHTTP.New(cardManager))
	addressmanagerHTTP.Router(users)(addressmanagerHTTP.New(addressManager))
	productmanagerHTTP.Router(products)(productmanagerHTTP.New(productManager))
	productpurchaserHTTP.Router(purchase)(productpurchaserHTTP.New(purchaser))

	return e
}

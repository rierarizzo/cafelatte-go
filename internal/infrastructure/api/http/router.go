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
	addressmanagerHttp "github.com/rierarizzo/cafelatte/internal/infrastructure/api/http/addressmanager"
	authenticatorHttp "github.com/rierarizzo/cafelatte/internal/infrastructure/api/http/authenticator"
	cardmanagerHttp "github.com/rierarizzo/cafelatte/internal/infrastructure/api/http/cardmanager"
	errorHttp "github.com/rierarizzo/cafelatte/internal/infrastructure/api/http/error"
	productmanagerHttp "github.com/rierarizzo/cafelatte/internal/infrastructure/api/http/productmanager"
	productpurchaserHttp "github.com/rierarizzo/cafelatte/internal/infrastructure/api/http/productpurchaser"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/api/http/saverequestid"
	usermanagerHttp "github.com/rierarizzo/cafelatte/internal/infrastructure/api/http/usermanager"
)

func Router(
	userManager usermanagerDomain.Manager,
	authenticator authenticatorDomain.Authenticator,
	addressManager addressmanagerDomain.Manager,
	cardManager cardmanagerDomain.Manager,
	productManager productmanagerDomain.Manager,
	purchaser productpurchaserDomain.Purchaser,
) *echo.Echo {
	e := echo.New()

	/* Middlewares */
	e.HTTPErrorHandler = errorHttp.CustomHttpErrorHandler

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}, id=${id}, latency=${latency}\n",
	}))
	e.Use(middleware.CORS())
	e.Use(middleware.RequestID())
	e.Use(saverequestid.Middleware)

	/* Groups */
	auth := e.Group("/auth")
	products := e.Group("/products")
	users := e.Group("/users", authenticatorHttp.Middleware)
	purchase := e.Group("/purchase", authenticatorHttp.Middleware)

	/* Routers */
	authenticatorHttp.Router(auth)(authenticatorHttp.New(authenticator))
	usermanagerHttp.Router(users)(usermanagerHttp.New(userManager))
	cardmanagerHttp.Router(users)(cardmanagerHttp.New(cardManager))
	addressmanagerHttp.Router(users)(addressmanagerHttp.New(addressManager))
	productmanagerHttp.Router(products)(productmanagerHttp.New(productManager))
	productpurchaserHttp.Router(purchase)(productpurchaserHttp.New(purchaser))

	return e
}

package http

import (
	"github.com/labstack/echo/v4"
	addressmanagerDomain "github.com/rierarizzo/cafelatte/internal/domain/addressmanager"
	authenticatorDomain "github.com/rierarizzo/cafelatte/internal/domain/authenticator"
	cardmanagerDomain "github.com/rierarizzo/cafelatte/internal/domain/cardmanager"
	productmanagerDomain "github.com/rierarizzo/cafelatte/internal/domain/productmanager"
	productpurchaserDomain "github.com/rierarizzo/cafelatte/internal/domain/productpurchaser"
	usermanagerDomain "github.com/rierarizzo/cafelatte/internal/domain/usermanager"
	addressmanagerHttp "github.com/rierarizzo/cafelatte/internal/infrastructure/api/http/addressmanager"
	authenticatorHttp "github.com/rierarizzo/cafelatte/internal/infrastructure/api/http/authenticator"
	cardmanagerHttp "github.com/rierarizzo/cafelatte/internal/infrastructure/api/http/cardmanager"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/api/http/cors"
	errorHttp "github.com/rierarizzo/cafelatte/internal/infrastructure/api/http/error"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/api/http/logger"
	productmanagerHttp "github.com/rierarizzo/cafelatte/internal/infrastructure/api/http/productmanager"
	productpurchaserHttp "github.com/rierarizzo/cafelatte/internal/infrastructure/api/http/productpurchaser"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/api/http/requestid"
	usermanagerHttp "github.com/rierarizzo/cafelatte/internal/infrastructure/api/http/usermanager"
)

func Router(userManager usermanagerDomain.Manager,
	authenticator authenticatorDomain.Authenticator,
	addressManager addressmanagerDomain.Manager,
	cardManager cardmanagerDomain.Manager,
	productManager productmanagerDomain.Manager,
	purchaser productpurchaserDomain.Purchaser) *echo.Echo {
	e := echo.New()

	/* Middlewares */
	e.HTTPErrorHandler = errorHttp.CustomHttpErrorHandler

	e.Use(cors.CustomMiddleware())
	e.Use(logger.CustomMiddleware())
	e.Use(requestid.CustomMiddleware())

	/* Groups */
	auth := e.Group("/auth")
	products := e.Group("/products")
	users := e.Group("/users", authenticatorHttp.Middleware)
	purchase := e.Group("/purchase", authenticatorHttp.Middleware)

	/* Routers */
	authenticatorHttp.ConfigureRouting(auth)(authenticator)
	usermanagerHttp.ConfigureRouting(users)(userManager)
	cardmanagerHttp.ConfigureRouting(users)(cardManager)
	addressmanagerHttp.ConfigureRouting(users)(addressManager)
	productmanagerHttp.ConfigureRouting(products)(productManager)
	productpurchaserHttp.ConfigureRouting(purchase)(purchaser)

	return e
}

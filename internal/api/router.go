package api

import (
	"github.com/labstack/echo/v4"
	addressmanagerHttp "github.com/rierarizzo/cafelatte/internal/api/handlers/addressmanager"
	authenticatorHttp "github.com/rierarizzo/cafelatte/internal/api/handlers/authenticator"
	cardmanagerHttp "github.com/rierarizzo/cafelatte/internal/api/handlers/cardmanager"
	productmanagerHttp "github.com/rierarizzo/cafelatte/internal/api/handlers/productmanager"
	productpurchaserHttp "github.com/rierarizzo/cafelatte/internal/api/handlers/productpurchaser"
	usermanagerHttp "github.com/rierarizzo/cafelatte/internal/api/handlers/usermanager"
	authenticator3 "github.com/rierarizzo/cafelatte/internal/api/middlewares/authenticator"
	"github.com/rierarizzo/cafelatte/internal/api/middlewares/cors"
	errorHttp "github.com/rierarizzo/cafelatte/internal/api/middlewares/error"
	"github.com/rierarizzo/cafelatte/internal/api/middlewares/logger"
	"github.com/rierarizzo/cafelatte/internal/api/middlewares/requestid"
	addressmanagerDomain "github.com/rierarizzo/cafelatte/internal/usecases/addressmanager"
	authenticatorDomain "github.com/rierarizzo/cafelatte/internal/usecases/authenticator"
	cardmanagerDomain "github.com/rierarizzo/cafelatte/internal/usecases/cardmanager"
	productmanagerDomain "github.com/rierarizzo/cafelatte/internal/usecases/productmanager"
	productpurchaserDomain "github.com/rierarizzo/cafelatte/internal/usecases/productpurchaser"
	usermanagerDomain "github.com/rierarizzo/cafelatte/internal/usecases/usermanager"
)

func Router(userManager usermanagerDomain.Manager, authenticator authenticatorDomain.Authenticator,
	addressManager addressmanagerDomain.Manager, cardManager cardmanagerDomain.Manager,
	productManager productmanagerDomain.Manager, purchaser productpurchaserDomain.Purchaser) *echo.Echo {
	e := echo.New()

	/* Middlewares */
	e.HTTPErrorHandler = errorHttp.CustomHttpErrorHandler

	e.Use(requestid.CustomMiddleware())
	e.Use(cors.CustomMiddleware())
	e.Use(logger.CustomMiddleware())

	/* Groups */
	auth := e.Group("/auth")
	products := e.Group("/products")
	users := e.Group("/users", authenticator3.Middleware)
	purchase := e.Group("/purchase", authenticator3.Middleware)

	/* Routers */
	authenticatorHttp.Routes(auth)(authenticator)
	usermanagerHttp.Routes(users)(userManager)
	cardmanagerHttp.Routes(users)(cardManager)
	addressmanagerHttp.Routes(users)(addressManager)
	productmanagerHttp.Routes(products)(productManager)
	productpurchaserHttp.Routes(purchase)(purchaser)

	return e
}

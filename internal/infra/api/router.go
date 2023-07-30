package api

import (
	"github.com/rierarizzo/cafelatte/internal/infra/api/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rierarizzo/cafelatte/internal/infra/api/handlers"
)

func Router(userHandler *handlers.UserHandler,
	addressHandler *handlers.AddressHandler,
	cardHandler *handlers.PaymentCardHandler) http.Handler {

	router := gin.New()

	// Error handling (always has to be the first middleware)
	router.Use(middlewares.ErrorMiddleware())

	router.Use(middlewares.RequestIDMiddleware())
	router.Use(middlewares.LoggingMiddleware())

	authGroup := router.Group("/auth")
	{
		authGroup.POST("/signup", userHandler.SignUp)
		authGroup.POST("/signin", userHandler.SignIn)
	}

	usersGroup := router.Group("/users")
	usersGroup.Use(middlewares.AuthenticateMiddleware())

	{
		usersGroup.GET("/find", userHandler.GetAllUsers)
		usersGroup.GET("/find/:userID", userHandler.FindUserByID)
		usersGroup.POST("/add-addresses", addressHandler.AddUserAddresses)
		usersGroup.POST("/add-paymentcards", cardHandler.AddUserPaymentCards)
	}

	return router
}

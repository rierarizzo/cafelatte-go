package api

import (
	"github.com/gin-contrib/cors"
	"github.com/rierarizzo/cafelatte/internal/infra/api/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rierarizzo/cafelatte/internal/infra/api/handlers"
)

func Router(userHandler *handlers.UserHandler,
	authHandler *handlers.AuthenticateHandler,
	addressHandler *handlers.AddressHandler,
	cardHandler *handlers.PaymentCardHandler,
	productHandler *handlers.ProductHandler) http.Handler {

	router := gin.New()

	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowCredentials: true,
	}))
	router.Use(middlewares.RequestIDMiddleware())
	router.Use(middlewares.LoggingMiddleware())
	router.Use(middlewares.ErrorMiddleware())

	authGroup := router.Group("/auth")
	{
		authGroup.POST("/signup", authHandler.SignUp)
		authGroup.POST("/signin", authHandler.SignIn)
	}

	usersGroup := router.Group("/users")
	usersGroup.Use(middlewares.AuthenticateMiddleware())
	{
		usersGroup.GET("/find", userHandler.GetAllUsers)
		usersGroup.GET("/find/:userID", userHandler.FindUserByID)
		usersGroup.POST("/add-addresses", addressHandler.AddUserAddresses)
		usersGroup.POST("/add-paymentcards", cardHandler.AddUserPaymentCards)
	}

	productsGroup := router.Group("/products")
	productsGroup.Use(middlewares.AuthenticateMiddleware())
	{
		productsGroup.GET("/find", productHandler.GetProducts)
		productsGroup.GET("/find/categories",
			productHandler.GetProductCategories)
	}

	return router
}

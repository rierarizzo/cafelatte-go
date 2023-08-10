package api

import (
	"github.com/gin-contrib/cors"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/api/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/api/handlers"
)

func Router(userHandler *handlers.UserHandler,
	authHandler *handlers.AuthenticateHandler,
	addressHandler *handlers.AddressHandler,
	cardHandler *handlers.PaymentCardHandler,
	productHandler *handlers.ProductHandler,
	purchaseHandler *handlers.PurchaseHandler) http.Handler {

	router := gin.New()

	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowCredentials: true,
		AllowHeaders:     []string{"Origin, X-Requested-With, Content-Type, Accept, Authorization"},
		AllowMethods:     []string{"POST, GET, OPTIONS, PUT, DELETE, UPDATE"},
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
		usersGroup.GET("/find", userHandler.GetUsers)
		usersGroup.GET("/find/:userID", userHandler.FindUserByID)

		usersGroup.POST("/add-addresses/:userID",
			addressHandler.AddUserAddresses)
		usersGroup.GET("/find-addresses/:userID",
			addressHandler.GetAddressesByUserID)

		usersGroup.POST("/add-cards/:userID", cardHandler.AddUserCards)
		usersGroup.GET("/find-cards/:userID", cardHandler.GetCardsByUserID)

		usersGroup.PUT("/update/:userID", userHandler.UpdateUser)
		usersGroup.DELETE("/delete/:userID", userHandler.DeleteUser)
	}

	productsGroup := router.Group("/products")
	{
		productsGroup.GET("/find", productHandler.GetProducts)
		productsGroup.GET("/find/categories",
			productHandler.GetProductCategories)
	}

	purchaseGroup := router.Group("/purchase")
	purchaseGroup.Use(middlewares.AuthenticateMiddleware())
	{
		purchaseGroup.POST("/", purchaseHandler.Purchase)
	}

	return router
}

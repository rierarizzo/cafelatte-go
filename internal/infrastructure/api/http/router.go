package http

import (
	"github.com/gin-contrib/cors"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/api/http/addressmanager"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/api/http/authenticator"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/api/http/cardmanager"
	error2 "github.com/rierarizzo/cafelatte/internal/infrastructure/api/http/error"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/api/http/logger"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/api/http/productmanager"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/api/http/productpurchaser"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/api/http/requestid"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/api/http/usermanager"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Router(userHandler *usermanager.Handler,
	authHandler *authenticator.Handler,
	addressHandler *addressmanager.Handler,
	cardHandler *cardmanager.Handler,
	productHandler *productmanager.Handler,
	purchaseHandler *productpurchaser.Handler) http.Handler {

	router := gin.New()

	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowCredentials: true,
		AllowHeaders:     []string{"Origin, X-Requested-With, Content-Type, Accept, Authorization"},
		AllowMethods:     []string{"POST, GET, OPTIONS, PUT, DELETE, UPDATE"},
	}))
	router.Use(requestid.Middleware())
	router.Use(logger.Middleware())
	router.Use(error2.Middleware())

	authGroup := router.Group("/auth")
	{
		authGroup.POST("/signup", authHandler.SignUp)
		authGroup.POST("/signin", authHandler.SignIn)
	}

	usersGroup := router.Group("/users")
	usersGroup.Use(authenticator.Middleware())
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

	purchaseGroup := router.Group("/productpurchaser")
	purchaseGroup.Use(authenticator.Middleware())
	{
		purchaseGroup.POST("/", purchaseHandler.Purchase)
	}

	return router
}

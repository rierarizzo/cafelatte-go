package http

import (
	"github.com/gin-contrib/cors"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/api/http/address"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/api/http/authenticate"
	error2 "github.com/rierarizzo/cafelatte/internal/infrastructure/api/http/error"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/api/http/logging"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/api/http/paymentcard"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/api/http/product"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/api/http/purchase"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/api/http/requestid"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/api/http/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Router(userHandler *user.Handler,
	authHandler *authenticate.Handler,
	addressHandler *address.Handler,
	cardHandler *paymentcard.Handler,
	productHandler *product.Handler,
	purchaseHandler *purchase.Handler) http.Handler {

	router := gin.New()

	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowCredentials: true,
		AllowHeaders:     []string{"Origin, X-Requested-With, Content-Type, Accept, Authorization"},
		AllowMethods:     []string{"POST, GET, OPTIONS, PUT, DELETE, UPDATE"},
	}))
	router.Use(requestid.Middleware())
	router.Use(logging.Middleware())
	router.Use(error2.Middleware())

	authGroup := router.Group("/auth")
	{
		authGroup.POST("/signup", authHandler.SignUp)
		authGroup.POST("/signin", authHandler.SignIn)
	}

	usersGroup := router.Group("/users")
	usersGroup.Use(authenticate.Middleware())
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
	purchaseGroup.Use(authenticate.Middleware())
	{
		purchaseGroup.POST("/", purchaseHandler.Purchase)
	}

	return router
}

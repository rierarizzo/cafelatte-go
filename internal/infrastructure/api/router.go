package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/api/handlers"
)

func Router(userHandler *handlers.UserHandler) http.Handler {
	router := gin.Default()

	authGroup := router.Group("/auth")
	{
		authGroup.POST("/signup", userHandler.SignUp)
		authGroup.POST("/signin", userHandler.SignIn)
	}

	usersGroup := router.Group("/users")
	{
		usersGroup.GET("/find", userHandler.GetAllUsers)
		usersGroup.GET("/find/:userID", userHandler.FindUser)
	}

	return router
}

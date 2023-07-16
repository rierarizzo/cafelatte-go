package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/api/handlers"
)

func Router(userHandler *handlers.UserHandler) http.Handler {
	r := gin.Default()

	authGroup := r.Group("/auth")
	{
		authGroup.POST("/signup", userHandler.SignUp)
		authGroup.POST("/signin", userHandler.SignIn)
	}

	usersGroup := r.Group("/users")
	{
		usersGroup.GET("/find/:userID", userHandler.FindUser)
	}

	return r
}

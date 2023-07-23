package cmd

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rierarizzo/cafelatte/internal/core/services"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/api"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/api/handlers"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/data"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/data/repositories"
)

func Server() {
	// Map config environment variable to struct
	config := LoadConfig()
	// Gin mode
	gin.SetMode(config.GinMode)
	// Connect to database
	db := data.Connect(config.DSN)

	// Addresses instance
	addressRepo := repositories.NewAddressRepository(db)
	addressService := services.NewAddressService(addressRepo)
	addressHandler := handlers.NewAddressHandler(addressService)

	// PaymentCards instance
	paymentCardRepo := repositories.NewPaymentCardRepository(db)
	paymentCardService := services.NewPaymentCardService(paymentCardRepo)
	paymentCardHandler := handlers.NewPaymentCardHandler(paymentCardService)

	// Users instance
	userRepo := repositories.NewUserRepository(db, addressRepo, paymentCardRepo)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	// Initialize router with all paths
	router := api.Router(userHandler, addressHandler, paymentCardHandler)

	logrus.Infof("Listening server on port %s", config.ServerPort)
	err := http.ListenAndServe(fmt.Sprintf(":%s", config.ServerPort), router)
	if err != nil {
		logrus.Panic(err)
	}
}

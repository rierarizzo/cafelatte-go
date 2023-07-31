package cmd

import (
	"fmt"
	config2 "github.com/rierarizzo/cafelatte/cmd/config"
	"github.com/sirupsen/logrus"
	"net/http"

	"github.com/rierarizzo/cafelatte/internal/domain/services"
	"github.com/rierarizzo/cafelatte/internal/infra/api"
	"github.com/rierarizzo/cafelatte/internal/infra/api/handlers"
	"github.com/rierarizzo/cafelatte/internal/infra/data"
	"github.com/rierarizzo/cafelatte/internal/infra/data/repos"
)

func Server() {
	// Map config environment variable to struct
	config := config2.GetConfig()
	config2.LoadInitConfig(config)

	// Connect to database
	db := data.Connect(config.DSN)

	// Users instance
	userRepo := repos.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	// Authentication instance
	authService := services.NewAuthService(userRepo)
	authHandler := handlers.NewAuthHandler(authService)

	// Addresses instance
	addressRepo := repos.NewAddressRepository(db)
	addressService := services.NewAddressService(addressRepo)
	addressHandler := handlers.NewAddressHandler(addressService)

	// PaymentCards instance
	paymentCardRepo := repos.NewPaymentCardRepository(db)
	paymentCardService := services.NewPaymentCardService(paymentCardRepo)
	paymentCardHandler := handlers.NewPaymentCardHandler(paymentCardService)

	// Initialize router with all paths
	router := api.Router(userHandler, authHandler, addressHandler,
		paymentCardHandler)

	logrus.WithField("port", config.ServerPort).Info("Starting server")
	if err := http.ListenAndServe(fmt.Sprintf(":%s", config.ServerPort),
		router); err != nil {
		logrus.Panic(err)
	}
}

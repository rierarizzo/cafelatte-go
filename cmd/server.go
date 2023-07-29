package cmd

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/rierarizzo/cafelatte/internal/domain/services"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/api"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/api/handlers"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/data"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/data/repositories"
)

func Server() {
	// Map config environment variable to struct
	config := GetConfig()
	LoadInitConfig(config)

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
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	// Initialize router with all paths
	router := api.Router(userHandler, addressHandler, paymentCardHandler)

	slog.Info("Starting server", "port", config.ServerPort)
	if err := http.ListenAndServe(
		fmt.Sprintf(":%s", config.ServerPort),
		router,
	); err != nil {
		slog.Error(err.Error())
		panic(err)
	}
}

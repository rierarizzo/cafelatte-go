package cmd

import (
	"fmt"
	"github.com/rierarizzo/cafelatte/cmd/config"
	"github.com/rierarizzo/cafelatte/internal/domain/usecases"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"

	"github.com/rierarizzo/cafelatte/internal/domain/services"
	"github.com/rierarizzo/cafelatte/internal/infra/api"
	"github.com/rierarizzo/cafelatte/internal/infra/api/handlers"
	"github.com/rierarizzo/cafelatte/internal/infra/data"
	"github.com/rierarizzo/cafelatte/internal/infra/data/repos"
)

func Server() {
	start := time.Now()

	// Map config environment variable to struct
	cf := config.GetConfig()
	config.LoadInitConfig(cf)

	// Connect to database
	db := data.Connect(cf.DSN)

	// Users instance
	userRepo := repos.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	// Authentication instance
	authUsecase := usecases.NewAuthenticateUsecase(userService)
	authHandler := handlers.NewAuthHandler(authUsecase)

	// Addresses instance
	addressRepo := repos.NewAddressRepository(db)
	addressService := services.NewAddressService(addressRepo)
	addressHandler := handlers.NewAddressHandler(addressService)

	// PaymentCards instance
	paymentCardRepo := repos.NewPaymentCardRepository(db)
	paymentCardService := services.NewPaymentCardService(paymentCardRepo)
	paymentCardHandler := handlers.NewPaymentCardHandler(paymentCardService)

	// Products instance
	productRepo := repos.NewProductRepo(db)
	productService := services.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productService)

	// Initialize router with all paths
	router := api.Router(userHandler, authHandler, addressHandler,
		paymentCardHandler, productHandler)

	elapsed := time.Since(start).Seconds()

	logrus.WithFields(logrus.Fields{
		"port":      cf.ServerPort,
		"startedAt": fmt.Sprintf("%.7fs", elapsed),
	}).Info("Starting server")

	if err := http.ListenAndServe(fmt.Sprintf(":%s", cf.ServerPort),
		router); err != nil {
		logrus.Panic(err)
	}
}

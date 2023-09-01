package main

import (
	"fmt"
	httpRouter "github.com/rierarizzo/cafelatte/internal/api"
	"github.com/rierarizzo/cafelatte/internal/config"
	"github.com/rierarizzo/cafelatte/internal/data"
	addressRepo "github.com/rierarizzo/cafelatte/internal/data/address"
	cardRepo "github.com/rierarizzo/cafelatte/internal/data/card"
	orderRepo "github.com/rierarizzo/cafelatte/internal/data/order"
	productRepo "github.com/rierarizzo/cafelatte/internal/data/product"
	userRepo "github.com/rierarizzo/cafelatte/internal/data/user"
	userfilesRepo "github.com/rierarizzo/cafelatte/internal/data/userfiles"
	"github.com/rierarizzo/cafelatte/internal/usecases/addressmanager"
	"github.com/rierarizzo/cafelatte/internal/usecases/authenticator"
	"github.com/rierarizzo/cafelatte/internal/usecases/cardmanager"
	"github.com/rierarizzo/cafelatte/internal/usecases/productmanager"
	"github.com/rierarizzo/cafelatte/internal/usecases/productpurchaser"
	"github.com/rierarizzo/cafelatte/internal/usecases/usermanager"
)

func main() {
	Server()
}

func Server() {
	// Map config environment variable to struct
	cf := config.GetConfig()
	config.LoadInitConfig(cf)

	// Connect to database
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", cf.DBUser,
		cf.DBPassword, cf.DBHost, cf.DBPort, cf.DBName)
	dbConn := data.ConnectToMySQL(dsn)

	// Get S3 client
	s3Client := data.ConnectToS3("us-east-1")

	// Repositories
	userRepository := userRepo.New(dbConn)
	addressRepository := addressRepo.New(dbConn)
	cardRepository := cardRepo.New(dbConn)
	productRepository := productRepo.New(dbConn)
	orderRepository := orderRepo.New(dbConn)
	userFilesRepo := userfilesRepo.New(s3Client)

	// Usecases
	defaultUserManager := usermanager.New(userRepository, userFilesRepo)
	defaultAuthenticator := authenticator.New(userRepository)
	defaultAddressManager := addressmanager.New(addressRepository)
	defaultCardManager := cardmanager.New(cardRepository)
	defaultProductManager := productmanager.New(productRepository)
	defaultPurchaser := productpurchaser.New(orderRepository)

	// Initialize router with all paths
	router := httpRouter.Router(defaultUserManager, defaultAuthenticator,
		defaultAddressManager, defaultCardManager, defaultProductManager,
		defaultPurchaser)

	router.Logger.Fatal(router.Start(fmt.Sprintf(":%s", cf.ServerPort)))
}

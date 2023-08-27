package server

import (
	"fmt"
	"github.com/rierarizzo/cafelatte/internal/domain/addressmanager"
	"github.com/rierarizzo/cafelatte/internal/domain/authenticator"
	"github.com/rierarizzo/cafelatte/internal/domain/cardmanager"
	"github.com/rierarizzo/cafelatte/internal/domain/productmanager"
	"github.com/rierarizzo/cafelatte/internal/domain/productpurchaser"
	"github.com/rierarizzo/cafelatte/internal/domain/usermanager"
	httpRouter "github.com/rierarizzo/cafelatte/internal/infrastructure/api/http"
	mysqlInfra "github.com/rierarizzo/cafelatte/internal/infrastructure/data/mysql"
	addressRepo "github.com/rierarizzo/cafelatte/internal/infrastructure/data/mysql/address"
	cardRepo "github.com/rierarizzo/cafelatte/internal/infrastructure/data/mysql/card"
	orderRepo "github.com/rierarizzo/cafelatte/internal/infrastructure/data/mysql/order"
	productRepo "github.com/rierarizzo/cafelatte/internal/infrastructure/data/mysql/product"
	userRepo "github.com/rierarizzo/cafelatte/internal/infrastructure/data/mysql/user"
	s3Infra "github.com/rierarizzo/cafelatte/internal/infrastructure/storage/s3"
	userfilesRepo "github.com/rierarizzo/cafelatte/internal/infrastructure/storage/s3/userfiles"
)

func Server() {
	// Map config environment variable to struct
	cf := GetConfig()
	LoadInitConfig(cf)

	// Connect to database
	dbConn := mysqlInfra.Connect(cf)

	// Get S3 client
	s3Client := s3Infra.Connect("us-east-1")

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

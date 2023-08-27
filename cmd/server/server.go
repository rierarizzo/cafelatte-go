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
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", cf.DBUser,
		cf.DBPassword, cf.DBHost, cf.DBPort, cf.DBName)
	db := mysqlInfra.Connect(dsn)

	// Get S3 client
	s3Client := s3Infra.Connect("us-east-1")

	// Users instance
	userRepository := userRepo.NewUserRepository(db)
	userFilesRepo := userfilesRepo.NewUserFilesRepository(s3Client)
	defaultUserManager := usermanager.New(userRepository,
		userFilesRepo)

	// Authentication instance
	defaultAuthenticator := authenticator.New(userRepository)

	// Addresses instance
	addressRepository := addressRepo.NewAddressRepository(db)
	defaultAddressManager := addressmanager.New(addressRepository)

	// PaymentCards instance
	cardRepository := cardRepo.NewPaymentCardRepository(db)
	defaultCardManager := cardmanager.New(cardRepository)

	// Products instance
	productRepository := productRepo.NewProductRepository(db)
	defaultProductManager := productmanager.New(productRepository)

	// Purchase instance
	orderRepository := orderRepo.NewOrderRepository(db)
	defaultPurchaser := productpurchaser.New(orderRepository)

	// Initialize router with all paths
	router := httpRouter.Router(defaultUserManager, defaultAuthenticator,
		defaultAddressManager, defaultCardManager, defaultProductManager,
		defaultPurchaser)

	router.Logger.Fatal(router.Start(fmt.Sprintf(":%s", cf.ServerPort)))
}

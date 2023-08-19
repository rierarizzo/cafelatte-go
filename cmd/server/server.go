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
	addressHTTP "github.com/rierarizzo/cafelatte/internal/infrastructure/api/http/addressmanager"
	authenticateHTTP "github.com/rierarizzo/cafelatte/internal/infrastructure/api/http/authenticator"
	paymentcardHTTP "github.com/rierarizzo/cafelatte/internal/infrastructure/api/http/cardmanager"
	productHTTP "github.com/rierarizzo/cafelatte/internal/infrastructure/api/http/productmanager"
	purchaseHTTP "github.com/rierarizzo/cafelatte/internal/infrastructure/api/http/productpurchaser"
	userHTTP "github.com/rierarizzo/cafelatte/internal/infrastructure/api/http/usermanager"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/data/mysql"
	addressData "github.com/rierarizzo/cafelatte/internal/infrastructure/data/mysql/address"
	orderData "github.com/rierarizzo/cafelatte/internal/infrastructure/data/mysql/order"
	paymentcardData "github.com/rierarizzo/cafelatte/internal/infrastructure/data/mysql/paymentcard"
	productData "github.com/rierarizzo/cafelatte/internal/infrastructure/data/mysql/product"
	userData "github.com/rierarizzo/cafelatte/internal/infrastructure/data/mysql/user"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/storage/s3"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/storage/s3/userfiles"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func Server() {
	start := time.Now()

	// Map config environment variable to struct
	cf := GetConfig()
	LoadInitConfig(cf)

	// Connect to database
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", cf.DBUser,
		cf.DBPassword, cf.DBHost, cf.DBPort, cf.DBName)
	db := mysql.Connect(dsn)

	// Get S3 client
	s3Client := s3.Connect("us-east-1")

	// Users instance
	userRepository := userData.NewUserRepository(db)
	userFilesRepo := userfiles.NewUserFilesRepository(s3Client)
	defaultUserManager := usermanager.NewDefaultManager(userRepository,
		userFilesRepo)
	userHandler := userHTTP.NewUserHandler(defaultUserManager)

	// Authentication instance
	defaultAuthenticator := authenticator.NewDefaultAuthenticator(userRepository)
	authHandler := authenticateHTTP.NewAuthHandler(defaultAuthenticator)

	// Addresses instance
	addressRepository := addressData.NewAddressRepository(db)
	defaultAddressManager := addressmanager.NewDefaultManager(addressRepository)
	addressHandler := addressHTTP.NewAddressHandler(defaultAddressManager)

	// PaymentCards instance
	cardRepository := paymentcardData.NewPaymentCardRepository(db)
	defaultCardManager := cardmanager.NewDefaultManager(cardRepository)
	paymentCardHandler := paymentcardHTTP.NewPaymentCardHandler(defaultCardManager)

	// Products instance
	productRepo := productData.NewProductRepository(db)
	defaultProductManager := productmanager.NewDefaultManager(productRepo)
	productHandler := productHTTP.NewProductHandler(defaultProductManager)

	// Purchase instance
	orderRepo := orderData.NewOrderRepository(db)
	defaultPurchaser := productpurchaser.NewDefaultPurchaser(orderRepo)
	purchaserHandler := purchaseHTTP.NewPurchaseHandler(defaultPurchaser)

	// Initialize router with all paths
	router := httpRouter.Router(userHandler, authHandler, addressHandler,
		paymentCardHandler, productHandler, purchaserHandler)

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

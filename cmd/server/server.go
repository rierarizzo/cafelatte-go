package server

import (
	"fmt"
	"github.com/rierarizzo/cafelatte/internal/domain/address"
	"github.com/rierarizzo/cafelatte/internal/domain/authenticate"
	"github.com/rierarizzo/cafelatte/internal/domain/order"
	"github.com/rierarizzo/cafelatte/internal/domain/paymentcard"
	"github.com/rierarizzo/cafelatte/internal/domain/product"
	"github.com/rierarizzo/cafelatte/internal/domain/purchase"
	"github.com/rierarizzo/cafelatte/internal/domain/user"
	httpRouter "github.com/rierarizzo/cafelatte/internal/infrastructure/api/http"
	addressHTTP "github.com/rierarizzo/cafelatte/internal/infrastructure/api/http/address"
	authenticateHTTP "github.com/rierarizzo/cafelatte/internal/infrastructure/api/http/authenticate"
	paymentcardHTTP "github.com/rierarizzo/cafelatte/internal/infrastructure/api/http/paymentcard"
	productHTTP "github.com/rierarizzo/cafelatte/internal/infrastructure/api/http/product"
	purchaseHTTP "github.com/rierarizzo/cafelatte/internal/infrastructure/api/http/purchase"
	userHTTP "github.com/rierarizzo/cafelatte/internal/infrastructure/api/http/user"
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
	userRepo := userData.NewUserRepository(db)
	userFilesRepo := userfiles.NewUserFilesRepository(s3Client)
	userService := user.NewUserService(userRepo, userFilesRepo)
	userHandler := userHTTP.NewUserHandler(userService)

	// Authentication instance
	authUsecase := authenticate.NewAuthenticateUsecase(userService)
	authHandler := authenticateHTTP.NewAuthHandler(authUsecase)

	// Addresses instance
	addressRepo := addressData.NewAddressRepository(db)
	addressService := address.NewAddressService(addressRepo)
	addressHandler := addressHTTP.NewAddressHandler(addressService)

	// PaymentCards instance
	paymentCardRepo := paymentcardData.NewPaymentCardRepository(db)
	paymentCardService := paymentcard.NewPaymentCardService(paymentCardRepo)
	paymentCardHandler := paymentcardHTTP.NewPaymentCardHandler(paymentCardService)

	// Products instance
	productRepo := productData.NewProductRepository(db)
	productService := product.NewProductService(productRepo)
	productHandler := productHTTP.NewProductHandler(productService)

	// Purchase instance
	orderRepo := orderData.NewOrderRepository(db)
	orderService := order.NewOrderService(orderRepo)
	purchaseUsecase := purchase.NewPurchaseUsecase(orderService)
	purchaseHandler := purchaseHTTP.NewPurchaseHandler(purchaseUsecase)

	// Initialize router with all paths
	router := httpRouter.Router(userHandler, authHandler, addressHandler,
		paymentCardHandler, productHandler, purchaseHandler)

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

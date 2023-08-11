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
	http2 "github.com/rierarizzo/cafelatte/internal/infrastructure/api/http"
	address3 "github.com/rierarizzo/cafelatte/internal/infrastructure/api/http/address"
	authenticate2 "github.com/rierarizzo/cafelatte/internal/infrastructure/api/http/authenticate"
	paymentcard3 "github.com/rierarizzo/cafelatte/internal/infrastructure/api/http/paymentcard"
	product3 "github.com/rierarizzo/cafelatte/internal/infrastructure/api/http/product"
	purchase2 "github.com/rierarizzo/cafelatte/internal/infrastructure/api/http/purchase"
	user3 "github.com/rierarizzo/cafelatte/internal/infrastructure/api/http/user"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/data/mysql"
	address2 "github.com/rierarizzo/cafelatte/internal/infrastructure/data/mysql/address"
	order2 "github.com/rierarizzo/cafelatte/internal/infrastructure/data/mysql/order"
	paymentcard2 "github.com/rierarizzo/cafelatte/internal/infrastructure/data/mysql/paymentcard"
	product2 "github.com/rierarizzo/cafelatte/internal/infrastructure/data/mysql/product"
	user2 "github.com/rierarizzo/cafelatte/internal/infrastructure/data/mysql/user"
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

	// Users instance
	userRepo := user2.NewUserRepository(db)
	userService := user.NewUserService(userRepo)
	userHandler := user3.NewUserHandler(userService)

	// Authentication instance
	authUsecase := authenticate.NewAuthenticateUsecase(userService)
	authHandler := authenticate2.NewAuthHandler(authUsecase)

	// Addresses instance
	addressRepo := address2.NewAddressRepository(db)
	addressService := address.NewAddressService(addressRepo)
	addressHandler := address3.NewAddressHandler(addressService)

	// PaymentCards instance
	paymentCardRepo := paymentcard2.NewPaymentCardRepository(db)
	paymentCardService := paymentcard.NewPaymentCardService(paymentCardRepo)
	paymentCardHandler := paymentcard3.NewPaymentCardHandler(paymentCardService)

	// Products instance
	productRepo := product2.NewProductRepository(db)
	productService := product.NewProductService(productRepo)
	productHandler := product3.NewProductHandler(productService)

	// Purchase instance
	orderRepo := order2.NewOrderRepository(db)
	orderService := order.NewOrderService(orderRepo)
	purchaseUsecase := purchase.NewPurchaseUsecase(orderService)
	purchaseHandler := purchase2.NewPurchaseHandler(purchaseUsecase)

	// Initialize router with all paths
	router := http2.Router(userHandler, authHandler, addressHandler,
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

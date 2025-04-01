package main

import (
	"go-ps-adv-homework/configs"
	"go-ps-adv-homework/internal/auth"
	"go-ps-adv-homework/internal/carts"
	"go-ps-adv-homework/internal/orders"
	"go-ps-adv-homework/internal/products"
	"go-ps-adv-homework/internal/sessions"
	"go-ps-adv-homework/internal/smsru"
	"go-ps-adv-homework/internal/users"
	"go-ps-adv-homework/pkg/db"
	"go-ps-adv-homework/pkg/middleware"
	"log"
	"net/http"
)

func App() http.Handler {
	config := configs.LoadConfig()
	database := db.Connect(config)
	// Repositories
	productsRepository := products.NewProductsRepository(database)
	userRepository := users.NewUserRepository(database)
	sessionsRepository := sessions.NewSessionRepository(database)
	ordersRepository := orders.NewOrdersRepository(database)
	// Services
	smsService := smsru.NewSmsRuService(config.Sms.ApiId)
	authService := auth.NewAuthService(auth.AuthServiceDependencies{
		Config:            config,
		UserRepository:    userRepository,
		SessionRepository: sessionsRepository,
		SmsService:        smsService,
	})
	ordersService := orders.NewOrdersService(orders.OrdersServiceDependencies{
		UserRepository:   userRepository,
		OrdersRepository: ordersRepository,
	})
	// Handlers
	router := http.NewServeMux()
	auth.NewHandler(router, auth.AuthHandlerDependencies{
		AuthService: authService,
	})
	products.NewProductsHandler(router, products.ProductsHandlerDependencies{
		Config:             config,
		ProductsRepository: productsRepository,
	})
	carts.NewCartHandler(router, carts.CartHandlerDependencies{Config: config})
	orders.NewOrdersHandler(router, orders.OrdersHandlerDependencies{
		Config:        config,
		OrdersService: ordersService,
	})
	// Middleware
	stack := middleware.Chain(
		middleware.CORS,
		middleware.Logger,
	)
	return stack(router)
}

func main() {
	app := App()
	server := http.Server{
		Addr:    ":8080",
		Handler: app,
	}
	log.Printf("Server is listening on localhost:8080")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
}

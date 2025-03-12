package main

import (
	"go-ps-adv-homework/configs"
	"go-ps-adv-homework/internal/auth"
	"go-ps-adv-homework/internal/carts"
	"go-ps-adv-homework/internal/products"
	"go-ps-adv-homework/pkg/db"
	"go-ps-adv-homework/pkg/middleware"
	"log"
	"net/http"
)

func main() {
	config := configs.LoadConfig()
	database := db.Connect(config)

	// Repositories
	productsRepository := products.NewProductsRepository(database)

	// Handlers
	router := http.NewServeMux()
	auth.NewHandler(router, auth.AuthHandlerDependencies{Config: config})
	products.NewProductsHandler(router, products.ProductsHandlerDependencies{
		Config:             config,
		ProductsRepository: productsRepository,
	})
	carts.NewHandler(router, carts.CartHandlerDependencies{Config: config})

	// Middleware
	stack := middleware.Chain(
		middleware.CORS,
		middleware.Logger,
	)
	server := http.Server{
		Addr:    config.Server.Host + ":" + config.Server.Port,
		Handler: stack(router),
	}

	log.Printf("Server is listening on %s:%s", config.Server.Host, config.Server.Port)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
}

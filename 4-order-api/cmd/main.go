package main

import (
	"go-ps-adv-homework/configs"
	"go-ps-adv-homework/internal/auth"
	"go-ps-adv-homework/internal/carts"
	"go-ps-adv-homework/internal/products"
	"go-ps-adv-homework/pkg/db"
	"log"
	"net/http"
)

func main() {
	config := configs.LoadConfig()
	_ = db.Connect(config)

	router := http.NewServeMux()
	auth.NewHandler(router, auth.HandlerDependencies{Config: config})
	products.NewHandler(router, products.HandlerDependencies{Config: config})
	carts.NewHandler(router, carts.HandlerDependencies{Config: config})

	server := http.Server{
		Addr:    config.Server.Host + ":" + config.Server.Port,
		Handler: router,
	}
	log.Printf("Server is listening on %s:%s", config.Server.Host, config.Server.Port)
	defer log.Println("Server stopped")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
}

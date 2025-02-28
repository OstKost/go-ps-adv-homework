package carts

import (
	"fmt"
	"go-ps-adv-homework/configs"
	"net/http"
)

type handler struct {
	Config *configs.Config
}

type HandlerDependencies struct {
	*configs.Config
}

func NewHandler(router *http.ServeMux, dependencies HandlerDependencies) {
	handler := &handler{
		Config: dependencies.Config,
	}
	router.HandleFunc("POST /carts", handler.createCart())
	router.HandleFunc("GET /carts/{cartId}", handler.getById())
	router.HandleFunc("GET /carts/user/{userId}", handler.getByUserId())
}

func (handler *handler) createCart() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("createCart")
	}
}

func (handler *handler) getById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("getById")
	}
}

func (handler *handler) getByUserId() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("getByUserId")
	}
}

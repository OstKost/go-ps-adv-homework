package carts

import (
	"fmt"
	"go-ps-adv-homework/configs"
	"go-ps-adv-homework/pkg/middleware"
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
	router.Handle("POST /carts", middleware.IsAuthed(handler.createCart()))
	router.Handle("GET /carts/{cartId}", middleware.IsAuthed(handler.getById()))
	router.Handle("GET /carts/user/{userId}", middleware.IsAuthed(handler.getByUserId()))
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

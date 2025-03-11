package carts

import (
	"fmt"
	"go-ps-adv-homework/configs"
	"go-ps-adv-homework/pkg/middleware"
	"net/http"
)

type cartHandler struct {
	Config *configs.Config
}

type CartHandlerDependencies struct {
	*configs.Config
}

func NewHandler(router *http.ServeMux, dependencies CartHandlerDependencies) {
	handler := &cartHandler{
		Config: dependencies.Config,
	}
	router.Handle("POST /carts", middleware.IsAuthed(handler.createCart(), dependencies.Config))
	router.Handle("GET /carts/{cartId}", middleware.IsAuthed(handler.getById(), dependencies.Config))
	router.Handle("GET /carts/user/{userId}", middleware.IsAuthed(handler.getByUserId(), dependencies.Config))
}

func (handler *cartHandler) createCart() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("createCart")
	}
}

func (handler *cartHandler) getById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("getById")
	}
}

func (handler *cartHandler) getByUserId() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("getByUserId")
	}
}

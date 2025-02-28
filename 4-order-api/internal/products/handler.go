package products

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
	router.HandleFunc("POST /products", handler.createProduct())
	router.HandleFunc("GET /products/{productId}", handler.getById())
	router.HandleFunc("GET /products", handler.getList())
	router.HandleFunc("PATCH /products/{productId}", handler.updateProduct())
	router.HandleFunc("DELETE /products/{productId}", handler.deleteProduct())
}

func (handler *handler) createProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("createProduct")
	}
}

func (handler *handler) getById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("getById")
	}
}

func (handler *handler) getList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("getList")
	}
}

func (handler *handler) updateProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("updateProduct")
	}
}

func (handler *handler) deleteProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("deleteProduct")
	}
}

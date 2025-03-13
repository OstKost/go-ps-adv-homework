package products

import (
	"fmt"
	"go-ps-adv-homework/configs"
	"go-ps-adv-homework/pkg/middleware"
	"go-ps-adv-homework/pkg/request"
	"go-ps-adv-homework/pkg/response"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type productsHandler struct {
	*configs.Config
	*ProductsRepository
}

type ProductsHandlerDependencies struct {
	*configs.Config
	*ProductsRepository
}

func NewProductsHandler(router *http.ServeMux, dependencies ProductsHandlerDependencies) {
	handler := &productsHandler{
		Config:             dependencies.Config,
		ProductsRepository: dependencies.ProductsRepository,
	}
	router.Handle("POST /products", middleware.IsAuthed(handler.createProduct(), dependencies.Config))
	router.Handle("PATCH /products/{productId}", middleware.IsAuthed(handler.updateProduct(), dependencies.Config))
	router.Handle("DELETE /products/{productId}", middleware.IsAuthed(handler.deleteProduct(), dependencies.Config))
	router.HandleFunc("GET /products", handler.findProducts())
	router.HandleFunc("GET /products/{productId}", handler.getProductById())
}

func (handler *productsHandler) createProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Context
		ctx := r.Context()
		phone, _ := ctx.Value(middleware.ContextPhoneKey).(string)
		session, _ := ctx.Value(middleware.ContextSessionKey).(string)
		fmt.Println("ContextPhoneKey: ", phone)
		fmt.Println("ContextSessionKey: ", session)
		// Validate
		body, err := request.HandleBody[CreateProductRequest](&w, r)
		if err != nil {
			response.Json(w, err.Error(), http.StatusBadRequest)
			return
		}
		// Prepare data
		product := &Product{
			Name:        body.Name,
			Description: body.Description,
			Images:      body.Images,
			Price:       body.Price,
		}
		// Action
		createdProduct, err := handler.ProductsRepository.Create(product)
		if err != nil {
			response.Json(w, err.Error(), http.StatusInternalServerError)
			return
		}
		response.Json(w, createdProduct, http.StatusCreated)
	}
}

func (handler *productsHandler) updateProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Context
		ctx := r.Context()
		phone, _ := ctx.Value(middleware.ContextPhoneKey).(string)
		session, _ := ctx.Value(middleware.ContextSessionKey).(string)
		fmt.Println("ContextPhoneKey: ", phone)
		fmt.Println("ContextSessionKey: ", session)
		// Validate id
		idString := r.PathValue("productId")
		if idString == "" {
			response.Json(w, "product ID is required", http.StatusBadRequest)
			return
		}
		id, err := strconv.ParseUint(idString, 10, 32)
		if err != nil {
			response.Json(w, err.Error(), http.StatusBadRequest)
			return
		}
		// Validate body
		body, err := request.HandleBody[UpdateProductRequest](&w, r)
		if err != nil {
			response.Json(w, err.Error(), http.StatusBadRequest)
			return
		}
		// Action
		product, err := handler.ProductsRepository.Update(&Product{
			Model:       gorm.Model{ID: uint(id)},
			Name:        body.Name,
			Description: body.Description,
			Images:      body.Images,
			Price:       body.Price,
		})
		response.Json(w, product, http.StatusOK)
	}
}

func (handler *productsHandler) deleteProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Validate id
		idString := r.PathValue("productId")
		if idString == "" {
			response.Json(w, "product ID is required", http.StatusBadRequest)
			return
		}
		id, err := strconv.ParseUint(idString, 10, 32)
		if err != nil {
			response.Json(w, err.Error(), http.StatusBadRequest)
			return
		}
		// Check if exists
		_, err = handler.ProductsRepository.GetById(uint(id))
		if err != nil {
			response.Json(w, err.Error(), http.StatusNotFound)
			return
		}
		// Delete
		err = handler.ProductsRepository.Delete(uint(id))
		if err != nil {
			response.Json(w, err.Error(), http.StatusInternalServerError)
			return
		}
		response.Json(w, id, http.StatusOK)
	}
}

func (handler *productsHandler) findProducts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		limit, offset := request.GetPaginationParams(query)
		products, err := handler.ProductsRepository.Find(query.Get("name"), limit, offset)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		total, err := handler.ProductsRepository.Count(query.Get("name"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		responseData := response.PreparePaginatedResponse[Product](*products, total, offset, limit)
		response.Json(w, responseData, http.StatusOK)
	}
}

func (handler *productsHandler) getProductById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Validate id
		idString := r.PathValue("productId")
		if idString == "" {
			response.Json(w, "product ID is required", http.StatusBadRequest)
			return
		}
		id, err := strconv.ParseUint(idString, 10, 32)
		if err != nil {
			response.Json(w, err.Error(), http.StatusBadRequest)
			return
		}
		// Action
		product, err := handler.ProductsRepository.GetById(uint(id))
		if err != nil {
			response.Json(w, err.Error(), http.StatusNotFound)
			return
		}
		response.Json(w, product, http.StatusOK)
	}
}

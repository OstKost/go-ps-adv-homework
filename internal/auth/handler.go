package auth

import (
	"fmt"
	"go-ps-adv-homework/configs"
	"go-ps-adv-homework/pkg/request"
	"go-ps-adv-homework/pkg/response"
	"net/http"
)

type handler struct {
	Config *configs.Config
}

type HandlerDependencies struct {
	*configs.Config
}

func NewAuthHandler(router *http.ServeMux, dependencies HandlerDependencies) {
	handler := &handler{
		Config: dependencies.Config,
	}
	router.HandleFunc("POST /auth/register", handler.Register())
	router.HandleFunc("POST /auth/login", handler.Login())
}

func (handler *handler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		body, err := request.HandleBody[RegisterRequest](&w, req)
		if err != nil {
			response.Json(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Println(body)
		res := LoginResponse{Token: handler.Config.Auth.Secret}
		response.Json(w, res, http.StatusCreated)
	}
}

func (handler *handler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := request.HandleBody[LoginRequest](&w, r)
		if err != nil {
			response.Json(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Println(body)
		res := LoginResponse{Token: handler.Config.Auth.Secret}
		response.Json(w, res, http.StatusCreated)
	}
}

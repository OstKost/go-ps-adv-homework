package auth

import (
	"fmt"
	"go-ps-adv-homework/pkg/request"
	"go-ps-adv-homework/pkg/response"
	"net/http"
)

type authHandler struct{}

func NewAuthHandler(router *http.ServeMux) {
	handler := &authHandler{}
	router.HandleFunc("POST /auth/register", handler.Register())
	router.HandleFunc("POST /auth/login", handler.Login())
}

func (handler *authHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		body, err := request.HandleBody[RegisterRequest](&w, req)
		if err != nil {
			return
		}
		fmt.Println(body) // пока не используем
		res := LoginResponse{Token: "1234567890"}
		response.Json(w, res, http.StatusOK)
	}
}

func (handler *authHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := request.HandleBody[LoginRequest](&w, r)
		if err != nil {
			return
		}
		fmt.Println(body)
		res := LoginResponse{Token: "1234567890"}
		response.Json(w, res, http.StatusCreated)
	}
}

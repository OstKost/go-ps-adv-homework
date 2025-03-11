package auth

import (
	"go-ps-adv-homework/configs"
	"go-ps-adv-homework/pkg/jwt"
	"go-ps-adv-homework/pkg/request"
	"go-ps-adv-homework/pkg/response"
	"net/http"
)

type authHandler struct {
	*configs.Config
	*AuthService
}

type AuthHandlerDependencies struct {
	*configs.Config
	*AuthService
}

func NewAuthHandler(router *http.ServeMux, dependencies AuthHandlerDependencies) {
	handler := &authHandler{
		Config:      dependencies.Config,
		AuthService: dependencies.AuthService,
	}
	router.HandleFunc("POST /auth/register", handler.Register())
	router.HandleFunc("POST /auth/login", handler.Login())
}

func (handler *authHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		body, err := request.HandleBody[RegisterRequest](&w, req)
		if err != nil {
			return
		}
		email, err := handler.AuthService.Register(body.Email, body.Password, body.Name)
		if err != nil {
			response.Json(w, err.Error(), http.StatusInternalServerError)
			return
		}
		token, err := jwt.NewJWT(handler.Config.Auth.Secret).SignToken(jwt.JWTData{Email: email})
		if err != nil {
			response.Json(w, err.Error(), http.StatusInternalServerError)
			return
		}
		res := RegisterResponse{Token: token}
		response.Json(w, res, http.StatusCreated)
	}
}

func (handler *authHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := request.HandleBody[LoginRequest](&w, r)
		if err != nil {
			return
		}
		email, err := handler.AuthService.Login(body.Email, body.Password)
		if err != nil {
			response.Json(w, err.Error(), http.StatusUnauthorized)
			return
		}
		token, err := jwt.NewJWT(handler.Config.Auth.Secret).SignToken(jwt.JWTData{Email: email})
		if err != nil {
			response.Json(w, err.Error(), http.StatusInternalServerError)
			return
		}
		res := LoginResponse{Token: token}
		response.Json(w, res, http.StatusOK)
	}
}

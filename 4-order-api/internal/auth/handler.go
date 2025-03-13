package auth

import (
	"go-ps-adv-homework/pkg/request"
	"go-ps-adv-homework/pkg/response"
	"net/http"
)

type authHandler struct {
	*AuthService
}

type AuthHandlerDependencies struct {
	*AuthService
}

func NewHandler(router *http.ServeMux, dependencies AuthHandlerDependencies) {
	handler := &authHandler{
		AuthService: dependencies.AuthService,
	}
	router.HandleFunc("POST /auth/authByPhone", handler.authByPhone())
	router.HandleFunc("POST /auth/authByCall", handler.authByCall())
	router.HandleFunc("POST /auth/verifyCode", handler.verifyCode())
}

func (handler *authHandler) authByPhone() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Validate body
		body, err := request.HandleBody[SendSmsRequest](&w, r)
		if err != nil {
			response.Json(w, err.Error(), http.StatusBadRequest)
			return
		}
		// Check user, create session, send code
		session, err := handler.AuthService.SendCode(body.Phone, false)
		if err != nil {
			response.Json(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Response
		res := SendSmsResponse{Message: "OK", Session: session.Session, Code: session.Code}
		response.Json(w, res, http.StatusCreated)
	}
}

func (handler *authHandler) verifyCode() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Validate body
		body, err := request.HandleBody[LoginRequest](&w, r)
		if err != nil {
			response.Json(w, err.Error(), http.StatusBadRequest)
			return
		}
		// Check session and code
		token, err, code := handler.AuthService.CheckCode(body.Session, body.Code)
		if err != nil {
			response.Json(w, err.Error(), code)
			return
		}
		// Response
		res := LoginResponse{Message: "OK", Token: token}
		response.Json(w, res, code)
	}
}

func (handler *authHandler) authByCall() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Validate body
		body, err := request.HandleBody[SendSmsRequest](&w, r)
		if err != nil {
			response.Json(w, err.Error(), http.StatusBadRequest)
			return
		}
		// Check user, create session, send code
		session, err := handler.AuthService.SendCode(body.Phone, true)
		if err != nil {
			response.Json(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Response
		res := SendSmsResponse{Message: "OK", Session: session.Session, Code: session.Code}
		response.Json(w, res, http.StatusCreated)
	}
}

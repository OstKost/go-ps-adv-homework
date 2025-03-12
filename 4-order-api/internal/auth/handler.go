package auth

import (
	"fmt"
	"go-ps-adv-homework/configs"
	"go-ps-adv-homework/pkg/request"
	"go-ps-adv-homework/pkg/response"
	"go-ps-adv-homework/pkg/smsru"
	"math/rand"
	"net/http"
)

type authHandler struct {
	Config *configs.Config
}

type AuthHandlerDependencies struct {
	*configs.Config
}

var codes []string

func NewHandler(router *http.ServeMux, dependencies AuthHandlerDependencies) {
	handler := &authHandler{
		Config: dependencies.Config,
	}
	router.HandleFunc("POST /auth/authByPhone", handler.authByPhone())
	router.HandleFunc("POST /auth/authByCall", handler.authByCall())
	router.HandleFunc("POST /auth/verifyCode", handler.verifyCode())
}

func (handler *authHandler) authByPhone() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := request.HandleBody[SendSmsRequest](&w, r)
		if err != nil {
			response.Json(w, err.Error(), http.StatusBadRequest)
			return
		}
		apiId := handler.Config.Sms.ApiId
		code := generateCode(6)
		codes = append(codes, code)
		smsStatus, err := smsru.Send(apiId, body.Phone, code)
		if err != nil {
			response.Json(w, err.Error(), http.StatusInternalServerError)
			return
		}
		res := SendSmsResponse{Message: smsStatus}
		response.Json(w, res, http.StatusCreated)
	}
}

func (handler *authHandler) verifyCode() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := request.HandleBody[LoginRequest](&w, r)
		if err != nil {
			response.Json(w, err.Error(), http.StatusBadRequest)
			return
		}

		found := false
		for i, h := range codes {
			if h == body.Code {
				found = true
				codes = append(codes[:i], codes[i+1:]...)
			}
		}

		var message string
		var statusCode int
		switch found {
		case true:
			message = "Code Verified"
			statusCode = http.StatusOK
		case false:
			message = "Wrong Code"
			statusCode = http.StatusNotFound
		}
		token := generateToken(16)
		res := LoginResponse{Message: message, Token: token}
		response.Json(w, res, statusCode)
	}
}

func (handler *authHandler) authByCall() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := request.HandleBody[SendSmsRequest](&w, r)
		if err != nil {
			response.Json(w, err.Error(), http.StatusBadRequest)
			return
		}
		apiId := handler.Config.Sms.ApiId
		code, err := smsru.Call(apiId, body.Phone)
		if err != nil {
			response.Json(w, err.Error(), http.StatusInternalServerError)
			return
		}
		codes = append(codes, code)
		fmt.Println(code)
		res := SendSmsResponse{Message: "success", Code: code}
		response.Json(w, res, http.StatusCreated)
	}
}

func generateToken(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	return generator(n, letterRunes)
}

func generateCode(n int) string {
	var letterRunes = []rune("1234567890")
	return generator(n, letterRunes)
}

func generator(n int, letterRunes []rune) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

package verify

import (
	"fmt"
	"go-ps-adv-homework/configs"
	"go-ps-adv-homework/pkg/hashes"
	"go-ps-adv-homework/pkg/mail"
	"go-ps-adv-homework/pkg/request"
	"go-ps-adv-homework/pkg/response"
	"net/http"
)

type verifyHandler struct {
	Config *configs.Config
	Hashes *hashes.Hashes
}

type VerifyHandlerDependencies struct {
	*configs.Config
	*hashes.Hashes
}

func NewVerifyHandler(router *http.ServeMux, dependencies VerifyHandlerDependencies) {
	handler := &verifyHandler{
		Config: dependencies.Config,
		Hashes: dependencies.Hashes,
	}
	router.HandleFunc("POST /verify/send", handler.SendVerifyEmail())
	router.HandleFunc("GET /verify/{hash}", handler.CheckHash())
}

func (handler *verifyHandler) SendVerifyEmail() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		// Проверяем body
		body, err := request.HandleBody[SendEmailRequest](&w, req)
		if err != nil {
			return
		}
		// Создаем хэш
		hash := handler.Hashes.Add()
		// Настройки почты
		config := handler.Config
		subject := "Подтверждение почты"
		emailName := "Verify Email"
		from := fmt.Sprintf("%s <%s>", emailName, config.Email.Address)
		text := fmt.Sprintf(`Привет, подтверди почту!
http://%s:%s/verify/%s`, config.Server.Host, config.Server.Port, hash)
		// Создаем новый email
		err = mail.SendMail(config.Email, from, body.Email, subject, text)
		if err != nil {
			fmt.Println(err)
			response.Json(w, "Ошибка отправки письма", 500)
		}
		res := SendEmailResponse{
			Hash: hash,
		}
		response.Json(w, res, 201)
	}
}

func (handler *verifyHandler) CheckHash() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		hash := req.PathValue("hash")
		hashFound := handler.Hashes.Exists(hash)

		var (
			message    string
			success    bool
			statusCode int
		)
		switch hashFound {
		case true:
			message = "Hash Found"
			success = true
			statusCode = http.StatusOK
			handler.Hashes.Remove(hash)
		case false:
			message = "Hash Not Found"
			success = false
			statusCode = http.StatusNotFound
		}

		res := CheckHashResponse{
			Message: message,
			Success: success,
		}
		response.Json(w, res, statusCode)
	}
}

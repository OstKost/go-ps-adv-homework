package verify

import (
	"fmt"
	"github.com/google/uuid"
	"go-ps-adv-homework/configs"
	"go-ps-adv-homework/pkg/mail"
	"go-ps-adv-homework/pkg/request"
	"go-ps-adv-homework/pkg/response"
	"net/http"
)

var hashes []string // временное хранилище

type handler struct {
	Config *configs.Config
}

type HandlerDependencies struct {
	*configs.Config
}

func NewVerifyHandler(router *http.ServeMux, dependencies HandlerDependencies) {
	handler := &handler{
		Config: dependencies.Config,
	}
	router.HandleFunc("POST /verify/send", handler.SendVerifyEmail())
	router.HandleFunc("GET /verify/{hash}", handler.CheckHash())
}

func (handler *handler) SendVerifyEmail() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		// Проверяем body
		body, err := request.HandleBody[SendEmailRequest](&w, req)
		// Создаем хэш
		hash := uuid.New()
		hashes = append(hashes, hash.String())
		// Настройки почты
		config := handler.Config
		subject := "Подтверждение почты"
		emailName := "Verify Email"
		from := fmt.Sprintf("%s <%s>", emailName, config.Email.Address)
		text := fmt.Sprintf(`Привет, подтверди почту!
http://%s:%s/verify/%s`, config.Server.Host, config.Server.Port, hash.String())
		// Создаем новый email
		err = mail.SendMail(config.Email, from, body.Email, subject, text)
		if err != nil {
			fmt.Println(err)
			response.Json(w, "Ошибка отправки письма", 500)
		}
		res := SendEmailResponse{
			Hash: hash.String(),
		}
		response.Json(w, res, 201)
	}
}

func (handler *handler) CheckHash() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		hash := req.PathValue("hash")

		found := false
		for i, h := range hashes {
			if h == hash {
				found = true
				hashes = append(hashes[:i], hashes[i+1:]...)
			}
		}

		var message string
		var success bool
		var statusCode int
		switch found {
		case true:
			message = "Hash Found"
			success = true
			statusCode = http.StatusOK
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

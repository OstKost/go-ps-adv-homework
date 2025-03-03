package request

import (
	"go-ps-adv-homework/pkg/response"
	"net/http"
)

func HandleBody[T any](writer *http.ResponseWriter, request *http.Request) (*T, error) {
	body, err := Decode[T](request.Body)
	if err != nil {
		response.Json(*writer, err.Error(), http.StatusBadRequest)
		return nil, err
	}

	err = isValid(body)
	if err != nil {
		response.Json(*writer, err.Error(), http.StatusBadRequest)
		return nil, err
	}

	return &body, nil
}

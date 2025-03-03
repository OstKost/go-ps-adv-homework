package response

import (
	"encoding/json"
	"log"
	"net/http"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func Json(writer http.ResponseWriter, data interface{}, statusCode int) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(statusCode)
	err := json.NewEncoder(writer).Encode(data)
	if err != nil {
		log.Fatalln(err)
	}
}

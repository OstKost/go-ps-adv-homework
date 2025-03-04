package response

import (
	"encoding/json"
	"log"
	"net/http"
)

type PaginatedResponse[T any] struct {
	Data       []T   `json:"data"`
	Total      int64 `json:"total"`
	Page       int   `json:"page"`
	PageSize   int   `json:"pageSize"`
	TotalPages int   `json:"totalPages"`
}

func Json(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Fatalln(err)
	}
}

func PreparePaginatedResponse[T any](data []T, total int64, offset int, limit int) PaginatedResponse[T] {
	page := offset/limit + 1

	totalPages := int(total) / limit
	if int(total)%limit != 0 {
		totalPages++
	}

	return PaginatedResponse[T]{
		Data:       data,
		Total:      total,
		Page:       page,
		PageSize:   limit,
		TotalPages: totalPages,
	}
}

package request

import (
	"encoding/json"
	"io"
	"net/url"
	"strconv"
)

func DecodeBody[T any](body io.ReadCloser) (T, error) {
	var payload T
	err := json.NewDecoder(body).Decode(&payload)
	if err != nil {
		return payload, err
	}
	return payload, nil
}

func GetPaginationParams(query url.Values) (limit, offset int) {
	// Получаем параметры пагинации
	page, err := strconv.Atoi(query.Get("page"))
	if err != nil || page < 1 {
		page = 1 // Default page value
	}

	pageSize, err := strconv.Atoi(query.Get("pageSize"))
	if err != nil || pageSize < 1 {
		pageSize = 10 // Default page size
	}

	// Вычисляем limit и offset
	limit = pageSize
	offset = (page - 1) * pageSize

	return limit, offset
}

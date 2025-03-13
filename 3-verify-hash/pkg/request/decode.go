package request

import (
	"encoding/json"
	"io"
	"net/url"
	"strconv"
)

func Decode[T any](body io.ReadCloser) (T, error) {
	var payload T
	err := json.NewDecoder(body).Decode(&payload)
	if err != nil {
		return payload, err
	}
	return payload, nil
}

func GetPaginationParams(query url.Values) (limit, offset int) {
	offset, err := strconv.Atoi(query.Get("offset"))
	if err != nil || offset < 0 {
		offset = 0
	}
	limit, err = strconv.Atoi(query.Get("limit"))
	if err != nil || limit < 1 {
		limit = 10
	}
	return limit, offset
}

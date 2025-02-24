package main

import (
	"log"
	"math/rand/v2"
	"net/http"
	"strconv"
	"strings"
)

type RandomHandler struct {
}

func NewRandomHandler(router *http.ServeMux) {
	handler := &RandomHandler{}
	router.HandleFunc("/random/", handler.randomMax())
	router.HandleFunc("/random", handler.random())
}

func (handler RandomHandler) random() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
                 rand.Seed(time.Now().UnixNano())
		const maxRandom = 6
		number := rand.IntN(maxRandom) + 1 // min 1
		strNumber := strconv.Itoa(number)
		_, err := writer.Write([]byte(strNumber))
		if err != nil {
			log.Println(err)
		}

	}
}

func (handler RandomHandler) randomMax() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		path := strings.TrimPrefix(request.URL.Path, "/random/")

		parts := strings.Split(path, "/")
		if len(parts) == 0 {
			http.Error(writer, "Invalid path", http.StatusBadRequest)
			return
		}

		// Преобразуем строку в int
		maxRandom, err := strconv.Atoi(parts[0])
		if err != nil {
			http.Error(writer, "Invalid max random number", http.StatusBadRequest)
			return
		}

		number := rand.IntN(maxRandom) + 1 // min 1
		strNumber := strconv.Itoa(number)
		_, err = writer.Write([]byte(strNumber))
		if err != nil {
			log.Println(err)
		}
	}
}

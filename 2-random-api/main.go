package main

import (
	"log"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	NewRandomHandler(router)

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	log.Println("Server is listening on port 8081")
	err := server.ListenAndServe()
	if err != nil {
		log.Println(err)
	}
}

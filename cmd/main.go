package main

import (
	"fmt"
	"go-ps-adv-homework/configs"
	"go-ps-adv-homework/internal/auth"
	"go-ps-adv-homework/internal/verify"
	"go-ps-adv-homework/pkg/hashes"
	"log"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	h := hashes.New()

	router := http.NewServeMux()
	auth.NewAuthHandler(router)
	verify.NewVerifyHandler(router, verify.VerifyHandlerDependencies{
		Config: conf,
		Hashes: h,
	})

	server := http.Server{
		Addr:    fmt.Sprintf("%s:%s", conf.Server.Host, conf.Server.Port),
		Handler: router,
	}

	log.Printf("Server is listening on %s:%s", conf.Server.Host, conf.Server.Port)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
}

package main

import (
	"fmt"
	"go-ps-adv-homework/configs"
	"go-ps-adv-homework/internal/auth"
	"go-ps-adv-homework/internal/verify"
	"go-ps-adv-homework/pkg/db"
	"log"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	_ = db.NewDB(conf)

	router := http.NewServeMux()
	auth.NewAuthHandler(router, auth.HandlerDependencies{
		Config: conf,
	})
	verify.NewVerifyHandler(router, verify.HandlerDependencies{
		Config: conf,
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

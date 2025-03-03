package main

import (
	"fmt"
	"go-ps-adv-homework/configs"
	"go-ps-adv-homework/internal/auth"
	"go-ps-adv-homework/internal/link"
	"go-ps-adv-homework/internal/verify"
	"go-ps-adv-homework/pkg/db"
	"log"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	database := db.NewDB(conf)

	// Repos
	linkRepository := link.NewLinkRepository(database)

	//Handlers
	router := http.NewServeMux()
	auth.NewAuthHandler(router, auth.HandlerDependencies{Config: conf})
	verify.NewVerifyHandler(router, verify.HandlerDependencies{Config: conf})
	link.NewLinkHandler(router, link.HandlerDependencies{Repository: linkRepository})

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

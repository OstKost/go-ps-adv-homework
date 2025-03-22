package main

import (
	"go-ps-adv-homework/configs"
	"go-ps-adv-homework/internal/auth"
	"go-ps-adv-homework/internal/link"
	"go-ps-adv-homework/internal/stat"
	"go-ps-adv-homework/internal/user"
	"go-ps-adv-homework/internal/verify"
	"go-ps-adv-homework/pkg/db"
	"go-ps-adv-homework/pkg/event"
	"go-ps-adv-homework/pkg/middleware"
	"log"
	"net/http"
)

func App() http.Handler {
	conf := configs.LoadConfig()
	database := db.NewDB(conf)
	eventBus := event.NewEventBus()
	// Repos
	linkRepository := link.NewLinkRepository(database)
	userRepository := user.NewUserRepository(database)
	statRepository := stat.NewStatRepository(database)
	// Services
	authService := auth.NewAuthService(userRepository)
	statService := stat.NewStatService(stat.StatServiceDependencies{
		EventBus:       eventBus,
		StatRepository: statRepository,
	})
	// Handlers
	router := http.NewServeMux()
	auth.NewAuthHandler(router, auth.AuthHandlerDependencies{Config: conf, AuthService: authService})
	verify.NewVerifyHandler(router, verify.VerifyHandlerDependencies{Config: conf})
	link.NewLinkHandler(router, link.HandlerDependencies{
		Config:         conf,
		LinkRepository: linkRepository,
		EventBus:       eventBus,
	})
	stat.NewStatHandler(router, stat.StatHandlerDependencies{
		StatRepository: statRepository,
		Config:         conf,
	})
	// Middlewares
	stack := middleware.Chain(
		middleware.Logger,
		middleware.CORS,
	)

	go statService.AddClick()

	return stack(router)
}

func main() {
	app := App()
	server := http.Server{
		Addr:    ":8080",
		Handler: app,
	}

	log.Printf("Server is listening on localhost:8080")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
}

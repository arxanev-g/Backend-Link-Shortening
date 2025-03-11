package main

import (
	"fmt"
	"net/http"

	"github.com/arxanev/adv/config"
	"github.com/arxanev/adv/internal/auth"
	"github.com/arxanev/adv/internal/link"
	"github.com/arxanev/adv/internal/stat"
	"github.com/arxanev/adv/internal/user"
	"github.com/arxanev/adv/middleware"
	"github.com/arxanev/adv/pkg/db"
	"github.com/arxanev/adv/pkg/event"
)

func App() http.Handler {
	conf := config.LoadConfig()
	db := db.NewDb(conf)
	router := http.NewServeMux()
	eventBus := event.NewEventBus()

	linkRepository := link.NewLinkRepository(db)
	userRepository := user.NewUserRepository(db)
	statRepository := stat.NewStatRepository(db)

	authService := auth.NewAuthService(userRepository)
	statService := stat.NewStatService(&stat.StatServiceDeps{
		EventBus:       eventBus,
		StatRepository: statRepository,
	})

	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config:      conf,
		AuthService: authService,
	})
	link.NewLinkHandler(router, link.LinkHandlerDeps{
		LinkRepository: linkRepository,
		Config:         conf,
		EventBus:       eventBus,
	})
	stat.NewStatHandler(router, stat.StatHandlerDeps{
		StatRepository: statRepository,
		Config:         conf,
	})

	go statService.AddClick()

	stack := middleware.Chain(
		middleware.CORS,
		middleware.Logging,
	)
	return stack(router)
}

func main() {
	app := App()
	server := http.Server{
		Addr:    ":8082",
		Handler: app,
	}
	fmt.Println("8081")
	server.ListenAndServe()
}

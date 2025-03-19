package app

import (
	"log"
	"net/http"

	"github.com/faruqii/msvc/authservice/internal/config"
	"github.com/faruqii/msvc/authservice/internal/handlers"
	"github.com/faruqii/msvc/authservice/internal/repository"
	"github.com/faruqii/msvc/authservice/internal/routes"
	"github.com/faruqii/msvc/authservice/internal/service"
)

func StartApplication() {
	// run db

	db, err := config.Connect()
	if err != nil {
		panic(err)
	}

	// repository
	userRepo := repository.NewUserRepository(db)
	// service
	authSvc := service.NewAuthService(userRepo)

	// handler
	authHandler := handlers.NewAuthHandler(authSvc)

	mux := http.NewServeMux()
	routes.SetupRoutes(mux, authHandler)

	log.Println("Starting server at :3000")
	if err := http.ListenAndServe(":3000", mux); err != nil {
		panic(err)
	}

}

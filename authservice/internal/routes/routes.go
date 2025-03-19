package routes

import (
	"net/http"

	"github.com/faruqii/msvc/authservice/internal/handlers"
)

func SetupRoutes(mux *http.ServeMux, authHandler *handlers.AuthHandler) {
	mux.HandleFunc("/api/v1/auth/register", authHandler.Register)
	mux.HandleFunc("/api/v1/auth/login", authHandler.Login)

	// You can add middleware here if needed
}

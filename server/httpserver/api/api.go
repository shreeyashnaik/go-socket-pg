package api

import (
	"chat-app/httpserver/controller"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rs/cors"
)

func InitRoutes() http.Handler {

	// Init handlers
	chatHandler := controller.InitChatHandler()

	// Create a new Mux instance
	multiplexerChi := chi.NewMux()

	// Register the handlers from controller to the routes
	multiplexerChi.Route("/api/v1", func(r chi.Router) {

		r.Post("/chats", chatHandler.CreateChatRoom)
		r.Post("/chats/{chatID}/users", chatHandler.AddUser)
	})

	multiplexer := SetupCORS().Handler(multiplexerChi)

	// Return the ServeMux instance
	return multiplexer
}

// Set up CORS options
func SetupCORS() *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"}, // Only allow frontend origin
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodOptions},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	})
}

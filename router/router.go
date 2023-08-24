package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"net/http"
	"url-shortener/handler"
)

func Routes() http.Handler {
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "DELETE"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	appHandler := handler.NewAppHandler()

	// App routes
	r.Post("/app", appHandler.Post)
	r.Get("/{shortUrl}", appHandler.Get)
	r.Delete("/{shortUrl}", appHandler.Delete)
	r.Get("/health", appHandler.HealthCheck)

	return r
}

package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"githib.com/ralvescosta/go-simple-http-server/internal/services"
	"githib.com/ralvescosta/go-simple-http-server/pkg/controllers"
	"githib.com/ralvescosta/go-simple-http-server/pkg/routes"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	server := &http.Server{
		Addr:         "localhost",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: time.Second,
		IdleTimeout:  time.Second,
		Handler:      r,
	}

	authorizationService := services.NewAuthorizationService()
	authorizationController := controllers.NewAuthorizationController(authorizationService)
	routes.RegisterAuthorizationRoutes(r, authorizationController)

	server.ListenAndServe()
}

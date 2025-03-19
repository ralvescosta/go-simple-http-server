package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/sirupsen/logrus"

	"githib.com/ralvescosta/go-simple-http-server/internal/services"
	"githib.com/ralvescosta/go-simple-http-server/pkg/configs"
	"githib.com/ralvescosta/go-simple-http-server/pkg/controllers"
	"githib.com/ralvescosta/go-simple-http-server/pkg/routes"
)

func main() {
	cfgs, err := configs.NewConfigs()
	if err != nil {
		logrus.Fatal(err)
	}

	logrus.Info("Go HTTP Simple Server")

	logrus.Info("creating router...")
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	addr := fmt.Sprintf("%s:%v", cfgs.Host, cfgs.Port)
	server := &http.Server{
		Addr:         addr,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: time.Second,
		IdleTimeout:  time.Second,
		Handler:      r,
	}

	logrus.Info("instantiating services, controllers and routers...")

	authorizationService := services.NewAuthorizationService()
	authorizationController := controllers.NewAuthorizationController(authorizationService)
	routes.RegisterAuthorizationRoutes(r, authorizationController)

	logrus.Infof("Stating HTTP server: %s", addr)
	server.ListenAndServe()
}

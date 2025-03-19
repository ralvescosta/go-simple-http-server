package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/sirupsen/logrus"

	"githib.com/ralvescosta/go-simple-http-server/internal/services"
	"githib.com/ralvescosta/go-simple-http-server/pkg/configs"
	"githib.com/ralvescosta/go-simple-http-server/pkg/controllers/financial"
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
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  10 * time.Second,
		Handler:      r,
	}

	logrus.Info("instantiating services, controllers and routers...")

	authorizationService := services.NewAuthorizationService()
	preAuthService := services.NewPreAuthorizationService()
	confirmationService := services.NewConfirmationService()
	cancellationService := services.NewCancellationService()
	reversalService := services.NewReversalService()

	authorizationController := financial.NewAuthorizationController(authorizationService)
	preAuthController := financial.NewPreAuthorizationController(preAuthService)
	confirmationController := financial.NewConfirmationController(confirmationService)
	cancellationController := financial.NewCancellationController(cancellationService)
	reversalController := financial.NewReversalController(reversalService)

	routes.RegisterFinancialRoutes(r, authorizationController, preAuthController, confirmationController, cancellationController, reversalController)

	go func() {
		logrus.Infof("Starting HTTP server: %s", addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.Fatalf("Could not listen on %s: %v", addr, err)
		}
	}()

	//=================================
	//===== Gracefully Shutdown =======
	//=================================
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	// Wait for interrupt signal
	<-signalChan
	logrus.Info("Shutting down server...")

	// Create a context with a timeout for the shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Shutdown the server gracefully
	if err := server.Shutdown(ctx); err != nil {
		logrus.Fatalf("Server forced to shutdown: %v", err)
	}

	logrus.Info("Server exiting")
}

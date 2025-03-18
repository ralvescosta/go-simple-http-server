package routes

import (
	"github.com/go-chi/chi"

	"githib.com/ralvescosta/go-simple-http-server/pkg/controllers"
)

func RegisterAuthorizationRoutes(r *chi.Mux, controller *controllers.AuthorizationController) {
	r.Post("/v1/authorizaton", controller.Post)
}

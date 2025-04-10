package routes

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/sirupsen/logrus"
	httpSwagger "github.com/swaggo/http-swagger"

	"githib.com/ralvescosta/go-simple-http-server/pkg/controllers/financial"
)

func RegisterFinancialRoutes(
	r *chi.Mux,
	auth *financial.AuthorizationController,
	preAuth *financial.PreAuthorizationController,
	confirmation *financial.ConfirmationController,
	cancellation *financial.CancellationController,
	reversal *financial.ReversalController,
) {
	logrus.Debug("GET /swagger/*")
	r.Mount("/swagger/", httpSwagger.WrapHandler)
	r.Use(middleware.Heartbeat("/ping"))

	logrus.Debug("POST /v1/payments/authorization")
	r.Post("/v1/payments/authorization", auth.Post)

	logrus.Debug("POST /v1/payments/pre_authorization")
	r.Post("/v1/payments/pre_authorization", preAuth.Post)

	logrus.Debug("POST /v1/payments/confirmation")
	r.Post("/v1/payments/confirmation", confirmation.Post)

	logrus.Debug("POST /v1/payments/cancellation")
	r.Post("/v1/payments/cancellation", cancellation.Post)

	logrus.Debug("POST /v1/payments/reversal")
	r.Post("/v1/payments/reversal", reversal.Post)
}

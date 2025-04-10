package financial

import (
	"encoding/json"
	"net/http"

	"githib.com/ralvescosta/go-simple-http-server/internal/models"
	"githib.com/ralvescosta/go-simple-http-server/internal/services"
	"githib.com/ralvescosta/go-simple-http-server/pkg/controllers"
)

type (
	PreAuthorizationController struct {
		service services.PreAuthorizationService
	}
)

func NewPreAuthorizationController(service services.PreAuthorizationService) *PreAuthorizationController {
	return &PreAuthorizationController{service}
}

// Post godoc
// @Summary Process pre-authorization
// @Description Process a payment pre-authorization request
// @Tags financial
// @Accept json
// @Produce json
// @Param body body models.PreAuthorizationRequest true "Pre-authorization request"
// @Success 200 {object} models.PreAuthorizationResponse
// @Failure 400 {object} controllers.HTTPResponse
// @Failure 500 {object} controllers.HTTPResponse
func (c *PreAuthorizationController) Post(w http.ResponseWriter, r *http.Request) {
	var body models.PreAuthorizationRequest

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		controllers.NewResponseBuilder(w).UnformattedBody().Build()
		return
	}

	if validationErr := controllers.BodyValidator(&body); validationErr != nil {
		controllers.NewResponseBuilder(w).InvalidBody().ErrMessage(validationErr.Message).Build()
		return
	}

	resp, err := c.service.Process(r.Context(), &body)
	if err != nil {
		controllers.NewResponseBuilder(w).InternalError().ErrMessage(err.Error()).Build()
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		controllers.NewResponseBuilder(w).InternalError().ErrMessage(err.Error()).Build()
		return
	}
}

package controllers

import (
	"encoding/json"
	"net/http"

	"githib.com/ralvescosta/go-simple-http-server/internal/models"
	"githib.com/ralvescosta/go-simple-http-server/internal/services"
)

type (
	AuthorizationController struct {
		service services.AuthorizationService
	}
)

func NewAuthorizationController(service services.AuthorizationService) *AuthorizationController {
	return &AuthorizationController{service}
}

func (c *AuthorizationController) Post(w http.ResponseWriter, r *http.Request) {
	var body models.AuthorizationRequest

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		NewResponseBuilder(w).UnformattedBody().Build()
		return
	}

	if validationErr := BodyValidator(&body); validationErr != nil {
		NewResponseBuilder(w).InvalidBody().ErrMessage(validationErr.Message).Build()
		return
	}

	resp, err := c.service.Auth(r.Context(), &body)
	if err != nil {
		NewResponseBuilder(w).InternalError().ErrMessage(err.Error()).Build()
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		NewResponseBuilder(w).InternalError().ErrMessage(err.Error()).Build()
		return
	}
}

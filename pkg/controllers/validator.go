package controllers

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func BodyValidator(body any) *HTTPError {
	val := validator.New()
	err := val.Struct(body)

	if err == nil {
		return nil
	}

	validationErrors := err.(validator.ValidationErrors)
	messages := make(map[string]string, len(validationErrors))

	for _, validationErr := range validationErrors {
		message := ""
		if validationErr.Tag() == "required" {
			message = fmt.Sprintf("%s is required", validationErr.Field())
		} else {
			message = fmt.Sprintf("%s invalid %s", validationErr.Field(), validationErr.Tag())
		}
		messages[validationErr.Field()] = message
	}

	return &HTTPError{
		StatusCode: http.StatusBadRequest,
		Message:    "invalid body",
		Details:    messages,
	}
}

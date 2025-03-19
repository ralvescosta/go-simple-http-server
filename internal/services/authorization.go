package services

import (
	"context"

	"githib.com/ralvescosta/go-simple-http-server/internal/models"
)

type (
	AuthorizationService interface {
		Process(ctx context.Context, req *models.AuthorizationRequest) (*models.AuthorizationResponse, error)
	}

	authorizationService struct{}
)

func NewAuthorizationService() AuthorizationService {
	return &authorizationService{}
}

func (s *authorizationService) Process(ctx context.Context, req *models.AuthorizationRequest) (*models.AuthorizationResponse, error) {
	return nil, nil
}

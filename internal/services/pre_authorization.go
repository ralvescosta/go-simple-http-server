package services

import "context"

type (
	PreAuthorizationService interface {
		Process(ctx context.Context, req any) (any, error)
	}

	preAuthorizationService struct{}
)

func NewPreAuthorizationService() PreAuthorizationService {
	return &preAuthorizationService{}
}

func (s *preAuthorizationService) Process(ctx context.Context, req any) (any, error) {
	return nil, nil
}

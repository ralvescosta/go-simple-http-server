package services

import "context"

type (
	ConfirmationService interface {
		Process(ctx context.Context, req any) (any, error)
	}

	confirmationService struct{}
)

func NewConfirmationService() ConfirmationService {
	return &confirmationService{}
}

func (s *confirmationService) Process(ctx context.Context, req any) (any, error) {
	return nil, nil
}

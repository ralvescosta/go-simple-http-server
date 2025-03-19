package services

import (
	"context"
)

type (
	CancellactionService interface {
		Process(ctx context.Context, req any) (any, error)
	}

	cancellationService struct {
	}
)

func NewCancellationService() CancellactionService {
	return &cancellationService{}
}

func (s *cancellationService) Process(ctx context.Context, req any) (any, error) {
	return nil, nil
}

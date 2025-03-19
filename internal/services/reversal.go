package services

import "context"

type (
	ReversalService interface {
		Process(ctx context.Context, req any) (any, error)
	}

	reversalService struct{}
)

func NewReversalService() ReversalService {
	return &reversalService{}
}

func (s *reversalService) Process(ctx context.Context, req any) (any, error) {
	return nil, nil
}

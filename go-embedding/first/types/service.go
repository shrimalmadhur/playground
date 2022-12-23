package types

import (
	"context"
)

type Service struct {
	client Client
}

func NewService(client Client) *Service {
	return &Service{client: client}
}

func(s *Service) MethodA(ctx context.Context, message string) {
	s.client.MethodA(ctx, message)
}

func(s *Service) MethodB(ctx context.Context, message string) {
	s.client.MethodB(ctx, message)
}

package webhooks

import (
	"context"
	"fmt"
	"os"
)

type Service interface {
	GetValidateEntry(ctx context.Context, token string) bool
	PostReceiveMessage(ctx context.Context, data string)
}

type service struct {
}

func NewService() Service {
	return &service{}
}

func (s *service) GetValidateEntry(ctx context.Context, token string) bool {
	dotEnv := os.Getenv("TOKEN")

	if token == dotEnv {
		return true
	}
	return false
}

func (s *service) PostReceiveMessage(ctx context.Context, data string) {
	fmt.Println(data)
}

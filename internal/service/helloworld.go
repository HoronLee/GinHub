package service

import (
	"context"

	"github.com/horonlee/ginhub/internal/data"
	"github.com/horonlee/ginhub/internal/model/helloworld"
)

type HelloWorldService interface {
	PostHelloWorld(ctx context.Context, message string) error
	GetDatabaseInfo(ctx context.Context) (string, error)
}

type helloWorldService struct {
	repo data.HelloWorldRepo
}

func NewHelloWorldService(repo data.HelloWorldRepo) HelloWorldService {
	return &helloWorldService{repo: repo}
}

func (s *helloWorldService) PostHelloWorld(ctx context.Context, message string) error {
	hw := &helloworld.HelloWorld{Message: message}
	return s.repo.CreateHelloWorld(ctx, hw)
}

func (s *helloWorldService) GetDatabaseInfo(ctx context.Context) (string, error) {
	return s.repo.GetDatabaseInfo(ctx)
}

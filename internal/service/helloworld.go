package service

import (
	"context"

	"github.com/HoronLee/GinHub/internal/model/helloworld"
)

// HelloWorldRepo 定义HelloWorld数据访问接口
type HelloWorldRepo interface {
	CreateHelloWorld(ctx context.Context, hw *helloworld.HelloWorld) error
	GetDatabaseInfo(ctx context.Context) (string, error)
}

// HelloWorldService 定义HelloWorld服务接口
type HelloWorldService interface {
	PostHelloWorld(ctx context.Context, message string) error
	GetDatabaseInfo(ctx context.Context) (string, error)
}

// helloWorldService HelloWorld服务实现
type helloWorldService struct {
	repo HelloWorldRepo
}

// NewHelloWorldService 创建HelloWorldService实例
func NewHelloWorldService(repo HelloWorldRepo) HelloWorldService {
	return &helloWorldService{repo: repo}
}

func (s *helloWorldService) PostHelloWorld(ctx context.Context, message string) error {
	hw := &helloworld.HelloWorld{Message: message}
	return s.repo.CreateHelloWorld(ctx, hw)
}

func (s *helloWorldService) GetDatabaseInfo(ctx context.Context) (string, error) {
	return s.repo.GetDatabaseInfo(ctx)
}

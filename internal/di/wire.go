//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/horonlee/ginhub/internal/config"
	data "github.com/horonlee/ginhub/internal/data"
	handler "github.com/horonlee/ginhub/internal/handler"
	"github.com/horonlee/ginhub/internal/server"
	service "github.com/horonlee/ginhub/internal/service"
)

// InitServer 初始化服务器（Wire自动生成依赖注入代码）
func InitServer(cfg *config.AppConfig) (*server.HTTPServer, error) {
	wire.Build(
		// Data层
		data.ProviderSet,

		// Service层
		service.ProviderSet,

		// Handler层
		handler.ProviderSet,

		// Server层
		server.ProviderSet,
	)
	return nil, nil
}

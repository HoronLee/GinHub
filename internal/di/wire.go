//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/horonlee/ginhub/internal/config"
	"github.com/horonlee/ginhub/internal/data"
	"github.com/horonlee/ginhub/internal/handler"
	"github.com/horonlee/ginhub/internal/server"
	"github.com/horonlee/ginhub/internal/service"
)

// InitServer 初始化服务器
func InitServer(cfg *config.AppConfig) (*server.HTTPServer, error) {
	wire.Build(
		data.ProviderSet,
		service.ProviderSet,
		handler.ProviderSet,
		server.ProviderSet,
	)
	return nil, nil
}

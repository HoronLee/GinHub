//go:build wireinject
// +build wireinject

package di

import (
	"github.com/HoronLee/GinHub/internal/config"
	"github.com/HoronLee/GinHub/internal/data"
	"github.com/HoronLee/GinHub/internal/handler"
	"github.com/HoronLee/GinHub/internal/server"
	"github.com/HoronLee/GinHub/internal/service"
	"github.com/google/wire"
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

package handler

import "github.com/google/wire"

// ProviderSet is handler providers.
var ProviderSet = wire.NewSet(NewHandlers, NewHelloWorldHandler)

// Handlers 聚合各个模块的Handler
type Handlers struct {
	HelloWorldHandler *HelloWorldHandler
}

// NewHandlers 创建Handlers实例
func NewHandlers(hwHandler *HelloWorldHandler) *Handlers {
	return &Handlers{
		HelloWorldHandler: hwHandler,
	}
}

package router

import "github.com/horonlee/ginhub/internal/handler"

// setupHelloWorldRoutes 设置HelloWorld路由
func setupHelloWorldRoutes(appRouterGroup *AppRouterGroup, h *handler.Handlers) {
	// Public
	appRouterGroup.PublicRouterGroup.POST("/helloworld", h.HelloWorldHandler.PostHelloWorld())
}

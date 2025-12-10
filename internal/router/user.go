package router

import "github.com/HoronLee/GinHub/internal/handler"

// setupUserRoutes 设置用户路由
func setupUserRoutes(appRouterGroup *AppRouterGroup, h *handler.Handlers) {
	// Public routes - 公开路由，无需认证
	appRouterGroup.PublicRouterGroup.POST("/register", h.UserHandler.Register())
	appRouterGroup.PublicRouterGroup.POST("/login", h.UserHandler.Login())

	// Private routes - 私有路由，需要 JWT 认证
	appRouterGroup.PrivateRouterGroup.DELETE("/user", h.UserHandler.DeleteUser())
}

package router

import (
	"github.com/HoronLee/GinHub/internal/handler"
	"github.com/HoronLee/GinHub/internal/middleware"
	"github.com/gin-gonic/gin"
)

// AppRouterGroup 路由组
type AppRouterGroup struct {
	PublicRouterGroup  *gin.RouterGroup
	PrivateRouterGroup *gin.RouterGroup
}

// SetupRouter 配置路由
func SetupRouter(r *gin.Engine, h *handler.Handlers) {
	appRouterGroup := setupRouterGroup(r)

	// 设置各模块路由
	setupHelloWorldRoutes(appRouterGroup, h)
	setupUserRoutes(appRouterGroup, h)
}

// setupRouterGroup 初始化路由组
func setupRouterGroup(r *gin.Engine) *AppRouterGroup {
	public := r.Group("/api")
	private := r.Group("/api")
	private.Use(middleware.JWTAuthMiddleware()) // JWT认证中间件

	return &AppRouterGroup{
		PublicRouterGroup:  public,
		PrivateRouterGroup: private,
	}
}

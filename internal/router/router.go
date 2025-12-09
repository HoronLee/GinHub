package router

import (
	"github.com/gin-gonic/gin"
	"github.com/horonlee/ginhub/internal/handler"
)

type AppRouterGroup struct {
	PublicRouterGroup *gin.RouterGroup
	AuthRouterGroup   *gin.RouterGroup
}

// SetupRouter 配置路由
func SetupRouter(r *gin.Engine, h *handler.Handlers) {
	// 初始化路由组
	appRouterGroup := setupRouterGroup(r)

	// 设置各模块路由
	setupHelloWorldRoutes(appRouterGroup, h)
}

// setupRouterGroup 初始化路由组
func setupRouterGroup(r *gin.Engine) *AppRouterGroup {
	public := r.Group("/api")
	auth := r.Group("/api")
	// auth.Use(middleware.NoCache(), middleware.JWTAuthMiddleware()) // JWT认证暂不实现

	return &AppRouterGroup{
		PublicRouterGroup: public,
		AuthRouterGroup:   auth,
	}
}

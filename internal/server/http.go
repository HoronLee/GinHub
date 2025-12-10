// Package server
//
// @title GinHub API 文档
// @version 1.0
// @description 基于Gin、Gorm、Viper、Wire、Cobra的HTTP快速开发框架 API 文档
// @contact.name GinHub Team
// @contact.url https://github.com/HoronLee/GinHub
// @contact.email support@ginhub.dev
// @license.name MIT
// @license.url https://opensource.org/licenses/MIT
// @host localhost:8080
// @BasePath /api
// @schemes http https
package server

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/HoronLee/GinHub/internal/config"
	"github.com/HoronLee/GinHub/internal/handler"
	"github.com/HoronLee/GinHub/internal/middleware"
	"github.com/HoronLee/GinHub/internal/router"
	"github.com/HoronLee/GinHub/internal/swagger"
	util "github.com/HoronLee/GinHub/internal/util/log"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type HTTPServer struct {
	cfg        *config.AppConfig
	engine     *gin.Engine
	httpServer *http.Server
	handlers   *handler.Handlers
	db         *gorm.DB
	logger     *util.Logger
}

func NewHTTPServer(
	cfg *config.AppConfig,
	handlers *handler.Handlers,
	db *gorm.DB,
	logger *util.Logger,
) *HTTPServer {
	if cfg.Server.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	// 配置Swagger信息
	configureSwagger(cfg)

	engine := gin.New()
	engine.Use(middleware.Logger(logger))
	engine.Use(middleware.Recovery(logger))

	return &HTTPServer{
		cfg:      cfg,
		engine:   engine,
		handlers: handlers,
		db:       db,
		logger:   logger,
	}
}

func (s *HTTPServer) Start() error {
	router.SetupRouter(s.engine, s.handlers)

	addr := fmt.Sprintf("%s:%s", s.cfg.Server.Host, s.cfg.Server.Port)
	s.httpServer = &http.Server{
		Addr:    addr,
		Handler: s.engine,
	}

	s.logger.Info("Server starting", zap.String("addr", addr))

	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.logger.Fatal("Failed to start server", zap.Error(err))
		}
	}()

	return nil
}

func (s *HTTPServer) Stop(ctx context.Context) error {
	s.logger.Info("Shutting down server...")
	if s.httpServer != nil {
		return s.httpServer.Shutdown(ctx)
	}
	return nil
}

func (s *HTTPServer) GetEngine() *gin.Engine {
	return s.engine
}

// configureSwagger 根据配置和环境变量动态配置Swagger信息
func configureSwagger(cfg *config.AppConfig) {
	// 从环境变量或配置文件获取Swagger设置
	host := getEnvOrConfig("SWAGGER_HOST", cfg.Swagger.Host)
	if host == "" {
		// 如果没有配置，使用服务器配置构建默认host
		host = fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port)
		if cfg.Server.Host == "0.0.0.0" {
			host = fmt.Sprintf("localhost:%s", cfg.Server.Port)
		}
	}

	basePath := getEnvOrConfig("SWAGGER_BASE_PATH", cfg.Swagger.BasePath)
	if basePath == "" {
		basePath = "/api"
	}

	schemes := cfg.Swagger.Schemes
	if len(schemes) == 0 {
		if cfg.Server.Mode == "release" {
			schemes = []string{"https", "http"}
		} else {
			schemes = []string{"http", "https"}
		}
	}

	title := getEnvOrConfig("SWAGGER_TITLE", cfg.Swagger.Title)
	if title == "" {
		title = "GinHub API 文档"
	}

	description := getEnvOrConfig("SWAGGER_DESCRIPTION", cfg.Swagger.Description)
	if description == "" {
		description = "基于Gin、Gorm、Viper、Wire、Cobra的HTTP快速开发框架 API 文档"
	}

	version := getEnvOrConfig("SWAGGER_VERSION", cfg.Swagger.Version)
	if version == "" {
		version = "1.0"
	}

	// 更新Swagger文档信息
	swagger.SwaggerInfo.Host = host
	swagger.SwaggerInfo.BasePath = basePath
	swagger.SwaggerInfo.Schemes = schemes
	swagger.SwaggerInfo.Title = title
	swagger.SwaggerInfo.Description = description
	swagger.SwaggerInfo.Version = version
}

// getEnvOrConfig 优先从环境变量获取值，如果不存在则使用配置文件的值
func getEnvOrConfig(envKey, configValue string) string {
	if envValue := os.Getenv(envKey); envValue != "" {
		return envValue
	}
	return configValue
}

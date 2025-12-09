package server

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/horonlee/ginhub/internal/config"
	"github.com/horonlee/ginhub/internal/handler"
	"github.com/horonlee/ginhub/internal/model/helloworld"
	"gorm.io/gorm"
)

type HTTPServer struct {
	cfg        *config.AppConfig
	engine     *gin.Engine
	httpServer *http.Server
	hwHandler  *handler.HelloWorldHandler
	db         *gorm.DB
}

func NewHTTPServer(
	cfg *config.AppConfig,
	hwHandler *handler.HelloWorldHandler,
	db *gorm.DB,
) *HTTPServer {
	if cfg.Server.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	engine := gin.New()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())

	return &HTTPServer{
		cfg:       cfg,
		engine:    engine,
		hwHandler: hwHandler,
		db:        db,
	}
}

func (s *HTTPServer) Start() error {
	if err := s.db.AutoMigrate(&helloworld.HelloWorld{}); err != nil {
		return fmt.Errorf("failed to migrate database: %w", err)
	}

	s.registerRoutes()

	addr := fmt.Sprintf("%s:%s", s.cfg.Server.Host, s.cfg.Server.Port)
	s.httpServer = &http.Server{
		Addr:    addr,
		Handler: s.engine,
	}

	log.Printf("Server starting on %s\n", addr)

	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v\n", err)
		}
	}()

	return nil
}

func (s *HTTPServer) Stop(ctx context.Context) error {
	log.Println("Shutting down server...")
	if s.httpServer != nil {
		return s.httpServer.Shutdown(ctx)
	}
	return nil
}

func (s *HTTPServer) registerRoutes() {
	s.engine.POST("/helloworld", s.hwHandler.PostHelloWorld())
}

func (s *HTTPServer) GetEngine() *gin.Engine {
	return s.engine
}

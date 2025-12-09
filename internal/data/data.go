package data

import (
	"fmt"
	"time"

	"github.com/HoronLee/GinHub/internal/config"
	"github.com/HoronLee/GinHub/internal/model/helloworld"
	util "github.com/HoronLee/GinHub/internal/util/log"
	"github.com/google/wire"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewDB, NewHelloWorldRepo)

// NewDB 创建数据库连接
func NewDB(cfg *config.AppConfig, logger *util.Logger) (*gorm.DB, error) {
	var dialector gorm.Dialector

	// 根据配置选择数据库驱动
	switch cfg.Database.Driver {
	case "mysql":
		dialector = mysql.Open(cfg.Database.Source)
	case "sqlite":
		dialector = sqlite.Open(cfg.Database.Source)
	default:
		return nil, fmt.Errorf("unsupported database driver: %s", cfg.Database.Driver)
	}

	// 配置GORM日志
	gormLogger := util.NewGormLogger(logger)

	// 打开数据库连接
	db, err := gorm.Open(dialector, &gorm.Config{
		Logger: gormLogger,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}

	// 配置连接池
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get database instance: %w", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	logger.Info("Database connected successfully", zap.String("driver", cfg.Database.Driver))

	// 自动迁移数据库表
	if err = db.AutoMigrate(
		&helloworld.HelloWorld{},
		// 未来添加新模型只需在这里追加
	); err != nil {
		return nil, fmt.Errorf("failed to migrate database: %w", err)
	}

	return db, nil
}

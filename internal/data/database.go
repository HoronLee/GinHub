package data

import (
	"fmt"
	"log"
	"time"

	"github.com/horonlee/ginhub/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// NewDB 创建数据库连接（通过Wire注入）
func NewDB(cfg *config.AppConfig) (*gorm.DB, error) {
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
	var gormLogger logger.Interface
	if cfg.Database.LogMode == "debug" {
		gormLogger = logger.Default.LogMode(logger.Info)
	} else {
		gormLogger = logger.Default.LogMode(logger.Silent)
	}

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

	log.Println("Database connected successfully")
	return db, nil
}

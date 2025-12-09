package data

import (
	"context"

	"github.com/HoronLee/GinHub/internal/model/helloworld"
	"gorm.io/gorm"
)

// HelloWorldRepo 定义HelloWorld数据访问接口
type HelloWorldRepo interface {
	CreateHelloWorld(ctx context.Context, hw *helloworld.HelloWorld) error
	GetDatabaseInfo(ctx context.Context) (string, error)
}

// helloworldRepo HelloWorld数据访问实现
type helloworldRepo struct {
	db *gorm.DB
}

// NewHelloWorldRepo 创建HelloWorldRepo实例（通过Wire注入）
func NewHelloWorldRepo(db *gorm.DB) HelloWorldRepo {
	return &helloworldRepo{
		db: db,
	}
}

// CreateHelloWorld 创建HelloWorld记录
func (r *helloworldRepo) CreateHelloWorld(ctx context.Context, hw *helloworld.HelloWorld) error {
	return r.db.WithContext(ctx).Create(hw).Error
}

// GetDatabaseInfo 获取数据库连接信息
func (r *helloworldRepo) GetDatabaseInfo(ctx context.Context) (string, error) {
	dbName := r.db.Migrator().CurrentDatabase()
	if dbName == "" {
		dbName = "unknown"
	}
	return "Connected to: " + dbName, nil
}

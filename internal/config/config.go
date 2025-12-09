package config

import (
	"bytes"
	"crypto/rand"
	_ "embed"
	"encoding/hex"
	"log"
	"os"

	"github.com/spf13/viper"

	model "github.com/horonlee/ginhub/internal/model/common"
)

// Config 全局配置变量
var Config AppConfig

// JWT_SECRET 用于JWT签名的密钥
var JWT_SECRET []byte

// AppConfig 应用程序配置结构体
type AppConfig struct {
	Server struct {
		Port string `mapstructure:"port"` // 服务器端口
		Host string `mapstructure:"host"` // 服务器主机地址
		Mode string `mapstructure:"mode"` // 运行模式，可能的值为 "debug" 或 "release"
	} `mapstructure:"server"`
	Database struct {
		Driver  string `mapstructure:"type"`    // 数据库驱动
		Source  string `mapstructure:"source"`  // 数据库连接字符串
		LogMode string `mapstructure:"logmode"` // 数据库日志模式
	} `mapstructure:"database"`
	Auth struct {
		Jwt struct {
			Expires  int    `mapstructure:"expires"`  // JWT的过期时间，单位为秒
			Issuer   string `mapstructure:"issuer"`   // JWT的发行者
			Audience string `mapstructure:"audience"` // JWT的受众
		} `mapstructure:"jwt"`
	} `mapstructure:"auth"`
}

//go:embed config.yaml
var configData []byte

// LoadAppConfig 加载应用程序配置
// configPath: 外部配置文件路径，如果为空则只使用嵌入式配置
func LoadAppConfig(configPath string) {
	v := viper.New()
	v.SetConfigType("yaml")

	// 1. 先加载嵌入式配置作为默认配置
	err := v.ReadConfig(bytes.NewReader(configData))
	if err != nil {
		panic(model.READ_CONFIG_PANIC + ":" + err.Error())
	}

	// 2. 如果指定了外部配置文件，则用外部配置覆盖
	if configPath != "" {
		if _, err := os.Stat(configPath); err == nil {
			v.SetConfigFile(configPath)
			err = v.MergeInConfig() // 使用MergeInConfig合并配置
			if err != nil {
				log.Printf("Warning: failed to merge external config: %v, using embedded config\n", err)
			} else {
				log.Printf("Loaded external config from: %s\n", configPath)
			}
		} else {
			log.Printf("Warning: external config file not found: %s, using embedded config\n", configPath)
		}
	}

	// 3. 将配置反序列化到结构体
	err = v.Unmarshal(&Config)
	if err != nil {
		panic(model.READ_CONFIG_PANIC + ":" + err.Error())
	}

	// 4. 初始化 JWT_SECRET
	JWT_SECRET = GetJWTSecret()
}

// GetJWTSecret 加载JWT密钥
func GetJWTSecret() []byte {
	// 从环境变量中获取JWT密钥
	secret := os.Getenv("JWT_SECRET")
	if secret == "" { // 如果没有设置环境变量，则使用UUID生成默认密钥
		b := make([]byte, 16)
		_, err := rand.Read(b)
		if err != nil {
			log.Fatal("failed to generate random JWT secret:", err)
		}
		secret = hex.EncodeToString(b)
	}

	return []byte(secret)
}

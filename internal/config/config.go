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
		Port string `yaml:"port"` // 服务器端口
		Host string `yaml:"host"` // 服务器主机地址
		Mode string `yaml:"mode"` // 运行模式，可能的值为 "debug" 或 "release"
	} `yaml:"server"`
	Database struct {
		Type    string `yaml:"type"`    // 数据库类型
		Source  string `yaml:"source"`  // 数据库文件路径
		LogMode string `yaml:"logmode"` // 数据库日志模式
	} `yaml:"database"`
	Auth struct {
		Jwt struct {
			Expires  int    `yaml:"expires"`  // JWT的过期时间，单位为秒
			Issuer   string `yaml:"issuer"`   // JWT的发行者
			Audience string `yaml:"audience"` // JWT的受众
		} `yaml:"jwt"`
	} `yaml:"auth"`
}

//go:embed config.yaml
var configData []byte

// LoadAppConfig 加载应用程序配置
func LoadAppConfig() {
	// viper.SetConfigFile("config/config.yaml")
	viper.SetConfigType("yaml")
	// 使用嵌入的配置数据而不是从文件系统读取
	err := viper.ReadConfig(bytes.NewReader(configData))
	if err != nil {
		panic(model.READ_CONFIG_PANIC + ":" + err.Error())
	}

	// 将配置文件内容反序列化到结构体 Config 中
	err = viper.Unmarshal(&Config)
	if err != nil {
		panic(model.READ_CONFIG_PANIC + ":" + err.Error())
	}

	// 初始化 JWT_SECRET
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

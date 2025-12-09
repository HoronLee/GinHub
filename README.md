# GinHub

> 基于Gin、Gorm、Viper、Wire、Cobra的HTTP快速开发框架

## 特性

- 严格依赖注入（Wire）- 除config外所有组件均通过DI管理
- 清晰的分层架构 - Handler/Service/Data三层分离
- 统一响应封装 - 参考Echo项目的Response设计
- 命令行支持 - 基于Cobra的CLI工具

## 目录架构

```
ginhub/
├── cmd/                    # 命令行入口
│   └── root.go            # Cobra根命令
├── internal/
│   ├── cli/               # CLI命令实现
│   ├── config/            # 配置管理（不使用DI）
│   ├── server/            # HTTP服务器
│   ├── router/            # 路由注册
│   ├── handler/           # HTTP处理层
│   ├── service/           # 业务逻辑层
│   ├── data/              # 数据访问层
│   ├── model/             # 数据模型
│   ├── response/          # 响应封
│   ├── middleware/        # 中间件
│   ├── database/          # 数据库连接
│   ├── client/            # 外部客户端
│   ├── util/              # 工具函数
│   └── di/                # 依赖注入配置
│       ├── wire.go        # Wire配置
│       ├── wire_gen.go    # Wire生成代码
│       └── providers.go   # Provider定义
└── Makefile               # 构建脚本
```

## 构建流程

### 1. 环境准备

```bash
# 安装依赖
go mod download

# 安装Wire工具
go install github.com/google/wire/cmd/wire@latest
```

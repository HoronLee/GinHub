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
│   ├── cli/               # CLI命令实现（参考Echo）
│   ├── config/            # 配置管理（不使用DI）
│   ├── server/            # HTTP服务器
│   ├── router/            # 路由注册
│   ├── handler/           # HTTP处理层（对应Echo的handler）
│   ├── service/           # 业务逻辑层（对应Echo的service）
│   ├── data/              # 数据访问层（对应Echo的repository）
│   ├── model/             # 数据模型
│   ├── response/          # 响应封装（参考Echo的handler/response）
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

## 层级对应关系

完全模仿Echo项目的架构，GinHub采用以下三层分离设计：

### Handler层（HTTP处理层）
- **对应Echo**: `internal/handler/*`
- **职责**: 处理HTTP请求/响应，参数验证，业务流程编排，调用Service层
- **依赖注入**: ✅ 通过Wire注入Service依赖

### Service层（业务逻辑层）
- **对应Echo**: `internal/service/*`
- **职责**: 核心业务逻辑，领域规则，调用Data层
- **依赖注入**: ✅ 通过Wire注入Data层依赖

### Data层（数据访问层）
- **对应Echo**: `internal/repository/*`
- **职责**: 数据库操作，缓存操作，外部API调用
- **依赖注入**: ✅ 通过Wire注入DB、Cache等基础设施

### Response封装
- **对应Echo**: `internal/handler/response/response.go`
- **特点**: 
  - 统一的Response结构（Code/Data/Msg/Err）
  - Execute包装器自动处理响应
  - 支持自定义状态码

## 构建流程

### 1. 环境准备

```bash
# 安装依赖
go mod download

# 安装Wire工具
go install github.com/google/wire/cmd/wire@latest
```

### 2. 配置管理

```bash
# config包不使用依赖注入，直接初始化
# 使用Viper加载配置文件
```

### 3. 定义模型和接口

```bash
# 1. 在internal/model中定义数据模型
# 2. 在internal/data中定义数据访问接口和实现
# 3. 在internal/service中定义业务逻辑接口和实现
```

### 4. 实现业务逻辑

```bash
# 按照 Data -> Service -> Handler 的顺序实现
# 每层都定义interface，通过Wire注入依赖
```

### 5. 配置依赖注入

```bash
# 编辑 internal/di/wire.go
# 定义各层的ProviderSet
# 运行 wire 生成代码
cd internal/di && wire
```

### 6. 注册路由

```bash
# 在 internal/router 中注册路由
# Handler通过DI注入到Router
```

### 7. 实现CLI命令

```bash
# 参考Echo的CLI设计
# 在 internal/cli 中实现命令
# 在 cmd/root.go 中注册Cobra命令
```

### 8. 构建和运行

```bash
# 开发模式
make dev

# 生成Wire代码
make wire

# 构建
make build

# 运行
./ginhub serve
```

## 依赖注入示例

```go
// internal/di/wire.go
//go:build wireinject

package di

import (
    "github.com/google/wire"
)

// DataSet 数据层
var DataSet = wire.NewSet(
    ProvideDB,
    data.NewUserData,
)

// ServiceSet 业务逻辑层
var ServiceSet = wire.NewSet(
    DataSet,
    service.NewUserService,
)

// HandlerSet 处理层
var HandlerSet = wire.NewSet(
    ServiceSet,
    handler.NewUserHandler,
)

func BuildHandlers() (*Handlers, error) {
    wire.Build(HandlerSet, NewHandlers)
    return &Handlers{}, nil
}
```

## 响应封装示例

```go
// internal/response/response.go
type Response struct {
    Code int    `json:"code"`
    Data any    `json:"data,omitempty"`
    Msg  string `json:"msg"`
    Err  error  `json:"-"`
}

func Execute(fn func(*gin.Context) Response) gin.HandlerFunc {
    return func(ctx *gin.Context) {
        res := fn(ctx)
        if res.Err != nil {
            ctx.JSON(http.StatusBadRequest, gin.H{
                "code": res.Code,
                "msg":  res.Msg,
            })
            return
        }
        ctx.JSON(http.StatusOK, gin.H{
            "code": res.Code,
            "data": res.Data,
            "msg":  res.Msg,
        })
    }
}
```

## 开发规范

1. **除config外，所有组件必须通过Wire进行依赖注入**
2. **每层定义清晰的interface，面向接口编程**
3. **Handler层处理HTTP逻辑和业务流程编排**
4. **Service层实现核心业务逻辑和领域规则**
5. **Data层只负责数据访问，不包含业务逻辑**
6. **使用Response.Execute统一处理响应**
7. **参考Echo项目的CLI设计实现命令行工具**

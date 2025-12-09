# GinHub 启动指南

## 方式一：使用外部配置文件（推荐）

```bash
# 使用 debug-config.yaml
./ginhub serve --config configs/debug-config.yaml

# 或使用短参数
./ginhub serve -c configs/debug-config.yaml
```

## 方式二：使用嵌入的默认配置

```bash
# 不指定配置文件，使用内嵌的 config.yaml
./ginhub serve
```

**注意：** 默认配置使用的数据库连接为：
- Host: 127.0.0.1:3306
- User: root
- Password: password
- Database: ginhub

## 其他命令

```bash
# 启动 TUI 交互界面
./ginhub tui

# 查看版本
./ginhub version

# 查看信息
./ginhub info

# 显示 Logo
./ginhub hello
```

## 配置文件说明

配置文件支持以下字段：

```yaml
server:
  port: "6000"           # 服务端口
  host: "0.0.0.0"        # 监听地址
  mode: "debug"          # 运行模式: debug/release

database:
  type: "mysql"          # 数据库类型: mysql/sqlite
  source: "..."          # 数据库连接字符串
  logmode: "debug"       # 日志模式: debug/release

auth:
  jwt:
    expires: 2592000     # JWT过期时间（秒）
    issuer: "ginhub"     # JWT签发者
    audience: "ginhub"   # JWT受众
```

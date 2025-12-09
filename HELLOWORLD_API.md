# HelloWorld API 使用说明

## 功能说明

POST `/helloworld` 接口实现以下功能：
1. 接收带有 `message` 字段的 JSON 请求
2. 将 message 保存到数据库
3. 返回包含 message、版本号和数据库连接信息的响应

## 启动服务

```bash
# 使用 debug-config.yaml 配置文件启动
./ginhub serve --config configs/debug-config.yaml
```

## API 接口

### POST /helloworld

**请求示例：**

```bash
curl -X POST http://localhost:6000/helloworld \
  -H "Content-Type: application/json" \
  -d '{"message": "Hello from GinHub!"}'
```

**请求体：**

```json
{
  "message": "Hello from GinHub!"
}
```

**成功响应：**

```json
{
  "code": 0,
  "data": {
    "message": "Hello from GinHub!",
    "version": "v0.1.0",
    "database": "Connected to: ginhub"
  },
  "msg": "success"
}
```

**错误响应（缺少 message 字段）：**

```json
{
  "code": 400,
  "msg": "Invalid request",
  "err": "Key: 'HelloWorldRequest.Message' Error:Field validation for 'Message' failed on the 'required' tag"
}
```

## 数据库表结构

```sql
CREATE TABLE `hello_worlds` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `message` text NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`)
);
```

## 项目架构

按照 Echo0 项目规范，采用三层架构：

### 1. Model 层 (`internal/model/helloworld/`)
- `HelloWorld`: 数据库实体模型
- `HelloWorldRequest`: 请求 DTO
- `HelloWorldResponse`: 响应 DTO

### 2. Data 层 (`internal/data/helloworld/`)
- `HelloWorldRepo`: 数据访问接口
- `CreateHelloWorld()`: 创建记录
- `GetDatabaseInfo()`: 获取数据库信息

### 3. Service 层 (`internal/service/helloworld/`)
- `HelloWorldService`: 业务逻辑接口
- `PostHelloWorld()`: 处理业务逻辑

### 4. Handler 层 (`internal/handler/helloworld/`)
- `HelloWorldHandler`: HTTP 处理器接口
- `PostHelloWorld()`: 处理 HTTP 请求

### 5. Response 封装 (`internal/response/`)
- `Response`: 统一响应结构
- `Execute()`: 自动处理响应包装

## 依赖注入

使用 Wire 进行依赖注入，所有组件通过 DI 管理：

```
Database -> Data -> Service -> Handler -> Server
```

## 测试

使用提供的测试脚本：

```bash
./test_helloworld.sh
```

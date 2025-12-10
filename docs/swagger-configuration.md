# Swagger Configuration Guide

## Overview

GinHub supports environment-specific Swagger configuration through configuration files and environment variables. This allows you to customize the Swagger documentation for different deployment environments (development, staging, production).

## Configuration Structure

The Swagger configuration is defined in the `swagger` section of the configuration file:

```yaml
swagger:
  host: "localhost:8080"              # Swagger文档的主机地址
  basepath: "/api"                    # API基础路径
  schemes: ["http", "https"]          # 支持的协议方案
  title: "GinHub API 文档"            # API文档标题
  description: "API 文档描述"          # API文档描述
  version: "1.0"                      # API版本
  contact_name: "GinHub Team"         # 联系人姓名
  contact_url: "https://github.com/HoronLee/GinHub"  # 联系人URL
  contact_email: "support@ginhub.dev" # 联系人邮箱
  license_name: "MIT"                 # 许可证名称
  license_url: "https://opensource.org/licenses/MIT"  # 许可证URL
```

## Environment Variables

You can override configuration values using environment variables:

- `SWAGGER_HOST`: Override the host setting
- `SWAGGER_BASE_PATH`: Override the base path setting
- `SWAGGER_TITLE`: Override the API title
- `SWAGGER_DESCRIPTION`: Override the API description
- `SWAGGER_VERSION`: Override the API version

## Environment-Specific Configurations

### Development Environment

For development, use the default configuration or `configs/debug-config.yaml`:

```bash
make dev
# or
make debug
```

This will use:
- Host: `localhost:8080`
- Schemes: `["http", "https"]`
- Mode: debug

### Production Environment

For production, use `configs/production-config.yaml`:

```bash
make prod
# or
go run main.go serve -c configs/production-config.yaml
```

This will use:
- Host: `api.ginhub.com` (or your production domain)
- Schemes: `["https"]` (HTTPS only for security)
- Mode: release

### Custom Environment

You can create custom configuration files for different environments:

```bash
# Create staging configuration
cp configs/production-config.yaml configs/staging-config.yaml
# Edit staging-specific settings
vim configs/staging-config.yaml

# Run with staging config
go run main.go serve -c configs/staging-config.yaml
```

## Build Process Integration

The Swagger documentation is automatically generated during the build process:

```bash
# Development build
make build

# Production build
make build-prod

# Generate Swagger docs only
make swagger
```

## Environment Variable Examples

```bash
# Override host for Docker deployment
export SWAGGER_HOST="api.example.com"

# Override for local development with custom port
export SWAGGER_HOST="localhost:3000"

# Set custom API version
export SWAGGER_VERSION="2.0"

# Run the application
go run main.go serve
```

## Docker Deployment

When deploying with Docker, you can pass environment variables:

```bash
docker run -e SWAGGER_HOST="api.example.com" -e SWAGGER_VERSION="1.0" ginhub
```

## Accessing Swagger UI

Once the application is running, access the Swagger UI at:

- Development: `http://localhost:8080/api/v1/swagger/index.html`
- Production: `https://your-domain.com/api/v1/swagger/index.html`

The Swagger JSON specification is available at:
- `http://localhost:8080/api/v1/swagger/doc.json`
# Quickstart: 用户登录功能

**Branch**: `001-user-login` | **Date**: 2026-04-17

## 功能介绍

用户登录功能是系统的基础功能，采用邀请制，只有用户表中的用户才能登录。

## 技术栈

- **语言**: Go 1.21.4
- **数据库**: PostgreSQL 15.0+
- **密码加密**: bcrypt
- **Web 框架**: 标准库 net/http (或 gorilla/mux)

## 快速启动

### 1. 环境准备

- 安装 Go 1.21.4 或更高版本
- 安装 PostgreSQL 15.0 或更高版本
- 创建数据库和用户表

### 2. 数据库初始化

```sql
-- 创建用户表
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    password_hash VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 创建索引
CREATE UNIQUE INDEX idx_users_username ON users(username);

-- 插入测试用户（密码为 "password123"）
INSERT INTO users (username, password_hash)
VALUES ('admin', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy');
```

### 3. 配置

在 `cmd/app/main.go` 中配置数据库连接：

```go
// 数据库连接配置
const (
    host     = "localhost"
    port     = 5432
    user     = "postgres"
    password = "your_password"
    dbname   = "stock"
)
```

### 4. 运行

```bash
# 构建并运行
cd cmd/app
go build -o app
./app

# 或直接运行
cd cmd/app
go run main.go
```

### 5. API 接口

#### 登录接口

- **URL**: `/api/auth/login`
- **方法**: POST
- **请求体**:
  ```json
  {
    "username": "admin",
    "password": "password123"
  }
  ```
- **响应**:
  ```json
  {
    "success": true,
    "user": {
      "id": 1,
      "username": "admin",
      "created_at": "2026-04-17T00:00:00Z",
      "updated_at": "2026-04-17T00:00:00Z"
    }
  }
  ```

#### 健康检查接口

- **URL**: `/api/health`
- **方法**: GET
- **响应**:
  ```json
  {
    "status": "ok"
  }
  ```

## 测试

### 单元测试

```bash
go test ./internal/...
```

### 集成测试

使用 Postman 或 curl 测试 API 接口。

## 故障排除

### 常见问题

1. **数据库连接失败**
   - 检查数据库配置是否正确
   - 确保 PostgreSQL 服务正在运行

2. **登录失败**
   - 检查用户名和密码是否正确
   - 检查用户是否存在于用户表中

3. **密码加密错误**
   - 确保安装了 `golang.org/x/crypto/bcrypt` 包

## 部署

### 生产环境部署

1. 配置环境变量存储数据库连接信息
2. 使用 HTTPS 协议
3. 配置适当的日志记录
4. 设置定期数据库备份

# Data Model: 用户登录功能

**Branch**: `001-user-login` | **Date**: 2026-04-17

## 数据模型

### 用户表 (users)

| 字段名 | 数据类型 | 约束 | 描述 |
|-------|---------|------|------|
| `id` | `SERIAL` | `PRIMARY KEY` | 用户ID |
| `username` | `VARCHAR(50)` | `UNIQUE NOT NULL` | 用户名 |
| `password_hash` | `VARCHAR(100)` | `NOT NULL` | 加密后的密码 |
| `created_at` | `TIMESTAMP` | `DEFAULT CURRENT_TIMESTAMP` | 注册时间 |
| `updated_at` | `TIMESTAMP` | `DEFAULT CURRENT_TIMESTAMP` | 更新时间 |

### 索引

- `username` 字段创建唯一索引，加速用户名查询

## Go 数据模型

```go
// internal/model/user.go
type User struct {
    ID           int       `json:"id" db:"id"`
    Username     string    `json:"username" db:"username"`
    PasswordHash string    `json:"-" db:"password_hash"` // 不在 JSON 中返回密码
    CreatedAt    time.Time `json:"created_at" db:"created_at"`
    UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}

// 用户登录请求
type LoginRequest struct {
    Username string `json:"username" binding:"required"`
    Password string `json:"password" binding:"required"`
}

// 用户登录响应
type LoginResponse struct {
    Success bool   `json:"success"`
    Message string `json:"message,omitempty"`
    User    *User  `json:"user,omitempty"`
}
```

## 数据访问层

- `UserRepository` 接口定义用户数据访问方法
- 实现 PostgreSQL 版本的用户仓库
- 提供用户查询、创建、更新等方法

## 数据流程

1. 用户登录时，通过用户名查询用户信息
2. 验证密码是否正确
3. 登录成功后，生成会话信息
4. 邀请用户时，在用户表中插入新记录

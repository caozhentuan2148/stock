# Research: 用户登录功能

**Branch**: `001-user-login` | **Date**: 2026-04-17

## 研究目标

本研究文档旨在为用户登录功能的实现提供技术参考和最佳实践。

## 技术研究

### 1. PostgreSQL 数据库

- **版本**: 15.0+
- **连接方式**: 使用 `database/sql` 包和 `github.com/lib/pq` 驱动
- **优势**: 稳定可靠，支持复杂查询，适合存储用户数据

### 2. 密码加密

- **库**: `golang.org/x/crypto/bcrypt`
- **工作原理**: 使用 bcrypt 算法对密码进行哈希处理
- **优势**: 安全性高，自动处理盐值，抵抗彩虹表攻击

### 3. Go 认证最佳实践

- **会话管理**: 可以使用 JWT 或 cookie 进行会话管理
- **安全头部**: 设置适当的 HTTP 安全头部
- **CSRF 保护**: 实现 CSRF 令牌验证
- **速率限制**: 防止暴力破解攻击

### 4. 项目结构

- 遵循 Go 项目标准结构
- 使用依赖注入模式
- 实现分层架构：handler → service → repository → model

## 参考资料

1. [Go 官方文档](https://golang.org/doc/)
2. [PostgreSQL 文档](https://www.postgresql.org/docs/)
3. [bcrypt 文档](https://pkg.go.dev/golang.org/x/crypto/bcrypt)
4. [Go 项目结构最佳实践](https://github.com/golang-standards/project-layout)

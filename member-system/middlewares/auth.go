package middlewares

import (
    "net/http"
    "strings"
    "github.com/gin-gonic/gin"
    "member-system/utils"
)

// AuthMiddleware 验证 JWT，并把用户信息存入上下文
func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "未提供认证令牌"})
            c.Abort()
            return
        }
        // 格式：Bearer <token>
        parts := strings.SplitN(authHeader, " ", 2)
        if len(parts) != 2 || parts[0] != "Bearer" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "令牌格式错误"})
            c.Abort()
            return
        }

        claims, err := utils.ParseToken(parts[1])
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "令牌无效或已过期"})
            c.Abort()
            return
        }
        // 将用户信息保存在上下文中，供后续控制器使用
        c.Set("userID", claims.UserID)
        c.Set("role", claims.Role)
        c.Next()
    }
}

// AdminMiddleware 校验是否为管理员（必须放在 AuthMiddleware 之后）
func AdminMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        role, exists := c.Get("role")
        if !exists || role != "admin" {
            c.JSON(http.StatusForbidden, gin.H{"error": "权限不足，需要管理员身份"})
            c.Abort()
            return
        }
        c.Next()
    }
}
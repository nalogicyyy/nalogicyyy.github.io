package routes

import (
    "github.com/gin-gonic/gin"
    "member-system/controllers"
    "member-system/middlewares"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()

    // 公开路由（无需登录）
    auth := r.Group("/api/auth")
    {
        auth.POST("/register", controllers.Register)
        auth.POST("/login", controllers.Login)
    }

    // 需要登录 + 管理员权限的路由
    admin := r.Group("/api/admin")
    admin.Use(middlewares.AuthMiddleware())    // 先验证登录
    admin.Use(middlewares.AdminMiddleware())   // 再校验管理员
    {
        admin.GET("/users", controllers.GetUsers)          // 获取所有成员
        admin.GET("/users/:id", controllers.GetUser)       // 获取单个成员
        admin.DELETE("/users/:id", controllers.DeleteUser) // 删除普通用户
    }

    return r
}
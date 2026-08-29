package controllers

import (
	"member-system/config"
	"member-system/models"
	"member-system/utils"
	"net/http"

	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
)

// RegisterRequest 注册请求参数
type RegisterRequest struct {
	Password string `json:"password" binding:"required,min=6"`
	Nickname string `json:"nickname"`
	Username string `json:"username" binding:"required"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
}

// LoginRequest 登录请求参数
type LoginRequest struct {
	UserID   string `json:"user_id"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Register 用户注册
func Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误: " + err.Error()})
		return
	}

	// 创建用户（默认角色为 'user'，如需管理员需手动在数据库修改）
	user := models.User{
		UserID:   uuid.NewString(), // 后端自动生成user_id
		Password: req.Password,
		Username: req.Username,
		Nickname: req.Nickname,
		Phone:    req.Phone,
		Email:    req.Email,
		Role:     "user",
	}
	if result := config.DB.Create(&user); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "注册失败: " + result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "注册成功",
		"user_id": user.UserID,
	})
}

// / Login 用户登录
func Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误: " + err.Error()})
		return
	}

	var user models.User
	// ✅ 使用 username 查询用户
	if result := config.DB.Where("username = ?", req.Username).First(&user); result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "用户不存在或密码错误"})
		return
	}

	if !user.CheckPassword(req.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "用户不存在或密码错误"})
		return
	}

	// 生成 JWT
	token, err := utils.GenerateToken(user.ID, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "生成令牌失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "登录成功",
		"token":   token,
		"role":    user.Role,
	})
}

// GetUsers 获取所有成员信息（仅管理员）
func GetUsers(c *gin.Context) {
	var users []models.User
	// 只查询需要的字段，隐藏密码
	result := config.DB.Select("id, user_id, nickname, phone, email, role, created_at").Find(&users)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "查询失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": users})
}

// GetUser 获取单个成员信息（仅管理员）
func GetUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	result := config.DB.Select("id, user_id, nickname, phone, email, role, created_at").First(&user, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "成员不存在"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}

// DeleteUser 删除普通用户（仅管理员）
func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	// 先查询要删除的用户信息，防止误删管理员（增强安全性）
	var user models.User
	if result := config.DB.First(&user, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "用户不存在"})
		return
	}
	if user.Role == "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "不允许删除管理员账号"})
		return
	}

	config.DB.Delete(&user)
	c.JSON(http.StatusOK, gin.H{"message": "删除成功"})
}

package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User 结构体（对应数据库表）
type User struct {
	gorm.Model        // 包含 ID, CreatedAt, UpdatedAt, DeletedAt
	UserID     string `gorm:"unique;not null" json:"user_id"` // 自定义唯一ID（如学号/工号）
	Password   string `json:"-"`
	Username   string `gorm:"column:username;unique"` // json:"-" 表示返回JSON时忽略密码
	Nickname   string `json:"nickname"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
	Role       string `gorm:"default:user" json:"role"` // admin 或 user
}

// BeforeCreate 是 Gorm 的钩子，在创建记录前自动加密密码
func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.Password == "" {
		return nil
	}
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPwd)
	return nil
}

// CheckPassword 校验密码是否正确
func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

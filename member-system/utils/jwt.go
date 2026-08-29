package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("your_secret_key_123456") // 生产环境建议从环境变量读取

type Claims struct {
	UserID uint   `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

// GenerateToken 生成 JWT
func GenerateToken(userID uint, role string) (string, error) {
	claims := Claims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // 有效期24小时
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// ParseToken 解析 JWT
func ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("无效的Token")
}

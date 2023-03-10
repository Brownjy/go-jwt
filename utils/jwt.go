package utils

import (
	"github.com/go-jwt/model"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

// TokenExpireDuration 有效时间
const TokenExpireDuration = time.Hour * 1

// MyClaims 自己的Claims结构体
type MyClaims struct {
	User                 model.User
	jwt.RegisteredClaims // 标准Claims结构体，可设置7个标准字段
}

// 生成签名秘钥
var jwtSecret = []byte("JWT_SECRET")

// GenerateToken 签发用户Token
func GenerateToken(user model.User) (signedString string, err error) {
	// 设置有效期
	expirationTime := time.Now().Add(TokenExpireDuration)
	// 初始化标准Claims
	registeredClaims := jwt.RegisteredClaims{
		Issuer:    "brown", // 签名颁发者
		Subject:   "",      // 签名主题
		Audience:  nil,
		ExpiresAt: jwt.NewNumericDate(expirationTime), // 过期时间
		NotBefore: nil,
		IssuedAt:  jwt.NewNumericDate(time.Now()), // 签名时间
		ID:        "",                             // 签名ID
	}
	//创建Token结构体
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, MyClaims{
		User:             user,
		RegisteredClaims: registeredClaims,
	})
	//调用加密方法,签发Token字符串
	signedString, err = token.SignedString(jwtSecret)
	return
}

// ParseToken 校验用户Token
func ParseToken(token string) (*MyClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*MyClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

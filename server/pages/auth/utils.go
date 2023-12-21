package auth

import (
	"github.com/LanceLRQ/cloud-clipboard/server/conf"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

// NewJWTToken 创建JWT令牌
func NewJWTToken() (string, int64, error) {
	CommonJWTTokenExpire := time.Hour * time.Duration(conf.SecurityConfig.JWTExpire)
	factory := jwt.New(jwt.SigningMethodHS256)

	expireAt := time.Now().Add(CommonJWTTokenExpire).Unix()

	claims := factory.Claims.(jwt.MapClaims)
	claims["exp"] = expireAt
	claims["iss"] = "cloud@clipboard.app"
	claims["nbf"] = time.Now().Unix()
	claims["iat"] = claims["nbf"]
	claims["sub"] = "login"
	claims["aud"] = "web"

	token, err := factory.SignedString([]byte(conf.SecurityConfig.JWTSecret))
	if err != nil {
		return "", 0, err
	}
	return "Bearer " + token, expireAt, nil
}

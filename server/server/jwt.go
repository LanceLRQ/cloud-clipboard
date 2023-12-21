package server

import (
	"github.com/LanceLRQ/cloud-clipboard/server/conf"
	"github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func CreateJWTMiddleware() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{
			Key: []byte(conf.SecurityConfig.JWTSecret),
		},
		AuthScheme:   "Bearer",
		ContextKey:   "jwt_token",
		ErrorHandler: jwtErrorHandler,
	})
}

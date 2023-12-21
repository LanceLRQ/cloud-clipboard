package routes

import (
	"github.com/LanceLRQ/cloud-clipboard/server/pages/auth"
	"github.com/LanceLRQ/cloud-clipboard/server/server"
	"github.com/gofiber/fiber/v2"
)

func SetupMainRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Get("", ApiIndex)

	apiAuth := app.Group("/api/auth")
	apiAuth.Get("", auth.AuthLogin)
	apiAuth.Get("otp_info", auth.OtpInfoView)

	clipBoard := app.Group("/api/clipboard")
	clipBoard.Use(server.CreateJWTMiddleware())
	clipBoard.Get("", ApiIndex)
}

func ApiIndex(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Hello World!"})
}

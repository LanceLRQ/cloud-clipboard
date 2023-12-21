package cli

import (
	"fmt"
	"github.com/LanceLRQ/cloud-clipboard/server/conf"
	"github.com/LanceLRQ/cloud-clipboard/server/routes"
	"github.com/LanceLRQ/cloud-clipboard/server/server"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func CreateServer(debugMode bool) error {
	var err error
	validatorEntity := validator.New()

	if debugMode {
		conf.ServerDebugging = true
		fmt.Println("[Server] Debug mode enable")
	}

	// 初始化logger
	releaseLogFile, err := server.InitServerLogger()
	if err != nil {
		return err
	}
	defer releaseLogFile()

	// 初始化Redis
	err = server.InitRedisDB()
	if err != nil {
		return err
	}

	app := fiber.New(fiber.Config{
		ErrorHandler: server.HTTPServerErrorHandler,
	})

	// recover中间件
	app.Use(server.ServerPanicLogMiddleware)

	// 注册全局validator
	app.Use(func(c *fiber.Ctx) error {
		c.Locals("validator", validatorEntity)
		return c.Next()
	})

	// 初始化路由
	routes.SetupMainRoutes(app)

	return app.Listen(fmt.Sprintf("%s:%d", conf.ServerConfig.HttpHost, conf.ServerConfig.HttpPort))
}

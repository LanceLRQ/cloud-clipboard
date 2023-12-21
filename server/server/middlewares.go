package server

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	log "github.com/sirupsen/logrus"
	"runtime/debug"
)

// ServerPanicLogMiddleware Server发生Panic时候的捕捉器
var ServerPanicLogMiddleware = recover.New(recover.Config{
	StackTraceHandler: func(c *fiber.Ctx, e interface{}) {
		log.Errorf("[Server] %s\n%s\n", e.(error).Error(), debug.Stack())
		return
	},
	EnableStackTrace: true,
})

// HTTPServerErrorHandler that process return errors from handlers
var HTTPServerErrorHandler = func(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	var httpError *HTTPServerError
	var fiberError *fiber.Error
	c.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)

	if errors.As(err, &httpError) {
		var err1 *HTTPServerError
		var ok bool
		if err1, ok = err.(*HTTPServerError); !ok {
			if err1, ok = errors.Unwrap(err).(*HTTPServerError); ok {
				return c.Status(err1.Code).JSON(&HTTPServerError{
					Code:    err1.Code,
					Errno:   err1.Errno,
					Message: err.Error(),
					Payload: err1.Payload,
				})
			} else {
				return c.Status(fiber.StatusInternalServerError).JSON(&HTTPServerError{
					Code:    code,
					Errno:   -1,
					Message: err.Error(),
				})
			}
		}
		return c.Status(err1.Code).JSON(err1)
	} else if errors.As(err, &fiberError) {
		e1 := err.(*fiber.Error)
		return c.Status(e1.Code).JSON(&HTTPServerError{
			Code:    e1.Code,
			Message: err.Error(),
		})
	}

	return c.Status(code).JSON(&HTTPServerError{
		Code:    code,
		Message: err.Error(),
	})
}

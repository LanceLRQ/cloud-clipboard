package server

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// HTTPServerError HTTP服务异常信息
// code为状态码，errno为这次异常对应的识别代码
// payload 为可选的内容，一般不返回
type HTTPServerError struct {
	Code    int         `json:"code"`
	Errno   int         `json:"errno,omitempty"`
	Message string      `json:"message"`
	Payload interface{} `json:"payload,omitempty"`
}

func (e *HTTPServerError) Error() string {
	return e.Message
}

func (e *HTTPServerError) WithPayload(payload interface{}) *HTTPServerError {
	return &HTTPServerError{
		Code:    e.Code,
		Errno:   e.Errno,
		Message: e.Message,
		Payload: payload,
	}
}

func (e *HTTPServerError) WithValidatorError(err error) *HTTPServerError {
	if vErrs, ok := err.(validator.ValidationErrors); ok {
		vMsg := make([]string, 0, 1)
		for _, vErr := range vErrs {
			vMsg = append(vMsg, vErr.Error())
		}
		return &HTTPServerError{
			Code:    e.Code,
			Errno:   e.Errno,
			Message: e.Message,
			Payload: struct {
				Errors []string `json:"errors"`
			}{
				Errors: vMsg,
			},
		}
	}
	return e
}

// NewHTTPServerError 创建一个500异常
func NewHTTPServerError(errno int, message string) *HTTPServerError {
	err := &HTTPServerError{
		Code:    fiber.StatusInternalServerError,
		Errno:   errno,
		Message: message,
	}
	return err
}

// NewHTTPServerRequestError 创建一个400异常
func NewHTTPServerRequestError(errno int, message string) *HTTPServerError {
	err := &HTTPServerError{
		Code:    fiber.StatusBadRequest,
		Errno:   errno,
		Message: message,
	}
	return err
}

var jwtTokenWrongErrMsg = fiber.Map{"code": "400", "message": "Missing or malformed JWT"}
var jwtTokenInvalidErrMsg = fiber.Map{"code": "401", "message": "Invalid or expired JWT"}

func jwtErrorHandler(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).JSON(jwtTokenWrongErrMsg)
	}
	return c.Status(fiber.StatusUnauthorized).JSON(jwtTokenInvalidErrMsg)
}

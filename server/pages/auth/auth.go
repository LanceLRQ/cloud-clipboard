package auth

import (
	"fmt"
	"github.com/LanceLRQ/cloud-clipboard/server/conf"
	"github.com/LanceLRQ/cloud-clipboard/server/exceptions"
	"github.com/LanceLRQ/cloud-clipboard/server/structs/requests"
	"github.com/LanceLRQ/cloud-clipboard/server/structs/responses"
	"github.com/gofiber/fiber/v2"
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
	"strings"
	"time"
)

func AuthLogin(c *fiber.Ctx) error {
	if strings.TrimSpace(conf.SecurityConfig.OTPUrl) == "" {
		return exceptions.OTPSecretError
	}
	otpKey, err := otp.NewKeyFromURL(conf.SecurityConfig.OTPUrl)
	if err != nil {
		return fmt.Errorf("%w: %s", exceptions.OTPSecretError, err)
	}

	params := requests.AuthRequestParams{}
	if err := c.BodyParser(&params); err != nil {
		return fmt.Errorf("%w: %s", exceptions.ParseJSONError, err)
	}

	ok, err := totp.ValidateCustom(params.OTPPassword, otpKey.Secret(), time.Now(), totp.ValidateOpts{
		Period:    uint(otpKey.Period()),
		Algorithm: otpKey.Algorithm(),
		Digits:    otpKey.Digits(),
	})
	if err != nil {
		return fmt.Errorf("%w: %s", exceptions.OTPPasswordValidateError, err)
	}
	if !ok {
		return exceptions.OTPPasswordValidateError
	}

	// 创建令牌
	token, expireAt, err := NewJWTToken()
	if err != nil {
		return fmt.Errorf("%w: %s", exceptions.InternalServerError, err)
	}
	return c.JSON(responses.AuthLogin{
		Token:    token,
		ExpireAt: expireAt,
	})
}

func OtpInfoView(c *fiber.Ctx) error {
	if strings.TrimSpace(conf.SecurityConfig.OTPUrl) == "" {
		return exceptions.OTPSecretError
	}
	otpKey, err := otp.NewKeyFromURL(conf.SecurityConfig.OTPUrl)
	if err != nil {
		return fmt.Errorf("%w: %s", exceptions.OTPSecretError, err)
	}
	code, _ := totp.GenerateCodeCustom(otpKey.Secret(), time.Now(), totp.ValidateOpts{
		Period:    uint(otpKey.Period()),
		Algorithm: otpKey.Algorithm(),
		Digits:    otpKey.Digits(),
	})
	return c.JSON(responses.OTPInfo{
		Secret:   otpKey.Secret(),
		Url:      conf.SecurityConfig.OTPUrl,
		TestCode: code,
	})
}

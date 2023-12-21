package auth

import (
	"fmt"
	"github.com/LanceLRQ/cloud-clipboard/server/conf"
	"github.com/LanceLRQ/cloud-clipboard/server/exceptions"
	"github.com/LanceLRQ/cloud-clipboard/server/structs/responses"
	"github.com/gofiber/fiber/v2"
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
	"strings"
	"time"
)

func AuthLogin(c *fiber.Ctx) error {
	// 发放令牌
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
		return fmt.Errorf("%w: %s", exceptions.InternalServerError, "请配置OTP密钥")
	}
	otpKey, err := otp.NewKeyFromURL(conf.SecurityConfig.OTPUrl)
	if err != nil {
		return fmt.Errorf("%w: %s", exceptions.InternalServerError, err)
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

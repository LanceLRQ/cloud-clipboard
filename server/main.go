package main

import "github.com/LanceLRQ/cloud-clipboard/server/cli"

func main() {
	//app := fiber.New()
	//
	//app.Get("/", func(c *fiber.Ctx) error {
	//	return c.SendString("Hello, World 👋!")
	//})
	//// Test OTP:
	//app.Get("/otp-gen", func(c *fiber.Ctx) error {
	//	key, err := totp.Generate(totp.GenerateOpts{
	//		Issuer:      "test",
	//		AccountName: "test-account",
	//		Period:      32,
	//	})
	//	if err != nil {
	//		return c.SendString(err.Error())
	//	}
	//	c.Type("html", "utf-8")
	//	return c.SendString(key.Secret() + "<br />" + key.URL())
	//})
	//
	//app.Get("/otp", func(c *fiber.Ctx) error {
	//	code, _ := totp.GenerateCode("FUS2CYV2ENFXGQFBUG7PBU3UTGNFWHT7", time.Now())
	//	return c.SendString(code)
	//})
	//// END
	//
	//app.Listen(":3000")

	cli.RunCli()
}

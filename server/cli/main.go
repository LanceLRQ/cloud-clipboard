package cli

import (
	"fmt"
	"github.com/LanceLRQ/cloud-clipboard/server/conf"
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yamlv3"
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"strings"
	"time"
)

func RunCli() {
	app := &cli.App{
		Name:  "Cloud Clipboard",
		Usage: "",
		Before: func(context *cli.Context) error {
			config.AddDriver(yamlv3.Driver)
			err := config.LoadFiles("./server.yaml")
			if err != nil {
				return err
			}
			// 载入服务端配置
			return conf.LoadServerConfig()
		},
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "debug",
				Usage: "Enable debug mode.",
			},
		},
		Action: func(ctx *cli.Context) error {
			return CreateServer(ctx.Bool("debug"))
		},
		Commands: []*cli.Command{
			{
				Name:  "otp_gen",
				Usage: "Generate otp url",
				Subcommands: []*cli.Command{
					{
						Name:  "secret",
						Usage: "Generate OTP secret url",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:    "issuer",
								Aliases: []string{"i"},
								Value:   "test",
							},
							&cli.StringFlag{
								Name:    "account",
								Aliases: []string{"a"},
								Value:   "test-account",
							},
							&cli.IntFlag{
								Name:    "period",
								Aliases: []string{"l", "p"},
								Value:   32,
							},
						},
						Action: func(c *cli.Context) error {
							otpKey, err := totp.Generate(totp.GenerateOpts{
								Issuer:      c.String("issuer"),
								AccountName: c.String("account"),
								Period:      32,
							})
							fmt.Println("otp_secret: " + otpKey.Secret())
							fmt.Println("otp_url: " + otpKey.URL())
							return err
						},
					},
					{
						Name:  "code",
						Usage: "Generate OTP code from secret url",
						Action: func(c *cli.Context) error {
							optUrl := c.Args().Get(0)
							if strings.TrimSpace(optUrl) == "" {
								fmt.Println("缺少OTP的URL")
								return nil
							}
							otpKey, err := otp.NewKeyFromURL(optUrl)
							if err != nil {
								return err
							}
							code, err := totp.GenerateCodeCustom(otpKey.Secret(), time.Now(), totp.ValidateOpts{
								Period:    uint(otpKey.Period()),
								Algorithm: otpKey.Algorithm(),
								Digits:    otpKey.Digits(),
							})
							fmt.Println(code)
							return err
						},
					},
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

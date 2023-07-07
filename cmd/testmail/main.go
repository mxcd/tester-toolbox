package main

import (
	"log"
	"os"

	"github.com/mxcd/go-config/config"
	"github.com/mxcd/testmail/internal/mail"
	"github.com/mxcd/testmail/internal/server"
	"github.com/mxcd/testmail/internal/util"
	"github.com/urfave/cli/v2"
)

func main() {
	initConfig()
	util.InitLogger()

	app := &cli.App{
		Name:        "testmail",
		Description: "Testmail - sending test emails to validate your infrastructure",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "verbose",
				Aliases: []string{"v"},
				Usage:   "debug output",
				EnvVars: []string{"VERBOSE"},
			},
			&cli.BoolFlag{
				Name:    "very-verbose",
				Aliases: []string{"vv"},
				Usage:   "trace output",
				EnvVars: []string{"VERY_VERBOSE"},
			},
		},
		Commands: []*cli.Command{
			{
				Name:        "send",
				Usage:       "testmail send <target address>",
				Description: "Send a single test mail",
				Action: func(c *cli.Context) error {
					return sendSingleMail(c)
				},
			},
			{
				Name:        "serve",
				Usage:       "testmail serve",
				Description: "serve http server with /send/:address endpoint",
				Action: func(c *cli.Context) error {
					server.StartServer()
					return nil
				},
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func sendSingleMail(c *cli.Context) error {
	targetAddress := c.Args().First()
	if targetAddress == "" {
		return cli.Exit("Please provide a target address", 1)
	}
	return mail.SendMail(targetAddress)
}

func initConfig() {
	err := config.LoadConfig([]config.Value{
		config.String("LOG_LEVEL").NotEmpty().Default("info"),

		config.String("SMTP_HOST").NotEmpty().Default("localhost"),
		config.Int("SMTP_PORT").Default(25),
		config.String("SMTP_USERNAME").Default(""),
		config.String("SMTP_PASSWORD").Default("").Sensitive(),
		config.Bool("SMTP_TLS").Default(true),
		config.String("FROM_ADDRESS").NotEmpty(),
		config.Int("PORT").Default(8080),
	})
	if err != nil {
		panic(err)
	}
	config.Print()
}

package main

import (
	"log"
	"os"

	"github.com/mxcd/go-config/config"
	"github.com/mxcd/testtimeout/internal/server"
	"github.com/mxcd/testtimeout/internal/util"
	"github.com/urfave/cli/v2"
)

func main() {
	initConfig()
	util.InitLogger()

	app := &cli.App{
		Name:        "testtimeout",
		Description: "testtimeout - testing connection timeouts",
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
				Name:        "serve",
				Usage:       "testtimeout serve",
				Description: "serve http server with /hold/:duration endpoint",
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

func initConfig() {
	err := config.LoadConfig([]config.Value{
		config.String("LOG_LEVEL").NotEmpty().Default("info"),
		config.Int("PORT").Default(8080),
	})
	if err != nil {
		panic(err)
	}
	config.Print()
}

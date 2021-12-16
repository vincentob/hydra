package template

func init() {
	WebProjectFiles["main.go"] = `package main

import (
	"os"
	"time"

	"github.com/dantin-s/hydra/logger"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"

	"{{ .ProjectPath }}/cmd"
	"{{ .ProjectPath }}/config"
)

func main() {
	app := cli.NewApp()
	app.Name = "{{ .ProjectName }}"
	app.Usage = "{{ .ProjectUsage }}"
	app.Description = "Project {{ .ProjectName }}"
	app.Compiled = time.Now()
	app.Version = "v1.0.0"

	app.Authors = []cli.Author{
		{
			Name:  "Vincent",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "server",
			Usage: "start a web server",
			Action: func(c *cli.Context) error {
				return cmd.Server(c)
			},
		},
		{
			Name:  "migrations",
			Usage: "run db migrations",
			Action: func(c *cli.Context) error {
				return cmd.DoMigrate(c)
			},
		},
	}

	logger.InitLogrusConsoleLogger()

	if lvl, err := logrus.ParseLevel(config.Env.LogLevel); err == nil {
		logrus.SetLevel(lvl)
	}

	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}
`
}

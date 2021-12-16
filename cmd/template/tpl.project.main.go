package template

var (
	CommonProjectFiles = map[string]string{}
	WebProjectFiles    = map[string]string{}
)

func init() {
	CommonProjectFiles["main.go"] = `package main

import (
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"github.com/dantin-s/hydra/logger"

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
			Name:  "default",
			Usage: "default cmd",
			Action: func(c *cli.Context) error {
				return cmd.Default(c)
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

package main

import (
	"os"
	"time"

	"github.com/dantin-s/hydra/cmd"
	_ "github.com/dantin-s/hydra/logger"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Hydra"
	app.Usage = "Tools for golang project building."
	app.Description = "Helper for golang project development"
	app.Compiled = time.Now()
	app.Version = "v0.0.1"

	app.Authors = []cli.Author{
		{
			Name: "Vincent",
		},
	}

	app.Commands = []cli.Command{
		// init golang project
		{
			Name:  "golang",
			Usage: "Hydra init new-golang-project with go.",
			Subcommands: []cli.Command{
				{
					Name: "init",
					Action: func(c *cli.Context) error {
						return cmd.InitGolangProject(c)
					},
				},
			},
		},
		// init golang migrations
		{
			Name:  "migrations",
			Usage: "Hydra init golang migrations with go.",
			Subcommands: []cli.Command{
				{
					Name: "init",
					Action: func(c *cli.Context) error {
						return cmd.InitGolangMigrations(c)
					},
				},
				{
					Name: "add",
					Action: func(c *cli.Context) error {
						return cmd.AddGolangMigrations(c)
					},
				},
			},
		},
		// convert
		{
			Name:  "convert",
			Usage: "convert everything",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  cmd.FlagFrom,
					Value: "bytes",
					Usage: "the raw format",
				},
				cli.StringFlag{
					Name:  cmd.FlagTo,
					Value: "string",
					Usage: "the target format",
				},
				cli.IntFlag{
					Name:  cmd.FlagBase,
					Value: 16,
					Usage: "the int base",
				},
			},
			Action: func(c *cli.Context) error {
				return cmd.Convert(c)
			},
		},
		// whois
		{
			Name:  "whois",
			Usage: "check domain whois info",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  cmd.FlagDomainFile,
					Value: "",
					Usage: "the domain list file, only one domain each line, tha max line is 100,000.",
				},
			},
			Action: func(c *cli.Context) error {
				return cmd.Whois(c)
			},
		},
		// certchk
		{
			Name:  "certchk",
			Usage: "check domain certs",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  cmd.FlagDomainFile,
					Value: "",
					Usage: "the domain list file, only one domain each line, tha max line is 100,000.",
				},
			},
			Action: func(c *cli.Context) error {
				return cmd.CertChk(c)
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}

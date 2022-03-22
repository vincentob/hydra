package template

func init() {
	CommonProjectFiles["cmd/default.go"] = `// Generate By Template
package cmd

import (
	"github.com/urfave/cli"
)

func Default(c *cli.Context) error {
	return nil
}`

	WebProjectFiles["cmd/server.go"] = `// Generate By Template
package cmd

import (
	"github.com/vincentob/hydra/middlewares"
	"github.com/vincentob/hydra/signals"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"

	"{{ .ProjectPath }}/config"
	"{{ .ProjectPath }}/route"
)

func Server(c *cli.Context) error {
	var engine *gin.Engine

	// Set gin release mode.
	if config.Env.Env == config.EnvProduction {
		gin.SetMode(gin.ReleaseMode)
		engine = gin.New()
		engine.Use(middlewares.AccessLog(config.Env.LogLevel))
	} else {
		engine = gin.Default()
	}

	// init router
	route.InitAPIRouter(engine)

	// print env
	config.PrintENV()

	go func() {
		defer close(signals.SigCompleted)
		if err := engine.Run(":80"); err != nil {
			logrus.Fatal(err)
		}
	}()

	signals.WaitForExit()

	return nil
}`

	WebProjectFiles["cmd/migrations.go"] = `// Generate By Template
package cmd

import (
	"github.com/urfave/cli"

	"{{ .ProjectPath }}/config"
	"{{ .ProjectPath }}/migrations"
)

func DoMigrate(c *cli.Context) error {
	return migrations.DoMigrate(config.DB)
}`

}

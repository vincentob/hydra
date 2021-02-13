package template

func init() {
	CommonProjectFiles["cmd/server.go"] = `// Generate By Template
package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"github.com/dantin-s/hydra/middlewares"
	"github.com/dantin-s/hydra/signals"
	"github.com/dantin-s/{{ .ProjectName }}/config"
	"github.com/dantin-s/{{ .ProjectName }}/route
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

	CommonProjectFiles["cmd/migrations.go"] = `// Generate By Template
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

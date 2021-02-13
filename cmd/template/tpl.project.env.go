package template

func init() {
	CommonProjectFiles["config/env.go"] = `package config

import (
	"github.com/caarlos0/env/v6"
	"github.com/sirupsen/logrus"

	"github.com/dantin-s/hydra/json"
	_ "github.com/dantin-s/hydra/logger"
)

const (
	EnvStaging    = "staging"
	EnvProduction = "production"
)

type config struct {
	
}

var Env = &config{}

func init() {
	if err := env.Parse(Env); err != nil {
		logrus.Panic(err)
	}

	PrintENV()
}

func PrintENV() {
	logrus.Infof("\nRunning with configuration: \n%v\n", Env)
}

func (c *config) String() string {
	v, _ := json.FastJJ.MarshalIndent(c, "", "  ")
	return string(v)
}
`
}

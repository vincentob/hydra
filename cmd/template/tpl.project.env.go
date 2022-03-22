package template

func init() {
	CommonProjectFiles["config/env.go"] = "package config\n\n" +
		"import (\n" +
		"	\"github.com/caarlos0/env/v6\"\n" +
		"	\"github.com/vincentob/hydra/json\"\n" +
		"	\"github.com/sirupsen/logrus\"\n" +
		")\n\n" +

		"const (\n" +
		"	EnvStaging    = \"staging\"\n" +
		"	EnvProduction = \"production\"\n" +
		")\n\n" +

		"type config struct {\n" +
		"	Env      string `env:\"ENV\"         envDefault:\"staging\"`\n" +
		"	LogLevel string `env:\"LOG_LEVEL\"   envDefault:\"info\"`\n" +
		"}\n\n" +

		"var Env = &config{}\n\n" +

		"func init() {\n" +
		"	if err := env.Parse(Env); err != nil {\n" +
		"		logrus.Panic(err)\n" +
		"	}\n\n" +

		"	PrintENV()\n" +
		"}\n\n" +

		"func PrintENV() {\n" +
		"	logrus.Infof(\"\\nRunning with configuration: \\n%v\\n\", Env)\n" +
		"}\n\n" +

		"func (c *config) String() string {\n" +
		"	v, _ := json.FastJJ.MarshalIndent(c, \"\", \"  \")\n" +
		"	return string(v)\n" +
		"}\n\n"
}

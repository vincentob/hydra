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
		"	DBUser   string `env:\"DB_USER\"     envDefault:\"root\"`\n" +
		"	DBPwd    string `env:\"DB_PASSWORD\" envDefault:\"test\"`\n" +
		"	DBHost   string `env:\"DB_HOST\"     envDefault:\"127.0.0.1\"`\n" +
		"	DBPort   string `env:\"DB_Port\"     envDefault:\"3306\"`\n" +
		"	DBName   string `env:\"DB_Name\"     envDefault:\"test\"`\n" +
		"}\n\n" +

		"var Env = &config{}\n\n" +

		"func init() {\n" +
		"	if err := env.Parse(Env); err != nil {\n" +
		"		logrus.Panic(err)\n" +
		"	}\n\n" +
		"}\n\n" +

		"func PrintENV() {\n" +
		"	logrus.Infof(\"\\nRunning with configuration: \\n%v\\n\", Env)\n" +
		"}\n\n" +

		"func (c *config) String() string {\n" +
		"	v, _ := json.FastJJ.MarshalIndent(c, \"\", \"  \")\n" +
		"	return string(v)\n" +
		"}\n\n" +

		"func (c *config) isProduction() bool {\n" +
		"	if c.Env == EnvProduction {\n" +
		"		return true\n" +
		"	}\n\n" +
		"	return false\n" +
		"}\n\n"
}

package cmd

import (
	"fmt"
	"os"
	"text/template"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"

	tpl "github.com/vincentob/hydra/cmd/template"
)

var (
	MigrationsDir = "migrations"
)

// InitGolangProject create new golang project with go mod init and urfave/cli framework.
func InitGolangMigrations(c *cli.Context) error {
	if err := os.MkdirAll(MigrationsDir, 0755); err != nil {
		return err
	}

	return initMigrationsTplFiles()
}

func initMigrationsTplFiles() error {
	for k, v := range tpl.MigrationInitTpl {
		filePath := fmt.Sprintf("%s/%s", MigrationsDir, k)

		logrus.Info("Create template file: ", filePath)

		tmpl, err := template.New(k).Parse(v)
		if err != nil {
			return err
		}

		f, err := os.Create(filePath)
		if err != nil {
			return err
		}

		if err := tmpl.Execute(f, nil); err != nil {
			return err
		}

		_ = f.Close()
	}

	return nil
}

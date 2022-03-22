package cmd

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"text/template"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"

	tpl "github.com/vincentob/hydra/cmd/template"
)

var (
	timeFormat = "20060102150405"
)

// AddGolangProject create new golang project with go mod init and urfave/cli framework.
func AddGolangMigrations(c *cli.Context) error {
	if c.NArg() != 1 {
		return errors.New("missing migration file description")
	}

	return addMigrationsTplFiles(c.Args().First())
}

func addMigrationsTplFiles(migrationDesc string) error { // create project dir
	timeFormatStr := time.Now().Format(timeFormat)
	filePath := fmt.Sprintf("%s_%s.go", timeFormatStr, migrationDesc)

	// if not under dir migrations
	if !hasMigrationsDir() {
		if err := os.MkdirAll(MigrationsDir, 0755); err != nil {
			return err
		}
		filePath = fmt.Sprintf("%s/%s_%s.go", MigrationsDir, timeFormatStr, migrationDesc)
	}

	logrus.Info("Create template file: ", filePath)

	tmpl, err := template.New("migration").Parse(tpl.MigrationAddTpl)
	if err != nil {
		return err
	}

	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	return tmpl.Execute(f, timeFormatStr)
}

func hasMigrationsDir() bool {
	return strings.HasSuffix(os.Getenv("PWD"), MigrationsDir)
}

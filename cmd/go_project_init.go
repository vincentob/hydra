package cmd

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"

	tpl "github.com/dantin-s/hydra/cmd/template"
)

var (
	initDir = []string{"/actions", "/cmd", "/config", "/migrations", "/model", "/playground", "/route"}
)

// TemplateData
type TemplateData struct {
	ProjectName  string
	ProjectPath  string
	ProjectUsage string
}

// InitGolangProject create new golang project with go mod init and urfave/cli framework.
func InitGolangProject(c *cli.Context) error {
	if len(c.Args()) != 1 {
		return errors.New("too many parameters")
	}

	projectPath := c.Args().First()
	projectPathSplit := strings.Split(projectPath, "/")

	if len(projectPathSplit) != 3 {
		return errors.New("Project path should like 'github.com/user/projectName' ")
	}

	if !strings.HasPrefix(projectPathSplit[0], "gitlab") && !strings.HasPrefix(projectPathSplit[0], "github") {
		return errors.New("project path should start with 'gitlab' or 'github'")
	}

	data := &TemplateData{
		ProjectName: projectPathSplit[2],
		ProjectPath: projectPath,
	}

	return createProjectTplFiles(data)
}

func createProjectTplFiles(data *TemplateData) error {
	if _, err := os.Stat(data.ProjectName); err == nil {
		return fmt.Errorf("Project dir exist: %s ", data.ProjectName)
	}

	// create project dir
	for _, dir := range initDir {
		if err := os.MkdirAll(data.ProjectName+dir, 0755); err != nil {
			return err
		}
	}

	// generate template files
	for k, v := range tpl.ProjectTpl {
		filePath := fmt.Sprintf("%s/%s", data.ProjectName, k)

		logrus.Info("Create template file: ", filePath)

		tmpl, err := template.New(k).Parse(v)
		if err != nil {
			return err
		}

		f, err := os.Create(filePath)
		if err != nil {
			return err
		}

		if err := tmpl.Execute(f, data); err != nil {
			return err
		}

		_ = f.Close()
	}

	// generate migrations template files
	MigrationsDir = fmt.Sprintf("%s/%s", data.ProjectName, MigrationsDir)
	if err := initMigrationsTplFiles(); err != nil {
		return err
	}

	logrus.Info("Generate template files done. You should now run 'go mod tidy' to gen the go.sum file.")

	return nil
}

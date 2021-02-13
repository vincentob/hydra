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
	commonDir = []string{"/cmd", "/config", "/playground"}
	webDir    = []string{"/actions", "/migrations", "/model", "/route"}
)

// TemplateData
type TemplateData struct {
	ProjectName  string
	ProjectPath  string
	ProjectUsage string
}

const (
	ProjectTypeWeb = "web"
	ProjectTypeCMD = "cmd"
)

func InitGolangWebProject(c *cli.Context) error {
	return initGolangProject(c, ProjectTypeWeb)
}

func InitGolangCMDProject(c *cli.Context) error {
	return initGolangProject(c, ProjectTypeCMD)
}

// initGolangProject create new golang project with go mod init and urfave/cli framework.
func initGolangProject(c *cli.Context, projectType string) error {
	if len(c.Args()) != 1 {
		return errors.New("too many parameters")
	}

	projectPath := c.Args().First()
	projectPathSplit := strings.Split(projectPath, "/")

	if len(projectPathSplit) != 3 {
		return errors.New("Project path should like 'github.com/user/projectName' ")
	}

	if !strings.HasPrefix(projectPathSplit[0], "gitlab") &&
		!strings.HasPrefix(projectPathSplit[0], "github") {
		return errors.New("project path should start with 'gitlab' or 'github'")
	}

	data := &TemplateData{
		ProjectName: projectPathSplit[2],
		ProjectPath: projectPath,
	}

	return createProjectTplFiles(data, projectType)
}

func createProjectTplFiles(data *TemplateData, projectType string) error {
	if _, err := os.Stat(data.ProjectName); err == nil {
		return fmt.Errorf("Project dir exist: %s ", data.ProjectName)
	}

	dirs := commonDir
	tpls := tpl.CommonProjectFiles

	if projectType == ProjectTypeWeb {
		dirs = append(commonDir, webDir...)

		for k, v := range tpl.WebProjectFiles {
			tpls[k] = v
		}
	}

	for _, dir := range dirs {
		if err := os.MkdirAll(data.ProjectName+dir, 0755); err != nil {
			return err
		}
	}

	// generate template files
	for k, v := range tpls {
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

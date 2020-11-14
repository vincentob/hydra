package template

func init() {
	ProjectTpl["config/gorm.go"] = `// Generate By Template
package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func init() {
	var err error

	DB, err = gorm.Open("mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			Env.DBUser,
			Env.DBPwd,
			Env.DBHost,
			Env.DBPort,
			Env.DBName))

	if err != nil {
		logrus.Error(errors.Wrap(err, "open db failed"))
	}
}
`
}

package template

func init() {
	WebProjectFiles["config/gorm.go"] = `// Generate By Template
package config

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	var err error

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		Env.DBUser,
		Env.DBPwd,
		Env.DBHost,
		Env.DBPort,
		Env.DBName)

	DB, err = gorm.Open(mysql.New(mysql.Config{
		DSN:               dsn,
		DefaultStringSize: 255,
	}), &gorm.Config{})

	if err != nil {
		logrus.Error(errors.Wrap(err, "open db failed"))
	}
}
`
}

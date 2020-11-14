package template

func init() {
	ProjectTpl["actions/actions.go"] = `// Generate By Template
package actions

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Data struct{}

func Confit(c *gin.Context) {
	reqData := &Data{}
	if err := c.ShouldBindJSON(reqData); err != nil {
		logrus.Error(err)
	}

	logrus.Info(reqData)
}
`
}

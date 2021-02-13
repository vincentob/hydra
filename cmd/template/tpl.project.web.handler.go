package template

func init() {
	WebProjectFiles["handler/ping.go"] = `// Generate By Template
package handler

import (
	"github.com/dantin-s/hydra/formutil"
	"github.com/gin-gonic/gin"
)

type PingForm struct{}

func (f *PingForm) Handle(c *gin.Context) (interface{}, error){
	c.Writer.Write([]byte("pong"))
	return nil, nil
}
`
}

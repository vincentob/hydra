package template

func init() {
	WebProjectFiles["api/ping.go"] = `// Generate By Template
package api

import (
	"github.com/dantin-s/hydra/formutil"
	"github.com/gin-gonic/gin"
)

func HandlePing(c *gin.Context) {
	formutil.Handle(&handler.PingForm{}, c)
}
`
}

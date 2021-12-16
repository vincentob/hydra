package template

func init() {
	WebProjectFiles["route/route.go"] = `// Generate By Template
package route

import (
	"github.com/gin-gonic/gin"

	"{{ .ProjectPath }}/api"
)

// InitAPIRouter register api path.
func InitAPIRouter(engine *gin.Engine) {
	engine.GET("/ping", api.Ping)
}
`
}

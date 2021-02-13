package template

func init() {
	WebProjectFiles["route/route.go"] = `// Generate By Template
package route

import (
	"github.com/gin-gonic/gin"
	"github.com/dantin-s/{{ .ProjectName }}/actions"
)

// InitAPIRouter register api path.
func InitAPIRouter(engine *gin.Engine) {
	api := engine.Group("/api/v1")
	{
		api.POST("/model", actions.Save)
	}
}

`
}

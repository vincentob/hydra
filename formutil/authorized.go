package formutil

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

const (
	DefaultUserKey = "username"
)

// Authorization check user session
func Authorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		if sessions.Default(c).Get(DefaultUserKey) == nil {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}

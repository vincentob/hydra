package formutil

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authorized(c *gin.Context) {
	if _, ok := GetSessionID(c); !ok {
		c.JSON(http.StatusUnauthorized, "Unauthorized")
		c.Abort()
	}
}

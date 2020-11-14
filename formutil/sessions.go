package formutil

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

const (
	DefaultSessionKey = "sessionid"
)

func GetSessionID(c *gin.Context) (int, bool) {
	id, ok := sessions.Default(c).Get(DefaultSessionKey).(int)
	return id, ok
}

func GetSession(c *gin.Context) sessions.Session {
	return sessions.Default(c)
}

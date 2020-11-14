package formutil

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

const (
	DefaultSessionKey = "sessionid"
)

func Sessions(store sessions.Store) gin.HandlerFunc {
	return sessions.Sessions(DefaultSessionKey, store)
}

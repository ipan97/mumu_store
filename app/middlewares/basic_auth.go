package middlewares

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type BasicAuth struct {
}

func NewBasicAuth() *BasicAuth {
	return &BasicAuth{}
}

func (basicAuth *BasicAuth) Authentication(ctx *gin.Context) {
	session := sessions.Default(ctx)
	user := session.Get("user")
	if user == nil {
		ctx.Redirect(http.StatusUnauthorized, "/login")
		return
	}
	ctx.Next()
}

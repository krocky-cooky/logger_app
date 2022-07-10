package model 

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	"fmt"
)

type SessionInfo struct {
	Username string 
}

func SessionSetup(r *gin.Engine) {
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("logger_app",store))
}

func SessionCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		username := session.Get("username")

		if username == nil {
			fmt.Print("not logged in")
			c.Redirect(302,"/login")
			c.Abort()
		} else {
			c.Set("username",username)
			c.Next()
			fmt.Print("logged in")
		}
	}
}
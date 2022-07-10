package controller 

import (
	"github.com/krocky-cooky/logger_app/model"
	"github.com/krocky-cooky/logger_app/crypto"
	"github.com/gin-contrib/sessions"

	"net/http"
	"fmt"
	"github.com/gin-gonic/gin"
)

func IndexPageGet(c *gin.Context) {
	c.HTML(200,"index.html",gin.H{})
}


func SignUpGet(c *gin.Context) {
	c.HTML(200,"signup.html",gin.H{})
}

func SignUpPost(c *gin.Context) {
	var form model.User

	if err := c.Bind(&form); err != nil {
		c.HTML(http.StatusBadRequest, "signup.html",gin.H{"err":err})
		c.Abort()

	}else {
		username := c.PostForm("username")
		password := c.PostForm("password")

		if err := model.CreateUser(username,password);err != nil {
			c.HTML(http.StatusBadRequest,"signup.html",gin.H{"err":err})
		}
		c.Redirect(302,"/")
	}
}

func LoginGet(c *gin.Context) {
	c.HTML(200,"login.html",gin.H{})
} 

func LoginPost(c *gin.Context) {
	username := c.PostForm("username")
	user := model.GetUser(username)
	dbPass := user.Password 
	formPass := c.PostForm("password")

	if err := crypto.CompareHashAndPassword(dbPass, formPass); err != nil {
		fmt.Print("login failed\n")
		c.HTML(http.StatusBadRequest, "login.html", gin.H{"err": err})
		c.Abort()
	} else {
		fmt.Print("successfully logined")
		session := sessions.Default(c)
		session.Set("username",username)
		session.Save()
		c.Redirect(302, "/app")
	}
}


func LogoutGet(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.Redirect(302,"/")
}

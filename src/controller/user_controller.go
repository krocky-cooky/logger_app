package controller 

import (
	"github.com/krocky-cooky/logger_app/model"
	"github.com/krocky-cooky/logger_app/crypto"
	"github.com/gin-contrib/sessions"

	"net/http"
	"fmt"
	"strconv"
	_"reflect"
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
		c.Redirect(302,"/app/create")
	}
}

func LoginGet(c *gin.Context) {
	c.HTML(200,"login.html",gin.H{})
} 

func LoginPost(c *gin.Context) {
	username := c.PostForm("username")
	user, err := model.GetUser(username);
	if  err != nil {
		fmt.Print("login failed\n")
		c.HTML(http.StatusBadRequest, "login.html", gin.H{"err": err})
		c.Abort()
	}
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

func LogCreateGet(c *gin.Context) {
	session := sessions.Default(c)
	username := session.Get("username").(string)
	logs, _ := model.GetLogs(username)
	/*
	var lognames []string
	for _, log := range logs {
		lognames = append(lognames,log.Logname)
	}
	*/
	c.HTML(200,"log_create.html",gin.H{
		"logs": logs,
	})
}

func LogCreatePost(c *gin.Context) {
	session := sessions.Default(c)
	username:= session.Get("username").(string)
	
	logname := c.PostForm("logname")
	err := model.CreateLog(logname,username)
	if err != nil {
		fmt.Print("fail registering\n")
		c.HTML(http.StatusBadRequest, "log_create.html",gin.H{})
	} else {
		fmt.Print("successfully registered\n")
		c.Redirect(302,"/app/create")
	}
}

func LogRegisterPost(c *gin.Context) {
	username := c.Param("username")
	logid64, _ := strconv.ParseUint(c.Param("logid"),10,64)
	logid := uint(logid64)
	guid := c.Param("strid")
	postData := c.PostForm("data")
	data, _ := strconv.ParseFloat(postData,64)
	fmt.Print("API requests received\n")
	err := model.RegisterLog(username,logid,guid,data)

	if err != nil {
		fmt.Print("fail registering\n")
	}
}

func LogViewGet(c *gin.Context) {
	session := sessions.Default(c)
	username := session.Get("username").(string)

	logid64, _ := strconv.ParseUint(c.Param("logid"),10,64)
	logid := uint(logid64)
	
	logdatas := model.GetLogDatas(username,logid)
	guid := model.GetLogGuid(username,logid)

	
	c.HTML(200,"log_view.html",gin.H{
		"logdatas": logdatas,
		"guid": guid,
		"username": username,
		"logid": logid,
	})

}

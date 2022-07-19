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
	c.HTML(200,"index.html",gin.H{
		"base": model.GetBaseTemplateData(c),
	})
}


func SignUpGet(c *gin.Context) {
	users, _ := model.GetAllUser()
	c.HTML(200,"signup.html",gin.H{
		"base": model.GetBaseTemplateData(c),
		"users": users,
	})
}

func SignUpPost(c *gin.Context) {
	var form model.User

	if err := c.Bind(&form); err != nil {
		c.HTML(http.StatusBadRequest, "signup.html",gin.H{
			"err":err,
			"base": model.GetBaseTemplateData(c),
		})
		c.Abort()

	}else {
		username := c.PostForm("username")
		password := c.PostForm("password")

		if err := model.CreateUser(username,password);err != nil {
			c.HTML(http.StatusBadRequest,"signup.html",gin.H{
				"err":err,
				"base": model.GetBaseTemplateData(c),
	})
		}
		c.Redirect(302,"/app/create")
	}
}

func LoginGet(c *gin.Context) {
	c.HTML(200,"login.html",gin.H{
		"base": model.GetBaseTemplateData(c),
	})
} 

func LoginPost(c *gin.Context) {
	username := c.PostForm("username")
	user, err := model.GetUser(username);
	if  err != nil {
		fmt.Print("login failed\n")
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"err": err,
			"base": model.GetBaseTemplateData(c),
		})
		c.Abort()
	}
	dbPass := user.Password 
	formPass := c.PostForm("password")

	if err := crypto.CompareHashAndPassword(dbPass, formPass); err != nil {
		fmt.Print("login failed\n")
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"err": err,
			"base": model.GetBaseTemplateData(c),
		})
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
	c.Redirect(302,"/login")
}

func LogCreateGet(c *gin.Context) {
	session := sessions.Default(c)
	usernamesession := session.Get("username")
	if usernamesession == nil {
		fmt.Print("not logged in\n")
		c.Redirect(302,"/login")
		c.Abort()
		return 
	}
	db := model.GormConnect()
	db_ret, err := db.DB()
	if err != nil {
		panic(err)
	}
	defer db_ret.Close()
	username := usernamesession.(string)
	logs, _ := model.GetLogs(username, db)
	/*
	var lognames []string
	for _, log := range logs {
		lognames = append(lognames,log.Logname)
	}
	*/
	c.HTML(200,"log_create.html",gin.H{
		"logs": logs,
		"base": model.GetBaseTemplateData(c),
	})
}

func LogCreatePost(c *gin.Context) {
	session := sessions.Default(c)
	username:= session.Get("username").(string)
	
	logname := c.PostForm("logname")
	var varNames []string
	var varTypes []int
	for i := 0;i < 2; i++ {
		name := c.PostForm("variable_" + strconv.Itoa(i))
		typeidstr := c.PostForm("type_" + strconv.Itoa(i))
		/*
		if reflect.ValueOf(name).IsNil() || reflect.ValueOf(typeidstr).IsNil(){
			break
		}
		*/
		typeid, _ := strconv.Atoi(typeidstr)
		varNames = append(varNames, name)
		varTypes = append(varTypes,typeid)
	}
	db := model.GormConnect()
	db_ret, err := db.DB()
	if err != nil {
		panic(err)
	}
	defer db_ret.Close()
	err = model.CreateLog(logname,username,varNames,varTypes, db)
	if err != nil {
		fmt.Print("fail registering\n")
		c.HTML(http.StatusBadRequest, "log_create.html",gin.H{
			"base": model.GetBaseTemplateData(c),
		})
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
	
	fmt.Print("API requests received\n")
	
	db := model.GormConnect()
	db_ret, err := db.DB()
	if err != nil {
		panic(err)
	}
	defer db_ret.Close()
	variables, _ := model.GetVariables(logid,db)
	postdata := map[string]string{} 
	fmt.Print(len(variables))
	fmt.Print("\n")
	for _, variable := range variables {
		postdata[variable.Name] = c.PostForm(variable.Name)
	}
	err = model.RegisterLog(username,logid,guid,postdata,db)

	if err != nil {
		fmt.Print("fail registering\n")
	}
}

func LogViewGet(c *gin.Context) {
	session := sessions.Default(c)
	username := session.Get("username").(string)
	
	logid64, _ := strconv.ParseUint(c.Param("logid"),10,64)
	logid := uint(logid64)
	
	db := model.GormConnect()
	db_ret, err := db.DB()
	if err != nil {
		panic(err)
	}
	defer db_ret.Close()
	log, _ := model.GetLog(logid, db)

	mapVariables := map[string][]string{}

	for _, dataset := range log.LogDataSets {
		logDataValues, _ := model.GetLogDataValues(dataset.ID,db)

		for _, logDataValue := range logDataValues {
			fmt.Print(logDataValue.Data)
			variable, _ := model.GetVariable(logDataValue.VariableID,db)
			mapVariables[variable.Name] = append(mapVariables[variable.Name],logDataValue.Data)
		}
	}

	guid := model.GetLogGuid(username,logid,db)

	
	c.HTML(200,"log_view.html",gin.H{
		"base": model.GetBaseTemplateData(c),
		"logdatas": mapVariables,
		"guid": guid,
		"username": username,
		"logid": logid,
	})

}

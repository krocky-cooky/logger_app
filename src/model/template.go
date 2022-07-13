package model 

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	_"fmt"
)

type BaseTemplateData struct {
	Is_authenticated bool 
	Username string 
}

func GetBaseTemplateData(c *gin.Context) BaseTemplateData {
	var tempdata BaseTemplateData
	session := sessions.Default(c)
	usernamesession := session.Get("username")
	if usernamesession == nil {
		tempdata.Is_authenticated = false
		tempdata.Username = "not authenticated"
		return tempdata
	}
	username := usernamesession.(string)
	_, err := GetUser(username)

	if err != nil {
		tempdata.Is_authenticated = false
		tempdata.Username = "not authenticated"
	} else {
		tempdata.Is_authenticated = true 
		tempdata.Username = username
	}

	return tempdata

}
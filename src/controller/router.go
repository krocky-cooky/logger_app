package controller

import (
	"github.com/krocky-cooky/logger_app/model"

	_"github.com/gin-contrib/sessions"
    _"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	r := gin.Default()
	model.SessionSetup(r)
	r.LoadHTMLGlob("view/*.html")

	r.GET("/login",LoginGet)
	r.POST("/login",LoginPost)

	r.GET("/signup",SignUpGet)
	r.POST("/signup",SignUpPost)

	r_in := r.Group("app/")
	{
		r_in.Use(model.SessionCheck())

		//root
		r_in.GET("/",IndexPageGet)

		//signup
		
		
		//login 
		

	}
	


	return r
}
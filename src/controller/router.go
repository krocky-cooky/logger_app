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
	r.Static("/css","./view/css")
	r.Static("/js","./view/js")


	//login
	r.GET("/login",LoginGet)
	r.POST("/login",LoginPost)

	


	//signup
	r.GET("/signup",SignUpGet)
	r.POST("/signup",SignUpPost)

	r.GET("/logout",LogoutGet)

	r_in := r.Group("app/")
	{
		r_in.Use(model.SessionCheck())

		//root
		r_in.GET("/",IndexPageGet)

		r_in.GET("/create",LogCreateGet)
		r_in.POST("/create",LogCreatePost)

		r_in.GET("/view/:logid",LogViewGet)

		

	}

	r_api := r.Group("api/")
	{
		r_api.POST("/:username/:logid/:strid",LogRegisterPost)
	}
	



	return r
}
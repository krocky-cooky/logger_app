package main

import (
	"github.com/krocky-cooky/logger_app/controller"
	"github.com/krocky-cooky/logger_app/model"
	_"github.com/gin-gonic/gin"
)

func main() {
	model.DbInit()
	router := controller.GetRouter()
	router.Run()
}
package app

import (
	"miniewallet/api/controller"
	"miniewallet/api/middlewares"
)

func route() {
	router.GET("/", controller.Index)
	router.POST("/register", controller.CreateUser)
	router.POST("/login", controller.Login)
	router.POST("/logout", middlewares.TokenAuthMiddleware(), controller.LogOut)
}

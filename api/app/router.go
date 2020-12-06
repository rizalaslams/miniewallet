package app

import (
	"miniewallet/api/controller"
)

func route() {
	router.GET("/", controller.Index)
}

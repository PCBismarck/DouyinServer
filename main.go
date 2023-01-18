package main

import (
	"github.com/PCBismarck/DouyinServer/service"
	"github.com/PCBismarck/DouyinServer/toolkit"
	"github.com/gin-gonic/gin"
)

func main() {
	toolkit.InitDB()
	go service.RunMessageServer()

	r := gin.Default()

	initRouter(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

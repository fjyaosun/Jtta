package main

import (
	"Jtta/api"

	"github.com/gin-gonic/gin"
)

func main() {
	//连接数据库，获取本身所需的配置
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	api.ApiRegister(router)
	router.Run()
}

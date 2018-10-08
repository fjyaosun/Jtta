package api

import (
	"T-future/midd"

	"github.com/gin-gonic/gin"
)

const accountCode = 10000
const configsCode = 11000

func ApiRegister(router *gin.Engine) {
	//注册接口
	web_api := router.Group("api")
	{
		//不用登陆的接口
		web_api.POST("/account/login", loginUser)
		//需要登陆的接口
		web_api.Use(midd.Auth())
		web_api.GET("/center/info", userInfo)
		web_api.POST("/account/logout", logoutUser)

		web_api.GET("/configs/get", getData)
		web_api.POST("/configs/set", setData)
		web_api.DELETE("/configs/del", delData)
	}
}

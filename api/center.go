package api

import (
	"T-future/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func userInfo(c *gin.Context) {
	value := c.MustGet("User")
	user := value.(*model.User)
	c.JSON(
		http.StatusOK,
		gin.H{
			"status": 1,
			"data":   gin.H{"userId": user.Id, "token": user.Token},
		},
	)
}

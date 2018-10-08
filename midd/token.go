package midd

import (
	"Jtta/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func TokenAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token == "" {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{"status": 0, "msg": "Token Unauthorized"},
			)
			return
		}

		err, user := model.GetUserByToken(token)
		if err != nil {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{"status": 0, "msg": "Token Unauthorized"},
			)
			return
		}
		c.Set("User", user)
		return
	}
}

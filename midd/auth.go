package midd

import (
	"T-future/model"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionId, err := c.Cookie("sessionId")
		if err != nil {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{"status": 0, "msg": "Unauthorized"},
			)
		}

		err, session := model.GetSession(bson.ObjectIdHex(sessionId))
		if err != nil {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{"status": 0, "msg": "Unauthorized"},
			)
			return
		}
		uerr, user := model.GetUser(session.UserId)
		if uerr != nil {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{"status": 0, "msg": "Unauthorized"},
			)
			return
		}
		c.Set("User", user)
	}
}

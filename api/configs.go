package api

import (
	"T-future/model"
	"fmt"
	"net/http"

	mgo "gopkg.in/mgo.v2"

	"github.com/gin-gonic/gin"
)

var configsErrorCode map[string]map[string]interface{} = map[string]map[string]interface{}{
	"paramsError": {
		"status": 0,
		"code":   configsCode + 0,
		"msg":    "params error",
	},
	"serverError": {
		"status": 0,
		"code":   configsCode + 1,
		"msg":    "server error",
	},
}

func setData(c *gin.Context) {
	v := c.MustGet("User")
	user := v.(*model.User)
	key := c.PostForm("key")
	value := c.PostForm("value")
	if key == "" {
		c.JSON(http.StatusForbidden, configsErrorCode["paramsError"])
		return
	}
	err := model.UpdateConfigsByUser(user.Id, key, value)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, configsErrorCode["serverError"])
		return
	}
	c.SecureJSON(http.StatusOK, gin.H{"status": 1})
}

func getData(c *gin.Context) {
	value := c.MustGet("User")
	user := value.(*model.User)
	err, configs := model.GetConfigsByUser(user.Id)
	if err == mgo.ErrNotFound {
		c.SecureJSON(http.StatusOK, gin.H{"configs": gin.H{}, "status": 1})
		return
	}
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, configsErrorCode["serverError"])
		return
	}
	c.SecureJSON(http.StatusOK, gin.H{"status": 1, "data": configs.Data})
}

func delData(c *gin.Context) {
	value := c.MustGet("User")
	user := value.(*model.User)
	key := c.PostForm("key")
	if key == "" {
		c.JSON(http.StatusForbidden, configsErrorCode["paramsError"])
		return
	}
	err := model.DelConfigsByUser(user.Id, key)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, configsErrorCode["serverError"])
		return
	}
	c.SecureJSON(http.StatusOK, gin.H{"status": 1})
}

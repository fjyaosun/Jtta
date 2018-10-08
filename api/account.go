package api

import (
	"T-future/logic"
	"T-future/midd"
	"T-future/model"
	"fmt"
	"net/http"

	mgo "gopkg.in/mgo.v2"

	"github.com/gin-gonic/gin"
)

var accountErrorCode map[string]map[string]interface{} = map[string]map[string]interface{}{
	"paramsError": {
		"status": 0,
		"code":   accountCode + 0,
		"msg":    "params error",
	},
	"serverError": {
		"status": 0,
		"code":   accountCode + 1,
		"msg":    "server error",
	},
	"userExist": {
		"status": 0,
		"code":   accountCode + 2,
		"msg":    "user exist",
	},
	"userError": {
		"status": 0,
		"code":   accountCode + 3,
		"msg":    "user or password error",
	},
}

func registerUser(c *gin.Context) {
	//注册用户
	//panic("error test")
	phone := c.PostForm("phone")
	password := c.PostForm("password")
	name := c.PostForm("name")
	email := c.PostForm("email")
	genderStr := c.PostForm("gender")
	//判断参数
	if logic.CheckPhone(phone) == false {
		c.JSON(http.StatusBadRequest, accountErrorCode["paramsError"])
		fmt.Println("phone")
		return
	}
	if password == "" {
		c.JSON(http.StatusBadRequest, accountErrorCode["paramsError"])
		fmt.Println("password")
		return
	}
	if name == "" {
		c.JSON(http.StatusBadRequest, accountErrorCode["paramsError"])
		fmt.Println("name")
		return
	}
	if logic.CheckEmail(email) == false {
		c.JSON(http.StatusBadRequest, accountErrorCode["paramsError"])
		fmt.Println("email")
		return
	}

	gender, ok := model.GenderMap[genderStr]
	if !ok {
		c.JSON(http.StatusBadRequest, accountErrorCode["paramsError"])
		fmt.Println("gender")
		return
	}

	// 判断用户是否已经存在
	uerr, _ := model.GetUserByPhone(phone)
	if uerr != mgo.ErrNotFound {
		c.JSON(http.StatusForbidden, accountErrorCode["userExist"])
		return
	}

	err, _ := logic.CreateUser(phone, password, name, email, gender)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, accountErrorCode["serverError"])
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": 1})
}

func loginUser(c *gin.Context) {
	//校验用户密码是否正确
	phone := c.PostForm("phone")
	password := c.PostForm("password")

	uerr, user := model.GetUserByPhone(phone)
	if uerr != nil {
		c.JSON(http.StatusForbidden, accountErrorCode["userError"])
		return
	}
	if user.Password != password {
		c.JSON(http.StatusForbidden, accountErrorCode["userError"])
		return
	}

	//创建新session
	cerr, configs := model.GetSysConfigs()
	if cerr != nil {
		fmt.Println(cerr.Error())
		c.JSON(http.StatusInternalServerError, accountErrorCode["serverError"])
		return
	}
	merr, sessionId := midd.CreateSession(&user.Id)
	if merr != nil {
		fmt.Println(merr.Error())
		c.JSON(http.StatusInternalServerError, accountErrorCode["serverError"])
		return
	}
	c.SetCookie("sessionId", sessionId, 3600*24*30, "/", configs.Domain, false, true)
	c.SecureJSON(http.StatusOK, gin.H{"status": 1})
}

func logoutUser(c *gin.Context) {
	sessionId, _ := c.Cookie("sessionId")
	_ = midd.ClearSession(sessionId)
	c.SetCookie("sessionId", "", -1, "/", "", false, true)
	c.SecureJSON(http.StatusOK, gin.H{"status": 1})
}

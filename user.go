package logic

import (
	"T-future/model"
	"strconv"
	"time"

	"gopkg.in/mgo.v2/bson"
)

func CreateUser(phone, password, name, email string, gender int) (
	err error, userId bson.ObjectId) {
	userId = bson.NewObjectId()
	var user = &model.User{
		Id:           userId,
		Phone:        phone,
		Password:     password,
		Name:         name,
		Gender:       gender,
		Email:        email,
		RegisterTime: time.Now(),
		Token:        bson.NewObjectId().Hex(),
	}
	err = model.CreateUser(user)
	return
}

func CheckPhone(phone string) bool {
	_, err := strconv.ParseInt(phone, 10, 64)
	if err != nil {
		return false
	}
	return true
}

func CheckEmail(email string) bool {
	if email == "" {
		return false
	}
	return true
}

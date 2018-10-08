package model

import (
	"Jtta/dbs/mongodb"
	"time"

	"gopkg.in/mgo.v2/bson"
)

var UserConfigsColName = "userConfigs"

type userConfigsField struct {
	Id         string
	UserId     string
	CreateTime string
	UpdateTime string
	Data       string
}

var UserConfigsField = userConfigsField{
	Id:         "_id",
	UserId:     "userId",
	CreateTime: "createTime",
	UpdateTime: "updateTime",
	Data:       "data",
}

type UserConfigs struct {
	Id         bson.ObjectId     `bson:"_id"`
	UserId     bson.ObjectId     `bson:"userId"`
	CreateTime time.Time         `bson:"createTime"`
	UpdateTime time.Time         `bson:"updateTime"`
	Data       map[string]string `bson: "data"`
}

func GetConfigsByUser(userId bson.ObjectId) (err error, configs *UserConfigs) {
	mgoSession := mongodb.GetMongoSession()
	defer mgoSession.Close()

	collection := mgoSession.DB(mongodb.MongoConf.DB).C(UserConfigsColName)
	configs = new(UserConfigs)
	err = collection.Find(bson.M{UserConfigsField.UserId: userId}).One(&configs)
	return
}

func UpdateConfigsByUser(userId bson.ObjectId, key string, value string) (err error) {
	mgoSession := mongodb.GetMongoSession()
	defer mgoSession.Close()

	now := time.Now()
	utc, _ := time.LoadLocation("UTC")
	collection := mgoSession.DB(mongodb.MongoConf.DB).C(UserConfigsColName)
	_, err = collection.Upsert(
		bson.M{UserConfigsField.UserId: userId},
		bson.M{
			"$set": bson.M{
				UserConfigsField.Data + "." + key: value,
				UserConfigsField.UpdateTime:       now.In(utc),
			},
			"$setOnInsert": bson.M{UserConfigsField.CreateTime: now.In(utc)},
		},
	)
	return
}

func DelConfigsByUser(userId bson.ObjectId, key string) (err error) {
	mgoSession := mongodb.GetMongoSession()
	defer mgoSession.Close()

	collection := mgoSession.DB(mongodb.MongoConf.DB).C(UserConfigsColName)
	now := time.Now()
	utc, _ := time.LoadLocation("UTC")
	_, err = collection.Upsert(
		bson.M{UserConfigsField.UserId: userId},
		bson.M{
			"$set":   bson.M{UserConfigsField.UpdateTime: now.In(utc)},
			"$unset": bson.M{UserConfigsField.Data + "." + key: 1},
		},
	)
	return
}

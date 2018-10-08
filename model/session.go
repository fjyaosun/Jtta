package model

import (
	"Jtta/dbs/mongodb"
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Session struct {
	Id         bson.ObjectId `bson:"_id"`
	Type       string        `bson:"name"`
	CreateTime time.Time     `bson:"createTime"`
	UserId     bson.ObjectId `bson:"userId"`
}

func GetSession(sessionId bson.ObjectId) (err error, session *Session) {
	mgoSession := mongodb.GetMongoSession()
	defer mgoSession.Close()

	collection := mgoSession.DB(mongodb.MongoConf.DB).C("session")
	session = new(Session)
	err = collection.FindId(sessionId).One(&session)
	return
}

func CreateSession(session *Session) (err error) {
	mgoSession := mongodb.GetMongoSession()
	defer mgoSession.Close()
	collection := mgoSession.DB(mongodb.MongoConf.DB).C("session")
	err = collection.Insert(session)
	return
}

func CleanSession(sessionId bson.ObjectId) (err error) {
	mgoSession := mongodb.GetMongoSession()
	defer mgoSession.Close()

	collection := mgoSession.DB(mongodb.MongoConf.DB).C("session")
	err = collection.RemoveId(sessionId)
	return
}

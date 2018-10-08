package model

import (
	"Jtta/dbs/mongodb"
	"time"

	"gopkg.in/mgo.v2/bson"
)

var UserColName = "user"

type userField struct {
	Id            string
	Code          string
	Name          string
	AvatarUrl     string
	Country       string
	Province      string
	City          string
	Language      string
	Gender        string
	RegisterTime  string
	UpdateTime    string
	PhoneVerified string
	EmailVerified string
}

var UserField = userField{
	Id:            "_id",
	Code:          "code",
	Name:          "name",
	AvatarUrl:     "AvatarUrl",
	Country:       "country",
	Province:      "province",
	City:          "city",
	Language:      "language",
	Gender:        "gender",
	RegisterTime:  "registerTime",
	UpdateTime:    "updateTime",
	PhoneVerified: "phoneVerified",
	EmailVerified: "emailVerified",
}

type User struct {
	Id            bson.ObjectId `bson:"_id"`
	Code          string        `bson:"code"`
	Name          string        `bson:"name"`
	AvatarUrl     string        `bson:"avatarUrl"`
	Country       string        `bson:"country"`
	Province      string        `bson:"province"`
	City          string        `bson:"city"`
	Language      string        `bson:"language"`
	Gender        int           `bson:"gender"`
	RegisterTime  time.Time     `bson:"registerTime"`
	UpdateTime    time.Time     `bson:"updateTime"`
	PhoneVerified bool          `bson:"phoneVerified"`
	EmailVerified bool          `bson:"emailVerified"`
}

func CreateUser(user *User) (err error) {
	mgoSession := mongodb.GetMongoSession()
	defer mgoSession.Close()
	collection := mgoSession.DB(mongodb.MongoConf.DB).C(UserColName)
	err = collection.Insert(user)
	return
}

func GetUserByCode(code string) (err error, user *User) {
	mgoSession := mongodb.GetMongoSession()
	defer mgoSession.Close()

	collection := mgoSession.DB(mongodb.MongoConf.DB).C(UserColName)
	user = new(User)
	err = collection.Find(bson.M{UserField.Code: code}).One(&user)
	return
}

func GetUser(userId bson.ObjectId) (err error, user *User) {
	mgoSession := mongodb.GetMongoSession()
	defer mgoSession.Close()

	collection := mgoSession.DB(mongodb.MongoConf.DB).C(UserColName)
	user = new(User)
	err = collection.FindId(userId).One(&user)
	return
}

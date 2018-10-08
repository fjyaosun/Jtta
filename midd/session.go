package midd

import (
	"T-future/model"
	"time"

	"gopkg.in/mgo.v2/bson"
)

func CreateSession(userId *bson.ObjectId) (err error, sessionId string) {
	var session = &model.Session{
		Id:         bson.NewObjectId(),
		Type:       "normal",
		CreateTime: time.Now(),
		UserId:     *userId,
	}
	err = model.CreateSession(session)
	if err != nil {
		return err, ""
	}
	return err, session.Id.Hex()
}

func ClearSession(sessionId string) (err error) {
	err = model.CleanSession(bson.ObjectIdHex(sessionId))
	return
}

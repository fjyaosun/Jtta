package model

import (
	"Jtta/dbs/mongodb"

	"gopkg.in/mgo.v2/bson"
)

type SysConfigs struct {
	Id     bson.ObjectId `bson:"_id"`
	Domain string        `bson:"domain"`
	Port   int           `bson: "port"`
}

func GetSysConfigs() (err error, configs *SysConfigs) {

	mgoSession := mongodb.GetMongoSession()
	defer mgoSession.Close()

	collection := mgoSession.DB(mongodb.MongoConf.DB).C("sysConfigs")
	configs = new(SysConfigs)
	err = collection.Find(nil).One(&configs)

	return
}

package mongodb

import (
	"fmt"

	mgo "gopkg.in/mgo.v2"
)

var GlobalMgoSession *mgo.Session

func init() {
	configInit()
	var err error

	GlobalMgoSession, err = mgo.Dial(MongoConf.URL)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	GlobalMgoSession.SetPoolLimit(MongoConf.POOL_LIMIT)
}

func GetMongoSession() *mgo.Session {
	return GlobalMgoSession.Copy()
}

package mongodb

var MongoConf struct {
	URL        string
	POOL_LIMIT int
	DB         string
}

func configInit() {
	MongoConf.URL = "127.0.0.1:27017"
	MongoConf.POOL_LIMIT = 256
	MongoConf.DB = "cloudConfig"
}

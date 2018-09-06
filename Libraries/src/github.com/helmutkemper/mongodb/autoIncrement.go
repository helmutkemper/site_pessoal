package mongodb

import "github.com/helmutkemper/mgo/bson"

type autoIncrementStt struct {
	Id   bson.ObjectId `bson:"_id"`
	Auto int64         `bson:"auto"`
}

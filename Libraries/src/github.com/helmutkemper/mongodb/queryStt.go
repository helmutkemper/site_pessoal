package mongodb

import (
	"github.com/helmutkemper/mgo/bson"
	"time"
)

type QueryStt struct {
	Id         interface{}
	Query      interface{}
	Update     interface{}
	Insert     interface{}
	Limit      int
	Skip       int
	Order      []string
	Parse      map[string]interface{}
	Select     interface{}
	Collection string
}

func (el *QueryStt) GetId() string {
	switch el.Id.(type) {
	case bson.ObjectId:
	}
	if !el.Id.(bson.ObjectId).Valid() {
		el.Id = bson.NewObjectIdWithTime(time.Now())
	}

	return el.Id.(bson.ObjectId).String()
}

package db

import (
	"encoding/json"
	"github.com/helmutkemper/mgo/bson"
	"runtime"
)

// point struct based on osm file
type ErrorStt struct {
	IdMongo  bson.ObjectId `bson:"_id,omitempty"`
	Error    error
	List     string
	File     string
	FileName string
	Line     int
	DataJSon string
	*DbStt   `bson:"-"` //mongodb class
}

func (errorAStt *ErrorStt) Add(errorAErr error, dataATfc interface{}) {
	var dataLAByt []byte
	dataLAByt, _ = json.Marshal(dataATfc)
	errorAStt.Error = errorAErr
	errorAStt.DataJSon = string(dataLAByt)
}

func (errorAStt *ErrorStt) Insert(collectionNameAStr string) error {
	var err error

	err = errorAStt.DbStt.TestConnection()
	if err != nil {
		return err
	}

	pc := make([]uintptr, 10) // at least 1 entry needed
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	errorAStt.File, errorAStt.Line = f.FileLine(pc[0])
	errorAStt.FileName = f.Name()

	errorAStt.Line -= 3

	errorAStt.IdMongo, err = errorAStt.DbStt.GetMongoId()
	if err != nil {
		return err
	}

	errorAStt.List = string(errorAStt.Error.Error())

	return errorAStt.DbStt.Insert(collectionNameAStr, errorAStt)
}

func (errorAStt *ErrorStt) FindOne(collectionNameAStr string, queryAObj bson.M) error {
	err := errorAStt.DbStt.TestConnection()
	if err != nil {
		return err
	}

	return errorAStt.DbStt.FindOne(collectionNameAStr, &errorAStt, queryAObj)
}

package mongodb

import (
	"github.com/helmutkemper/mgo"
	log "github.com/helmutkemper/seelog"
)

type MongoDb struct {
	session *mgo.Session
	dbName  string
}

func (el *MongoDb) Connect(dbUrl, dbName, sessionName string) (error, *MongoDb) {
	var err error

	el.session, err = mgo.Dial(dbUrl)
	if err != nil {
		log.Criticalf("mongodb.connection.session.error: %v", err.Error())
		return err, nil
	}

	el.dbName = dbName

	if len(mongoDbSessions) == 0 {
		mongoDbSessions = make(map[string]MongoDb)
	}
	mongoDbSessions[sessionName] = *el

	return err, el
}

func (el *MongoDb) Disconnect() {
	err := el.session.Ping()

	if err != nil {
		return
	}

	el.session.Close()
}

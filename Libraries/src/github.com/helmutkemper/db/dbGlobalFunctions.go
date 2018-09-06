package db

import (
	"fmt"
	"github.com/helmutkemper/gOsm/consts"
	"github.com/helmutkemper/mgo"
	"github.com/helmutkemper/mgo/bson"
	log "github.com/helmutkemper/seelog"
	"time"
)

var errorMErr error
var connectionIsOpenMBoo bool
var addressMStr string
var nameDbMStr string
var connectionMObj *mgo.Session

// archives data keys after connection
// this var is used to check if index key 2d location exists into collection before insert first point
// todo usar isto na inserção de dados
// o que acontece se tentar ler uma chave que não existe?
var IndexListMMap map[string]map[string]mgo.Index

func init() {
	connectionIsOpenMBoo = false
}

func Connect(addressAStr, dbNameAStr string) (error, *mgo.Session) {
	connectionMObj, errorMErr = mgo.Dial(addressAStr)

	addressMStr = addressAStr
	nameDbMStr = dbNameAStr

	if errorMErr != nil {
		return errorMErr, nil
	}

	BuildIndexKeyList()

	connectionIsOpenMBoo = true

	return errorMErr, connectionMObj
}

func HasIndex(collectionNameAStr, indexNameAStr string) bool {
	return len(IndexListMMap[collectionNameAStr][indexNameAStr].Key) != 0
}

func BuildIndexKeyList() {
	collectionsLAStr, err := ListCollections()
	if err != nil {
		// todo colocar o log
	}

	IndexListMMap = make(map[string]map[string]mgo.Index, len(collectionsLAStr))
	for _, collectionNameLStr := range collectionsLAStr {
		indexLAStt, _ := IndexList(collectionNameLStr)
		IndexListMMap[collectionNameLStr] = make(map[string]mgo.Index, len(indexLAStt))
		for _, indexLStt := range indexLAStt {
			IndexListMMap[collectionNameLStr][indexLStt.Name] = indexLStt
		}
	}
}

func Disconnect() {
	errLErr := TestConnection()

	if errLErr != nil {
		return
	}

	connectionMObj.Close()
	connectionIsOpenMBoo = false

	if nameDbMStr == consts.DB_DATABASE_TEST {
		time.Sleep(consts.DELAY_TIME_BETWEEN_TESTS * time.Millisecond)
	}
}

func TestConnection() error {
	if connectionIsOpenMBoo == true {
		return nil
	} else {

		if nameDbMStr == "" && addressMStr == "" {
			log.Criticalf("Error: db is not connected")
			return fmt.Errorf("error: db is not connected")
		}

		log.Criticalf("Error: db is not connected to '%v' on address '%v'", nameDbMStr, addressMStr)
		return fmt.Errorf("error: db is not connected to '%v' on address '%v'", nameDbMStr, addressMStr)
	}
}

func DropCollection(collectionNameAStr string) error {
	return connectionMObj.DB(nameDbMStr).C(collectionNameAStr).DropCollection()
}

func DropDatabase(databaseNameAStr interface{}) error {
	if databaseNameAStr == nil {
		return connectionMObj.DB(nameDbMStr).DropDatabase()
	}

	return connectionMObj.DB(databaseNameAStr.(string)).DropDatabase()
}

func ListDatabases() ([]string, error) {
	return connectionMObj.DatabaseNames()
}

func ListCollections() ([]string, error) {
	return connectionMObj.DB(nameDbMStr).CollectionNames()
}

/** Monta um índice 2D de 26bits

 */
func Index2DMake(collectionNameAStr, fieldNameAStr string) error {
	var keyLStr string = "$2d:" + fieldNameAStr
	indexLStt := mgo.Index{
		Key:  []string{keyLStr},
		Bits: 26,
		Name: fieldNameAStr}
	err := connectionMObj.DB(nameDbMStr).C(collectionNameAStr).EnsureIndex(indexLStt)

	BuildIndexKeyList()

	return err
}

func IndexUniqueMake(collectionNameAStr, fieldNameAStr string) error {
	indexLStt := mgo.Index{
		Key:    []string{fieldNameAStr},
		Unique: true}
	err := connectionMObj.DB(nameDbMStr).C(collectionNameAStr).EnsureIndex(indexLStt)

	BuildIndexKeyList()

	return err
}

func IndexMake(collectionNameAStr string, index mgo.Index) error {
	err := connectionMObj.DB(nameDbMStr).C(collectionNameAStr).EnsureIndex(index)

	BuildIndexKeyList()

	return err
}

func IndexKeyMake(collectionNameAStr, fieldNameAStr string) error {
	indexLStt := mgo.Index{
		Key: []string{fieldNameAStr}}
	err := connectionMObj.DB(nameDbMStr).C(collectionNameAStr).EnsureIndex(indexLStt)

	BuildIndexKeyList()

	return err
}

func IndexTextMake(collectionNameAStr, fieldNameAStr, keyName string) error {
	indexLStt := mgo.Index{
		Key:  []string{"$text:" + keyName},
		Name: fieldNameAStr}
	err := connectionMObj.DB(nameDbMStr).C(collectionNameAStr).EnsureIndex(indexLStt)

	BuildIndexKeyList()

	return err
}

func IndexExpireAfterSeconds(collectionNameAStr, fieldNameAStr string, expireAfterSeconds time.Duration) error {
	indexLStt := mgo.Index{
		Key:         []string{fieldNameAStr},
		ExpireAfter: expireAfterSeconds}
	err := connectionMObj.DB(nameDbMStr).C(collectionNameAStr).EnsureIndex(indexLStt)

	BuildIndexKeyList()

	return err
}

func IndexDropByName(collectionNameAStr, indexNameAStr string) error {
	err := connectionMObj.DB(nameDbMStr).C(collectionNameAStr).DropIndexName(indexNameAStr)

	BuildIndexKeyList()

	return err
}

func IndexList(collectionNameAStr string) ([]mgo.Index, error) {
	return connectionMObj.DB(nameDbMStr).C(collectionNameAStr).Indexes()
}

func FindAll(collectionNameAStr string, dataATfc interface{}, queryAObj bson.M) error {
	errLErr := TestConnection()

	if errLErr != nil {
		return errLErr
	}

	return connectionMObj.DB(nameDbMStr).C(collectionNameAStr).Find(queryAObj).All(dataATfc)
}

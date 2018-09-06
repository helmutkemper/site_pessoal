package db

import (
	"github.com/helmutkemper/gOsm/consts"
	"github.com/helmutkemper/mgo"
	"github.com/helmutkemper/mgo/bson"
	log "github.com/helmutkemper/seelog"
	"math"
	"time"
)

var errorLog ErrorStt

type DbStt struct {
	conAddr            string
	conDbName          string
	Error              error
	Address            string
	Connection         *mgo.Session
	ConnectionIsOpen   bool
	DbName             string
	CollectionTags     *mgo.Collection
	CollectionWays     *mgo.Collection
	CollectionNodes    *mgo.Collection
	CollectionRelation *mgo.Collection
	CollectionPolygon  *mgo.Collection
	GridFSConnection   *mgo.GridFS
}

type autoIncrementStt struct {
	Id   bson.ObjectId `bson:"_id"`
	Auto int64         `bson:"auto"`
}

func (dbAStt *DbStt) Index2DMake(collectionNameAStr, fieldNameAStr string) error {
	return Index2DMake(collectionNameAStr, fieldNameAStr)
}

func (dbAStt *DbStt) IndexKeyMake(collectionNameAStr, fieldNameAStr string) error {
	return IndexKeyMake(collectionNameAStr, fieldNameAStr)
}

func (dbAStt *DbStt) IndexTextMake(collectionNameAStr, fieldNameAStr, keyName string) error {
	return IndexTextMake(collectionNameAStr, fieldNameAStr, keyName)
}

func (dbAStt *DbStt) IndexMake(collectionNameAStr string, index mgo.Index) error {
	return IndexMake(collectionNameAStr, index)
}

func (dbAStt *DbStt) HasIndex(collectionNameAStr, indexNameAStr string) bool {
	return HasIndex(collectionNameAStr, indexNameAStr)
}

func (dbAStt *DbStt) IndexExpireAfterSeconds(collectionNameAStr, indexNameAStr string, expireAfterSeconds time.Duration) error {
	return IndexExpireAfterSeconds(collectionNameAStr, indexNameAStr, expireAfterSeconds)
}

func (dbAStt *DbStt) BackupMake(data *MongoBackupData) error {
	var command MongoBackupCommand = MONGODB_BACKUP_DUMP
	return CreateBackup(command, data)
}

func (dbAStt *DbStt) BackupRestore(data *MongoBackupData) error {
	var command MongoBackupCommand = MONGODB_BACKUP_RESTORE
	return CreateBackup(command, data)
}

func (dbAStt *DbStt) TestConnection() error {
	err := TestConnection()
	if err != nil && dbAStt.conAddr != "" && dbAStt.conDbName != "" {
		return dbAStt.Connect(dbAStt.conAddr, dbAStt.conDbName)
	}
	return nil
}

func (dbAStt *DbStt) Connect(addressAStr, dbNameAStr string) error {
	dbAStt.conAddr = addressAStr
	dbAStt.conDbName = dbNameAStr
	dbAStt.Error, dbAStt.Connection = Connect(addressAStr, dbNameAStr)
	return dbAStt.Error
}

func (dbAStt *DbStt) Disconnect(addressAStr, dbNameAStr string) {
	Disconnect()
}

func (dbAStt *DbStt) GetMongoId() (bson.ObjectId, error) {
	errLErr := dbAStt.TestConnection()
	return bson.NewObjectId(), errLErr
}

func (dbAStt *DbStt) Insert(collectionNameAStr string, dataATfc interface{}) error {
	errLErr := dbAStt.TestConnection()

	if errLErr != nil {
		return errLErr
	}

	sessionCopy := connectionMObj.New()
	defer sessionCopy.Close()

	errLErr = sessionCopy.DB(nameDbMStr).C(collectionNameAStr).Insert(dataATfc)
	if errLErr != nil {
		log.Criticalf("gOsm.db.insert.error: %v", errLErr)
	}

	return errLErr
}

func goInsert(nameDbMStr, collectionNameAStr string, dataATfc interface{}) error {
	sessionCopy := connectionMObj.New()
	defer sessionCopy.Close()

	errLErr := sessionCopy.DB(nameDbMStr).C(collectionNameAStr).Insert(dataATfc)
	if errLErr != nil {
		log.Criticalf("gOsm.db.insert.error: %v", errLErr)
	}

	return errLErr
}

func (dbAStt *DbStt) Refresh() {
	connectionMObj.Refresh()
}

func (dbAStt *DbStt) MongoGridFSRemoveId(collectionNameAStr string, fileName string) error {
	errLErr := dbAStt.TestConnection()
	if errLErr != nil {
		return nil
	}

	errLErr = connectionMObj.DB(nameDbMStr).GridFS(collectionNameAStr + "Grid").RemoveId(fileName)

	return errLErr
}

func (dbAStt *DbStt) GridFSOpen(collectionNameAStr string, fileName string) (*mgo.GridFile, error) {
	var data *mgo.GridFile
	errLErr := dbAStt.TestConnection()
	if errLErr != nil {
		return nil, errLErr
	}

	data, errLErr = connectionMObj.DB(nameDbMStr).GridFS(collectionNameAStr + "Grid").Open(fileName)

	return data, errLErr
}

//todo: fala sobre o session.new e session.close
func (dbAStt *DbStt) GridFSCreate(collectionNameAStr string, fileName string) (*mgo.GridFile, error) {
	errLErr := dbAStt.TestConnection()
	if errLErr != nil {
		return nil, errLErr
	}

	return connectionMObj.DB(nameDbMStr).GridFS(collectionNameAStr + "Grid").Create(fileName)
}

func (dbAStt *DbStt) Update(collectionNameAStr string, dataATfc interface{}, queryAObj interface{}) error {
	errLErr := dbAStt.TestConnection()

	if errLErr != nil {
		return errLErr
	}

	sessionCopy := connectionMObj.New()
	defer sessionCopy.Close()

	return sessionCopy.DB(nameDbMStr).C(collectionNameAStr).Update(queryAObj, dataATfc)
}

func (dbAStt *DbStt) Delete(collectionNameAStr string, queryAObj bson.M) error {
	errLErr := dbAStt.TestConnection()

	if errLErr != nil {
		return errLErr
	}

	sessionCopy := connectionMObj.New()
	defer sessionCopy.Close()

	return sessionCopy.DB(nameDbMStr).C(collectionNameAStr).Remove(queryAObj)
}

func (dbAStt *DbStt) FindOne(collectionNameAStr string, dataATfc interface{}, queryAObj bson.M) error {
	errLErr := dbAStt.TestConnection()

	if errLErr != nil {
		return errLErr
	}

	sessionCopy := connectionMObj.New()
	defer sessionCopy.Close()

	return sessionCopy.DB(nameDbMStr).C(collectionNameAStr).Find(queryAObj).One(dataATfc)
}

func (dbAStt *DbStt) FindLast(collectionNameAStr string, dataATfc interface{}, queryAObj bson.M, order string) error {
	errLErr := dbAStt.TestConnection()

	if errLErr != nil {
		return errLErr
	}

	sessionCopy := connectionMObj.New()
	defer sessionCopy.Close()

	return sessionCopy.DB(nameDbMStr).C(collectionNameAStr).Find(queryAObj).Sort(order).One(dataATfc)
}

func (dbAStt *DbStt) Remove(collectionNameAStr string, queryAObj bson.M) error {
	errLErr := dbAStt.TestConnection()

	if errLErr != nil {
		return errLErr
	}

	sessionCopy := connectionMObj.New()
	defer sessionCopy.Close()

	return sessionCopy.DB(nameDbMStr).C(collectionNameAStr).Remove(queryAObj)
}

func (dbAStt *DbStt) RemoveAll(collectionNameAStr string, queryAObj bson.M) (*mgo.ChangeInfo, error) {
	errLErr := dbAStt.TestConnection()

	if errLErr != nil {
		return nil, errLErr
	}

	sessionCopy := connectionMObj.New()
	defer sessionCopy.Close()

	return sessionCopy.DB(nameDbMStr).C(collectionNameAStr).RemoveAll(queryAObj)
}

func (dbAStt *DbStt) RemoveId(collectionNameAStr string, queryAObj bson.M) error {
	errLErr := dbAStt.TestConnection()

	if errLErr != nil {
		return errLErr
	}

	sessionCopy := connectionMObj.New()
	defer sessionCopy.Close()

	return sessionCopy.DB(nameDbMStr).C(collectionNameAStr).RemoveId(queryAObj)
}

//bson.M{ "loc": bson.M{ "$geoWithin": bson.M{ "$box": boxAAFlt } } }
func (dbAStt *DbStt) Find(collectionNameAStr string, dataATfc interface{}, argsATfc ...interface{}) error {
	//log.Trace( "db.go" )
	//log.Trace( "DbFind()" )
	//log.Tracef( "database: %v", nameDbMStr )
	//log.Tracef( "collection: %v", collectionNameAStr )

	errLErr := dbAStt.TestConnection()
	if errLErr != nil {
		return errLErr
	}

	sessionCopy := connectionMObj.New()
	defer sessionCopy.Close()

	if len(argsATfc) == 3 {
		//log.Tracef( "query: %v", argsATfc[ consts.FIND_ALL_QUERY ].( bson.M ) )
		//log.Tracef( "skip: %v", argsATfc[ consts.FIND_ALL_SKIP ].( int ) )
		//log.Tracef( "limit: %v", argsATfc[ consts.FIND_ALL_LIMIT ].( int ) )
		return sessionCopy.DB(nameDbMStr).C(collectionNameAStr).Find(argsATfc[consts.FIND_ALL_QUERY].(bson.M)).Skip(argsATfc[consts.FIND_ALL_SKIP].(int)).Limit(argsATfc[consts.FIND_ALL_LIMIT].(int)).All(dataATfc)
	}

	if len(argsATfc) == 2 {
		//log.Tracef( "query: %v", argsATfc[ consts.FIND_ALL_QUERY ].( bson.M ) )
		//log.Tracef( "limit: %v", argsATfc[ consts.FIND_ALL_LIMIT ].( int ) )
		return sessionCopy.DB(nameDbMStr).C(collectionNameAStr).Find(argsATfc[consts.FIND_ALL_QUERY].(bson.M)).Limit(argsATfc[consts.FIND_ALL_LIMIT].(int)).All(dataATfc)
	}

	//log.Tracef( "query: %v", argsATfc[ consts.FIND_ALL_QUERY ].( bson.M ) )
	return sessionCopy.DB(nameDbMStr).C(collectionNameAStr).Find(argsATfc[consts.FIND_ALL_QUERY].(bson.M)).All(dataATfc)
}

func (dbAStt *DbStt) Count(collectionNameAStr string, argsATfc ...interface{}) (int, error) {
	errLErr := dbAStt.TestConnection()

	if errLErr != nil {
		return -1, errLErr
	}

	sessionCopy := connectionMObj.New()
	defer sessionCopy.Close()

	if len(argsATfc) == 3 {
		return sessionCopy.DB(nameDbMStr).C(collectionNameAStr).Find(argsATfc[consts.FIND_ALL_QUERY].(bson.M)).Skip(argsATfc[consts.FIND_ALL_SKIP].(int)).Limit(argsATfc[consts.FIND_ALL_LIMIT].(int)).Count()
	}

	if len(argsATfc) == 2 {
		return sessionCopy.DB(nameDbMStr).C(collectionNameAStr).Find(argsATfc[consts.FIND_ALL_QUERY].(bson.M)).Limit(argsATfc[consts.FIND_ALL_LIMIT].(int)).Count()
	}

	return sessionCopy.DB(nameDbMStr).C(collectionNameAStr).Find(argsATfc[consts.FIND_ALL_QUERY].(bson.M)).Count()
}

func (dbAStt *DbStt) AutoIncrementPrepare(collectionNameAStr string) error {
	errLErr := dbAStt.TestConnection()

	if errLErr != nil {
		return errLErr
	}

	var auto autoIncrementStt
	auto.Id = bson.NewObjectId()

	errLErr = connectionMObj.DB(nameDbMStr).C(collectionNameAStr).Find(bson.M{}).One(&auto)
	if errLErr == nil {
		return nil
	}

	auto.Auto = math.MaxInt64
	return connectionMObj.DB(nameDbMStr).C(collectionNameAStr).Insert(&auto)
}

func (dbAStt *DbStt) AutoIncrement(collectionNameAStr string) (error, int64) {
	errLErr := dbAStt.TestConnection()

	if errLErr != nil {
		return errLErr, 0
	}

	var auto autoIncrementStt
	change := mgo.Change{
		Update:    bson.M{"$inc": bson.M{"auto": -1}},
		ReturnNew: true,
	}

	_, errLErr = connectionMObj.DB(nameDbMStr).C(collectionNameAStr).Find(bson.M{}).Apply(change, &auto)

	return errLErr, auto.Auto
}

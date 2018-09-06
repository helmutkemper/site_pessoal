package mongodb

import (
	"fmt"
	"github.com/helmutkemper/mgo"
	"github.com/helmutkemper/mgo/bson"
	"math"
	"reflect"
	"regexp"
	"strings"
	"time"
)

type DbStt struct {
	conAddr       string
	conDbName     string
	conCollection string
	error         error
	connection    *mgo.Session
	dbName        string
	data          []map[string]interface{}
	dataPointer   int
	dataLength    int
}

func (el *DbStt) Init(sessionName string) {
	var c = mongoDbSessions[sessionName]
	el.connection = c.session
	el.dbName = c.dbName
}

func (el *DbStt) Collection(name string) {
	el.conCollection = name
}

func (el *DbStt) GetBuildInfo() string {
	var info mgo.BuildInfo
	info, el.error = el.connection.BuildInfo()
	if el.error != nil {
		fmt.Printf("error: %v\n", el.error.Error())
	}

	return info.Version
}

func (el *DbStt) Index2DMake(fieldNameAStr string) bool {
	var keyLStr = "$2d:" + fieldNameAStr
	indexLStt := mgo.Index{
		Key:  []string{keyLStr},
		Bits: 26,
		Name: fieldNameAStr,
	}
	el.error = el.connection.DB(el.dbName).C(el.conCollection).EnsureIndex(indexLStt)
	if el.error != nil {
		return false
	}
	return true
}

func (el *DbStt) IndexUniqueMake(fieldNameAStr []string) bool {
	indexLStt := mgo.Index{
		Key:    fieldNameAStr,
		Unique: true,
	}
	el.error = el.connection.DB(el.dbName).C(el.conCollection).EnsureIndex(indexLStt)

	if el.error != nil {
		return false
	}
	return true
}

func (el *DbStt) IndexKeyMake(fieldNameAStr string) bool {
	indexLStt := mgo.Index{
		Key: []string{fieldNameAStr},
	}
	el.error = el.connection.DB(el.dbName).C(el.conCollection).EnsureIndex(indexLStt)
	if el.error != nil {
		return false
	}
	return true
}

func (el *DbStt) IndexTextMake(fieldNameAStr, keyName string) bool {
	indexLStt := mgo.Index{
		Key:  []string{"$text:" + keyName},
		Name: fieldNameAStr,
	}
	el.error = el.connection.DB(el.dbName).C(el.conCollection).EnsureIndex(indexLStt)
	if el.error != nil {
		return false
	}
	return true
}

func (el *DbStt) IndexList() []mgo.Index {
	var index []mgo.Index
	index, el.error = el.connection.DB(el.dbName).C(el.conCollection).Indexes()

	return index
}

func (el *DbStt) IndexDropByName(indexNameAStr string) bool {
	el.error = el.connection.DB(el.dbName).C(el.conCollection).DropIndexName(indexNameAStr)
	if el.error != nil {
		return false
	}
	return true
}

func (el *DbStt) IndexMake(index mgo.Index) bool {
	el.error = el.connection.DB(el.dbName).C(el.conCollection).EnsureIndex(index)
	if el.error != nil {
		return false
	}
	return true
}

func (el *DbStt) HasIndex(indexNameAStr string) bool {
	var list []mgo.Index

	list, el.error = el.connection.DB(el.dbName).C(el.conCollection).Indexes()

	for _, data := range list {
		if data.Name == indexNameAStr {
			return true
		}
	}

	return false
}

func (el *DbStt) IndexExpireAfterSeconds(fieldNameAStr []string, expireAfterSeconds int64) bool {
	indexLStt := mgo.Index{
		Key:         fieldNameAStr,
		ExpireAfter: time.Duration(expireAfterSeconds * int64(time.Second)),
	}
	el.error = el.connection.DB(el.dbName).C(el.conCollection).EnsureIndex(indexLStt)
	if el.error != nil {
		return false
	}
	return true
}

func (el *DbStt) DropCollection() bool {
	el.error = el.connection.DB(el.dbName).C(el.conCollection).DropCollection()
	if el.error != nil {
		return false
	}
	return true
}

func (el *DbStt) DropDatabase() bool {
	el.error = el.connection.DB(el.dbName).DropDatabase()
	if el.error != nil {
		return false
	}
	return true
}

func (el *DbStt) ListDatabases() []string {
	var list []string
	list, el.error = el.connection.DatabaseNames()

	return list
}

func (el *DbStt) ListCollections() []string {
	var list []string
	list, el.error = el.connection.DB(el.dbName).CollectionNames()

	return list
}

func (el *DbStt) TestConnection() error {
	el.error = el.connection.Ping()

	return el.error
}

func (el *DbStt) TestConnectionJs() bool {
	el.error = el.TestConnection()
	if el.error != nil {
		return false
	}

	return true
}

func (el *DbStt) GetId() bson.ObjectId {
	el.error = el.TestConnection()
	return bson.NewObjectIdWithTime(time.Now())
}

// fixme: adicionar um campo mpara normalizar
func (el *DbStt) Insert(data interface{}) interface{} {
	el.error = el.TestConnection()

	if el.error != nil {
		return QueryStt{}
	}

	sessionCopy := el.connection.New()
	defer sessionCopy.Close()

	el.error = sessionCopy.DB(el.dbName).C(el.conCollection).Insert(data)

	return data
}

func (el *DbStt) Refresh() {
	el.connection.Refresh()
}

func (el *DbStt) GridFSRemoveId(data *QueryStt) error {
	el.error = el.TestConnection()
	if el.error != nil {
		return nil
	}

	el.error = el.connection.DB(el.dbName).GridFS(el.conCollection + "Grid").RemoveId(data.GetId())

	return el.error
}

func (el *DbStt) GridFSOpen(data *QueryStt) (*mgo.GridFile, error) {
	var dataGridFile *mgo.GridFile
	el.error = el.TestConnection()
	if el.error != nil {
		return nil, el.error
	}

	dataGridFile, el.error = el.connection.DB(el.dbName).GridFS(el.conCollection + "Grid").Open(data.GetId())

	return dataGridFile, el.error
}

//todo: fala sobre o session.new e session.close
func (el *DbStt) GridFSCreate(data *QueryStt) (*mgo.GridFile, error) {
	el.error = el.TestConnection()
	if el.error != nil {
		return nil, el.error
	}

	return el.connection.DB(el.dbName).GridFS(el.conCollection + "Grid").Create(data.GetId())
}

func (el *DbStt) idStringToObject(data interface{}) interface{} {
	switch data.(type) {
	case map[string]interface{}:
		if data.(map[string]interface{})["_id"] != nil {
			switch data.(map[string]interface{})["_id"].(type) {
			case string:
				data.(map[string]interface{})["_id"] = bson.ObjectIdHex(data.(map[string]interface{})["_id"].(string))
			}
		}
	}

	return data
}

func (el *DbStt) Update(data *QueryStt) bool {
	el.error = el.TestConnection()

	if el.error != nil {
		return false
	}

	sessionCopy := el.connection.New()
	defer sessionCopy.Close()

	data.Query = el.idStringToObject(data.Query)

	el.error = sessionCopy.DB(el.dbName).C(el.conCollection).Update(data.Query, data.Update)

	if el.error != nil {
		return false
	}

	return true
}

func (el *DbStt) Upset(data *QueryStt) bool {
	el.error = el.TestConnection()

	if el.error != nil {
		return false
	}

	sessionCopy := el.connection.New()
	defer sessionCopy.Close()

	data.Query = el.idStringToObject(data.Query)

	_, el.error = sessionCopy.DB(el.dbName).C(el.conCollection).Upsert(data.Query, data.Update)

	if el.error != nil {
		return false
	}

	return true
}

func (el *DbStt) Remove(data *QueryStt) bool {
	el.error = el.TestConnection()

	if el.error != nil {
		return false
	}

	sessionCopy := el.connection.New()
	defer sessionCopy.Close()

	data.Query = el.idStringToObject(data.Query)
	el.error = sessionCopy.DB(el.dbName).C(el.conCollection).Remove(data.Query)
	if el.error != nil {
		return false
	}

	return true
}

func (el *DbStt) RemoveAll(data *QueryStt) bool {
	el.error = el.TestConnection()

	if el.error != nil {
		return false
	}

	sessionCopy := el.connection.New()
	defer sessionCopy.Close()

	data.Query = el.idStringToObject(data.Query)
	_, el.error = sessionCopy.DB(el.dbName).C(el.conCollection).RemoveAll(data.Query)
	if el.error != nil {
		return false
	}

	return true
}

func (el *DbStt) FindOne(data *QueryStt) bool {
	el.error = el.TestConnection()

	if el.error != nil {
		el.dataLength = 0
		el.dataPointer = 0
		return false
	}

	sessionCopy := el.connection.New()
	defer sessionCopy.Close()

	el.data = make([]map[string]interface{}, 1)

	data.Query = el.idStringToObject(data.Query)
	el.error = sessionCopy.DB(el.dbName).C(el.conCollection).Find(data.Query).One(&el.data[0])
	if el.error != nil {
		el.dataLength = 0
		el.dataPointer = 0
		el.data = make([]map[string]interface{}, 0)

		return false
	}

	el.dataLength = len(el.data)
	el.dataPointer = 0

	return true
}

func (el *DbStt) FindLast(data *QueryStt) bool {
	el.error = el.TestConnection()

	if el.error != nil {
		return false
	}

	sessionCopy := el.connection.New()
	defer sessionCopy.Close()

	data.Query = el.idStringToObject(data.Query)
	el.data = make([]map[string]interface{}, 1)

	sessionCopy.DB(el.dbName).C(el.conCollection).Find(data.Query).Sort(data.Order...).One(&el.data[0])
	//fixme: on error?
	el.dataLength = 1
	el.dataPointer = 0

	return true
}

func (el *DbStt) RemoveById(data *QueryStt) bool {
	el.error = el.TestConnection()

	if el.error != nil {
		return false
	}

	sessionCopy := el.connection.New()
	defer sessionCopy.Close()

	data.Id = el.idStringToObject(data.Id)
	el.error = sessionCopy.DB(el.dbName).C(el.conCollection).RemoveId(data.Id)
	if el.error != nil {
		return false
	}
	return true
}

//bson.M{ "loc": bson.M{ "$geoWithin": bson.M{ "$box": boxAAFlt } } }
func (el *DbStt) Find(data *QueryStt) bool {
	el.error = el.TestConnection()
	if el.error != nil {
		el.dataLength = 0
		el.dataPointer = 0
		return false
	}

	sessionCopy := el.connection.New()
	defer sessionCopy.Close()

	data.Query = el.idStringToObject(data.Query)

	if data.Collection != "" {
		el.error = sessionCopy.DB(el.dbName).C(data.Collection).Find(data.Query).Skip(data.Skip).Limit(data.Limit).Sort(data.Order...).Select(data.Select).All(&el.data)
	} else {
		el.error = sessionCopy.DB(el.dbName).C(el.conCollection).Find(data.Query).Skip(data.Skip).Limit(data.Limit).Sort(data.Order...).Select(data.Select).All(&el.data)
	}

	if el.error != nil {
		el.dataLength = 0
		el.dataPointer = 0
		return false
	}

	el.dataLength = len(el.data)
	el.dataPointer = 0

	return true
}

func (el *DbStt) DataGet() map[string]interface{} {
	if el.dataPointer < el.dataLength {
		return el.data[el.dataPointer]
	}

	return nil
}

func (el *DbStt) DataGetAll() []map[string]interface{} {
	return el.data
}

func (el *DbStt) incDataPointer() {
	el.dataPointer += 1
}

func (el *DbStt) decDataPointer() {
	el.dataPointer -= 1
}

func (el *DbStt) ResetDataPointer() {
	el.dataPointer = 0
}

func (el *DbStt) ToLastDataPointer() {
	el.dataPointer = len(el.data)
}

func (el *DbStt) DataGetNext() map[string]interface{} {
	defer el.incDataPointer()

	if el.dataPointer < el.dataLength {
		return el.data[el.dataPointer]
	}

	return map[string]interface{}{}
}

func (el *DbStt) DataGetPre() map[string]interface{} {
	defer el.decDataPointer()

	if el.dataPointer > -1 {
		return el.data[el.dataPointer]
	}

	return map[string]interface{}{}
}

func (el *DbStt) Count(data *QueryStt) int {
	el.error = el.TestConnection()
	if el.error != nil {
		return 0
	}

	var count int

	sessionCopy := el.connection.New()
	defer sessionCopy.Close()

	data.Query = el.idStringToObject(data.Query)

	count, el.error = sessionCopy.DB(el.dbName).C(el.conCollection).Find(data.Query).Skip(data.Skip).Limit(data.Limit).Sort(data.Order...).Count()

	return count
}

func (el *DbStt) AutoIncrementPrepare(data QueryStt) bool {
	el.error = el.TestConnection()

	if el.error != nil {
		return false
	}

	var auto autoIncrementStt
	auto.Id = bson.NewObjectIdWithTime(time.Now())

	// fixme: rever isto
	el.error = el.connection.DB(el.dbName).C(el.conCollection).Find(bson.M{}).One(&auto)
	if el.error == nil {
		return true
	}

	if auto.Auto != 0 {
		return true
	}

	auto.Auto = math.MaxInt64
	el.error = el.connection.DB(el.dbName).C(el.conCollection).Insert(&auto)
	if el.error != nil {
		return false
	}
	return true
}

func (el *DbStt) AutoIncrement(data QueryStt) (bool, int64) {
	el.error = el.TestConnection()

	if el.error != nil {
		return false, 0
	}

	var auto autoIncrementStt
	change := mgo.Change{
		Update:    bson.M{"$inc": bson.M{"auto": -1}},
		ReturnNew: true,
	}

	_, el.error = el.connection.DB(el.dbName).C(el.conCollection).Find(bson.M{}).Apply(change, &auto)
	if el.error != nil {
		return false, 0
	}
	return true, auto.Auto
}

func (el *DbStt) copyInterface(obj interface{}) interface{} {
	// Wrap the original in a reflect.Value
	original := reflect.ValueOf(obj)

	copy := reflect.New(original.Type()).Elem()
	el.copyRecursive(copy, original)

	// Remove the reflection wrapper
	return copy.Interface()
}

func (el *DbStt) copyRecursive(copy, original reflect.Value) {
	switch original.Kind() {
	// The first cases handle nested structures and translate them recursively

	// If it is a pointer we need to unwrap and call once again
	case reflect.Ptr:
		// To get the actual value of the original we have to call Elem()
		// At the same time this unwraps the pointer so we don't end up in
		// an infinite recursion
		originalValue := original.Elem()
		// Check if the pointer is nil
		if !originalValue.IsValid() {
			return
		}
		// Allocate a new object and set the pointer to it
		copy.Set(reflect.New(originalValue.Type()))
		// Unwrap the newly created pointer
		el.copyRecursive(copy.Elem(), originalValue)

		// If it is an interface (which is very similar to a pointer), do basically the
		// same as for the pointer. Though a pointer is not the same as an interface so
		// note that we have to call Elem() after creating a new object because otherwise
		// we would end up with an actual pointer
	case reflect.Interface:
		// Get rid of the wrapping interface
		originalValue := original.Elem()
		// Create a new object. Now new gives us a pointer, but we want the value it
		// points to, so we have to call Elem() to unwrap it
		copyValue := reflect.New(originalValue.Type()).Elem()
		el.copyRecursive(copyValue, originalValue)
		copy.Set(copyValue)

		// If it is a struct we translate each field
	case reflect.Struct:
		for i := 0; i < original.NumField(); i += 1 {
			el.copyRecursive(copy.Field(i), original.Field(i))
		}

		// If it is a slice we create a new slice and translate each element
	case reflect.Slice:
		copy.Set(reflect.MakeSlice(original.Type(), original.Len(), original.Cap()))
		for i := 0; i < original.Len(); i += 1 {
			el.copyRecursive(copy.Index(i), original.Index(i))
		}

		// If it is a map we create a new map and translate each value
	case reflect.Map:
		copy.Set(reflect.MakeMap(original.Type()))
		for _, key := range original.MapKeys() {
			originalValue := original.MapIndex(key)
			// New gives us a pointer, but again we want the value
			copyValue := reflect.New(originalValue.Type()).Elem()
			el.copyRecursive(copyValue, originalValue)
			copy.SetMapIndex(key, copyValue)
		}

		// Otherwise we cannot traverse anywhere so this finishes the the recursion

		// If it is a string translate it (yay finally we're doing what we came for)
	case reflect.String:
		translatedString := original.Interface().(string)
		copy.SetString(translatedString)

		// And everything else will simply be taken from the original
	default:
		copy.Set(original)
	}

}

func (el *DbStt) ParserQuery(query interface{}, dataToParse interface{}) interface{} {
	query = el.copyInterface(query)

	out := el.parserQuerySupport(query, dataToParse)
	return out
}

func (el *DbStt) parserQuerySupport(data interface{}, parse interface{}) interface{} {

	switch converted := data.(type) {
	case bson.M:

		for queryKey, queryValue := range converted {

			switch convertedValue := queryValue.(type) {
			case string:

				if strings.HasPrefix(convertedValue, "#") && strings.HasSuffix(convertedValue, "#") {

					regExpSub := regexp.MustCompile("^#(.*?)#$")
					formKeyToReplace := regExpSub.ReplaceAllString(convertedValue, "$1")

					switch parse.(type) {
					case map[string]interface{}:

						data.(bson.M)[queryKey] = parse.(map[string]interface{})[formKeyToReplace]

					case map[string]string:

						data.(bson.M)[queryKey] = parse.(map[string]string)[formKeyToReplace]
					}
				}

			case bson.M:

				data.(bson.M)[queryKey] = el.parserQuerySupport(data, parse)
			}
		}

	case map[string]interface{}:

		for queryKey, queryValue := range converted {

			switch convertedValue := queryValue.(type) {
			case string:

				if strings.HasPrefix(convertedValue, "#") && strings.HasSuffix(convertedValue, "#") {

					regExpSub := regexp.MustCompile("^#(.*?)#$")
					formKeyToReplace := regExpSub.ReplaceAllString(convertedValue, "$1")

					switch parse.(type) {
					case map[string]interface{}:

						data.(map[string]interface{})[queryKey] = parse.(map[string]interface{})[formKeyToReplace]

					case map[string]string:

						data.(map[string]interface{})[queryKey] = parse.(map[string]string)[formKeyToReplace]
					}
				}

			case bson.M:

				data.(bson.M)[queryKey] = el.parserQuerySupport(data.(bson.M)[queryKey], parse)

			case map[string]interface{}:

				data.(map[string]interface{})[queryKey] = el.parserQuerySupport(data.(map[string]interface{})[queryKey], parse)
			}
		}
	}

	return data
}

func (el *DbStt) GetLastErrorAsString() string {
	if el.error != nil {
		return el.error.Error()
	}
	return ""
}

func (el *DbStt) GetLastError() error {
	return el.error
}

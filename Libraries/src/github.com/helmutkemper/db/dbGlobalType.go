package db

type GenericData struct {
	Id int64 `bson:"_id"`

	Content       []byte            `bson:"content"`
	Tag           map[string]string `bson:"tag"`
	International map[string]string `bson:"international"`

	Md5  [16]byte `bson:"md5"`
	Size int      `bson:"size"`
}

package db

/*
import (
  log "github.com/helmutkemper/seelog"
  "github.com/helmutkemper/gorocksdb"
  "errors"
  "strconv"
  "encoding/binary"
  "os"
)

const folders_count int = 10

type RocksDB struct {
  Path                  string                              `bson:"path" json:"path"`
  Threads               int                                 `bson:"threads" json:"threads"`
  ThreadsCounter        int                                 `bson:"threadsCounter" json:"threadsCounter"`
  ThreadsCounterInc     int                                 `bson:"threadsCounterInc" json:"threadsCounterInc"`
  Next                  int                                 `bson:"threadsCounterInc" json:"threadsCounterInc"`
  Busy                  bool                                `bson:"busy" json:"busy"`
  Options               *gorocksdb.Options                  `bson:"options" json:"options"`
  Db                    []*gorocksdb.DB                     `bson:"db" json:"db"`
  WriteOptions          *gorocksdb.WriteOptions             `bson:"writeOptions" json:"writeOptions"`
  ReadOptions           *gorocksdb.ReadOptions              `bson:"readOptions" json:"readOptions"`
  HasInitialized        bool                                `bson:"hasInitialized" json:"hasInitialized"`

  Md5                   [16]byte                            `bson:"md5" json:"-"`
  Size                  int                                 `bson:"size" json:"-"`
}

func( el *RocksDB ) Init( path string ) error {
  var err error

  if path == "" {
    log.Critical( "rocksdb.Path must be a directory" )
    return errors.New( "rocksdb.Path must be a directory" )
  }

  el.Path = path

  el.Options = gorocksdb.NewDefaultOptions()
  el.Options.SetCreateIfMissing(true)

  el.Db = make( []*gorocksdb.DB, folders_count )

  for i := 0; i != folders_count; i += 1 {
    if _, err = os.Stat( path + "/" + strconv.FormatInt( int64( i ), 10 ) ); os.IsNotExist( err ) {
      err = os.Mkdir( path + "/" + strconv.FormatInt( int64( i ), 10 ), 0777 )
      if err != nil {
        return err
      }
    }

    el.Db[i], err = gorocksdb.OpenDb( el.Options, path + "/" + strconv.FormatInt( int64( i ), 10 ) )
    if err != nil {
      return err
    }
  }

  el.WriteOptions   = gorocksdb.NewDefaultWriteOptions()
  el.ReadOptions    = gorocksdb.NewDefaultReadOptions()

  el.HasInitialized = true

  return nil
}

func( el *RocksDB ) Put( id int64, data []byte ) error {
  var err error

  key := id % int64( folders_count )
  var idByte = make( []byte, 8 )
  binary.LittleEndian.PutUint64( idByte, uint64( id ) )

  err = el.Db[key].Put( el.WriteOptions, idByte, data )
  //if err != nil {
  //  log.Criticalf( "TmpPointStt.DbKeyValueInsert.Db.Put.error: %v", err.Error() )
  //}

  return err
}

func( el *RocksDB ) Get( id int64 ) ( *gorocksdb.Slice, error ) {
  var err error
  var slice *gorocksdb.Slice

  key := id % int64( folders_count )
  var idByte = make( []byte, 8 )
  binary.LittleEndian.PutUint64( idByte, uint64( id ) )

  slice, err = el.Db[key].Get( el.ReadOptions, idByte )
  //if err != nil {
  //  log.Criticalf( "TmpPointStt.DbKeyValueFind.Db.Get.error: %v", err.Error() )
  //  return err, []byte{}
  //}

  return slice, err
}

func( el *RocksDB ) Delete( id int64 ) error {
  var err error

  key := id % int64( folders_count )
  var idByte = make( []byte, 8 )
  binary.LittleEndian.PutUint64( idByte, uint64( id ) )

  err = el.Db[key].Delete( el.WriteOptions, idByte )
  //if err != nil {
  //  log.Criticalf( "tmpPoint.DbKeyValueRemove.Db.Delete.error: %v", err.Error() )
  //}

  return err
}
*/

package database

import (
	"graph_robot/config"
	"log"
	"reflect"
	"strings"

	"github.com/dgraph-io/badger/v4"
)

var db *badger.DB
var seqMap *map[string]*badger.Sequence

func InitDb(dbPath string, seqBandwidth int) {
	db = getDB(dbPath)
	seqMap = getSequenceObject(seqBandwidth)
}

func getDB(dbPath string) *badger.DB {
	db, err := badger.Open(badger.DefaultOptions(dbPath))
	if err != nil {
		log.Panic(err)
	}
	return db
}

func getSequenceObject(seqBandwidth int) *map[string]*badger.Sequence {
	// run this function will add a data which key name is the same name with prefix
	seqMap := make(map[string]*badger.Sequence)
	prefix := config.GetAllPrefix()
	prefix = append(prefix, config.TestPrefix) // use when testing
	for i := 0; i < len(prefix); i++ {
		keyPrefix := prefix[i]
		seq, err := db.GetSequence([]byte(keyPrefix), uint64(seqBandwidth))
		if err != nil {
			log.Panic(err)
		}
		seqMap[keyPrefix] = seq
	}
	return &seqMap
}

func CloseDb() {
	for k := range *seqMap {
		err := (*seqMap)[k].Release()
		if err != nil {
			log.Println("*****", reflect.TypeOf(err))
			log.Panic(err)
		}
	}
	err := db.Close()
	if err != nil && !strings.Contains(err.Error(), "resource temporarily unavailable") {
		log.Panic(err)
	}
}

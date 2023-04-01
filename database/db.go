package database

import (
	"graph_robot/config"

	"github.com/dgraph-io/badger/v4"
)

var db *badger.DB
var seqMap map[string]*badger.Sequence

func InitDb(dbPath string) {
	db = getDB(dbPath)
	seqMap = getSequenceObject()
}

func getDB(dbPath string) *badger.DB {
	db, err := badger.Open(badger.DefaultOptions(dbPath))
	if err != nil {
		panic(err)
	}
	return db
}

func getSequenceObject() map[string]*badger.Sequence {
	seqMap := make(map[string]*badger.Sequence)
	for i := 0; i < len(config.NeurePrefix); i++ {
		keyPrefix := config.NeurePrefix[i]
		seq, err := db.GetSequence([]byte(keyPrefix), 1000)
		if err != nil {
			panic(err)
		}
		seqMap[keyPrefix] = seq
	}
	return seqMap
}

func CloseDb() {
	db.Close()
}

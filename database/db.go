package database

import (
	"graph_robot/config"

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
		panic(err)
	}
	return db
}

func getSequenceObject(seqBandwidth int) *map[string]*badger.Sequence {
	seqMap := make(map[string]*badger.Sequence)
	for i := 0; i < len(config.NeurePrefix); i++ {
		keyPrefix := config.NeurePrefix[i]
		seq, err := db.GetSequence([]byte(keyPrefix), uint64(seqBandwidth))
		if err != nil {
			panic(err)
		}
		seqMap[keyPrefix] = seq
	}
	return &seqMap
}

func CloseDb() {
	for k := range *seqMap {
		err := (*seqMap)[k].Release()
		if err != nil {
			panic(err)
		}
	}
	err := db.Close()
	if err != nil {
		panic(err)
	}
}

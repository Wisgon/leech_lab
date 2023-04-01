package database

import (
	"fmt"

	"github.com/dgraph-io/badger/v4"
)

func CreateNeure(neureByte []byte, keyPrefix string) (key string) {
	uniqueNum := GetSeqNum(keyPrefix)
	key = keyPrefix + fmt.Sprint(uniqueNum)
	_ = db.Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte(key), neureByte)
		if err != nil {
			panic(err)
		}
		return nil
	})
	return
}

func UpdateNeure(neureByte []byte, neureId string) {
	_ = db.Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte(neureId), neureByte)
		if err != nil {
			panic(err)
		}
		return nil
	})
}

func DeleteNeure(neureId string) {
	_ = db.Update(func(txn *badger.Txn) error {
		err := txn.Delete([]byte(neureId))
		if err != nil {
			panic(err)
		}
		return nil
	})
}

func GetNeure(neureId string) []byte {
	var neure []byte
	_ = db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(neureId))
		if err != nil {
			panic(err)
		}
		_ = item.Value(func(val []byte) error {
			neure = append([]byte{}, val...)
			return nil
		})
		return nil
	})
	return neure
}

func GetSeqNum(keyPrefix string) string {
	uniqueNum, err := seqMap[keyPrefix].Next()
	if err != nil {
		panic(err)
	}
	return fmt.Sprint(uniqueNum)
}

func KeyOnlyPrefixScan(keyPrefix string) (resultKeys []string) {
	return
}

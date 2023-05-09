package database

import (
	"errors"
	"graph_robot/config"
	"log"
	"strings"

	"github.com/dgraph-io/badger/v4"
)

type ConditionFilter func(result []byte) bool

func CreateData(dataByte []byte, key string) {
	_ = db.Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte(key), dataByte)
		if err != nil {
			log.Panic(err)
		}
		return nil
	})
}

func UpdateNeure(neureByte []byte, neureId string) {
	_ = db.Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte(neureId), neureByte)
		if err != nil {
			log.Panic(err)
		}
		return nil
	})
}

func DeleteNeure(neureId string) {
	_ = db.Update(func(txn *badger.Txn) error {
		err := txn.Delete([]byte(neureId))
		if err != nil {
			log.Panic(err)
		}
		return nil
	})
}

func GetDataById(id string) (data []byte) {
	_ = db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(id))
		if err != nil {
			if strings.Contains(err.Error(), "Key not found") {
				// key not found will not be an error any more,but caller should judge that whether data is empty
				return nil
			}
			log.Panic(err.Error() + " key is: " + id)
		}
		_ = item.Value(func(val []byte) error {
			data = append([]byte{}, val...) // can't directly data = val according to doc of badger
			return nil
		})
		return nil
	})
	return
}

func KeyOnlyPrefixScan(keyPrefix string) *[][]byte {
	var resultKeys [][]byte
	// key only scan is more and more faster than normal scan, so if only need keys, use it.
	prefix := []byte(keyPrefix)
	_ = db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchValues = false
		it := txn.NewIterator(opts)
		defer it.Close()
		for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {
			var key []byte
			item := it.Item()
			key = append(key, item.Key()...) // key is the same with value, must copy, can't directly append
			resultKeys = append(resultKeys, key)

		}
		return nil
	})
	return &resultKeys
}

func ValueAndPrefixScan(keyPrefix string, filterFn ConditionFilter, firstFlag bool) map[string][]byte {
	prefix := []byte(keyPrefix)
	results := make(map[string][]byte)
	_ = db.View(func(txn *badger.Txn) error {
		it := txn.NewIterator(badger.DefaultIteratorOptions)
		defer it.Close()
		for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {
			item := it.Item()
			k := item.Key()
			var key []byte
			key = append(key, k...)
			err := item.Value(func(v []byte) error {
				if filterFn(v) { // if filter work, than get in results
					valueCopy := append([]byte{}, v...)
					results[string(key)] = valueCopy
					if firstFlag {
						return errors.New("get a result")
					}
				}
				return nil
			})
			if err != nil {
				// the only situation that err is not nil is that firstFlag is True and got a value
				break
			}
		}
		return nil
	})
	return results
}

func ValueAllDbScan(filterFn ConditionFilter, firstFlag bool) map[string][]byte {
	results := make(map[string][]byte)
	_ = db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchSize = config.PrefetchSize
		it := txn.NewIterator(opts)
		defer it.Close()
		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			k := item.Key()
			var key []byte
			key = append(key, k...)
			err := item.Value(func(v []byte) error {
				if filterFn(v) { // if filter work, than get in results
					valueCopy := append([]byte{}, v...)
					results[string(key)] = valueCopy
					if firstFlag {
						return errors.New("get a result")
					}
				}
				return nil
			})
			if err != nil {
				// the only situation that err is not nil is that firstFlag is True and got a value
				break
			}
		}
		return nil
	})
	return results
}

func CheckAllKey() {
	err := db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchSize = 10
		it := txn.NewIterator(opts)
		defer it.Close()
		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			k := item.Key()
			err := item.Value(func(v []byte) error {
				log.Println("key=", string(k))
				return nil
			})
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		log.Panic(err)
	}
}

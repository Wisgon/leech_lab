package main

import (
	"encoding/json"
	"graph_robot/utils"
	"log"

	"github.com/dgraph-io/badger/v4"
)

type ObjectTest struct {
	BBB string `json:"bbb"`
}

type TestStruct struct {
	Abc int        `json:"abc"`
	Obj ObjectTest `json:"obj"`
}

func (n *TestStruct) Struct2Byte() []byte {
	nb, err := json.Marshal(n)
	if err != nil {
		log.Panic("json marshal error: " + err.Error())
	}
	return nb
}

func (n *TestStruct) Byte2Struct(neureByte []byte) {
	err := json.Unmarshal(neureByte, n)
	if err != nil {
		log.Panic("json unmarshal error: " + err.Error())
	}
}

func main() {
	// Open the Badger database located in the /path/to/project/experiments/badger_lab/datas directory.
	// It will be created if it doesn't exist.
	db, err := badger.Open(badger.DefaultOptions(utils.GetProjectRoot() + "/experiments/badger_lab/datas"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// Your code here…

	// Start a writable transaction.
	txn := db.NewTransaction(true)
	defer txn.Discard()

	// Use the transaction...
	testStr := TestStruct{
		Abc: 55,
	}
	err = txn.Set([]byte("aaa1"), []byte(testStr.Struct2Byte()))
	if err != nil {
		log.Panic(err)
	}

	err = txn.Set([]byte("empty"), []byte{})
	if err != nil {
		log.Panic(err)
	}
	// err = db.Update(func(txn *badger.Txn) error {
	// 	err := txn.Delete([]byte("aaa2"))
	// 	if err != nil {
	// 		log.Panic(err)
	// 	}
	// 	return nil
	// })
	// if err != nil {
	// 	log.Panic(err)
	// }

	// Commit the transaction and check for error.
	if err := txn.Commit(); err != nil {
		log.Panic(err)
	}

	txn2 := db.NewTransaction(true)
	defer txn2.Discard()
	item, err := txn2.Get([]byte("aaa1"))
	if err != nil {
		log.Panic(err)
	}
	log.Println("before get value~~~")
	_ = item.Value(func(val []byte) error {
		st := TestStruct{}
		st.Byte2Struct(val)
		log.Println("item:", st.Abc)
		return nil
	})
	log.Println("after get value~~~")

	txn2.Set([]byte("not_commit3"), []byte("aaa"))
	item, err = txn2.Get([]byte("not_commit3"))
	if err != nil {
		log.Panic(err)
	}
	log.Println("before get value~~~")
	_ = item.Value(func(val []byte) error {
		log.Println("val:@@@", string(val))
		return nil
	})
	log.Println("after get value~~~")

	_ = db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchSize = 10
		it := txn.NewIterator(opts)
		defer it.Close()
		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			k := item.Key()
			err := item.Value(func(v []byte) error {
				log.Printf("key=%s, value=%s\n", k, v)
				return nil
			})
			if err != nil {
				return err
			}
		}
		return nil
	})

	_ = db.View(func(txn *badger.Txn) error {
		it := txn.NewIterator(badger.DefaultIteratorOptions)
		defer it.Close()
		prefix := []byte("empty")
		for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {
			item := it.Item()
			k := item.Key()
			err := item.Value(func(v []byte) error {
				log.Println("~~~~~empty:", len(v))
				log.Printf("key=%s, value=%s\n", k, v)
				return nil
			})
			if err != nil {
				return err
			}
		}
		return nil
	})

	seq, err := db.GetSequence([]byte("aaa"), 1024)
	if err != nil {
		log.Panic(err)
	}
	defer seq.Release()
	num, err := seq.Next()
	_, _ = seq.Next()
	if err != nil {
		log.Panic(err)
	}
	log.Println("num:!!!!:", num)
	// var i = 0
	// for {
	// 	num, err := seq.Next()
	// 	if err != nil {
	// 		log.Panic(err)
	// 	}
	// 	if err = txn2.Set([]byte("aaa"+fmt.Sprint(num)), []byte("aaa")); err == badger.ErrTxnTooBig {
	// 		log.Println("~~~i:", i) // test max i, result is more than 100000
	// 		_ = txn2.Commit()
	// 		break
	// 	}
	// 	i += 1
	// }

}

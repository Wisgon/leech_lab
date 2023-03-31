package database

import (
	"graph_robot/config"

	"github.com/dgraph-io/badger/v4"
)

type ManuallyTransaction struct {
	transactionNumber int         // the operation number of writes/delete transaction
	txn               *badger.Txn // it must be the same txn that can get value which not be commited yet.
}

func (m *ManuallyTransaction) Init() {
	m.txn = db.NewTransaction(true)
}

func (m *ManuallyTransaction) Create(keyPrefix string, data []byte) string {
	if m.transactionNumber > config.FixedTransactionNum {
		err := m.txn.Commit()
		if err != nil {
			panic(err)
		}
	}
	uniqueNum := GetSeqNum(keyPrefix)
	key := keyPrefix + uniqueNum
	err := m.txn.Set([]byte(key), data)
	if err != nil {
		panic(err)
	}
	m.transactionNumber += 1
	return key
}

func (m *ManuallyTransaction) Update(key string, data []byte) {
	if m.transactionNumber > config.FixedTransactionNum {
		err := m.txn.Commit()
		if err != nil {
			panic(err)
		}
	}
	err := m.txn.Set([]byte(key), data)
	if err != nil {
		panic(err)
	}
	m.transactionNumber += 1
}

func (m *ManuallyTransaction) Delete(key string) {
	if m.transactionNumber > config.FixedTransactionNum {
		err := m.txn.Commit()
		if err != nil {
			panic(err)
		}
	}
	err := m.txn.Delete([]byte(key))
	if err != nil {
		panic(err)
	}
	m.transactionNumber += 1
}

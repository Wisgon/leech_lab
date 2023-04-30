package database

// for now, these struct is not using

// type ManuallyTransaction struct {
// 	transactionNumber int         // the operation number of writes/delete transaction
// 	txn               *badger.Txn // it must be the same txn that can get value which not be commited yet.
// }

// func (m *ManuallyTransaction) Init() {
// 	m.txn = db.NewTransaction(true)
// }

// func (m *ManuallyTransaction) Create(key string, data []byte) string {
// 	if m.transactionNumber > config.FixedTransactionNum {
// 		err := m.txn.Commit()
// 		if err != nil {
// 			log.Panic(err)
// 		}
// 		m.Init()
// 		m.transactionNumber = 0
// 	}
// 	if err := m.txn.Set([]byte(key), data); err == badger.ErrTxnTooBig {
// 		_ = m.txn.Commit()
// 		m.Init()
// 		err = m.txn.Set([]byte(key), data)
// 		if err != nil {
// 			log.Panic(err)
// 		}
// 	}
// 	m.transactionNumber += 1
// 	return key
// }

// func (m *ManuallyTransaction) Update(key string, data []byte) {
// 	if m.transactionNumber > config.FixedTransactionNum {
// 		err := m.txn.Commit()
// 		if err != nil {
// 			log.Panic(err)
// 		}
// 		m.Init()
// 		m.transactionNumber = 0
// 	}
// 	if err := m.txn.Set([]byte(key), data); err == badger.ErrTxnTooBig {
// 		_ = m.txn.Commit()
// 		m.Init()
// 		err = m.txn.Set([]byte(key), data)
// 		if err != nil {
// 			log.Panic(err)
// 		}
// 	}
// 	m.transactionNumber += 1
// }

// func (m *ManuallyTransaction) Delete(key string) {
// 	if m.transactionNumber > config.FixedTransactionNum {
// 		err := m.txn.Commit()
// 		if err != nil {
// 			log.Panic(err)
// 		}
// 		m.Init()
// 		m.transactionNumber = 0
// 	}
// 	if err := m.txn.Delete([]byte(key)); err == badger.ErrTxnTooBig {
// 		_ = m.txn.Commit()
// 		m.Init()
// 		err = m.txn.Delete([]byte(key))
// 		if err != nil {
// 			log.Panic(err)
// 		}
// 	}
// 	m.transactionNumber += 1
// }

// func (m *ManuallyTransaction) ManuallyCommit() {
// 	err := m.txn.Commit()
// 	if err != nil {
// 		log.Panic(err)
// 	}
// }

// func (m *ManuallyTransaction) Close() {
// 	m.txn.Discard()
// }

package iota

import (
     "fmt"
     "strconv"
 	   badger "github.com/dgraph-io/badger"
)

type readfunc func(value []byte) error

func Store() {
  db, err := badger.Open(badger.DefaultOptions("./iota/store"))
  if err != nil {
	  fmt.Println("error", err)
  }
  defer db.Close()

  const value = 77
  Update(db, "answer", strconv.Itoa(value))

  var result string
	if err := Read(db, "answer", func(value []byte) error {
		result = string(value)
		return nil
	}); err != nil {
		fmt.Println("error", err)
	}
	if result != "" {
		fmt.Println("result", result)
	}

}

func Update(db *badger.DB, key string, value string) error {
  err := db.Update(func(txn *badger.Txn) error {
    err := txn.Set([]byte(key), []byte(value))
    return err
  })

  if err != nil {
    return err
  }
  return nil
}

func Read(db *badger.DB, key string, readFunc readfunc) error {
	return db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))
		if err == badger.ErrKeyNotFound {
			return readFunc(nil)
		}
		if err != nil {
			return err
		}
    value, err := item.ValueCopy(nil)

		if err != nil {
			return err
		}
		return readFunc(value)
	})
}

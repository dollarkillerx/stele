package main

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/dgraph-io/badger/v2"
)

func main() {
	log.SetFlags(log.Llongfile | log.LstdFlags)

	db, err := badger.Open(badger.DefaultOptions("./badger.db"))
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	// 更新
	if err := db.Update(func(txn *badger.Txn) error {
		return txn.Set([]byte("Hello"), []byte("world"))
	}); err != nil {
		log.Fatalln(err)
	}

	get(db)
	// 批量更新
	batchSet(db)
	// 设置过期事件
	setTll(db)
	// 遍历所有的Key
	iteratingOverKeys(db)
	// 前缀扫描
	prefix(db)
	// all key
	allKey(db)
}

func get(db *badger.DB) {
	err := db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte("scp"))
		if err != nil {
			log.Println(err)
			return err
		}

		return item.Value(func(val []byte) error {
			log.Println(string(val))
			return nil
		})
	})
	log.Println(errors.Is(err, badger.ErrKeyNotFound))
	//log.Println(errors.As(err, badger.ErrKeyNotFound))
}

func prefix(db *badger.DB) {
	if err := db.View(func(txn *badger.Txn) error {
		it := txn.NewIterator(badger.DefaultIteratorOptions)
		defer it.Close()
		prefix := []byte("sp")
		for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {
			item := it.Item()
			k := it.Item()
			if err := item.Value(func(v []byte) error {
				fmt.Printf("key=%s, value=%s\n", k, v)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		log.Fatalln(err)
	}
}

func allKey(db *badger.DB) {
	if err := db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchValues = false
		it := txn.NewIterator(opts)
		defer it.Close()
		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			k := item.Key()
			fmt.Printf("key=%s\n", k)
		}
		return nil
	}); err != nil {
		log.Fatalln(err)
	}
}

func batchSet(db *badger.DB) {
	batch := db.NewWriteBatch()
	defer batch.Cancel()
	testP := map[string]string{
		"sp1":  "sssp1xsxs",
		"sp2":  "sssp2xsxs",
		"sp3":  "sssp3xxs",
		"sp4":  "sssp4xs",
		"sp5":  "sssp5",
		"sp6":  "sssp5",
		"sp7":  "sssp151515",
		"sp8":  "sssp5151",
		"sp9":  "sssp5sdsad",
		"sp10": "sssp5asdsad",
		"sp11": "sssp5sds",
		"sp12": "sssp5",
		"sp13": "sssp5",
	}
	for k, v := range testP {
		if k == "sp7" {
			return
		}
		if err := batch.Set([]byte(k), []byte(v)); err != nil {
			log.Println(err)
		}
	}
	if err := batch.Flush(); err != nil {
		log.Fatalln(err)
	}
}

func setTll(db *badger.DB) {
	if err := db.Update(func(txn *badger.Txn) error {
		ent := badger.NewEntry([]byte("this_is_key"), []byte("this is value")).WithTTL(time.Second * 3)
		return txn.SetEntry(ent)
	}); err != nil {
		log.Println(err)
	}
}

func iteratingOverKeys(db *badger.DB) {
	if err := db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchSize = 10
		it := txn.NewIterator(opts)
		defer it.Close()
		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			k := item.Key()
			if err := item.Value(func(v []byte) error {
				fmt.Printf("Key=%s,value=%s\n", k, v)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		log.Println(err)
	}
}

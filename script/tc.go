package script

import (
	"fmt"
	"github.com/dgraph-io/badger/v2"
	"log"
	"time"
)

func main() {
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
		"sp1":  "sssp1",
		"sp2":  "sssp2",
		"sp3":  "sssp3",
		"sp4":  "sssp4",
		"sp5":  "sssp5",
		"sp6":  "sssp5",
		"sp7":  "sssp5",
		"sp8":  "sssp5",
		"sp9":  "sssp5",
		"sp10": "sssp5",
		"sp11": "sssp5",
		"sp12": "sssp5",
		"sp13": "sssp5",
	}
	for k, v := range testP {
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

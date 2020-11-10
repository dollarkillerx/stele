package stele

import (
	"errors"
	"log"
	"time"

	"github.com/dgraph-io/badger/v2"
	"github.com/dollarkillerx/stele/rpc/generate"
)

type Local struct {
	db *badger.DB
}

func NewLocal(storagePath string) (*Local, error) {
	var local Local
	db, err := badger.Open(badger.DefaultOptions(storagePath))
	if err != nil {
		return nil, err
	}
	local.db = db
	return &local, nil
}

func (l *Local) Set(key, val []byte, tll int64) error {
	return l.db.Update(func(txn *badger.Txn) error {
		entry := badger.NewEntry(key, val)
		if tll != 0 {
			entry = entry.WithTTL(time.Duration(tll))
		}
		return txn.SetEntry(entry)
	})
}
func (l *Local) Delete(key []byte) error {
	return l.db.Update(func(txn *badger.Txn) error {
		return txn.Delete(key)
	})
}

func (l *Local) Get(key []byte) (value []byte, err error) {
	err = l.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get(key)
		if err != nil {
			return err
		}

		return item.Value(func(val []byte) error {
			value = val
			return nil
		})
	})
	if err != nil {
		if errors.Is(err, badger.ErrKeyNotFound) {
			return nil, NotFund
		}
	}

	return value, nil
}

// Batch Insertion Failure Rollback. 批量插入 失败 回滚
func (l *Local) BatchSet(kvs []*generate.SteleKV) error {
	batch := l.db.NewWriteBatch()
	defer batch.Cancel()

	for k := range kvs {
		entry := badger.NewEntry(kvs[k].Key, kvs[k].Val)
		if kvs[k].Ttl != 0 {
			entry.WithTTL(time.Duration(kvs[k].Ttl))
		}
		if err := batch.SetEntry(entry); err != nil {
			return err
		}
	}

	return batch.Flush()
}

// Batch Get. 批量查询
func (l *Local) BatchGet(keys [][]byte) (kvs []*generate.SteleKV, err error) {
	kvs = []*generate.SteleKV{}

	for k := range keys {
		val, err := l.Get(keys[k])
		if err != nil {
			log.Println("BatchGet Error: ", err)
			continue
		}
		if val == nil {
			continue
		}
		kvs = append(kvs, &generate.SteleKV{
			Key: keys[k],
			Val: val,
		})
	}

	return kvs, nil
}

// Iterate over all keys. 遍历所有Key
func (l *Local) IterateKeys() (keys [][]byte, err error) {
	keys = [][]byte{}

	l.db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchValues = false
		it := txn.NewIterator(opts)
		defer it.Close()
		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			k := item.Key()
			keys = append(keys, k)
		}
		return nil
	})

	return keys, nil
}

// Iterate over keys and values. 遍历Key和value
func (l *Local) IterateKeysAndValues() (kvs []*generate.SteleKV, err error) {
	kvs = []*generate.SteleKV{}

	l.db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchSize = 10
		it := txn.NewIterator(opts)
		defer it.Close()
		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			k := item.Key()
			if err := item.Value(func(v []byte) error {
				kvs = append(kvs, &generate.SteleKV{
					Key: k,
					Val: v,
				})
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	})

	return kvs, nil
}

// Prefix Scan. 前缀扫描
func (l *Local) PrefixScan(prefix []byte) (kvs []*generate.SteleKV, err error) {
	kvs = []*generate.SteleKV{}

	l.db.View(func(txn *badger.Txn) error {
		it := txn.NewIterator(badger.DefaultIteratorOptions)
		defer it.Close()
		for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {
			item := it.Item()
			k := it.Item()
			if err := item.Value(func(v []byte) error {
				kvs = append(kvs, &generate.SteleKV{
					Key: k.Key(),
					Val: v,
				})
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	})

	return kvs, nil
}

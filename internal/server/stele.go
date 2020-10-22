package server

import (
	"context"
	"errors"
	"github.com/dollarkillerx/stele/pkg/stele"
	"log"
	"time"

	"github.com/dgraph-io/badger/v2"
	"github.com/dollarkillerx/stele/internal/config"
	"github.com/dollarkillerx/stele/rpc/generate"
)

type SteleServer struct {
	cfg *config.BaseConfig
	db  *badger.DB
}

func NewStele(cfg *config.BaseConfig) *SteleServer {
	var server SteleServer
	db, err := badger.Open(badger.DefaultOptions(cfg.StoragePath))
	if err != nil {
		log.Fatalln(err)
	}
	server.db = db
	server.cfg = cfg
	return &server
}

func (s *SteleServer) Set(ctx context.Context, req *generate.SteleKV) (resp *generate.SteleStatus, err error) {
	return &generate.SteleStatus{}, s.db.Update(func(txn *badger.Txn) error {
		entry := badger.NewEntry(req.Key, req.Val)
		if req.Ttl != 0 {
			entry = entry.WithTTL(time.Duration(req.Ttl))
		}
		return txn.SetEntry(entry)
	})
}

func (s *SteleServer) Get(ctx context.Context, req *generate.SteleK) (resp *generate.SteleVal, err error) {
	resp = &generate.SteleVal{}
	val, err := s.get(req.Key)
	if err != nil {
		return nil, err
	}
	if val == nil {
		return nil, stele.NotFund
	}
	resp.Val = val
	return resp, nil
}

// Batch Insertion Failure Rollback. 批量插入 失败 回滚
func (s *SteleServer) BatchSet(ctx context.Context, req *generate.BatchSetKVs) (resp *generate.SteleStatus, err error) {
	batch := s.db.NewWriteBatch()
	defer batch.Cancel()

	for k := range req.Kvs {
		entry := badger.NewEntry(req.Kvs[k].Key, req.Kvs[k].Val)
		if req.Kvs[k].Ttl != 0 {
			entry.WithTTL(time.Duration(req.Kvs[k].Ttl))
		}
		if err := batch.SetEntry(entry); err != nil {
			return nil, err
		}
	}

	return &generate.SteleStatus{}, batch.Flush()
}

// Batch Get. 批量查询
func (s *SteleServer) BatchGet(ctx context.Context, req *generate.SteleKS) (resp *generate.BatchSetKVs, err error) {
	var kvs []*generate.SteleKV
	for k := range req.Keys {
		val, err := s.get(req.Keys[k])
		if err != nil {
			log.Println("BatchGet Error: ", err)
			continue
		}
		if val == nil {
			continue
		}
		kvs = append(kvs, &generate.SteleKV{
			Key: req.Keys[k],
			Val: val,
		})
	}

	return &generate.BatchSetKVs{Kvs: kvs}, nil
}

// Iterate over all keys. 遍历所有Key
func (s *SteleServer) IterateKeys(ctx context.Context, req *generate.SteleRequest) (resp *generate.SteleKS, err error) {
	var keys [][]byte
	s.db.View(func(txn *badger.Txn) error {
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

	return &generate.SteleKS{Keys: keys}, nil
}

// Iterate over keys and values. 遍历Key和value
func (s *SteleServer) IterateKeysAndValues(ctx context.Context, req *generate.SteleRequest) (resp *generate.BatchSetKVs, err error) {
	var kvs []*generate.SteleKV
	s.db.View(func(txn *badger.Txn) error {
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

	return &generate.BatchSetKVs{Kvs: kvs}, nil
}

// Prefix Scan. 前缀扫描
func (s *SteleServer) PrefixScan(ctx context.Context, req *generate.Prefix) (resp *generate.BatchSetKVs, err error) {
	var kvs []*generate.SteleKV

	s.db.View(func(txn *badger.Txn) error {
		it := txn.NewIterator(badger.DefaultIteratorOptions)
		defer it.Close()
		for it.Seek(req.Prefix); it.ValidForPrefix(req.Prefix); it.Next() {
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

	return &generate.BatchSetKVs{Kvs: kvs}, nil
}

func (s *SteleServer) get(key []byte) (value []byte, err error) {
	err = s.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get(key)
		if err != nil {
			return err
		}

		return item.Value(func(val []byte) error {
			value = val
			return nil
		})
	})
	if errors.Is(err, badger.ErrKeyNotFound) {
		return nil, nil
	}
	return value, nil
}

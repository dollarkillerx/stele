package server

import (
	"context"
	"log"

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
	s.db.Update(func(txn *badger.Txn) error {

	})
}

func (s *SteleServer) Get(ctx context.Context, req *generate.SteleK) (resp *generate.SteleVal, err error) {

}

// Batch Insertion Failure Rollback. 批量插入 失败 回滚
func (s *SteleServer) BatchSet(ctx context.Context, req *generate.BatchSetKVs) (resp *generate.SteleStatus, err error) {

}

// Batch Get. 批量查询
func (s *SteleServer) BatchGet(ctx context.Context, req *generate.SteleKS) (resp *generate.BatchSetKVs, err error) {

}

// Iterate over all keys. 遍历所有Key
func (s *SteleServer) IterateKeys(ctx context.Context, req *generate.SteleRequest) (resp *generate.SteleKS, err error) {

}

// Iterate over keys and values. 遍历Key和value
func (s *SteleServer) IterateKeysAndValues(ctx context.Context, req *generate.SteleRequest) (resp *generate.BatchSetKVs, err error) {

}

// Prefix Scan. 前缀扫描
func (s *SteleServer) PrefixScan(ctx context.Context, req *generate.Prefix) (resp *generate.BatchSetKVs, err error) {

}

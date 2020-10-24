package server

import (
	"context"
	"github.com/dollarkillerx/stele/internal/config"
	"github.com/dollarkillerx/stele/rpc/generate"
	"log"
)

type SteleServer struct {
	cfg *config.BaseConfig
	db  *Local
}

func NewStele(cfg *config.BaseConfig) *SteleServer {
	var server SteleServer

	local, err := NewLocal(cfg.StoragePath)
	if err != nil {
		log.Fatalln(err)
	}
	server.db = local
	server.cfg = cfg
	return &server
}

func (s *SteleServer) Set(ctx context.Context, req *generate.SteleKV) (resp *generate.SteleStatus, err error) {
	return &generate.SteleStatus{}, s.db.Set(req.Key, req.Val, req.Ttl)
}

func (s *SteleServer) Get(ctx context.Context, req *generate.SteleK) (resp *generate.SteleVal, err error) {
	resp = &generate.SteleVal{}
	val, err := s.db.Get(req.Key)
	if err != nil {
		return nil, err
	}

	resp.Val = val
	return resp, nil
}

// Batch Insertion Failure Rollback. 批量插入 失败 回滚
func (s *SteleServer) BatchSet(ctx context.Context, req *generate.BatchSetKVs) (resp *generate.SteleStatus, err error) {
	return &generate.SteleStatus{}, s.db.BatchSet(req.Kvs)
}

// Batch Get. 批量查询
func (s *SteleServer) BatchGet(ctx context.Context, req *generate.SteleKS) (resp *generate.BatchSetKVs, err error) {
	kvs, err := s.db.BatchGet(req.Keys)
	if err != nil {
		return nil, err
	}
	return &generate.BatchSetKVs{Kvs: kvs}, nil
}

// Iterate over all keys. 遍历所有Key
func (s *SteleServer) IterateKeys(ctx context.Context, req *generate.SteleRequest) (resp *generate.SteleKS, err error) {
	keys, err := s.db.IterateKeys()
	if err != nil {
		return nil, err
	}
	return &generate.SteleKS{Keys: keys}, nil
}

// Iterate over keys and values. 遍历Key和value
func (s *SteleServer) IterateKeysAndValues(ctx context.Context, req *generate.SteleRequest) (resp *generate.BatchSetKVs, err error) {
	kvs, err := s.db.IterateKeysAndValues()
	if err != nil {
		return nil, err
	}

	return &generate.BatchSetKVs{Kvs: kvs}, nil
}

// Prefix Scan. 前缀扫描
func (s *SteleServer) PrefixScan(ctx context.Context, req *generate.Prefix) (resp *generate.BatchSetKVs, err error) {
	kvs, err := s.db.PrefixScan(req.Prefix)
	if err != nil {
		return nil, err
	}

	return &generate.BatchSetKVs{Kvs: kvs}, nil
}

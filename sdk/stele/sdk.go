package stele

import (
	"context"
	"strings"
	"time"

	"github.com/dollarkillerx/stele/pkg/stele"
	"github.com/dollarkillerx/stele/rpc/generate"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type Stele struct {
	stele generate.SteleClient
}

func New(socketAddr string, username string, password string) (*Stele, error) {
	creds, err := credentials.NewClientTLSFromFile("cert/servermpe", "stele")
	if err != nil {
		return nil, err
	}

	dial, err := grpc.Dial(socketAddr, grpc.WithTransportCredentials(creds), grpc.WithPerRPCCredentials(&loginCreds{
		Username: username,
		Password: password,
	}))
	if err != nil {
		return nil, err
	}

	client := generate.NewSteleClient(dial)
	return &Stele{stele: client}, nil
}

// Set
// Params: Key, Val, Ttl: A zero is permanent.
func (s *Stele) Set(key, val []byte, ttl time.Duration) error {
	_, err := s.stele.Set(context.Background(), &generate.SteleKV{Key: key, Val: val, Ttl: int64(ttl)})
	return err
}

// Get
// Params: key
// Return: val: There is no such thing as a null.
func (s *Stele) Get(key []byte) (val []byte, err error) {
	resp, err := s.stele.Get(context.Background(), &generate.SteleK{Key: key})
	if err != nil {
		return nil, err
	}
	if resp == nil {
		return nil, nil
	}
	return resp.Val, nil
}

// Batch Insertion Failure Rollback. 批量插入 失败 回滚
// Params: datas: []map[key]val
func (s *Stele) NewBatchSet() *BatchSet {
	return &BatchSet{
		stele: s.stele,
		data:  []*generate.SteleKV{},
	}
}

// Batch Get. 批量查询
func (s *Stele) BatchGet(keys [][]byte) (kvs []*generate.SteleKV, err error) {
	result, err := s.stele.BatchGet(context.Background(), &generate.SteleKS{Keys: keys})
	if err != nil {
		return nil, err
	}

	return result.Kvs, nil
}

// Iterate over all keys. 遍历所有Key
func (s *Stele) IterateKeys() (keys [][]byte, err error) {
	values, err := s.stele.IterateKeys(context.Background(), &generate.SteleRequest{})
	if err != nil {
		return nil, err
	}
	return values.Keys, nil
}

// Iterate over keys and values. 遍历Key和value
func (s *Stele) IterateKeysAndValues() (kvs []*generate.SteleKV, err error) {
	values, err := s.stele.IterateKeysAndValues(context.Background(), &generate.SteleRequest{})
	if err != nil {
		return nil, err
	}

	return values.Kvs, nil
}

// Prefix Scan. 前缀扫描
// Params: prefix: Prefix of key
func (s *Stele) PrefixScan(prefix []byte) (kvs []*generate.SteleKV, err error) {
	values, err := s.stele.PrefixScan(context.Background(), &generate.Prefix{Prefix: prefix})
	if err != nil {
		return nil, err
	}

	return values.Kvs, nil
}

type loginCreds struct {
	Username, Password string
}

func (c *loginCreds) GetRequestMetadata(context.Context, ...string) (map[string]string, error) {
	return map[string]string{
		"username": c.Username,
		"password": c.Password,
	}, nil
}

func (c *loginCreds) RequireTransportSecurity() bool {
	return true
}

type BatchSet struct {
	stele generate.SteleClient
	data  []*generate.SteleKV
}

func (b *BatchSet) Set(key, val []byte, ttl time.Duration) {
	b.data = append(b.data, &generate.SteleKV{Key: key, Val: val, Ttl: int64(ttl)})
}

func (b *BatchSet) Flash() error {
	_, err := b.stele.BatchSet(context.Background(), &generate.BatchSetKVs{Kvs: b.data})
	return err
}

func IsNotFund(err error) bool {
	if strings.Index(err.Error(), stele.NotFund.Error()) == -1 {
		return false
	}
	return true
}

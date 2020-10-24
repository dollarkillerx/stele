# Stele (Stele Distributed KV based on badger)
### 目标是实现通过配置 灵活切换CP AP 模式的分布式KV

### 分支说明
- main 为分布式版本
- node 为单节点badger的grpc包装
- explore 系统设计的探索

### Node 单节点版本
example
```go
package main

import (
	"log"
	"time"

	"github.com/dollarkillerx/stele/sdk/stele"
)

func main() {
	log.SetFlags(log.Llongfile | log.LstdFlags)

	db, err := stele.New("127.0.0.1:9695", "root", "root")
	if err != nil {
		log.Fatalln(err)
	}

	if err := db.Set([]byte("hello"), []byte("world"), 0); err != nil {
		log.Fatalln(err)
	}

	if val, err := db.Get([]byte("hello")); err != nil {
		log.Fatalln(err)
	} else {
		log.Printf("val: %s \n", val)
	}

	batch := db.NewBatchSet()
	testMap := map[string]string{
		"day1":  "day1",
		"day2":  "day2",
		"day3":  "day3",
		"day4":  "day4",
		"day5":  "day5",
		"day6":  "day6",
		"day7":  "day7",
		"day8":  "day8",
		"day9":  "day9",
		"day10": "day10",
		"day11": "day11",
	}
	for k, v := range testMap {
		batch.Set([]byte(k), []byte(v), 0)
	}
	if err := batch.Flash(); err != nil {
		log.Fatalln(err)
	}

	if keys, err := db.IterateKeys(); err != nil {
		log.Fatalln(err)
	} else {
		for _, v := range keys {
			log.Printf("key: %s \n", v)
		}
	}

	if values, err := db.IterateKeysAndValues(); err != nil {
		log.Fatalln(err)
	} else {
		for _, v := range values {
			log.Printf("key: %s Val: %s \n", v.Key, v.Val)
		}
	}

	if scan, err := db.PrefixScan([]byte("day")); err != nil {
		log.Fatalln(err)
	} else {
		for _, v := range scan {
			log.Printf("key: %s Val: %s \n", v.Key, v.Val)
		}
	}

	k1 := []byte("k1")
	if err := db.Set(k1, []byte("a"), time.Second*3); err != nil {
		log.Fatalln(err)
	}

	time.Sleep(time.Second)
	if val, err := db.Get(k1); err != nil {
		log.Fatalln(err)
	} else {
		log.Printf("%s \n", val)
	}

	time.Sleep(time.Second * 3)
	val, err := db.Get(k1)
	if err != nil {
		if !stele.IsNotFund(err) {
			log.Fatalln(err)
		}
	}
	log.Println(val)
}
```

注意如果使用1.5+的go 需要配置`export GODEBUG=x509ignoreCN=0`  因为1.5证书的没有使用 Common Name 
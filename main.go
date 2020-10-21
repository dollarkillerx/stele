package main

import (
	"github.com/dgraph-io/badger"
	"log"
)

func main() {
	db, err := badger.Open(badger.DefaultOptions("./badger.db"))
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	db.
}

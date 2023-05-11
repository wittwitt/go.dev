package main

import (
	"fmt"

	"github.com/syndtr/goleveldb/leveldb"
)

func main() {
	var err error

	db, err := leveldb.OpenFile("a.db", nil)
	defer db.Close()

	fmt.Println(err)

	data, err := db.Get([]byte("key"), nil)
	if err != nil {

	}

	fmt.Println(data)

	err = db.Put([]byte("key"), []byte("123"), nil)

	//	err = db.Delete([]byte("key"), nil)

}

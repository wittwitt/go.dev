package main

import (
	"fmt"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
)

func main() {
	var err error

	db, err := leveldb.OpenFile("test.db", nil)
	defer db.Close()

	fmt.Println(err)

	db.Put([]byte("foo-ab"), []byte("abc"), nil)
	db.Put([]byte("fooab"), []byte("ddd"), nil)

	{
		data, err := db.Get([]byte("key"), nil)
		if err != nil {

		}
		fmt.Println(data)

		err = db.Put([]byte("key"), []byte("123"), nil)
	}

	{
		iter := db.NewIterator(util.BytesPrefix([]byte("foo-")), nil)
		for iter.Next() {
			fmt.Println("value:", string(iter.Value()))
		}
		iter.Release()
		err = iter.Error()
	}

	{
		iter := db.NewIterator(&util.Range{Start: []byte("foo"), Limit: []byte("xoo")}, nil)
		for iter.Next() {

		}
		iter.Release()
		err = iter.Error()
	}

}

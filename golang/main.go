package main

import (
	"fmt"
	"strconv"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
)

func main() {
	db, err := leveldb.OpenFile("./testdb", nil)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	for i := 0; i < 256; i++ {
		key := "TestKey:" + strconv.Itoa(i)
		value := "TestValue:" + strconv.Itoa(i)
		err = db.Put([]byte(key), []byte(value), nil)
		if err != nil {
			panic(err)
		}
	}

	data, err := db.Get([]byte("TestKey:1"), nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))

	iter := db.NewIterator(util.BytesPrefix([]byte("TestKey:")), nil)
	for iter.Next() {
		fmt.Printf("%s ---> %s\n", string(iter.Key()), string(iter.Value()))
	}
	iter.Release()
	err = iter.Error()
	if err != nil {
		panic(err)
	}
}

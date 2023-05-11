package main

import (
	"fmt"
	"log"
	"time"

	bolt "go.etcd.io/bbolt"
)

func main() {
	d1, err := bolt.Open(".test.db", 0666, nil)
	if err != nil {
		log.Printf("DataStore.Open: %v", err)
		return
	}

	d2, err := bolt.Open(".test.db", 0666, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Printf("DataStore.Open: %v", err)
		return
	}

	fmt.Println(d1, d2)

}

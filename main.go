package main

import (
	"fmt"
	"sync"
	"time"
)

var mutex = &sync.Mutex{}

type Db struct{}

var DbInstance *Db

func ConnectDb() *Db {
	if DbInstance == nil {
		mutex.Lock()
		defer mutex.Unlock()
		if DbInstance == nil {
			fmt.Println("Creating DBInstance now.")
			DbInstance = &Db{}
		} else {
			fmt.Println("DBInstance already created.")
		}
	} else {
		fmt.Println("DBInstance already created.")
	}
	return DbInstance
}

func main() {
	for i := 0; i < 5; i++ {
		go ConnectDb()
	}
	time.Sleep(1 * time.Second)
}

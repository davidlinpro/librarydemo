package main

import (
	"database/sql"
	"log"
	"math/rand"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	DBConnStr = "library:library@tcp(db:3306)/library?parseTime=true"
)

var (
	dbConn *sql.DB
	dbLock sync.Mutex
)

func GetDB() *sql.DB {
	var err error
	// lock and try to get connection
	dbLock.Lock()
	// if no connection then reset connection
	if dbConn == nil {
		dbConn, err = sql.Open("mysql", DBConnStr)
		if err != nil {
			log.Println(err.Error())
			dbLock.Unlock()
			time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
			return GetDB()
		}
	}
	// if yes connection but no ping, then reconnect
	err = dbConn.Ping()
	if err != nil {
		dbConn, err = sql.Open("mysql", DBConnStr)
		if err != nil {
			log.Println(err.Error())
			dbLock.Unlock()
			time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
			return GetDB()
		}
		// try one more time to ping, then give up
		err = dbConn.Ping()
		if err != nil {
			log.Println(err.Error())
			dbLock.Unlock()
			time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
			return GetDB()
		}
	}
	dbLock.Unlock()
	return dbConn
}

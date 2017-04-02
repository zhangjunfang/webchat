package mydb

import (
	"database/sql"
	"sync"
)

var Db *sql.DB = nil
var once sync.Once

func Instantiation(dataSourceName string, maxOpenConns, maxIdleConns int) {

	once.Do(func() {
		GetDB(dataSourceName, maxOpenConns, maxIdleConns)
	})
}

func GetDB(dataSourceName string, maxOpenConns, maxIdleConns int) (db *sql.DB, err error) {
	db, err = sql.Open("mysql", dataSourceName)
	if err == nil {
		db.SetMaxOpenConns(maxOpenConns)
		db.SetMaxIdleConns(maxIdleConns)
	}
	return db, err
}

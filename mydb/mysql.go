package mydb

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/zhangjunfang/webchat/myerror"
)

var db *sql.DB
var once sync.Once
var dbErr error

func init() {
	db, dbErr = sql.Open("mysql", "root:Cme0328@@tcp(10.0.4.245:3306)/im")
	if dbErr != nil {
		fmt.Println(" database  error :  ", dbErr)
	}
}

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
func DataStore(datas []string) {
	tx, err := db.Begin()
	if err != nil {
		fmt.Println(" tx  error : ", err)
		return
	}
	stmt, err := db.PrepareContext(context.Background(), "INSERT into im_message(senderId,receiverId,messageType,messageContent,createDate,expirationDate) values(?,?,?,?,?,?)")
	myerror.CheckError(err, "SQL 语法错误！！！")
	t := time.Now().Add(12 * 24 * time.Hour)
	//|001|002|003|004|005|006|
	senderId, err := strconv.Atoi(datas[3])
	myerror.CheckError(err, "数据转化错误！！！")
	receiverId, err := strconv.Atoi(datas[2])
	myerror.CheckError(err, "数据转化错误！！！")
	messageType, err := strconv.Atoi(datas[1])
	myerror.CheckError(err, "数据转化错误！！！")
	res, err := stmt.Exec(senderId, receiverId, byte(messageType), datas[5], time.Now(), t)
	myerror.CheckError(err, "数据执行错误！！！")
	id, err := res.LastInsertId()
	myerror.CheckError(err, "数据id标示错误！！！")
	fmt.Println(id)
	tx.Commit()
	stmt.Close()
}

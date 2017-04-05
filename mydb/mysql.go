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

var Db *sql.DB = nil
var once sync.Once
var err error

func init() {
	Db, dbErr = sql.Open("mysql", "root:Cme0328@@tcp(10.0.4.245:3306)/im")
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
	Db, err = sql.Open("mysql", dataSourceName)
	if err == nil {
		Db.SetMaxOpenConns(maxOpenConns)
		Db.SetMaxIdleConns(maxIdleConns)
	}
	return Db, err
}
func DataStore(datas []string) {
	tx, err := db.Begin()
	if err != nil {
		fmt.Println(" tx  error : ", err)
		return
	}
	//stmt, err := db.PrepareContext(context.Background(), "INSERT im_message SET senderId=?,receiverId=?,messageType=?,messageContent=?,createDate=?,expirationDate=? ")
	//插入数据
	//stmt, err := db.Prepare("INSERT im_message SET senderId=?,receiverId=?,messageType=?,messageContent=?,createDate=?,expirationDate=?")
	stmt, err := db.PrepareContext(context.Background(), "INSERT into im_message(senderId,receiverId,messageType,messageContent,createDate,expirationDate) values(?,?,?,?,?,?)")
	//stmt, err := db.Prepare("INSERT into im_message(name) values(?)")
	//stmt, err := db.PrepareContext(context.Background(), "INSERT into im_message(name) values(?)")
	myerror.CheckError(err, "SQL 语法错误！！！")
	t := time.Now().Add(12 * 24 * time.Hour)
	//|001|002|003|004|005|006|
	senderId, err := strconv.Atoi(datas[3])
	myerror.CheckError(err, "数据转化错误！！！")
	receiverId, err := strconv.Atoi(datas[2])
	myerror.CheckError(err, "数据转化错误！！！")
	messageType, err := strconv.Atoi(datas[1])
	myerror.CheckError(err, "数据转化错误！！！")
	//fmt.Println(senderId, receiverId, byte(messageType), datas[5], "===============", time.Now(), t)
	//stmt.Exec(datas[3], datas[2], datas[1], datas[5], time.Now(), t)
	//stmt.ExecContext(context.Background(), datas[3], datas[2], datas[1], datas[5], time.Now(), t)
	res, err := stmt.Exec(senderId, receiverId, byte(messageType), datas[5], time.Now(), t)
	//	res, err := stmt.Exec(senderId)
	//	stmt.Exec(senderId)
	myerror.CheckError(err, "数据执行错误！！！")
	id, err := res.LastInsertId()
	myerror.CheckError(err, "数据id标示错误！！！")
	fmt.Println(id)
	tx.Commit()
	stmt.Close()
}

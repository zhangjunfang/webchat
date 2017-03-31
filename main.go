package main

import (
	"context"
	"database/sql"
	"fmt"
	"net"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/garyburd/redigo/redis"
	_ "github.com/go-sql-driver/mysql"
	//"github.com/henrylee2cn/mahonia"
)

var (
	dbhostsip  = "127.0.0.1:3306" //IP地址
	dbusername = "root"           //用户名
	dbpassword = "123456"         //密码
	dbname     = "Test"           //表名
	pools      *redis.Pool
	dbErr      error
	db         *sql.DB
)

const MaxIdle = 100

func init() {
	pools = redis.NewPool(creatPool, MaxIdle)
	db, dbErr = sql.Open("mysql", "root:Cme0328@@tcp(10.0.4.245:3306)/im")
	if dbErr != nil {
		fmt.Println(" database  error :  ", dbErr)
	}
}
func creatPool() (c redis.Conn, err error) {
	c, err = redis.Dial("tcp", "10.0.6.222:6379", redis.DialConnectTimeout(10*time.Second), redis.DialDatabase(0))
	if err != nil {
		fmt.Println("error====: ", err)
		return
	}
	return c, err
}
func main() {
	runtime.GOMAXPROCS(4)
	l, err := net.Listen("tcp", ":9999")
	defer l.Close()
	if err != nil {
		fmt.Printf("Failure to listen: %s\n", err.Error())
		return
	}
	for {
		if conn, err := l.Accept(); err == nil {
			go Connnection(conn) //new thread
		}
	}

}
func Connnection(conn net.Conn) {
	if conn != nil {
		b := make([]byte, 64)
		//var buff bytes.Buffer
		for {
			n, err := conn.Read(b)
			//n, err := conn.Read(buff.Bytes())
			if n == 0 && err == nil {
				continue
			}
			if err != nil {
				fmt.Printf("%s  ==== %d  :  %T \n", string(b[:n]), n, err)
				conn.Close()
				return
			}
			var m sync.Mutex
			m.Lock()
			testStr := string(b[:n])
			array := strings.Split(testStr, "|")
			if len(array) != 8 {
				fmt.Println("len====", len(array))
				continue
			}
			m.Unlock()
			//fmt.Println("message  data length : ", len(testStr), "------------------", n)
			//dec := mahonia.NewDecoder("gbk")
			//if ret, ok := dec.ConvertStringOK(testStr); ok {
			//fmt.Println("GBK to UTF-8: ", ret, " bytes:", b)
			//go
			//ParseData(conn, ret)
			//}
			go ParseData(conn, testStr)
		}
	}

}
func ParseData(conn net.Conn, data string) {
	if data == "" || len(strings.Trim(data, " ")) == 0 {
		return
	} else {
		array := strings.Split(data, "|")
		switch len(array) {
		case 8:
			{
				RedisConn(conn, array[3]+"-"+array[2]) //redis  sendId+ revicerId==>conn
				go DataStore(array)
				go DataSend(conn, array[3])
				break
			}
		default:
			{
				fmt.Println("The data format is not correct:", data)
				break
			}
		}
	}
}

//
func RedisConn(conn net.Conn, id string) {
	c := pools.Get()
	c.Do("set", id, conn)
}
func GetConnById(id string) (conn net.Conn) {
	c := pools.Get()
	reply, err := c.Do("get", id)
	if err != nil {
		fmt.Println("  redis  get data  error : ", err)
		return
	}
	conn, ok := reply.(net.Conn)
	if ok {
		return conn
	} else {
		return nil
	}
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
	checkErr(err)
	t := time.Now()
	t.Add(12 * 24 * time.Hour)
	//|001|002|003|004|005|006|
	senderId, err := strconv.Atoi(datas[3])
	checkErr(err)
	receiverId, err := strconv.Atoi(datas[2])
	checkErr(err)
	messageType, err := strconv.Atoi(datas[1])
	checkErr(err)
	//fmt.Println(senderId, receiverId, byte(messageType), datas[5], "===============", time.Now(), t)
	//stmt.Exec(datas[3], datas[2], datas[1], datas[5], time.Now(), t)
	//stmt.ExecContext(context.Background(), datas[3], datas[2], datas[1], datas[5], time.Now(), t)
	res, err := stmt.Exec(senderId, receiverId, byte(messageType), datas[5], time.Now(), t)
	//	res, err := stmt.Exec(senderId)
	//	stmt.Exec(senderId)
	checkErr(err)
	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Println(id)
	tx.Commit()
	stmt.Close()
}
func DataSend(conn net.Conn, id string) {
	conn.Write([]byte(id + ":::::::::::::::::::::::ok\r\n"))
}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

package connect

import (
	"fmt"
	"net"
	"strings"
	"sync"
	"time"

	"github.com/zhangjunfang/webchat/myerror"
)

func DataSend(conn net.Conn, id string) {
	conn.Write([]byte(id + ":::::::::::::::::::::::ok\r\n"))
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
func Connnection(conn net.Conn) {
	if conn != nil {
		b := make([]byte, 64)
		for {
			n, err := conn.Read(b)
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

func TickConnnection(conn net.Conn, c chan string) {
	t1 := time.NewTimer(time.Second * 2)
	b := make([]byte, 8)
	n, err := conn.Read(b)
	myerror.CheckError(err, "读取数据出错了！！！")
	c <- string(b[:n])
	for {
		select {
		case msg := <-c:
			fmt.Println(msg)
		case <-t1.C:
			conn.Write([]byte("|4|1|"))
			t1.Reset(time.Second * 2)
		}
	}
}
func TickTime(wg sync.WaitGroup) {
	l, err := net.Listen("tcp", ":9998")
	defer l.Close()
	if err != nil {
		fmt.Printf("Failure to listen: %s\n", err.Error())
		return
	}
	wg.Done()
	c := make(chan string, 10)
	for {
		if conn, err := l.Accept(); err == nil {
			go TickConnnection(conn, c) //new thread
		}
	}

}
func MainService(wg sync.WaitGroup) {

	l, err := net.Listen("tcp", ":9999")
	defer l.Close()
	if err != nil {
		fmt.Printf("Failure to listen: %s\n", err.Error())
		return
	}
	wg.Done()
	for {
		if conn, err := l.Accept(); err == nil {
			go Connnection(conn) //new thread
		}
	}
}

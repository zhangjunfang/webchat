package main

import (
	"fmt"
	"net"
	"strings"

	"github.com/henrylee2cn/mahonia"
)

func main() {
	fmt.Println("fsdfsdfsdsfdf=====================!")
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
			testStr := string(b)
			dec := mahonia.NewDecoder("gbk")
			if ret, ok := dec.ConvertStringOK(testStr); ok {
				fmt.Println("GBK to UTF-8: ", ret, " bytes:", b)
				go ParseData(ret)
			}
		}
		conn.Close()
	}

}
func ParseData(data string) {
	if data == "" || len(strings.Trim(data, " ")) == 0 {
		return
	} else {
		array := strings.Split(data, "|")
		switch len(array) {
		case 8:
			{
				fmt.Println("--", array[0], "---------------", array[1], "====", array[2])
				break
			}
		default:
			{
				fmt.Println("============================================================")
			}
		}
		fmt.Println(len(array))
	}
}

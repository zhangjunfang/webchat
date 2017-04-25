package redis

import (
	"fmt"
	"net"
	"time"

	"github.com/garyburd/redigo/redis"
)

var (
	pools *redis.Pool
)

const MaxIdle = 100

func init() {
	pools = redis.NewPool(creatPool, MaxIdle)

}
func RedisConn(conn net.Conn, id string) {
	c := pools.Get()
	c.Do("set", id, conn)
}
func creatPool() (c redis.Conn, err error) {
	c, err = redis.Dial("tcp", "10.0.6.222:6379", redis.DialConnectTimeout(10*time.Second), redis.DialDatabase(0))
	if err != nil {
		fmt.Println("error====: ", err)
		return
	}
	return c, err
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

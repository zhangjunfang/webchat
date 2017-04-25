package redis

import (
	"fmt"
	"testing"
	"time"
)

func Test_AA(t *testing.T) {
	c := pools.Get()
	c.Do("SUBSCRIBE", "redisChat")
}

func Test_BB(t *testing.T) {
	c := pools.Get()
	c.Do("PUBLISH", "redisChat", "sdfsdfsdfsdf")
	c.Do("PUBLISH", "redisChat", "hhhhhhhhhhhh")
	c.Do("PUBLISH", "redisChat", "eeeeeeeeeeeeeeeeeeeee")
	go func() {
		for {
			m, _ := c.Receive()
			fmt.Println(m)
		}

	}()
	time.Sleep(time)
}

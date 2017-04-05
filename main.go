package main

import (
	"database/sql"
	"fmt"
	"os"
	"runtime"
	"sync"

	"github.com/garyburd/redigo/redis"

	//"github.com/henrylee2cn/mahonia"
	log "github.com/YoungPioneers/blog4go"
)

var (
	pools *redis.Pool
	dbErr error
	db    *sql.DB
)

const MaxIdle = 100

type ImLog struct {
}

func (ImLog) Fire(level log.LevelType, message ...interface{}) {
	fmt.Println(message, "---", level)
}
func LogFunc() {
	// init a file write using xml config file
	err := log.NewWriterFromConfigAsFile("config.xml")
	if nil != err {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	//defer log.Close()
	// initialize your hook instance
	hook := new(ImLog)
	log.SetHook(hook) // writersFromConfig can be replaced with writers
	log.SetHookLevel(log.INFO)
	log.SetHookAsync(true) // hook will be called in async mode
	// optionally set output colored
	log.SetColored(true)
}

func main() {
	runtime.GOMAXPROCS(4)
	//	Ticker4Second(4, func() {
	//		fmt.Println("Ticker4Second::=========", 11111)
	//	})

	LogFunc()
	var wg sync.WaitGroup
	wg.Add(2)
	go MainService(wg)
	go TickTime(wg)
	wg.Wait()
}

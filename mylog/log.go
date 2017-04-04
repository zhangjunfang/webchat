package mylog

import (
	"fmt"
	"os"

	log "github.com/YoungPioneers/blog4go"
)

type ImLog struct {
	User string
}

func (log ImLog) Fire(level log.LevelType, message ...interface{}) {
	fmt.Println(message, "---", level)
}
func init() {
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

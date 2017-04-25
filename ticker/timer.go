package ticker

import (
	"time"

	log "bloggithub.com/YoungPioneers/blog4go"
)

func AsynTickerForSecond(second int, function func()) {
	defer func() {
		if err := recover(); err != nil {
			log.Info(" asyn Ticker4Second error :", err)
		}
	}()
	time.Sleep(time.Duration(second) * time.Second)
	timer := time.NewTicker(time.Duration(second) * time.Second)
	for {
		select {
		case <-timer.C:
			go function()
		}
	}
}
func SynTickerForSecond(second int, function func()) {
	defer func() {
		if err := recover(); err != nil {

			log.Info("syn  Ticker4Second error :", err)
		}
	}()
	time.Sleep(time.Duration(second) * time.Second)
	timer := time.NewTicker(time.Duration(second) * time.Second)
	for {
		select {
		case <-timer.C:
			function()
		}
	}
}

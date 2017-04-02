package ticker

import (
	"fmt"
	"time"
)

func AsynTickerForSecond(second int, function func()) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Ticker4Second error :", err)
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
			fmt.Println("Ticker4Second error :", err)
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

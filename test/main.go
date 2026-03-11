package main

import (
	"github.com/zhangyiming748/GracefullyExit"
	"time"
)

func main() {
	go GracefullyExit.StartReceivedExit()

	for i := 0; i < 100; i++ {
		println("程序运行中...")
		time.Sleep(3 * time.Second)
		if GracefullyExit.ShouldExit() {
			println("程序退出")
			return
		}
	}
}

package main

import (
	"log"
	"time"

	"github.com/zhangyiming748/GracefullyExit"
)

func main() {
	go GracefullyExit.StartReceivedExit()
	for i := 0; i < 100; i++ {
		log.Println("程序运行中...")
		time.Sleep(3 * time.Second)
		if GracefullyExit.ShouldExit() {
			log.Println("程序退出")
			return
		}
	}
}

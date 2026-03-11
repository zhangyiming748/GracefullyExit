package GracefullyExit

import (
	"fmt"
	"log"
	"time"
)

var exit bool
var exited = make(chan struct{}, 1)

// StartReceivedExit 启动监听退出信号的 goroutine（使用默认的标准输入）
func StartReceivedExit() {
	// 这里单纯阻塞监听控制台输入
	var key string
	go alert()
	for {
		fmt.Scan(&key)
		if key == "q" {
			fmt.Println("收到退出信号")
			exit = true
			exited <- struct{}{}
			return // 直接退出函数，语义更清晰
		}
	}
}

func ShouldExit() bool {
	return exit
}

func alert() {
	for {
		if _, ok := <-exited; ok {
			log.Println("已经通过 q 安全退出")
			return
		} else {
			fmt.Println("按 q 可以安全退出")
			time.Sleep(30 * time.Second)
		}
	}
}

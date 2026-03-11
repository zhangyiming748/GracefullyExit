package GracefullyExit

import (
	"fmt"
	"time"
)

var exit bool

func StartReceivedExit() {
	// 这里单纯阻塞监听控制台输入
	var key string
	go alert()
	for {
		fmt.Scan(&key)
		if key == "q" {
			fmt.Println("收到退出信号")
			exit = true
			break
		}

	}
}

func ShouldExit() bool {
	return exit
}

func alert() {
	for {
		fmt.Println("按q可以安全退出")
		time.Sleep(30 * time.Second)
	}
}

package GracefullyExit

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

var exit bool

// StartReceivedExit 启动监听退出信号的 goroutine（使用默认的标准输入）
func StartReceivedExit() {
	startReceivedExitWithReader(os.Stdin)
}

// startReceivedExitWithReader 内部函数，接受 io.Reader 用于测试
func startReceivedExitWithReader(reader io.Reader) {
	bufReader := bufio.NewReader(reader)
	go alert()
	for {
		key, err := bufReader.ReadString('\n')
		if err != nil {
			fmt.Println("读取输入失败:", err)
			return
		}
		// 去除换行符和空格
		key = strings.TrimSpace(key)
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
		fmt.Println("按 q 可以安全退出")
		time.Sleep(30 * time.Second)
	}
}

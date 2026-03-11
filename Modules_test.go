// exitchecker_test.go
package GracefullyExit

import (
	"strings"
	"testing"
	"time"
	"log"
)

//go test -v -timeout 10m -run TestUsage
func TestUsage(t *testing.T) {
	// 使用 strings.Reader 模拟输入 'q\n'
	input := strings.NewReader("q\n")
	go startReceivedExitWithReader(input)
	Program()
}

func Program(){
	for i:=0;i<100;i++{
		log.Println("程序运行中")
		time.Sleep(3*time.Second)
		log.Printf("当前的exit为%v\n",exit)
		if ShouldExit() {
			log.Println("程序退出")
			return
		}
	}
}

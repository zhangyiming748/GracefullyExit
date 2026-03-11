// exitchecker_test.go
package GracefullyExit

import (
	"log"
	"testing"
	"time"
)

// go test -v -timeout 10m -run TestUsage
func TestUsage(t *testing.T) {
	go StartReceivedExit()
	Program()
}

func Program() {
	for i := 0; i < 100; i++ {
		log.Println("程序运行中")
		time.Sleep(3 * time.Second)
		log.Printf("当前的exit为%v\n", exit)
		if ShouldExit() {
			log.Println("程序退出")
			return
		}
	}
}

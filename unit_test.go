package GracefullyExit

import (
	"fmt"
	"github.com/zhangyiming748/log"
	"os"
	"testing"
	"time"
)

// go test -v -cover -run TestExitAfterRun unit_test.go GracefullyExit.go
func TestExitAfterRun(t *testing.T) {
	ExitAfterRun(circle, exitInfo)
}
func exitInfo() {
	log.Debug.Println("开始退出...")
	log.Debug.Println("执行清理...")
	log.Debug.Println("结束退出...")
	os.Exit(0)
}
func circle() {
	fmt.Println("启动了程序")
	sum := 0
	for {
		sum++
		fmt.Println("休眠了:", sum, "秒")
		time.Sleep(1 * time.Second)
	}
}

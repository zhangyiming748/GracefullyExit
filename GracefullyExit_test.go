// exitchecker_test.go
package GracefullyExit

import (
	"testing"
	"time"
	"log"
)

func TestExitChecker_ShouldExitAfterAtomicOperation(t *testing.T) {}

func main() {
	ec := New()
	defer ec.Stop() // 程序结束时清理

	// 模拟原子操作循环
	for i := 0; i < 10; i++ {
		// 执行原子操作（这里模拟耗时工作）
		log.Printf("Performing atomic operation %d...\n", i)
		time.Sleep(1 * time.Second) // 模拟工作

		// 操作结束后检查是否退出
		if ec.ShouldExit("q") {
			log.Println("Exit signal received. Quitting after current operation.")
			break
		}
	}
	log.Println("Program exited gracefully.")
}
// exitchecker_test.go
package GracefullyExit

import (
	"testing"
)

func TestExitChecker_ShouldExitAfterAtomicOperation(t *testing.T) {
	go StartReceivedExit()
	for {
		if ShouldExit() {
			t.Log("接收到了退出信号，程序结束")
			break
		}

	}
}

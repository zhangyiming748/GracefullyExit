package GracefullyExit

import (
	"bufio"
	"os"
	"strings"
	"sync"
)

// ExitChecker 结构体，持有通道和goroutine管理
type ExitChecker struct {
	ch   chan string // 通道用于传输输入字符串
	done chan struct{} // 用于停止输入goroutine
	wg   sync.WaitGroup
}

// New 创建一个ExitChecker实例，并启动输入监听goroutine
func New() *ExitChecker {
	ec := &ExitChecker{
		ch:   make(chan string),
		done: make(chan struct{}),
	}
	ec.wg.Add(1)
	go ec.listenInput() // 启动函数A：阻塞读取输入
	return ec
}

// listenInput 函数A：阻塞从stdin读取字符串，并发送到通道
func (ec *ExitChecker) listenInput() {
	defer ec.wg.Done()
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		select {
		case <-ec.done:
			return // 收到停止信号，退出
		default:
			input := strings.TrimSpace(scanner.Text())
			if input != "" {
				ec.ch <- input // 发送到通道
			}
		}
	}
}

// ShouldExit 函数B：非阻塞检查通道中是否有"q"（或指定字符串），返回true表示应退出
func (ec *ExitChecker) ShouldExit(quitStr string) bool {
	select {
	case input := <-ec.ch:
		return input == quitStr // 如果收到且匹配，返回true
	default:
		return false // 通道空，无需退出
	}
}

// Stop 停止输入监听（清理资源）
func (ec *ExitChecker) Stop() {
	close(ec.done)
	ec.wg.Wait()
	close(ec.ch)
}
package GracefullyExit

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

// ExitChecker 结构体，持有通道和 goroutine 管理
type ExitChecker struct {
	ch       chan string   // 通道用于传输输入字符串
	done     chan struct{} // 用于停止所有 goroutine
	wg       sync.WaitGroup
	stopOnce sync.Once // 确保 Stop 只被调用一次
}

// New 创建一个ExitChecker实例，并启动输入监听goroutine
func New() *ExitChecker {
	ec := &ExitChecker{
		ch:   make(chan string, 1), // 使用缓冲通道，避免阻塞
		done: make(chan struct{}),
	}
	ec.wg.Add(2)              // 两个 goroutine
	go ec.listenInput()       // 启动函数A：阻塞读取输入
	go ec.startReminder()     // 启动提示 goroutine
	return ec
}

// listenInput 函数A：阻塞从stdin读取字符串，并发送到通道
func (ec *ExitChecker) listenInput() {
	defer ec.wg.Done()
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := strings.TrimSpace(scanner.Text())
		if input == "" {
			continue
		}

		// 非阻塞发送输入到通道
		select {
		case ec.ch <- input:
			// 发送成功
		default:
			// 通道满，丢弃输入（避免阻塞）
		}

		// 检查是否需要停止
		select {
		case <-ec.done:
			return // 收到停止信号，退出
		default:
			// 继续循环
		}
	}

	// 检查 scanner 错误
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "读取输入错误：%v\n", err)
	}
}

// ShouldExit 函数 B：非阻塞检查通道中是否有"q"（或指定字符串），返回 true 表示应退出
func (ec *ExitChecker) ShouldExit(quitStr string) bool {
	select {
	case input := <-ec.ch:
		return input == quitStr // 如果收到且匹配，返回 true
	default:
		return false // 通道空，无需退出
	}
}

// Stop 停止输入监听（清理资源）
func (ec *ExitChecker) Stop() {
	ec.stopOnce.Do(func() {
		close(ec.done)
		ec.wg.Wait()
		close(ec.ch)
	})
}

/*
这里单独一个函数 准备使用 go 关键字以协程方式启动
循环在控制台打印fmt.Println("输入q安全退出")
输出一次sleep30秒
*/

// startReminder 启动提示 goroutine，周期性提醒用户如何退出
func (ec *ExitChecker) startReminder() {
	defer ec.wg.Done()
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ec.done:
			return // 收到停止信号，退出
		case <-ticker.C:
			fmt.Println("输入 q 安全退出")
		}
	}
}

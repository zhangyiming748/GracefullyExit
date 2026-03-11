# 程序功能

为其他程序提供原子操作形式的退出
即在原子操作过程中即使输入了退出信号 也在当前这轮原子操作之后再交由主程序判断后退出

## 程序大纲

1. 创建一个通道 C，用来传输字符串
2. 创建一个函数 A，阻塞接收字符串
3. 另一个函数 B 循环检测通道中是否有可以接收的字符串
4. 如果接收到字符串，这个函数返回 true 否则返回 false

## 代码逻辑说明

### 核心组件

1. **ExitChecker 结构体**
   - `ch`: 无缓冲通道，用于传输用户输入的字符串
   - `done`: 控制通道，用于通知所有 goroutine 停止
   - `wg`: WaitGroup，等待所有 goroutine 完成
   - `stopOnce`: sync.Once，确保 Stop() 只执行一次

2. **New() 函数**
   - 创建 ExitChecker 实例
   - 初始化通道和同步原语
   - 启动两个 goroutine：
     - `listenInput()`: 监听用户输入
     - `startReminder()`: 每 30 秒打印退出提示

3. **listenInput() 函数**
   - 使用 bufio.Scanner 从标准输入读取
   - 去除输入字符串的前后空格
   - 非阻塞发送输入到通道（通道满时丢弃）
   - 每次读取后检查 done 信号，实现优雅退出
   - 处理 scanner 错误

4. **ShouldExit() 函数**
   - 非阻塞检查通道
   - 如果通道中有输入且匹配退出字符串，返回 true
   - 否则立即返回 false

5. **startReminder() 函数**
   - 使用 Ticker 每 30 秒触发一次
   - 打印"输入 q 安全退出"提示
   - 监听 done 信号，可被中断

6. **Stop() 函数**
   - 使用 sync.Once 确保只关闭一次
   - 关闭 done 通道，通知所有 goroutine
   - 等待所有 goroutine 完成
   - 关闭数据通道

### 线程安全保证

- 所有通道操作由 Go runtime 保证线程安全
- sync.Once 防止重复关闭导致的 panic
- WaitGroup 确保资源正确释放

## 调用方法

```golang
package main

import (
    "fmt"
    "time"

    "github.com/zhangyiming748/GracefullyExit" // 替换为实际路径
)

func main() {
    ge := exitchecker.New()
    defer ge.Stop() // 程序结束时清理

    // 模拟原子操作循环
    for i := 0; i < 10; i++ {
        // 执行原子操作（这里模拟耗时工作）
        fmt.Printf("Performing atomic operation %d...\n", i)
        time.Sleep(1 * time.Second) // 模拟工作

        // 操作结束后检查是否退出
        if ge.ShouldExit("q") {
            fmt.Println("Exit signal received. Quitting after current operation.")
            break
        }
    }
    fmt.Println("Program exited gracefully.")
}
```

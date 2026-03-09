# 程序功能
为其他程序提供原子操作形式的退出
即在原子操作过程中即使输入了退出信号 也在当前这轮原子操作之后再交由主程序判断后退出

## 程序大纲

1. 创建一个通道C，用来传输字符串
2. 创建一个函数A，阻塞接收字符串
3. 另一个函数B循环检测通道中是否有可以接收的字符串
4. 如果接收到字符串，这个函数返回true 否则返回false

### 调用方法

```golang
package main

import (
	"fmt"
	"time"

	"yourmodule/exitchecker" // 替换为实际路径
)

func main() {
	ec := exitchecker.New()
	defer ec.Stop() // 程序结束时清理

	// 模拟原子操作循环
	for i := 0; i < 10; i++ {
		// 执行原子操作（这里模拟耗时工作）
		fmt.Printf("Performing atomic operation %d...\n", i)
		time.Sleep(1 * time.Second) // 模拟工作

		// 操作结束后检查是否退出
		if ec.ShouldExit("q") {
			fmt.Println("Exit signal received. Quitting after current operation.")
			break
		}
	}
	fmt.Println("Program exited gracefully.")
}
```
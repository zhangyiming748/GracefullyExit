# GracefullyExit - 优雅退出工具

## 功能说明

为 Go 程序提供简单优雅的退出机制，支持在原子操作过程中响应退出信号，确保当前操作完成后再退出，避免中途终止导致状态不一致。

## 核心特性

- **原子性保障**：即使收到退出信号，也等待当前原子操作完成后再退出
- **简单易用**：仅需两个函数即可实现退出检测
- **轻量级设计**：无外部依赖，仅使用 Go 标准库
- **控制台交互**：通过简单的键盘输入触发退出

## API 说明

### `StartReceivedExit()`

启动退出监听功能，在后台 goroutine 中阻塞监听控制台输入。

**特点：**
- 非阻塞调用，可独立启动
- 监听用户输入"q"触发退出信号
- 设置全局退出标志位

### `ShouldExit() bool`

检查是否收到退出信号。

**返回值：**
- `true`: 已收到退出信号
- `false`: 未收到退出信号

**使用建议：**
- 在原子操作完成后调用
- 配合循环结构实现优雅退出

## 使用示例

```go
package main

import (
    "fmt"
    "time"
    
    "github.com/zhangyiming748/GracefullyExit"
)

func main() {
    // 启动退出监听（后台运行）
    go GracefullyExit.StartReceivedExit()
    
    // 模拟原子操作循环
    for i := 0; i < 10; i++ {
        // 执行原子操作（模拟耗时工作）
        fmt.Printf("执行原子操作 %d...\n", i)
        time.Sleep(1 * time.Second)
        
        // 操作结束后检查是否退出
        if GracefullyExit.ShouldExit() {
            fmt.Println("收到退出信号，安全退出")
            break
        }
    }
    
    fmt.Println("程序已优雅退出")
}
```

## 运行说明

1. **运行程序**
   ```bash
   go run main.go
   ```

2. **查看提示**
   - 程序每 30 秒显示一次退出提示
   - 提示："按 q 可以安全退出"

3. **触发退出**
   - 在控制台输入 `q` 并按回车
   - 程序将在当前原子操作完成后退出

## 代码结构

```
GracefullyExit/
├── GracefullyExit.go      # 核心实现
│   ├── StartReceivedExit()  # 启动退出监听
│   ├── ShouldExit()         # 检查退出状态
│   └── alert()              # 定时提示函数
├── GracefullyExit_test.go # 测试文件
└── README.md              # 文档说明
```

## 注意事项

- 输入"q"后需要按回车键确认
- 退出信号是全局的，一旦触发将影响整个程序
- 建议在程序入口启动监听功能

package GracefullyExit

import (
	"github.com/zhangyiming748/log"
	"os"
	"os/signal"
	"syscall"
)

func ExitAfterRun(last func()) {
	c := make(chan os.Signal)
	// 监听信号
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGUSR1, syscall.SIGUSR2)
	go func() {
		for s := range c {
			switch s { // 终端控制进程结束(终端连接断开)|用户发送INTR字符(Ctrl+C)触发|结束程序
			case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM:
				log.Debug.Println("退出:", s)
				last()
			case syscall.SIGUSR1:
				log.Debug.Println("usr1", s)
			case syscall.SIGUSR2:
				log.Debug.Println("usr2", s)
			default:
				log.Debug.Println("其他信号:", s)
			}
		}
	}()
}

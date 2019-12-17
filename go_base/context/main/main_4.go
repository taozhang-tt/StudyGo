//使用 context 控制嵌套的 goroutine 停止
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go parentWatch(ctx,"【父监控】")

	time.Sleep(10 * time.Second)
	fmt.Println("可以了，通知监控停止")
	cancel()
	//为了检测监控过是否停止，如果没有监控输出，就表示停止了
	time.Sleep(5 * time.Second)
}

func parentWatch(ctx context.Context, name string) {
	ctxChild, _ := context.WithCancel(ctx)	//使用父亲 context 创建了 child context，当父 context 的 cancel 函数被调用时，会去调用孩子context 的 cancel 函数
	go childWatch(ctxChild, "【子监控】")
	for {
		select {
		case <-ctx.Done():	//当调用了 cancel 函数时这里检测到讯号，停止协程，优雅退出
			fmt.Println(name,"监控退出，停止了...")
			return
		default:
			fmt.Println(name,"goroutine监控中...")
			time.Sleep(2 * time.Second)
		}
	}
}

func childWatch(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name,"监控退出，停止了...")
			return
		default:
			fmt.Println(name,"goroutine监控中...")
			time.Sleep(2 * time.Second)
		}
	}
}
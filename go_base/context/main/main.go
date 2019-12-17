package main

import (
	"fmt"
	"time"
)

func main() {
	stop := make(chan bool)

	go func() {
		for {
			select {
			case value1, ok1:= <-stop:
				fmt.Println(value1, ok1)
				fmt.Println("监控退出，停止了...1")
				return
			default:
				fmt.Println("goroutine监控中...1")
				time.Sleep(2 * time.Second)
			}
		}
	}()

	go func() {
		for {
			select {
			case value2, ok2:= <-stop:
				fmt.Println(value2, ok2)
				fmt.Println("监控退出，停止了...2")
				return
			default:
				fmt.Println("goroutine监控中...2")
				time.Sleep(2 * time.Second)
			}
		}
	}()

	time.Sleep(10 * time.Second)
	fmt.Println("可以了，通知监控停止")
	close(stop)
	//stop<- true
	//为了检测监控过是否停止，如果没有监控输出，就表示停止了
	time.Sleep(5 * time.Second)
}


package main

import (
	"context"
	"fmt"
	"os/exec"
	"time"
)

type result struct {
	err error
	output []byte
}

func main()  {
	//执行一个cmd，在一个协程里执行
	var (
		ctx context.Context
		cancelFunc context.CancelFunc
		cmd *exec.Cmd
		resultChan chan *result
		res *result
	)
	//创建一个结果队列
	resultChan = make(chan *result, 1000)

	ctx, cancelFunc = context.WithCancel(context.TODO())
	go func() {
		var (
			err error
			output []byte
		)
		cmd = exec.CommandContext(ctx, "/bin/zsh", "-c", "sleep 3;echo hello world")
		output, err = cmd.CombinedOutput()
		resultChan <- &result{
			err: err,
			output:output,
		}
	}()

	time.Sleep(1*time.Second)

	//取消上下文
	cancelFunc()

	//在main协程里，等待子协程的退出，并打印任务执行结果
	res = <-resultChan

	fmt.Println(res.err, string(res.output))
}

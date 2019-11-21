package main

import (
	"fmt"
	"os/exec"
)

func main()  {
	var (
		cmd *exec.Cmd
		output []byte
		err error
	)

	//生成Cmd
	cmd = exec.Command("/bin/zsh", "-c", "ls -al; echo tt")
	//执行命令，捕获子进程的输出(pipe)
	if output,err = cmd.CombinedOutput();err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(output))
}

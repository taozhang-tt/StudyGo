package main

import (
	"bufio"
	"errors"
	"fmt"
	"os/exec"
)

func ExecCommand(command string) error {
	cmd := exec.Command("/bin/bash", "-c", command)
	stderr, _ := cmd.StderrPipe()
	err := cmd.Start() //运行脚本
	if nil == err { //持续获取输出
		scanner := bufio.NewScanner(stderr)
		scanner.Split(bufio.ScanLines)
		for scanner.Scan() {
			m := scanner.Text()
			fmt.Println(m)
			return errors.New(m)
		}
	}
	err = cmd.Wait() //等待执行完成
	if nil != err {
		return err
	}
	return nil
}

func main() {
	instancePath := "/data/app/Pressure_Test/current"
	description := "example"
	command := "source /etc/profile && cd " + instancePath + "/instance/" + description + " && go build main.go"
	err := ExecCommand(command)
	fmt.Println(err)
}

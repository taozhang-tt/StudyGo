package main

import (
	"fmt"
	"os/exec"
)

func main()  {
	var (
		output []byte
		err error
	)
	command := "go run main1.go"
	cmd := exec.Command("/bin/bash", "-c", command)
	if output, err = cmd.CombinedOutput(); err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(output))
}

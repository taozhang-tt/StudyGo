package main

import (
	"errors"
	"fmt"
)

func main() {
	//Test()

	m := map[string]interface{}{
		"a": "a",
		"b": "b",
	}
	v, ok := m["b"].(string)
	if ok {
		fmt.Println(v)
	}
	fmt.Println(errors.New("hello world"))
}

func Test() {
	panic("error")
}

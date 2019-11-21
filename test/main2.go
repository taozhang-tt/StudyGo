package main

import (
	"StudyGo/test/common"
	"io/ioutil"
)

//生成go代码
func main()  {
	header := common.GetHeader()
	footer := common.GetFooter()
	var body string

	arr := []string{"GamerAddRoomC2S", "GamerLoginC2S", "GamerLoginS2C", "GamerLoginGetDataS2C"}

	for _, structName := range arr {
		body =  body + common.GenerateBody(structName)
	}

	ioutil.WriteFile("main3.go", []byte(header+body+footer), 777)
}

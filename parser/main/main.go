package main

import (
	"StudyGo/parser/common"
	"fmt"
	"go/parser"
	"go/token"
	"io/ioutil"
	"path"
	"strings"
)

func main()  {
	//fileName := "../obj/user.go"
	//fset := token.NewFileSet()
	//f, err := parser.ParseFile(fset, fileName, nil, 0)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//arr := f.Scope.Objects
	//for name, object := range arr {
	//	fmt.Println(name)
	//	fmt.Println(object.Kind)
	//}

	fileDir := "/Users/tt/Code/GoCode/StudyGo/parser/pb"	//定义要读取的pb对应的go文件的路径
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, fileDir, nil, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	structNameSlice := make([]string, 0)
	for pkgName, pkg := range pkgs {
		pbMap := make(map[string][]string, 0)
		fmt.Println("pkg name----", pkgName) //包名
		for fileDir, file := range pkg.Files{
			fmt.Println("name----", fileDir) //文件路径
			fileName := strings.TrimSuffix(path.Base(fileDir), ".pb.go")//文件名
			pbMap[fileName] = make([]string, 0)
			for _, object := range file.Scope.Objects {
				if object.Kind.String() == "type" {
					pbMap[fileName] = append(pbMap[fileName], object.Name)
					structNameSlice = append(structNameSlice, object.Name)
					//fmt.Println(object.Name) //结构体名
				}
			}
		}
	}
	fmt.Println(structNameSlice)

	header := common.GetHeader()
	footer := common.GetFooter()
	var body string

	//arr := []string{"GamerAddRoomC2S", "GamerLoginC2S", "GamerLoginS2C", "GamerLoginGetDataS2C"}

	for _, structName := range structNameSlice {
		body =  body + common.GenerateBody(structName)
	}

	ioutil.WriteFile("main3.go", []byte(header+body+footer), 777)
}
package main

import (
	"go/ast"
	"go/parser"
	"go/token"
)

func main()  {
	src := `
package main
type GamerStatus int32
`	//定义要读取的pb对应的go文件的路径
	fset := token.NewFileSet() // positions are relative to fset
	f, err := parser.ParseFile(fset, "", src, 0)
	if err != nil {
		panic(err)
	}
	ast.Print(fset, f)
}
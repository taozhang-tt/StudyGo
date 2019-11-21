package common

import (
	"github.com/bitly/go-simplejson"
	"reflect"
	"strings"
)

func Struct2Json(objType reflect.Type) (json *simplejson.Json, err error){
	objName := objType.Name()
	json = simplejson.New()
	json.Set("title", objName)
	json.Set("name", objName)
	jsonSlice := make([]*simplejson.Json, 0)

	for i:=0; i<objType.NumField() ;i++  {
		jsonInner := simplejson.New()
		field := objType.Field(i)
		if !strings.Contains(field.Name, "XXX") && field.Type.Kind().String() == "struct" {
			jsonInner,_ = InnerStruct2Json(field.Type)
			jsonSlice = append(jsonSlice, jsonInner)
		} else {
			if !strings.Contains(field.Name, "XXX") {
				jsonInner.Set("field", field.Name)
				jsonInner.Set("type", field.Type.Kind().String())
				jsonSlice = append(jsonSlice, jsonInner)
			}
		}
	}
	if len(jsonSlice) > 0 {
		json.Set("properties", jsonSlice)
	}
	return
}

func InnerStruct2Json(objType reflect.Type) (json *simplejson.Json, err error){
	objName := objType.Name()
	json = simplejson.New()
	json.Set("field", objName)
	json.Set("type", "struct")
	jsonSlice := make([]*simplejson.Json, 0)

	for i:=0; i<objType.NumField() ;i++  {
		jsonInner := simplejson.New()
		field := objType.Field(i)
		if !strings.Contains(field.Name, "XXX") && field.Type.Kind().String() == "struct" {
			jsonInner,_ = InnerStruct2Json(field.Type)
			jsonSlice = append(jsonSlice, jsonInner)
		} else {
			if !strings.Contains(field.Name, "XXX") {
				jsonInner.Set("field", field.Name)
				jsonInner.Set("type", field.Type.Kind().String())
				jsonSlice = append(jsonSlice, jsonInner)
			}
		}
	}
	if len(jsonSlice) > 0 {
		json.Set("props", jsonSlice)
	}
	return
}

func GetHeader() string {
	return `
package main

import (
	"StudyGo/parser/pb"
	"StudyGo/test/common"
	"fmt"
	"github.com/bitly/go-simplejson"
	"io/ioutil"
	"reflect"

)

func main()  {
	//头部
	jsonConfig := simplejson.New()	//最外层
	jsonPb := simplejson.New() //存储pb
	jsonPb.Set("title", "pb")
	jsonPb.Set("name", "pb")
	jsonPbFiledSlice := make([]*simplejson.Json, 0)	//存储pb的properties
	jsonStructSlice := make([]*simplejson.Json, 0) //用于存储结构体
`
}

func GetFooter() string {
	return `
	jsonPb.Set("properties", jsonPbFiledSlice)
	jsonStructSlice = append(jsonStructSlice, jsonPb)
	jsonConfig.Set("config", jsonStructSlice)
	str, _ := jsonConfig.MarshalJSON()
	ioutil.WriteFile("main.json", str, 0777)
}
`
}

func GenerateBody(structName string) string {
	body := `

	structNameObj := new(pb.structName)
	structNameType := reflect.TypeOf(*structNameObj)
	if structNameType.Kind().String() == "struct" {
		structNameJson, _ := common.Struct2Json(structNameType)
		jsonStructSlice = append(jsonStructSlice, structNameJson)
		structNameJsonPbFiled := simplejson.New()
		structNameJsonPbFiled.Set("field", "structName")
		structNameJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, structNameJsonPbFiled)
	}

`
	return strings.Replace(body, "structName", structName, -1)
}

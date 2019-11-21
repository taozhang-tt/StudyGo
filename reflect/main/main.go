package main

import (
	"fmt"
	"github.com/bitly/go-simplejson"
	"reflect"
)

type User struct {
	Name    string         `json:"name"`
	Age     int            `json:"age"`
	Friends []string       `json:"friends"`
	Grade   map[string]int `json:"grade"`
	Brother Family         `json:"brother"`
}

type Family struct {
	Relation string `json:"relation"`
	Age      int    `json:"age"`
}

func main() {
	var user = new(User)
	userType := reflect.TypeOf(*user)

	//json, _ := Struct2Json(userType)
	//jsonBytes,_ := json.MarshalJSON()
	//fmt.Println(string(jsonBytes))

	fmt.Println(userType.Kind())

	for i := 0; i < userType.NumField(); i++ {
		field := userType.Field(i)
		fmt.Println("Name:", field.Name, "type.string:", field.Type.String(), "type.Name:", field.Type.Name(), "type.kind:", field.Type.Kind().String())
		if field.Type.Kind().String() == "struct" {
			fmt.Println("--------inner-----------")
			innerType := field.Type
			for j:=0; j<innerType.NumField(); j++ {
				innerField := innerType.Field(j)
				fmt.Println("Name:", innerField.Name, "type.string:", innerField.Type.String(), "type.Name:", innerField.Type.Name(), "type.kind:", innerField.Type.Kind().String())
			}
		}
	}
}

func Struct2Json(objType reflect.Type) (json *simplejson.Json, err error){
	objName := objType.Name()
	json = simplejson.New()
	json.Set("title", objName)
	json.Set("name", objName)
	jsonSlice := make([]*simplejson.Json, 0)

	for i:=0; i<objType.NumField() ;i++  {
		jsonInner := simplejson.New()
		field := objType.Field(i)
		if field.Type.Kind().String() == "struct" {
			jsonInner,_ = InnerStruct2Json(field.Type)
		} else {
			jsonInner.Set("field", field.Name)
			jsonInner.Set("type", field.Type.Kind().String())
		}
		jsonSlice = append(jsonSlice, jsonInner)
	}
	json.Set("properties", jsonSlice)
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
		if field.Type.Kind().String() == "struct" {
			jsonInner,_ = InnerStruct2Json(field.Type)
		} else {
			jsonInner.Set("field", field.Name)
			jsonInner.Set("type", field.Type.Kind().String())
		}
		jsonSlice = append(jsonSlice, jsonInner)
	}
	json.Set("props", jsonSlice)
	return
}

//func main()  {
//	simpleJsonInner := simplejson.New()
//	simpleJsonInner.Set("first", "TT")
//	simpleJsonInner.Set("last", "TT")
//
//	simpleJson := simplejson.New()
//	simpleJson.Set("name", simpleJsonInner)
//
//	str, _ := simpleJson.MarshalJSON()
//	fmt.Println(string(str))
//}
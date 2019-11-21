package main

import (
	"StudyGo/parser/pb"
	"StudyGo/test/common"
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

	//body
	GamerAddRoomC2SObj := new(pb.GamerAddRoomC2S)
	GamerAddRoomC2SType := reflect.TypeOf(*GamerAddRoomC2SObj)
	if GamerAddRoomC2SType.Kind().String() == "struct" {
		GamerAddRoomC2SJson, _ := common.Struct2Json(GamerAddRoomC2SType)
		jsonStructSlice = append(jsonStructSlice, GamerAddRoomC2SJson)
		GamerAddRoomC2SJsonPbFiled := simplejson.New()
		GamerAddRoomC2SJsonPbFiled.Set("field", "GamerAddRoomC2S")
		GamerAddRoomC2SJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerAddRoomC2SJsonPbFiled)
	}


	//尾部
	jsonPb.Set("properties", jsonPbFiledSlice)
	jsonStructSlice = append(jsonStructSlice, jsonPb)
	jsonConfig.Set("config", jsonStructSlice)
	str, _ := jsonConfig.MarshalJSON()
	ioutil.WriteFile("main.json", str, 0777)
}

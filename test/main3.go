
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


	GamerAddRoomC2SObj := new(pb.GamerAddRoomC2S)
	GamerAddRoomC2SType := reflect.TypeOf(*GamerAddRoomC2SObj)
	GamerAddRoomC2SJson, _ := common.Struct2Json(GamerAddRoomC2SType)
	jsonStructSlice = append(jsonStructSlice, GamerAddRoomC2SJson)
	GamerAddRoomC2SJsonPbFiled := simplejson.New()
	GamerAddRoomC2SJsonPbFiled.Set("field", "GamerAddRoomC2S")
	GamerAddRoomC2SJsonPbFiled.Set("type", "struct")
	jsonPbFiledSlice = append(jsonPbFiledSlice, GamerAddRoomC2SJsonPbFiled)



	GamerLoginC2SObj := new(pb.GamerLoginC2S)
	GamerLoginC2SType := reflect.TypeOf(*GamerLoginC2SObj)
	GamerLoginC2SJson, _ := common.Struct2Json(GamerLoginC2SType)
	jsonStructSlice = append(jsonStructSlice, GamerLoginC2SJson)
	GamerLoginC2SJsonPbFiled := simplejson.New()
	GamerLoginC2SJsonPbFiled.Set("field", "GamerLoginC2S")
	GamerLoginC2SJsonPbFiled.Set("type", "struct")
	jsonPbFiledSlice = append(jsonPbFiledSlice, GamerLoginC2SJsonPbFiled)



	GamerLoginS2CObj := new(pb.GamerLoginS2C)
	GamerLoginS2CType := reflect.TypeOf(*GamerLoginS2CObj)
	GamerLoginS2CJson, _ := common.Struct2Json(GamerLoginS2CType)
	jsonStructSlice = append(jsonStructSlice, GamerLoginS2CJson)
	GamerLoginS2CJsonPbFiled := simplejson.New()
	GamerLoginS2CJsonPbFiled.Set("field", "GamerLoginS2C")
	GamerLoginS2CJsonPbFiled.Set("type", "struct")
	jsonPbFiledSlice = append(jsonPbFiledSlice, GamerLoginS2CJsonPbFiled)



	GamerLoginGetDataS2CObj := new(pb.GamerLoginGetDataS2C)
	GamerLoginGetDataS2CType := reflect.TypeOf(*GamerLoginGetDataS2CObj)
	GamerLoginGetDataS2CJson, _ := common.Struct2Json(GamerLoginGetDataS2CType)
	jsonStructSlice = append(jsonStructSlice, GamerLoginGetDataS2CJson)
	GamerLoginGetDataS2CJsonPbFiled := simplejson.New()
	GamerLoginGetDataS2CJsonPbFiled.Set("field", "GamerLoginGetDataS2C")
	GamerLoginGetDataS2CJsonPbFiled.Set("type", "struct")
	jsonPbFiledSlice = append(jsonPbFiledSlice, GamerLoginGetDataS2CJsonPbFiled)


	jsonPb.Set("properties", jsonPbFiledSlice)
	jsonStructSlice = append(jsonStructSlice, jsonPb)
	jsonConfig.Set("config", jsonStructSlice)
	str, _ := jsonConfig.MarshalJSON()
	ioutil.WriteFile("main.json", str, 0777)
}

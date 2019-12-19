
package main

import (
	"StudyGo/parser/common"
	"StudyGo/parser/pb"
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


	ServerCmdObj := new(pb.GamerBuyItemC2S)
	ServerCmdType := reflect.TypeOf(*ServerCmdObj)
	ServerCmdJson, _ := common.Struct2Json(ServerCmdType)
	jsonStructSlice = append(jsonStructSlice, ServerCmdJson)
	ServerCmdJsonPbFiled := simplejson.New()
	ServerCmdJsonPbFiled.Set("field", "ServerCmd")
	ServerCmdJsonPbFiled.Set("type", "struct")
	jsonPbFiledSlice = append(jsonPbFiledSlice, ServerCmdJsonPbFiled)



	ServerInfoObj := new(pb.ServerInfo)
	ServerInfoType := reflect.TypeOf(*ServerInfoObj)
	ServerInfoJson, _ := common.Struct2Json(ServerInfoType)
	jsonStructSlice = append(jsonStructSlice, ServerInfoJson)
	ServerInfoJsonPbFiled := simplejson.New()
	ServerInfoJsonPbFiled.Set("field", "ServerInfo")
	ServerInfoJsonPbFiled.Set("type", "struct")
	jsonPbFiledSlice = append(jsonPbFiledSlice, ServerInfoJsonPbFiled)



	ServerWarnObj := new(pb.ServerWarn)
	ServerWarnType := reflect.TypeOf(*ServerWarnObj)
	ServerWarnJson, _ := common.Struct2Json(ServerWarnType)
	jsonStructSlice = append(jsonStructSlice, ServerWarnJson)
	ServerWarnJsonPbFiled := simplejson.New()
	ServerWarnJsonPbFiled.Set("field", "ServerWarn")
	ServerWarnJsonPbFiled.Set("type", "struct")
	jsonPbFiledSlice = append(jsonPbFiledSlice, ServerWarnJsonPbFiled)



	C2SObj := new(pb.C2S)
	C2SType := reflect.TypeOf(*C2SObj)
	C2SJson, _ := common.Struct2Json(C2SType)
	jsonStructSlice = append(jsonStructSlice, C2SJson)
	C2SJsonPbFiled := simplejson.New()
	C2SJsonPbFiled.Set("field", "C2S")
	C2SJsonPbFiled.Set("type", "struct")
	jsonPbFiledSlice = append(jsonPbFiledSlice, C2SJsonPbFiled)



	LeagueDataObj := new(pb.LeagueData)
	LeagueDataType := reflect.TypeOf(*LeagueDataObj)
	LeagueDataJson, _ := common.Struct2Json(LeagueDataType)
	jsonStructSlice = append(jsonStructSlice, LeagueDataJson)
	LeagueDataJsonPbFiled := simplejson.New()
	LeagueDataJsonPbFiled.Set("field", "LeagueData")
	LeagueDataJsonPbFiled.Set("type", "struct")
	jsonPbFiledSlice = append(jsonPbFiledSlice, LeagueDataJsonPbFiled)



	OnlineLeagueObj := new(pb.OnlineLeague)
	OnlineLeagueType := reflect.TypeOf(*OnlineLeagueObj)
	OnlineLeagueJson, _ := common.Struct2Json(OnlineLeagueType)
	jsonStructSlice = append(jsonStructSlice, OnlineLeagueJson)
	OnlineLeagueJsonPbFiled := simplejson.New()
	OnlineLeagueJsonPbFiled.Set("field", "OnlineLeague")
	OnlineLeagueJsonPbFiled.Set("type", "struct")
	jsonPbFiledSlice = append(jsonPbFiledSlice, OnlineLeagueJsonPbFiled)



	S2CObj := new(pb.S2C)
	S2CType := reflect.TypeOf(*S2CObj)
	S2CJson, _ := common.Struct2Json(S2CType)
	jsonStructSlice = append(jsonStructSlice, S2CJson)
	S2CJsonPbFiled := simplejson.New()
	S2CJsonPbFiled.Set("field", "S2C")
	S2CJsonPbFiled.Set("type", "struct")
	jsonPbFiledSlice = append(jsonPbFiledSlice, S2CJsonPbFiled)


	jsonPb.Set("properties", jsonPbFiledSlice)
	jsonStructSlice = append(jsonStructSlice, jsonPb)
	jsonConfig.Set("config", jsonStructSlice)
	str, _ := jsonConfig.MarshalJSON()
	ioutil.WriteFile("main.json", str, 0777)
}

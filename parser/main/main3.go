
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


	RedisNotifyNewPVPResultObj := new(pb.RedisNotifyNewPVPResult)
	RedisNotifyNewPVPResultType := reflect.TypeOf(*RedisNotifyNewPVPResultObj)
	if RedisNotifyNewPVPResultType.Kind().String() == "struct" {
		RedisNotifyNewPVPResultJson, _ := common.Struct2Json(RedisNotifyNewPVPResultType)
		jsonStructSlice = append(jsonStructSlice, RedisNotifyNewPVPResultJson)
		RedisNotifyNewPVPResultJsonPbFiled := simplejson.New()
		RedisNotifyNewPVPResultJsonPbFiled.Set("field", "RedisNotifyNewPVPResult")
		RedisNotifyNewPVPResultJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, RedisNotifyNewPVPResultJsonPbFiled)
	}



	RedisNotifyNewGamerChatObj := new(pb.RedisNotifyNewGamerChat)
	RedisNotifyNewGamerChatType := reflect.TypeOf(*RedisNotifyNewGamerChatObj)
	if RedisNotifyNewGamerChatType.Kind().String() == "struct" {
		RedisNotifyNewGamerChatJson, _ := common.Struct2Json(RedisNotifyNewGamerChatType)
		jsonStructSlice = append(jsonStructSlice, RedisNotifyNewGamerChatJson)
		RedisNotifyNewGamerChatJsonPbFiled := simplejson.New()
		RedisNotifyNewGamerChatJsonPbFiled.Set("field", "RedisNotifyNewGamerChat")
		RedisNotifyNewGamerChatJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, RedisNotifyNewGamerChatJsonPbFiled)
	}



	RedisNotifyNewWorldMailObj := new(pb.RedisNotifyNewWorldMail)
	RedisNotifyNewWorldMailType := reflect.TypeOf(*RedisNotifyNewWorldMailObj)
	if RedisNotifyNewWorldMailType.Kind().String() == "struct" {
		RedisNotifyNewWorldMailJson, _ := common.Struct2Json(RedisNotifyNewWorldMailType)
		jsonStructSlice = append(jsonStructSlice, RedisNotifyNewWorldMailJson)
		RedisNotifyNewWorldMailJsonPbFiled := simplejson.New()
		RedisNotifyNewWorldMailJsonPbFiled.Set("field", "RedisNotifyNewWorldMail")
		RedisNotifyNewWorldMailJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, RedisNotifyNewWorldMailJsonPbFiled)
	}



	RedisNotifyNewWorldChatObj := new(pb.RedisNotifyNewWorldChat)
	RedisNotifyNewWorldChatType := reflect.TypeOf(*RedisNotifyNewWorldChatObj)
	if RedisNotifyNewWorldChatType.Kind().String() == "struct" {
		RedisNotifyNewWorldChatJson, _ := common.Struct2Json(RedisNotifyNewWorldChatType)
		jsonStructSlice = append(jsonStructSlice, RedisNotifyNewWorldChatJson)
		RedisNotifyNewWorldChatJsonPbFiled := simplejson.New()
		RedisNotifyNewWorldChatJsonPbFiled.Set("field", "RedisNotifyNewWorldChat")
		RedisNotifyNewWorldChatJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, RedisNotifyNewWorldChatJsonPbFiled)
	}



	RedisNotifyObj := new(pb.RedisNotify)
	RedisNotifyType := reflect.TypeOf(*RedisNotifyObj)
	if RedisNotifyType.Kind().String() == "struct" {
		RedisNotifyJson, _ := common.Struct2Json(RedisNotifyType)
		jsonStructSlice = append(jsonStructSlice, RedisNotifyJson)
		RedisNotifyJsonPbFiled := simplejson.New()
		RedisNotifyJsonPbFiled.Set("field", "RedisNotify")
		RedisNotifyJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, RedisNotifyJsonPbFiled)
	}



	RedisNotifyNewGamerObj := new(pb.RedisNotifyNewGamer)
	RedisNotifyNewGamerType := reflect.TypeOf(*RedisNotifyNewGamerObj)
	if RedisNotifyNewGamerType.Kind().String() == "struct" {
		RedisNotifyNewGamerJson, _ := common.Struct2Json(RedisNotifyNewGamerType)
		jsonStructSlice = append(jsonStructSlice, RedisNotifyNewGamerJson)
		RedisNotifyNewGamerJsonPbFiled := simplejson.New()
		RedisNotifyNewGamerJsonPbFiled.Set("field", "RedisNotifyNewGamer")
		RedisNotifyNewGamerJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, RedisNotifyNewGamerJsonPbFiled)
	}



	RedisNotifyDelGamerFriendObj := new(pb.RedisNotifyDelGamerFriend)
	RedisNotifyDelGamerFriendType := reflect.TypeOf(*RedisNotifyDelGamerFriendObj)
	if RedisNotifyDelGamerFriendType.Kind().String() == "struct" {
		RedisNotifyDelGamerFriendJson, _ := common.Struct2Json(RedisNotifyDelGamerFriendType)
		jsonStructSlice = append(jsonStructSlice, RedisNotifyDelGamerFriendJson)
		RedisNotifyDelGamerFriendJsonPbFiled := simplejson.New()
		RedisNotifyDelGamerFriendJsonPbFiled.Set("field", "RedisNotifyDelGamerFriend")
		RedisNotifyDelGamerFriendJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, RedisNotifyDelGamerFriendJsonPbFiled)
	}



	RedisNotifyNewLogicAreaObj := new(pb.RedisNotifyNewLogicArea)
	RedisNotifyNewLogicAreaType := reflect.TypeOf(*RedisNotifyNewLogicAreaObj)
	if RedisNotifyNewLogicAreaType.Kind().String() == "struct" {
		RedisNotifyNewLogicAreaJson, _ := common.Struct2Json(RedisNotifyNewLogicAreaType)
		jsonStructSlice = append(jsonStructSlice, RedisNotifyNewLogicAreaJson)
		RedisNotifyNewLogicAreaJsonPbFiled := simplejson.New()
		RedisNotifyNewLogicAreaJsonPbFiled.Set("field", "RedisNotifyNewLogicArea")
		RedisNotifyNewLogicAreaJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, RedisNotifyNewLogicAreaJsonPbFiled)
	}



	RedisNotifyNewGamerFriendReqObj := new(pb.RedisNotifyNewGamerFriendReq)
	RedisNotifyNewGamerFriendReqType := reflect.TypeOf(*RedisNotifyNewGamerFriendReqObj)
	if RedisNotifyNewGamerFriendReqType.Kind().String() == "struct" {
		RedisNotifyNewGamerFriendReqJson, _ := common.Struct2Json(RedisNotifyNewGamerFriendReqType)
		jsonStructSlice = append(jsonStructSlice, RedisNotifyNewGamerFriendReqJson)
		RedisNotifyNewGamerFriendReqJsonPbFiled := simplejson.New()
		RedisNotifyNewGamerFriendReqJsonPbFiled.Set("field", "RedisNotifyNewGamerFriendReq")
		RedisNotifyNewGamerFriendReqJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, RedisNotifyNewGamerFriendReqJsonPbFiled)
	}



	RedisNotifyAccessTokenObj := new(pb.RedisNotifyAccessToken)
	RedisNotifyAccessTokenType := reflect.TypeOf(*RedisNotifyAccessTokenObj)
	if RedisNotifyAccessTokenType.Kind().String() == "struct" {
		RedisNotifyAccessTokenJson, _ := common.Struct2Json(RedisNotifyAccessTokenType)
		jsonStructSlice = append(jsonStructSlice, RedisNotifyAccessTokenJson)
		RedisNotifyAccessTokenJsonPbFiled := simplejson.New()
		RedisNotifyAccessTokenJsonPbFiled.Set("field", "RedisNotifyAccessToken")
		RedisNotifyAccessTokenJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, RedisNotifyAccessTokenJsonPbFiled)
	}



	RedisNotifyGamerLogoutObj := new(pb.RedisNotifyGamerLogout)
	RedisNotifyGamerLogoutType := reflect.TypeOf(*RedisNotifyGamerLogoutObj)
	if RedisNotifyGamerLogoutType.Kind().String() == "struct" {
		RedisNotifyGamerLogoutJson, _ := common.Struct2Json(RedisNotifyGamerLogoutType)
		jsonStructSlice = append(jsonStructSlice, RedisNotifyGamerLogoutJson)
		RedisNotifyGamerLogoutJsonPbFiled := simplejson.New()
		RedisNotifyGamerLogoutJsonPbFiled.Set("field", "RedisNotifyGamerLogout")
		RedisNotifyGamerLogoutJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, RedisNotifyGamerLogoutJsonPbFiled)
	}



	RedisNotifyGamerUpdateObj := new(pb.RedisNotifyGamerUpdate)
	RedisNotifyGamerUpdateType := reflect.TypeOf(*RedisNotifyGamerUpdateObj)
	if RedisNotifyGamerUpdateType.Kind().String() == "struct" {
		RedisNotifyGamerUpdateJson, _ := common.Struct2Json(RedisNotifyGamerUpdateType)
		jsonStructSlice = append(jsonStructSlice, RedisNotifyGamerUpdateJson)
		RedisNotifyGamerUpdateJsonPbFiled := simplejson.New()
		RedisNotifyGamerUpdateJsonPbFiled.Set("field", "RedisNotifyGamerUpdate")
		RedisNotifyGamerUpdateJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, RedisNotifyGamerUpdateJsonPbFiled)
	}



	RedisNotifyNewGamerFriendObj := new(pb.RedisNotifyNewGamerFriend)
	RedisNotifyNewGamerFriendType := reflect.TypeOf(*RedisNotifyNewGamerFriendObj)
	if RedisNotifyNewGamerFriendType.Kind().String() == "struct" {
		RedisNotifyNewGamerFriendJson, _ := common.Struct2Json(RedisNotifyNewGamerFriendType)
		jsonStructSlice = append(jsonStructSlice, RedisNotifyNewGamerFriendJson)
		RedisNotifyNewGamerFriendJsonPbFiled := simplejson.New()
		RedisNotifyNewGamerFriendJsonPbFiled.Set("field", "RedisNotifyNewGamerFriend")
		RedisNotifyNewGamerFriendJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, RedisNotifyNewGamerFriendJsonPbFiled)
	}



	RedisNotifyGamerLoginObj := new(pb.RedisNotifyGamerLogin)
	RedisNotifyGamerLoginType := reflect.TypeOf(*RedisNotifyGamerLoginObj)
	if RedisNotifyGamerLoginType.Kind().String() == "struct" {
		RedisNotifyGamerLoginJson, _ := common.Struct2Json(RedisNotifyGamerLoginType)
		jsonStructSlice = append(jsonStructSlice, RedisNotifyGamerLoginJson)
		RedisNotifyGamerLoginJsonPbFiled := simplejson.New()
		RedisNotifyGamerLoginJsonPbFiled.Set("field", "RedisNotifyGamerLogin")
		RedisNotifyGamerLoginJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, RedisNotifyGamerLoginJsonPbFiled)
	}



	RedisNotifyNewGamerMailObj := new(pb.RedisNotifyNewGamerMail)
	RedisNotifyNewGamerMailType := reflect.TypeOf(*RedisNotifyNewGamerMailObj)
	if RedisNotifyNewGamerMailType.Kind().String() == "struct" {
		RedisNotifyNewGamerMailJson, _ := common.Struct2Json(RedisNotifyNewGamerMailType)
		jsonStructSlice = append(jsonStructSlice, RedisNotifyNewGamerMailJson)
		RedisNotifyNewGamerMailJsonPbFiled := simplejson.New()
		RedisNotifyNewGamerMailJsonPbFiled.Set("field", "RedisNotifyNewGamerMail")
		RedisNotifyNewGamerMailJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, RedisNotifyNewGamerMailJsonPbFiled)
	}



	OrderNotifyObj := new(pb.OrderNotify)
	OrderNotifyType := reflect.TypeOf(*OrderNotifyObj)
	if OrderNotifyType.Kind().String() == "struct" {
		OrderNotifyJson, _ := common.Struct2Json(OrderNotifyType)
		jsonStructSlice = append(jsonStructSlice, OrderNotifyJson)
		OrderNotifyJsonPbFiled := simplejson.New()
		OrderNotifyJsonPbFiled.Set("field", "OrderNotify")
		OrderNotifyJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, OrderNotifyJsonPbFiled)
	}



	SuperAccountObj := new(pb.SuperAccount)
	SuperAccountType := reflect.TypeOf(*SuperAccountObj)
	if SuperAccountType.Kind().String() == "struct" {
		SuperAccountJson, _ := common.Struct2Json(SuperAccountType)
		jsonStructSlice = append(jsonStructSlice, SuperAccountJson)
		SuperAccountJsonPbFiled := simplejson.New()
		SuperAccountJsonPbFiled.Set("field", "SuperAccount")
		SuperAccountJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, SuperAccountJsonPbFiled)
	}



	SDKChannelObj := new(pb.SDKChannel)
	SDKChannelType := reflect.TypeOf(*SDKChannelObj)
	if SDKChannelType.Kind().String() == "struct" {
		SDKChannelJson, _ := common.Struct2Json(SDKChannelType)
		jsonStructSlice = append(jsonStructSlice, SDKChannelJson)
		SDKChannelJsonPbFiled := simplejson.New()
		SDKChannelJsonPbFiled.Set("field", "SDKChannel")
		SDKChannelJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, SDKChannelJsonPbFiled)
	}



	UserInfoObj := new(pb.UserInfo)
	UserInfoType := reflect.TypeOf(*UserInfoObj)
	if UserInfoType.Kind().String() == "struct" {
		UserInfoJson, _ := common.Struct2Json(UserInfoType)
		jsonStructSlice = append(jsonStructSlice, UserInfoJson)
		UserInfoJsonPbFiled := simplejson.New()
		UserInfoJsonPbFiled.Set("field", "UserInfo")
		UserInfoJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, UserInfoJsonPbFiled)
	}



	LineObj := new(pb.Line)
	LineType := reflect.TypeOf(*LineObj)
	if LineType.Kind().String() == "struct" {
		LineJson, _ := common.Struct2Json(LineType)
		jsonStructSlice = append(jsonStructSlice, LineJson)
		LineJsonPbFiled := simplejson.New()
		LineJsonPbFiled.Set("field", "Line")
		LineJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, LineJsonPbFiled)
	}



	BlackObj := new(pb.Black)
	BlackType := reflect.TypeOf(*BlackObj)
	if BlackType.Kind().String() == "struct" {
		BlackJson, _ := common.Struct2Json(BlackType)
		jsonStructSlice = append(jsonStructSlice, BlackJson)
		BlackJsonPbFiled := simplejson.New()
		BlackJsonPbFiled.Set("field", "Black")
		BlackJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, BlackJsonPbFiled)
	}



	UserRoleInfoObj := new(pb.UserRoleInfo)
	UserRoleInfoType := reflect.TypeOf(*UserRoleInfoObj)
	if UserRoleInfoType.Kind().String() == "struct" {
		UserRoleInfoJson, _ := common.Struct2Json(UserRoleInfoType)
		jsonStructSlice = append(jsonStructSlice, UserRoleInfoJson)
		UserRoleInfoJsonPbFiled := simplejson.New()
		UserRoleInfoJsonPbFiled.Set("field", "UserRoleInfo")
		UserRoleInfoJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, UserRoleInfoJsonPbFiled)
	}



	ServerStatusObj := new(pb.ServerStatus)
	ServerStatusType := reflect.TypeOf(*ServerStatusObj)
	if ServerStatusType.Kind().String() == "struct" {
		ServerStatusJson, _ := common.Struct2Json(ServerStatusType)
		jsonStructSlice = append(jsonStructSlice, ServerStatusJson)
		ServerStatusJsonPbFiled := simplejson.New()
		ServerStatusJsonPbFiled.Set("field", "ServerStatus")
		ServerStatusJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, ServerStatusJsonPbFiled)
	}



	LogicServerInfoObj := new(pb.LogicServerInfo)
	LogicServerInfoType := reflect.TypeOf(*LogicServerInfoObj)
	if LogicServerInfoType.Kind().String() == "struct" {
		LogicServerInfoJson, _ := common.Struct2Json(LogicServerInfoType)
		jsonStructSlice = append(jsonStructSlice, LogicServerInfoJson)
		LogicServerInfoJsonPbFiled := simplejson.New()
		LogicServerInfoJsonPbFiled.Set("field", "LogicServerInfo")
		LogicServerInfoJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, LogicServerInfoJsonPbFiled)
	}



	OrderStatusObj := new(pb.OrderStatus)
	OrderStatusType := reflect.TypeOf(*OrderStatusObj)
	if OrderStatusType.Kind().String() == "struct" {
		OrderStatusJson, _ := common.Struct2Json(OrderStatusType)
		jsonStructSlice = append(jsonStructSlice, OrderStatusJson)
		OrderStatusJsonPbFiled := simplejson.New()
		OrderStatusJsonPbFiled.Set("field", "OrderStatus")
		OrderStatusJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, OrderStatusJsonPbFiled)
	}



	ChatDataObj := new(pb.ChatData)
	ChatDataType := reflect.TypeOf(*ChatDataObj)
	if ChatDataType.Kind().String() == "struct" {
		ChatDataJson, _ := common.Struct2Json(ChatDataType)
		jsonStructSlice = append(jsonStructSlice, ChatDataJson)
		ChatDataJsonPbFiled := simplejson.New()
		ChatDataJsonPbFiled.Set("field", "ChatData")
		ChatDataJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, ChatDataJsonPbFiled)
	}



	RoomMemInfoObj := new(pb.RoomMemInfo)
	RoomMemInfoType := reflect.TypeOf(*RoomMemInfoObj)
	if RoomMemInfoType.Kind().String() == "struct" {
		RoomMemInfoJson, _ := common.Struct2Json(RoomMemInfoType)
		jsonStructSlice = append(jsonStructSlice, RoomMemInfoJson)
		RoomMemInfoJsonPbFiled := simplejson.New()
		RoomMemInfoJsonPbFiled.Set("field", "RoomMemInfo")
		RoomMemInfoJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, RoomMemInfoJsonPbFiled)
	}



	PVPStateSyncObj := new(pb.PVPStateSync)
	PVPStateSyncType := reflect.TypeOf(*PVPStateSyncObj)
	if PVPStateSyncType.Kind().String() == "struct" {
		PVPStateSyncJson, _ := common.Struct2Json(PVPStateSyncType)
		jsonStructSlice = append(jsonStructSlice, PVPStateSyncJson)
		PVPStateSyncJsonPbFiled := simplejson.New()
		PVPStateSyncJsonPbFiled.Set("field", "PVPStateSync")
		PVPStateSyncJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, PVPStateSyncJsonPbFiled)
	}



	PVPStateSync_CheckPointObj := new(pb.PVPStateSync_CheckPoint)
	PVPStateSync_CheckPointType := reflect.TypeOf(*PVPStateSync_CheckPointObj)
	if PVPStateSync_CheckPointType.Kind().String() == "struct" {
		PVPStateSync_CheckPointJson, _ := common.Struct2Json(PVPStateSync_CheckPointType)
		jsonStructSlice = append(jsonStructSlice, PVPStateSync_CheckPointJson)
		PVPStateSync_CheckPointJsonPbFiled := simplejson.New()
		PVPStateSync_CheckPointJsonPbFiled.Set("field", "PVPStateSync_CheckPoint")
		PVPStateSync_CheckPointJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, PVPStateSync_CheckPointJsonPbFiled)
	}



	PVPStateSync_JumpObj := new(pb.PVPStateSync_Jump)
	PVPStateSync_JumpType := reflect.TypeOf(*PVPStateSync_JumpObj)
	if PVPStateSync_JumpType.Kind().String() == "struct" {
		PVPStateSync_JumpJson, _ := common.Struct2Json(PVPStateSync_JumpType)
		jsonStructSlice = append(jsonStructSlice, PVPStateSync_JumpJson)
		PVPStateSync_JumpJsonPbFiled := simplejson.New()
		PVPStateSync_JumpJsonPbFiled.Set("field", "PVPStateSync_Jump")
		PVPStateSync_JumpJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, PVPStateSync_JumpJsonPbFiled)
	}



	GamerMiniObj := new(pb.GamerMini)
	GamerMiniType := reflect.TypeOf(*GamerMiniObj)
	if GamerMiniType.Kind().String() == "struct" {
		GamerMiniJson, _ := common.Struct2Json(GamerMiniType)
		jsonStructSlice = append(jsonStructSlice, GamerMiniJson)
		GamerMiniJsonPbFiled := simplejson.New()
		GamerMiniJsonPbFiled.Set("field", "GamerMini")
		GamerMiniJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerMiniJsonPbFiled)
	}



	MatchInfoObj := new(pb.MatchInfo)
	MatchInfoType := reflect.TypeOf(*MatchInfoObj)
	if MatchInfoType.Kind().String() == "struct" {
		MatchInfoJson, _ := common.Struct2Json(MatchInfoType)
		jsonStructSlice = append(jsonStructSlice, MatchInfoJson)
		MatchInfoJsonPbFiled := simplejson.New()
		MatchInfoJsonPbFiled.Set("field", "MatchInfo")
		MatchInfoJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, MatchInfoJsonPbFiled)
	}



	GamerHeroObj := new(pb.GamerHero)
	GamerHeroType := reflect.TypeOf(*GamerHeroObj)
	if GamerHeroType.Kind().String() == "struct" {
		GamerHeroJson, _ := common.Struct2Json(GamerHeroType)
		jsonStructSlice = append(jsonStructSlice, GamerHeroJson)
		GamerHeroJsonPbFiled := simplejson.New()
		GamerHeroJsonPbFiled.Set("field", "GamerHero")
		GamerHeroJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerHeroJsonPbFiled)
	}



	PVPStateSync_PillarObj := new(pb.PVPStateSync_Pillar)
	PVPStateSync_PillarType := reflect.TypeOf(*PVPStateSync_PillarObj)
	if PVPStateSync_PillarType.Kind().String() == "struct" {
		PVPStateSync_PillarJson, _ := common.Struct2Json(PVPStateSync_PillarType)
		jsonStructSlice = append(jsonStructSlice, PVPStateSync_PillarJson)
		PVPStateSync_PillarJsonPbFiled := simplejson.New()
		PVPStateSync_PillarJsonPbFiled.Set("field", "PVPStateSync_Pillar")
		PVPStateSync_PillarJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, PVPStateSync_PillarJsonPbFiled)
	}



	BagItemObj := new(pb.BagItem)
	BagItemType := reflect.TypeOf(*BagItemObj)
	if BagItemType.Kind().String() == "struct" {
		BagItemJson, _ := common.Struct2Json(BagItemType)
		jsonStructSlice = append(jsonStructSlice, BagItemJson)
		BagItemJsonPbFiled := simplejson.New()
		BagItemJsonPbFiled.Set("field", "BagItem")
		BagItemJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, BagItemJsonPbFiled)
	}



	PVPStateSync_GamerEnterObj := new(pb.PVPStateSync_GamerEnter)
	PVPStateSync_GamerEnterType := reflect.TypeOf(*PVPStateSync_GamerEnterObj)
	if PVPStateSync_GamerEnterType.Kind().String() == "struct" {
		PVPStateSync_GamerEnterJson, _ := common.Struct2Json(PVPStateSync_GamerEnterType)
		jsonStructSlice = append(jsonStructSlice, PVPStateSync_GamerEnterJson)
		PVPStateSync_GamerEnterJsonPbFiled := simplejson.New()
		PVPStateSync_GamerEnterJsonPbFiled.Set("field", "PVPStateSync_GamerEnter")
		PVPStateSync_GamerEnterJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, PVPStateSync_GamerEnterJsonPbFiled)
	}



	CountActivityObj := new(pb.CountActivity)
	CountActivityType := reflect.TypeOf(*CountActivityObj)
	if CountActivityType.Kind().String() == "struct" {
		CountActivityJson, _ := common.Struct2Json(CountActivityType)
		jsonStructSlice = append(jsonStructSlice, CountActivityJson)
		CountActivityJsonPbFiled := simplejson.New()
		CountActivityJsonPbFiled.Set("field", "CountActivity")
		CountActivityJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, CountActivityJsonPbFiled)
	}



	GamerMatchObj := new(pb.GamerMatch)
	GamerMatchType := reflect.TypeOf(*GamerMatchObj)
	if GamerMatchType.Kind().String() == "struct" {
		GamerMatchJson, _ := common.Struct2Json(GamerMatchType)
		jsonStructSlice = append(jsonStructSlice, GamerMatchJson)
		GamerMatchJsonPbFiled := simplejson.New()
		GamerMatchJsonPbFiled.Set("field", "GamerMatch")
		GamerMatchJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerMatchJsonPbFiled)
	}



	GameItemObj := new(pb.GameItem)
	GameItemType := reflect.TypeOf(*GameItemObj)
	if GameItemType.Kind().String() == "struct" {
		GameItemJson, _ := common.Struct2Json(GameItemType)
		jsonStructSlice = append(jsonStructSlice, GameItemJson)
		GameItemJsonPbFiled := simplejson.New()
		GameItemJsonPbFiled.Set("field", "GameItem")
		GameItemJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GameItemJsonPbFiled)
	}



	GamerFriendObj := new(pb.GamerFriend)
	GamerFriendType := reflect.TypeOf(*GamerFriendObj)
	if GamerFriendType.Kind().String() == "struct" {
		GamerFriendJson, _ := common.Struct2Json(GamerFriendType)
		jsonStructSlice = append(jsonStructSlice, GamerFriendJson)
		GamerFriendJsonPbFiled := simplejson.New()
		GamerFriendJsonPbFiled.Set("field", "GamerFriend")
		GamerFriendJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerFriendJsonPbFiled)
	}



	PVPStateSync_C2SObj := new(pb.PVPStateSync_C2S)
	PVPStateSync_C2SType := reflect.TypeOf(*PVPStateSync_C2SObj)
	if PVPStateSync_C2SType.Kind().String() == "struct" {
		PVPStateSync_C2SJson, _ := common.Struct2Json(PVPStateSync_C2SType)
		jsonStructSlice = append(jsonStructSlice, PVPStateSync_C2SJson)
		PVPStateSync_C2SJsonPbFiled := simplejson.New()
		PVPStateSync_C2SJsonPbFiled.Set("field", "PVPStateSync_C2S")
		PVPStateSync_C2SJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, PVPStateSync_C2SJsonPbFiled)
	}



	PVPFramesObj := new(pb.PVPFrames)
	PVPFramesType := reflect.TypeOf(*PVPFramesObj)
	if PVPFramesType.Kind().String() == "struct" {
		PVPFramesJson, _ := common.Struct2Json(PVPFramesType)
		jsonStructSlice = append(jsonStructSlice, PVPFramesJson)
		PVPFramesJsonPbFiled := simplejson.New()
		PVPFramesJsonPbFiled.Set("field", "PVPFrames")
		PVPFramesJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, PVPFramesJsonPbFiled)
	}



	PVPStateSync_ResultObj := new(pb.PVPStateSync_Result)
	PVPStateSync_ResultType := reflect.TypeOf(*PVPStateSync_ResultObj)
	if PVPStateSync_ResultType.Kind().String() == "struct" {
		PVPStateSync_ResultJson, _ := common.Struct2Json(PVPStateSync_ResultType)
		jsonStructSlice = append(jsonStructSlice, PVPStateSync_ResultJson)
		PVPStateSync_ResultJsonPbFiled := simplejson.New()
		PVPStateSync_ResultJsonPbFiled.Set("field", "PVPStateSync_Result")
		PVPStateSync_ResultJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, PVPStateSync_ResultJsonPbFiled)
	}



	PVPStateSync_MatchInfoObj := new(pb.PVPStateSync_MatchInfo)
	PVPStateSync_MatchInfoType := reflect.TypeOf(*PVPStateSync_MatchInfoObj)
	if PVPStateSync_MatchInfoType.Kind().String() == "struct" {
		PVPStateSync_MatchInfoJson, _ := common.Struct2Json(PVPStateSync_MatchInfoType)
		jsonStructSlice = append(jsonStructSlice, PVPStateSync_MatchInfoJson)
		PVPStateSync_MatchInfoJsonPbFiled := simplejson.New()
		PVPStateSync_MatchInfoJsonPbFiled.Set("field", "PVPStateSync_MatchInfo")
		PVPStateSync_MatchInfoJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, PVPStateSync_MatchInfoJsonPbFiled)
	}



	ServerTimeObj := new(pb.ServerTime)
	ServerTimeType := reflect.TypeOf(*ServerTimeObj)
	if ServerTimeType.Kind().String() == "struct" {
		ServerTimeJson, _ := common.Struct2Json(ServerTimeType)
		jsonStructSlice = append(jsonStructSlice, ServerTimeJson)
		ServerTimeJsonPbFiled := simplejson.New()
		ServerTimeJsonPbFiled.Set("field", "ServerTime")
		ServerTimeJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, ServerTimeJsonPbFiled)
	}



	PVPStateSync_GamerInfoObj := new(pb.PVPStateSync_GamerInfo)
	PVPStateSync_GamerInfoType := reflect.TypeOf(*PVPStateSync_GamerInfoObj)
	if PVPStateSync_GamerInfoType.Kind().String() == "struct" {
		PVPStateSync_GamerInfoJson, _ := common.Struct2Json(PVPStateSync_GamerInfoType)
		jsonStructSlice = append(jsonStructSlice, PVPStateSync_GamerInfoJson)
		PVPStateSync_GamerInfoJsonPbFiled := simplejson.New()
		PVPStateSync_GamerInfoJsonPbFiled.Set("field", "PVPStateSync_GamerInfo")
		PVPStateSync_GamerInfoJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, PVPStateSync_GamerInfoJsonPbFiled)
	}



	GamerMainObj := new(pb.GamerMain)
	GamerMainType := reflect.TypeOf(*GamerMainObj)
	if GamerMainType.Kind().String() == "struct" {
		GamerMainJson, _ := common.Struct2Json(GamerMainType)
		jsonStructSlice = append(jsonStructSlice, GamerMainJson)
		GamerMainJsonPbFiled := simplejson.New()
		GamerMainJsonPbFiled.Set("field", "GamerMain")
		GamerMainJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerMainJsonPbFiled)
	}



	NumberRecordObj := new(pb.NumberRecord)
	NumberRecordType := reflect.TypeOf(*NumberRecordObj)
	if NumberRecordType.Kind().String() == "struct" {
		NumberRecordJson, _ := common.Struct2Json(NumberRecordType)
		jsonStructSlice = append(jsonStructSlice, NumberRecordJson)
		NumberRecordJsonPbFiled := simplejson.New()
		NumberRecordJsonPbFiled.Set("field", "NumberRecord")
		NumberRecordJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, NumberRecordJsonPbFiled)
	}



	GamerFriendRequestObj := new(pb.GamerFriendRequest)
	GamerFriendRequestType := reflect.TypeOf(*GamerFriendRequestObj)
	if GamerFriendRequestType.Kind().String() == "struct" {
		GamerFriendRequestJson, _ := common.Struct2Json(GamerFriendRequestType)
		jsonStructSlice = append(jsonStructSlice, GamerFriendRequestJson)
		GamerFriendRequestJsonPbFiled := simplejson.New()
		GamerFriendRequestJsonPbFiled.Set("field", "GamerFriendRequest")
		GamerFriendRequestJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerFriendRequestJsonPbFiled)
	}



	PVPStateSync_S2CObj := new(pb.PVPStateSync_S2C)
	PVPStateSync_S2CType := reflect.TypeOf(*PVPStateSync_S2CObj)
	if PVPStateSync_S2CType.Kind().String() == "struct" {
		PVPStateSync_S2CJson, _ := common.Struct2Json(PVPStateSync_S2CType)
		jsonStructSlice = append(jsonStructSlice, PVPStateSync_S2CJson)
		PVPStateSync_S2CJsonPbFiled := simplejson.New()
		PVPStateSync_S2CJsonPbFiled.Set("field", "PVPStateSync_S2C")
		PVPStateSync_S2CJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, PVPStateSync_S2CJsonPbFiled)
	}



	MailObj := new(pb.Mail)
	MailType := reflect.TypeOf(*MailObj)
	if MailType.Kind().String() == "struct" {
		MailJson, _ := common.Struct2Json(MailType)
		jsonStructSlice = append(jsonStructSlice, MailJson)
		MailJsonPbFiled := simplejson.New()
		MailJsonPbFiled.Set("field", "Mail")
		MailJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, MailJsonPbFiled)
	}



	RankLengthObj := new(pb.RankLength)
	RankLengthType := reflect.TypeOf(*RankLengthObj)
	if RankLengthType.Kind().String() == "struct" {
		RankLengthJson, _ := common.Struct2Json(RankLengthType)
		jsonStructSlice = append(jsonStructSlice, RankLengthJson)
		RankLengthJsonPbFiled := simplejson.New()
		RankLengthJsonPbFiled.Set("field", "RankLength")
		RankLengthJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, RankLengthJsonPbFiled)
	}



	RoomInfoObj := new(pb.RoomInfo)
	RoomInfoType := reflect.TypeOf(*RoomInfoObj)
	if RoomInfoType.Kind().String() == "struct" {
		RoomInfoJson, _ := common.Struct2Json(RoomInfoType)
		jsonStructSlice = append(jsonStructSlice, RoomInfoJson)
		RoomInfoJsonPbFiled := simplejson.New()
		RoomInfoJsonPbFiled.Set("field", "RoomInfo")
		RoomInfoJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, RoomInfoJsonPbFiled)
	}



	PVPStateSync_ItemObj := new(pb.PVPStateSync_Item)
	PVPStateSync_ItemType := reflect.TypeOf(*PVPStateSync_ItemObj)
	if PVPStateSync_ItemType.Kind().String() == "struct" {
		PVPStateSync_ItemJson, _ := common.Struct2Json(PVPStateSync_ItemType)
		jsonStructSlice = append(jsonStructSlice, PVPStateSync_ItemJson)
		PVPStateSync_ItemJsonPbFiled := simplejson.New()
		PVPStateSync_ItemJsonPbFiled.Set("field", "PVPStateSync_Item")
		PVPStateSync_ItemJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, PVPStateSync_ItemJsonPbFiled)
	}



	GamerBlackObj := new(pb.GamerBlack)
	GamerBlackType := reflect.TypeOf(*GamerBlackObj)
	if GamerBlackType.Kind().String() == "struct" {
		GamerBlackJson, _ := common.Struct2Json(GamerBlackType)
		jsonStructSlice = append(jsonStructSlice, GamerBlackJson)
		GamerBlackJsonPbFiled := simplejson.New()
		GamerBlackJsonPbFiled.Set("field", "GamerBlack")
		GamerBlackJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerBlackJsonPbFiled)
	}



	PVPResultObj := new(pb.PVPResult)
	PVPResultType := reflect.TypeOf(*PVPResultObj)
	if PVPResultType.Kind().String() == "struct" {
		PVPResultJson, _ := common.Struct2Json(PVPResultType)
		jsonStructSlice = append(jsonStructSlice, PVPResultJson)
		PVPResultJsonPbFiled := simplejson.New()
		PVPResultJsonPbFiled.Set("field", "PVPResult")
		PVPResultJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, PVPResultJsonPbFiled)
	}



	GashaponObj := new(pb.Gashapon)
	GashaponType := reflect.TypeOf(*GashaponObj)
	if GashaponType.Kind().String() == "struct" {
		GashaponJson, _ := common.Struct2Json(GashaponType)
		jsonStructSlice = append(jsonStructSlice, GashaponJson)
		GashaponJsonPbFiled := simplejson.New()
		GashaponJsonPbFiled.Set("field", "Gashapon")
		GashaponJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GashaponJsonPbFiled)
	}



	PVPInputObj := new(pb.PVPInput)
	PVPInputType := reflect.TypeOf(*PVPInputObj)
	if PVPInputType.Kind().String() == "struct" {
		PVPInputJson, _ := common.Struct2Json(PVPInputType)
		jsonStructSlice = append(jsonStructSlice, PVPInputJson)
		PVPInputJsonPbFiled := simplejson.New()
		PVPInputJsonPbFiled.Set("field", "PVPInput")
		PVPInputJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, PVPInputJsonPbFiled)
	}



	PVPFrameObj := new(pb.PVPFrame)
	PVPFrameType := reflect.TypeOf(*PVPFrameObj)
	if PVPFrameType.Kind().String() == "struct" {
		PVPFrameJson, _ := common.Struct2Json(PVPFrameType)
		jsonStructSlice = append(jsonStructSlice, PVPFrameJson)
		PVPFrameJsonPbFiled := simplejson.New()
		PVPFrameJsonPbFiled.Set("field", "PVPFrame")
		PVPFrameJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, PVPFrameJsonPbFiled)
	}



	PrevC2SObj := new(pb.PrevC2S)
	PrevC2SType := reflect.TypeOf(*PrevC2SObj)
	if PrevC2SType.Kind().String() == "struct" {
		PrevC2SJson, _ := common.Struct2Json(PrevC2SType)
		jsonStructSlice = append(jsonStructSlice, PrevC2SJson)
		PrevC2SJsonPbFiled := simplejson.New()
		PrevC2SJsonPbFiled.Set("field", "PrevC2S")
		PrevC2SJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, PrevC2SJsonPbFiled)
	}



	TimeRecordObj := new(pb.TimeRecord)
	TimeRecordType := reflect.TypeOf(*TimeRecordObj)
	if TimeRecordType.Kind().String() == "struct" {
		TimeRecordJson, _ := common.Struct2Json(TimeRecordType)
		jsonStructSlice = append(jsonStructSlice, TimeRecordJson)
		TimeRecordJsonPbFiled := simplejson.New()
		TimeRecordJsonPbFiled.Set("field", "TimeRecord")
		TimeRecordJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, TimeRecordJsonPbFiled)
	}



	WXInfoObj := new(pb.WXInfo)
	WXInfoType := reflect.TypeOf(*WXInfoObj)
	if WXInfoType.Kind().String() == "struct" {
		WXInfoJson, _ := common.Struct2Json(WXInfoType)
		jsonStructSlice = append(jsonStructSlice, WXInfoJson)
		WXInfoJsonPbFiled := simplejson.New()
		WXInfoJsonPbFiled.Set("field", "WXInfo")
		WXInfoJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, WXInfoJsonPbFiled)
	}



	GamerNotifyExpLevelS2CObj := new(pb.GamerNotifyExpLevelS2C)
	GamerNotifyExpLevelS2CType := reflect.TypeOf(*GamerNotifyExpLevelS2CObj)
	if GamerNotifyExpLevelS2CType.Kind().String() == "struct" {
		GamerNotifyExpLevelS2CJson, _ := common.Struct2Json(GamerNotifyExpLevelS2CType)
		jsonStructSlice = append(jsonStructSlice, GamerNotifyExpLevelS2CJson)
		GamerNotifyExpLevelS2CJsonPbFiled := simplejson.New()
		GamerNotifyExpLevelS2CJsonPbFiled.Set("field", "GamerNotifyExpLevelS2C")
		GamerNotifyExpLevelS2CJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerNotifyExpLevelS2CJsonPbFiled)
	}



	GamerSellItemC2SObj := new(pb.GamerSellItemC2S)
	GamerSellItemC2SType := reflect.TypeOf(*GamerSellItemC2SObj)
	if GamerSellItemC2SType.Kind().String() == "struct" {
		GamerSellItemC2SJson, _ := common.Struct2Json(GamerSellItemC2SType)
		jsonStructSlice = append(jsonStructSlice, GamerSellItemC2SJson)
		GamerSellItemC2SJsonPbFiled := simplejson.New()
		GamerSellItemC2SJsonPbFiled.Set("field", "GamerSellItemC2S")
		GamerSellItemC2SJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerSellItemC2SJsonPbFiled)
	}



	GamerNotifyExpChangeS2CObj := new(pb.GamerNotifyExpChangeS2C)
	GamerNotifyExpChangeS2CType := reflect.TypeOf(*GamerNotifyExpChangeS2CObj)
	if GamerNotifyExpChangeS2CType.Kind().String() == "struct" {
		GamerNotifyExpChangeS2CJson, _ := common.Struct2Json(GamerNotifyExpChangeS2CType)
		jsonStructSlice = append(jsonStructSlice, GamerNotifyExpChangeS2CJson)
		GamerNotifyExpChangeS2CJsonPbFiled := simplejson.New()
		GamerNotifyExpChangeS2CJsonPbFiled.Set("field", "GamerNotifyExpChangeS2C")
		GamerNotifyExpChangeS2CJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerNotifyExpChangeS2CJsonPbFiled)
	}



	GamerSubChatChannelS2CObj := new(pb.GamerSubChatChannelS2C)
	GamerSubChatChannelS2CType := reflect.TypeOf(*GamerSubChatChannelS2CObj)
	if GamerSubChatChannelS2CType.Kind().String() == "struct" {
		GamerSubChatChannelS2CJson, _ := common.Struct2Json(GamerSubChatChannelS2CType)
		jsonStructSlice = append(jsonStructSlice, GamerSubChatChannelS2CJson)
		GamerSubChatChannelS2CJsonPbFiled := simplejson.New()
		GamerSubChatChannelS2CJsonPbFiled.Set("field", "GamerSubChatChannelS2C")
		GamerSubChatChannelS2CJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerSubChatChannelS2CJsonPbFiled)
	}



	GamerChangeNameS2CObj := new(pb.GamerChangeNameS2C)
	GamerChangeNameS2CType := reflect.TypeOf(*GamerChangeNameS2CObj)
	if GamerChangeNameS2CType.Kind().String() == "struct" {
		GamerChangeNameS2CJson, _ := common.Struct2Json(GamerChangeNameS2CType)
		jsonStructSlice = append(jsonStructSlice, GamerChangeNameS2CJson)
		GamerChangeNameS2CJsonPbFiled := simplejson.New()
		GamerChangeNameS2CJsonPbFiled.Set("field", "GamerChangeNameS2C")
		GamerChangeNameS2CJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerChangeNameS2CJsonPbFiled)
	}



	GamerNotifyNewPVPResultS2CObj := new(pb.GamerNotifyNewPVPResultS2C)
	GamerNotifyNewPVPResultS2CType := reflect.TypeOf(*GamerNotifyNewPVPResultS2CObj)
	if GamerNotifyNewPVPResultS2CType.Kind().String() == "struct" {
		GamerNotifyNewPVPResultS2CJson, _ := common.Struct2Json(GamerNotifyNewPVPResultS2CType)
		jsonStructSlice = append(jsonStructSlice, GamerNotifyNewPVPResultS2CJson)
		GamerNotifyNewPVPResultS2CJsonPbFiled := simplejson.New()
		GamerNotifyNewPVPResultS2CJsonPbFiled.Set("field", "GamerNotifyNewPVPResultS2C")
		GamerNotifyNewPVPResultS2CJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerNotifyNewPVPResultS2CJsonPbFiled)
	}



	GamerNewFriendReqC2SObj := new(pb.GamerNewFriendReqC2S)
	GamerNewFriendReqC2SType := reflect.TypeOf(*GamerNewFriendReqC2SObj)
	if GamerNewFriendReqC2SType.Kind().String() == "struct" {
		GamerNewFriendReqC2SJson, _ := common.Struct2Json(GamerNewFriendReqC2SType)
		jsonStructSlice = append(jsonStructSlice, GamerNewFriendReqC2SJson)
		GamerNewFriendReqC2SJsonPbFiled := simplejson.New()
		GamerNewFriendReqC2SJsonPbFiled.Set("field", "GamerNewFriendReqC2S")
		GamerNewFriendReqC2SJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerNewFriendReqC2SJsonPbFiled)
	}



	GamerCheckPVPBattleC2SObj := new(pb.GamerCheckPVPBattleC2S)
	GamerCheckPVPBattleC2SType := reflect.TypeOf(*GamerCheckPVPBattleC2SObj)
	if GamerCheckPVPBattleC2SType.Kind().String() == "struct" {
		GamerCheckPVPBattleC2SJson, _ := common.Struct2Json(GamerCheckPVPBattleC2SType)
		jsonStructSlice = append(jsonStructSlice, GamerCheckPVPBattleC2SJson)
		GamerCheckPVPBattleC2SJsonPbFiled := simplejson.New()
		GamerCheckPVPBattleC2SJsonPbFiled.Set("field", "GamerCheckPVPBattleC2S")
		GamerCheckPVPBattleC2SJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerCheckPVPBattleC2SJsonPbFiled)
	}



	ServerTimeS2CObj := new(pb.ServerTimeS2C)
	ServerTimeS2CType := reflect.TypeOf(*ServerTimeS2CObj)
	if ServerTimeS2CType.Kind().String() == "struct" {
		ServerTimeS2CJson, _ := common.Struct2Json(ServerTimeS2CType)
		jsonStructSlice = append(jsonStructSlice, ServerTimeS2CJson)
		ServerTimeS2CJsonPbFiled := simplejson.New()
		ServerTimeS2CJsonPbFiled.Set("field", "ServerTimeS2C")
		ServerTimeS2CJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, ServerTimeS2CJsonPbFiled)
	}



	GamerNotifyIconChangeS2CObj := new(pb.GamerNotifyIconChangeS2C)
	GamerNotifyIconChangeS2CType := reflect.TypeOf(*GamerNotifyIconChangeS2CObj)
	if GamerNotifyIconChangeS2CType.Kind().String() == "struct" {
		GamerNotifyIconChangeS2CJson, _ := common.Struct2Json(GamerNotifyIconChangeS2CType)
		jsonStructSlice = append(jsonStructSlice, GamerNotifyIconChangeS2CJson)
		GamerNotifyIconChangeS2CJsonPbFiled := simplejson.New()
		GamerNotifyIconChangeS2CJsonPbFiled.Set("field", "GamerNotifyIconChangeS2C")
		GamerNotifyIconChangeS2CJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerNotifyIconChangeS2CJsonPbFiled)
	}



	GamerChangeReadyInRoomC2SObj := new(pb.GamerChangeReadyInRoomC2S)
	GamerChangeReadyInRoomC2SType := reflect.TypeOf(*GamerChangeReadyInRoomC2SObj)
	if GamerChangeReadyInRoomC2SType.Kind().String() == "struct" {
		GamerChangeReadyInRoomC2SJson, _ := common.Struct2Json(GamerChangeReadyInRoomC2SType)
		jsonStructSlice = append(jsonStructSlice, GamerChangeReadyInRoomC2SJson)
		GamerChangeReadyInRoomC2SJsonPbFiled := simplejson.New()
		GamerChangeReadyInRoomC2SJsonPbFiled.Set("field", "GamerChangeReadyInRoomC2S")
		GamerChangeReadyInRoomC2SJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerChangeReadyInRoomC2SJsonPbFiled)
	}



	GamerClubRequestC2SObj := new(pb.GamerClubRequestC2S)
	GamerClubRequestC2SType := reflect.TypeOf(*GamerClubRequestC2SObj)
	if GamerClubRequestC2SType.Kind().String() == "struct" {
		GamerClubRequestC2SJson, _ := common.Struct2Json(GamerClubRequestC2SType)
		jsonStructSlice = append(jsonStructSlice, GamerClubRequestC2SJson)
		GamerClubRequestC2SJsonPbFiled := simplejson.New()
		GamerClubRequestC2SJsonPbFiled.Set("field", "GamerClubRequestC2S")
		GamerClubRequestC2SJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerClubRequestC2SJsonPbFiled)
	}



	GamerNotifyGamerMiniS2CObj := new(pb.GamerNotifyGamerMiniS2C)
	GamerNotifyGamerMiniS2CType := reflect.TypeOf(*GamerNotifyGamerMiniS2CObj)
	if GamerNotifyGamerMiniS2CType.Kind().String() == "struct" {
		GamerNotifyGamerMiniS2CJson, _ := common.Struct2Json(GamerNotifyGamerMiniS2CType)
		jsonStructSlice = append(jsonStructSlice, GamerNotifyGamerMiniS2CJson)
		GamerNotifyGamerMiniS2CJsonPbFiled := simplejson.New()
		GamerNotifyGamerMiniS2CJsonPbFiled.Set("field", "GamerNotifyGamerMiniS2C")
		GamerNotifyGamerMiniS2CJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerNotifyGamerMiniS2CJsonPbFiled)
	}



	GamerNotifyRoomInfoChangeS2CObj := new(pb.GamerNotifyRoomInfoChangeS2C)
	GamerNotifyRoomInfoChangeS2CType := reflect.TypeOf(*GamerNotifyRoomInfoChangeS2CObj)
	if GamerNotifyRoomInfoChangeS2CType.Kind().String() == "struct" {
		GamerNotifyRoomInfoChangeS2CJson, _ := common.Struct2Json(GamerNotifyRoomInfoChangeS2CType)
		jsonStructSlice = append(jsonStructSlice, GamerNotifyRoomInfoChangeS2CJson)
		GamerNotifyRoomInfoChangeS2CJsonPbFiled := simplejson.New()
		GamerNotifyRoomInfoChangeS2CJsonPbFiled.Set("field", "GamerNotifyRoomInfoChangeS2C")
		GamerNotifyRoomInfoChangeS2CJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerNotifyRoomInfoChangeS2CJsonPbFiled)
	}



	GamerFriendChatS2CObj := new(pb.GamerFriendChatS2C)
	GamerFriendChatS2CType := reflect.TypeOf(*GamerFriendChatS2CObj)
	if GamerFriendChatS2CType.Kind().String() == "struct" {
		GamerFriendChatS2CJson, _ := common.Struct2Json(GamerFriendChatS2CType)
		jsonStructSlice = append(jsonStructSlice, GamerFriendChatS2CJson)
		GamerFriendChatS2CJsonPbFiled := simplejson.New()
		GamerFriendChatS2CJsonPbFiled.Set("field", "GamerFriendChatS2C")
		GamerFriendChatS2CJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerFriendChatS2CJsonPbFiled)
	}



	GamerChoseHeroC2SObj := new(pb.GamerChoseHeroC2S)
	GamerChoseHeroC2SType := reflect.TypeOf(*GamerChoseHeroC2SObj)
	if GamerChoseHeroC2SType.Kind().String() == "struct" {
		GamerChoseHeroC2SJson, _ := common.Struct2Json(GamerChoseHeroC2SType)
		jsonStructSlice = append(jsonStructSlice, GamerChoseHeroC2SJson)
		GamerChoseHeroC2SJsonPbFiled := simplejson.New()
		GamerChoseHeroC2SJsonPbFiled.Set("field", "GamerChoseHeroC2S")
		GamerChoseHeroC2SJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerChoseHeroC2SJsonPbFiled)
	}



	GamerUploadWXInfoC2SObj := new(pb.GamerUploadWXInfoC2S)
	GamerUploadWXInfoC2SType := reflect.TypeOf(*GamerUploadWXInfoC2SObj)
	if GamerUploadWXInfoC2SType.Kind().String() == "struct" {
		GamerUploadWXInfoC2SJson, _ := common.Struct2Json(GamerUploadWXInfoC2SType)
		jsonStructSlice = append(jsonStructSlice, GamerUploadWXInfoC2SJson)
		GamerUploadWXInfoC2SJsonPbFiled := simplejson.New()
		GamerUploadWXInfoC2SJsonPbFiled.Set("field", "GamerUploadWXInfoC2S")
		GamerUploadWXInfoC2SJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerUploadWXInfoC2SJsonPbFiled)
	}



	GamerLeaveRoomS2CObj := new(pb.GamerLeaveRoomS2C)
	GamerLeaveRoomS2CType := reflect.TypeOf(*GamerLeaveRoomS2CObj)
	if GamerLeaveRoomS2CType.Kind().String() == "struct" {
		GamerLeaveRoomS2CJson, _ := common.Struct2Json(GamerLeaveRoomS2CType)
		jsonStructSlice = append(jsonStructSlice, GamerLeaveRoomS2CJson)
		GamerLeaveRoomS2CJsonPbFiled := simplejson.New()
		GamerLeaveRoomS2CJsonPbFiled.Set("field", "GamerLeaveRoomS2C")
		GamerLeaveRoomS2CJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerLeaveRoomS2CJsonPbFiled)
	}



	GamerDelFriendC2SObj := new(pb.GamerDelFriendC2S)
	GamerDelFriendC2SType := reflect.TypeOf(*GamerDelFriendC2SObj)
	if GamerDelFriendC2SType.Kind().String() == "struct" {
		GamerDelFriendC2SJson, _ := common.Struct2Json(GamerDelFriendC2SType)
		jsonStructSlice = append(jsonStructSlice, GamerDelFriendC2SJson)
		GamerDelFriendC2SJsonPbFiled := simplejson.New()
		GamerDelFriendC2SJsonPbFiled.Set("field", "GamerDelFriendC2S")
		GamerDelFriendC2SJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerDelFriendC2SJsonPbFiled)
	}



	GamerNewHeroS2CObj := new(pb.GamerNewHeroS2C)
	GamerNewHeroS2CType := reflect.TypeOf(*GamerNewHeroS2CObj)
	if GamerNewHeroS2CType.Kind().String() == "struct" {
		GamerNewHeroS2CJson, _ := common.Struct2Json(GamerNewHeroS2CType)
		jsonStructSlice = append(jsonStructSlice, GamerNewHeroS2CJson)
		GamerNewHeroS2CJsonPbFiled := simplejson.New()
		GamerNewHeroS2CJsonPbFiled.Set("field", "GamerNewHeroS2C")
		GamerNewHeroS2CJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerNewHeroS2CJsonPbFiled)
	}



	GamerLoginC2SObj := new(pb.GamerLoginC2S)
	GamerLoginC2SType := reflect.TypeOf(*GamerLoginC2SObj)
	if GamerLoginC2SType.Kind().String() == "struct" {
		GamerLoginC2SJson, _ := common.Struct2Json(GamerLoginC2SType)
		jsonStructSlice = append(jsonStructSlice, GamerLoginC2SJson)
		GamerLoginC2SJsonPbFiled := simplejson.New()
		GamerLoginC2SJsonPbFiled.Set("field", "GamerLoginC2S")
		GamerLoginC2SJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerLoginC2SJsonPbFiled)
	}



	GamerGetBackpackS2CObj := new(pb.GamerGetBackpackS2C)
	GamerGetBackpackS2CType := reflect.TypeOf(*GamerGetBackpackS2CObj)
	if GamerGetBackpackS2CType.Kind().String() == "struct" {
		GamerGetBackpackS2CJson, _ := common.Struct2Json(GamerGetBackpackS2CType)
		jsonStructSlice = append(jsonStructSlice, GamerGetBackpackS2CJson)
		GamerGetBackpackS2CJsonPbFiled := simplejson.New()
		GamerGetBackpackS2CJsonPbFiled.Set("field", "GamerGetBackpackS2C")
		GamerGetBackpackS2CJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerGetBackpackS2CJsonPbFiled)
	}



	GamerExtractGashaponS2CObj := new(pb.GamerExtractGashaponS2C)
	GamerExtractGashaponS2CType := reflect.TypeOf(*GamerExtractGashaponS2CObj)
	if GamerExtractGashaponS2CType.Kind().String() == "struct" {
		GamerExtractGashaponS2CJson, _ := common.Struct2Json(GamerExtractGashaponS2CType)
		jsonStructSlice = append(jsonStructSlice, GamerExtractGashaponS2CJson)
		GamerExtractGashaponS2CJsonPbFiled := simplejson.New()
		GamerExtractGashaponS2CJsonPbFiled.Set("field", "GamerExtractGashaponS2C")
		GamerExtractGashaponS2CJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerExtractGashaponS2CJsonPbFiled)
	}



	GamerBuyItemC2SObj := new(pb.GamerBuyItemC2S)
	GamerBuyItemC2SType := reflect.TypeOf(*GamerBuyItemC2SObj)
	if GamerBuyItemC2SType.Kind().String() == "struct" {
		GamerBuyItemC2SJson, _ := common.Struct2Json(GamerBuyItemC2SType)
		jsonStructSlice = append(jsonStructSlice, GamerBuyItemC2SJson)
		GamerBuyItemC2SJsonPbFiled := simplejson.New()
		GamerBuyItemC2SJsonPbFiled.Set("field", "GamerBuyItemC2S")
		GamerBuyItemC2SJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerBuyItemC2SJsonPbFiled)
	}



	GetActivityRewardS2CObj := new(pb.GetActivityRewardS2C)
	GetActivityRewardS2CType := reflect.TypeOf(*GetActivityRewardS2CObj)
	if GetActivityRewardS2CType.Kind().String() == "struct" {
		GetActivityRewardS2CJson, _ := common.Struct2Json(GetActivityRewardS2CType)
		jsonStructSlice = append(jsonStructSlice, GetActivityRewardS2CJson)
		GetActivityRewardS2CJsonPbFiled := simplejson.New()
		GetActivityRewardS2CJsonPbFiled.Set("field", "GetActivityRewardS2C")
		GetActivityRewardS2CJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GetActivityRewardS2CJsonPbFiled)
	}



	GamerNotifyChangeReadyInRoomS2CObj := new(pb.GamerNotifyChangeReadyInRoomS2C)
	GamerNotifyChangeReadyInRoomS2CType := reflect.TypeOf(*GamerNotifyChangeReadyInRoomS2CObj)
	if GamerNotifyChangeReadyInRoomS2CType.Kind().String() == "struct" {
		GamerNotifyChangeReadyInRoomS2CJson, _ := common.Struct2Json(GamerNotifyChangeReadyInRoomS2CType)
		jsonStructSlice = append(jsonStructSlice, GamerNotifyChangeReadyInRoomS2CJson)
		GamerNotifyChangeReadyInRoomS2CJsonPbFiled := simplejson.New()
		GamerNotifyChangeReadyInRoomS2CJsonPbFiled.Set("field", "GamerNotifyChangeReadyInRoomS2C")
		GamerNotifyChangeReadyInRoomS2CJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerNotifyChangeReadyInRoomS2CJsonPbFiled)
	}



	RoomListC2SObj := new(pb.RoomListC2S)
	RoomListC2SType := reflect.TypeOf(*RoomListC2SObj)
	if RoomListC2SType.Kind().String() == "struct" {
		RoomListC2SJson, _ := common.Struct2Json(RoomListC2SType)
		jsonStructSlice = append(jsonStructSlice, RoomListC2SJson)
		RoomListC2SJsonPbFiled := simplejson.New()
		RoomListC2SJsonPbFiled.Set("field", "RoomListC2S")
		RoomListC2SJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, RoomListC2SJsonPbFiled)
	}



	GamerGetRealTimeRankC2SObj := new(pb.GamerGetRealTimeRankC2S)
	GamerGetRealTimeRankC2SType := reflect.TypeOf(*GamerGetRealTimeRankC2SObj)
	if GamerGetRealTimeRankC2SType.Kind().String() == "struct" {
		GamerGetRealTimeRankC2SJson, _ := common.Struct2Json(GamerGetRealTimeRankC2SType)
		jsonStructSlice = append(jsonStructSlice, GamerGetRealTimeRankC2SJson)
		GamerGetRealTimeRankC2SJsonPbFiled := simplejson.New()
		GamerGetRealTimeRankC2SJsonPbFiled.Set("field", "GamerGetRealTimeRankC2S")
		GamerGetRealTimeRankC2SJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerGetRealTimeRankC2SJsonPbFiled)
	}



	GamerSellItemS2CObj := new(pb.GamerSellItemS2C)
	GamerSellItemS2CType := reflect.TypeOf(*GamerSellItemS2CObj)
	if GamerSellItemS2CType.Kind().String() == "struct" {
		GamerSellItemS2CJson, _ := common.Struct2Json(GamerSellItemS2CType)
		jsonStructSlice = append(jsonStructSlice, GamerSellItemS2CJson)
		GamerSellItemS2CJsonPbFiled := simplejson.New()
		GamerSellItemS2CJsonPbFiled.Set("field", "GamerSellItemS2C")
		GamerSellItemS2CJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerSellItemS2CJsonPbFiled)
	}



	GamerGetAllowanceS2CObj := new(pb.GamerGetAllowanceS2C)
	GamerGetAllowanceS2CType := reflect.TypeOf(*GamerGetAllowanceS2CObj)
	if GamerGetAllowanceS2CType.Kind().String() == "struct" {
		GamerGetAllowanceS2CJson, _ := common.Struct2Json(GamerGetAllowanceS2CType)
		jsonStructSlice = append(jsonStructSlice, GamerGetAllowanceS2CJson)
		GamerGetAllowanceS2CJsonPbFiled := simplejson.New()
		GamerGetAllowanceS2CJsonPbFiled.Set("field", "GamerGetAllowanceS2C")
		GamerGetAllowanceS2CJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerGetAllowanceS2CJsonPbFiled)
	}



	GamerNotifyRoomStateChangeS2CObj := new(pb.GamerNotifyRoomStateChangeS2C)
	GamerNotifyRoomStateChangeS2CType := reflect.TypeOf(*GamerNotifyRoomStateChangeS2CObj)
	if GamerNotifyRoomStateChangeS2CType.Kind().String() == "struct" {
		GamerNotifyRoomStateChangeS2CJson, _ := common.Struct2Json(GamerNotifyRoomStateChangeS2CType)
		jsonStructSlice = append(jsonStructSlice, GamerNotifyRoomStateChangeS2CJson)
		GamerNotifyRoomStateChangeS2CJsonPbFiled := simplejson.New()
		GamerNotifyRoomStateChangeS2CJsonPbFiled.Set("field", "GamerNotifyRoomStateChangeS2C")
		GamerNotifyRoomStateChangeS2CJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerNotifyRoomStateChangeS2CJsonPbFiled)
	}



	GamerUploadWXInfoS2CObj := new(pb.GamerUploadWXInfoS2C)
	GamerUploadWXInfoS2CType := reflect.TypeOf(*GamerUploadWXInfoS2CObj)
	if GamerUploadWXInfoS2CType.Kind().String() == "struct" {
		GamerUploadWXInfoS2CJson, _ := common.Struct2Json(GamerUploadWXInfoS2CType)
		jsonStructSlice = append(jsonStructSlice, GamerUploadWXInfoS2CJson)
		GamerUploadWXInfoS2CJsonPbFiled := simplejson.New()
		GamerUploadWXInfoS2CJsonPbFiled.Set("field", "GamerUploadWXInfoS2C")
		GamerUploadWXInfoS2CJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerUploadWXInfoS2CJsonPbFiled)
	}



	ServerTimeC2SObj := new(pb.ServerTimeC2S)
	ServerTimeC2SType := reflect.TypeOf(*ServerTimeC2SObj)
	if ServerTimeC2SType.Kind().String() == "struct" {
		ServerTimeC2SJson, _ := common.Struct2Json(ServerTimeC2SType)
		jsonStructSlice = append(jsonStructSlice, ServerTimeC2SJson)
		ServerTimeC2SJsonPbFiled := simplejson.New()
		ServerTimeC2SJsonPbFiled.Set("field", "ServerTimeC2S")
		ServerTimeC2SJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, ServerTimeC2SJsonPbFiled)
	}



	GamerTestChatC2SObj := new(pb.GamerTestChatC2S)
	GamerTestChatC2SType := reflect.TypeOf(*GamerTestChatC2SObj)
	if GamerTestChatC2SType.Kind().String() == "struct" {
		GamerTestChatC2SJson, _ := common.Struct2Json(GamerTestChatC2SType)
		jsonStructSlice = append(jsonStructSlice, GamerTestChatC2SJson)
		GamerTestChatC2SJsonPbFiled := simplejson.New()
		GamerTestChatC2SJsonPbFiled.Set("field", "GamerTestChatC2S")
		GamerTestChatC2SJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerTestChatC2SJsonPbFiled)
	}



	GamerUpgradeHeroC2SObj := new(pb.GamerUpgradeHeroC2S)
	GamerUpgradeHeroC2SType := reflect.TypeOf(*GamerUpgradeHeroC2SObj)
	if GamerUpgradeHeroC2SType.Kind().String() == "struct" {
		GamerUpgradeHeroC2SJson, _ := common.Struct2Json(GamerUpgradeHeroC2SType)
		jsonStructSlice = append(jsonStructSlice, GamerUpgradeHeroC2SJson)
		GamerUpgradeHeroC2SJsonPbFiled := simplejson.New()
		GamerUpgradeHeroC2SJsonPbFiled.Set("field", "GamerUpgradeHeroC2S")
		GamerUpgradeHeroC2SJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerUpgradeHeroC2SJsonPbFiled)
	}



	GamerNewRoomS2CObj := new(pb.GamerNewRoomS2C)
	GamerNewRoomS2CType := reflect.TypeOf(*GamerNewRoomS2CObj)
	if GamerNewRoomS2CType.Kind().String() == "struct" {
		GamerNewRoomS2CJson, _ := common.Struct2Json(GamerNewRoomS2CType)
		jsonStructSlice = append(jsonStructSlice, GamerNewRoomS2CJson)
		GamerNewRoomS2CJsonPbFiled := simplejson.New()
		GamerNewRoomS2CJsonPbFiled.Set("field", "GamerNewRoomS2C")
		GamerNewRoomS2CJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerNewRoomS2CJsonPbFiled)
	}



	GamerWorldChatS2CObj := new(pb.GamerWorldChatS2C)
	GamerWorldChatS2CType := reflect.TypeOf(*GamerWorldChatS2CObj)
	if GamerWorldChatS2CType.Kind().String() == "struct" {
		GamerWorldChatS2CJson, _ := common.Struct2Json(GamerWorldChatS2CType)
		jsonStructSlice = append(jsonStructSlice, GamerWorldChatS2CJson)
		GamerWorldChatS2CJsonPbFiled := simplejson.New()
		GamerWorldChatS2CJsonPbFiled.Set("field", "GamerWorldChatS2C")
		GamerWorldChatS2CJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerWorldChatS2CJsonPbFiled)
	}



	GamerSplitItemS2CObj := new(pb.GamerSplitItemS2C)
	GamerSplitItemS2CType := reflect.TypeOf(*GamerSplitItemS2CObj)
	if GamerSplitItemS2CType.Kind().String() == "struct" {
		GamerSplitItemS2CJson, _ := common.Struct2Json(GamerSplitItemS2CType)
		jsonStructSlice = append(jsonStructSlice, GamerSplitItemS2CJson)
		GamerSplitItemS2CJsonPbFiled := simplejson.New()
		GamerSplitItemS2CJsonPbFiled.Set("field", "GamerSplitItemS2C")
		GamerSplitItemS2CJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerSplitItemS2CJsonPbFiled)
	}



	GamerGetRewardC2SObj := new(pb.GamerGetRewardC2S)
	GamerGetRewardC2SType := reflect.TypeOf(*GamerGetRewardC2SObj)
	if GamerGetRewardC2SType.Kind().String() == "struct" {
		GamerGetRewardC2SJson, _ := common.Struct2Json(GamerGetRewardC2SType)
		jsonStructSlice = append(jsonStructSlice, GamerGetRewardC2SJson)
		GamerGetRewardC2SJsonPbFiled := simplejson.New()
		GamerGetRewardC2SJsonPbFiled.Set("field", "GamerGetRewardC2S")
		GamerGetRewardC2SJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerGetRewardC2SJsonPbFiled)
	}



	GamerChangeNameC2SObj := new(pb.GamerChangeNameC2S)
	GamerChangeNameC2SType := reflect.TypeOf(*GamerChangeNameC2SObj)
	if GamerChangeNameC2SType.Kind().String() == "struct" {
		GamerChangeNameC2SJson, _ := common.Struct2Json(GamerChangeNameC2SType)
		jsonStructSlice = append(jsonStructSlice, GamerChangeNameC2SJson)
		GamerChangeNameC2SJsonPbFiled := simplejson.New()
		GamerChangeNameC2SJsonPbFiled.Set("field", "GamerChangeNameC2S")
		GamerChangeNameC2SJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerChangeNameC2SJsonPbFiled)
	}



	GamerFriendChatC2SObj := new(pb.GamerFriendChatC2S)
	GamerFriendChatC2SType := reflect.TypeOf(*GamerFriendChatC2SObj)
	if GamerFriendChatC2SType.Kind().String() == "struct" {
		GamerFriendChatC2SJson, _ := common.Struct2Json(GamerFriendChatC2SType)
		jsonStructSlice = append(jsonStructSlice, GamerFriendChatC2SJson)
		GamerFriendChatC2SJsonPbFiled := simplejson.New()
		GamerFriendChatC2SJsonPbFiled.Set("field", "GamerFriendChatC2S")
		GamerFriendChatC2SJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerFriendChatC2SJsonPbFiled)
	}



	GamerNotifyLeaveRoomS2CObj := new(pb.GamerNotifyLeaveRoomS2C)
	GamerNotifyLeaveRoomS2CType := reflect.TypeOf(*GamerNotifyLeaveRoomS2CObj)
	if GamerNotifyLeaveRoomS2CType.Kind().String() == "struct" {
		GamerNotifyLeaveRoomS2CJson, _ := common.Struct2Json(GamerNotifyLeaveRoomS2CType)
		jsonStructSlice = append(jsonStructSlice, GamerNotifyLeaveRoomS2CJson)
		GamerNotifyLeaveRoomS2CJsonPbFiled := simplejson.New()
		GamerNotifyLeaveRoomS2CJsonPbFiled.Set("field", "GamerNotifyLeaveRoomS2C")
		GamerNotifyLeaveRoomS2CJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerNotifyLeaveRoomS2CJsonPbFiled)
	}



	GamerNotifyMatchInfoS2CObj := new(pb.GamerNotifyMatchInfoS2C)
	GamerNotifyMatchInfoS2CType := reflect.TypeOf(*GamerNotifyMatchInfoS2CObj)
	if GamerNotifyMatchInfoS2CType.Kind().String() == "struct" {
		GamerNotifyMatchInfoS2CJson, _ := common.Struct2Json(GamerNotifyMatchInfoS2CType)
		jsonStructSlice = append(jsonStructSlice, GamerNotifyMatchInfoS2CJson)
		GamerNotifyMatchInfoS2CJsonPbFiled := simplejson.New()
		GamerNotifyMatchInfoS2CJsonPbFiled.Set("field", "GamerNotifyMatchInfoS2C")
		GamerNotifyMatchInfoS2CJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerNotifyMatchInfoS2CJsonPbFiled)
	}



	GamerDelHaveReadMailC2SObj := new(pb.GamerDelHaveReadMailC2S)
	GamerDelHaveReadMailC2SType := reflect.TypeOf(*GamerDelHaveReadMailC2SObj)
	if GamerDelHaveReadMailC2SType.Kind().String() == "struct" {
		GamerDelHaveReadMailC2SJson, _ := common.Struct2Json(GamerDelHaveReadMailC2SType)
		jsonStructSlice = append(jsonStructSlice, GamerDelHaveReadMailC2SJson)
		GamerDelHaveReadMailC2SJsonPbFiled := simplejson.New()
		GamerDelHaveReadMailC2SJsonPbFiled.Set("field", "GamerDelHaveReadMailC2S")
		GamerDelHaveReadMailC2SJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerDelHaveReadMailC2SJsonPbFiled)
	}



	GamerNotifyEnergyChangeS2CObj := new(pb.GamerNotifyEnergyChangeS2C)
	GamerNotifyEnergyChangeS2CType := reflect.TypeOf(*GamerNotifyEnergyChangeS2CObj)
	if GamerNotifyEnergyChangeS2CType.Kind().String() == "struct" {
		GamerNotifyEnergyChangeS2CJson, _ := common.Struct2Json(GamerNotifyEnergyChangeS2CType)
		jsonStructSlice = append(jsonStructSlice, GamerNotifyEnergyChangeS2CJson)
		GamerNotifyEnergyChangeS2CJsonPbFiled := simplejson.New()
		GamerNotifyEnergyChangeS2CJsonPbFiled.Set("field", "GamerNotifyEnergyChangeS2C")
		GamerNotifyEnergyChangeS2CJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerNotifyEnergyChangeS2CJsonPbFiled)
	}



	GamerOneKeyRcvMailRewardS2CObj := new(pb.GamerOneKeyRcvMailRewardS2C)
	GamerOneKeyRcvMailRewardS2CType := reflect.TypeOf(*GamerOneKeyRcvMailRewardS2CObj)
	if GamerOneKeyRcvMailRewardS2CType.Kind().String() == "struct" {
		GamerOneKeyRcvMailRewardS2CJson, _ := common.Struct2Json(GamerOneKeyRcvMailRewardS2CType)
		jsonStructSlice = append(jsonStructSlice, GamerOneKeyRcvMailRewardS2CJson)
		GamerOneKeyRcvMailRewardS2CJsonPbFiled := simplejson.New()
		GamerOneKeyRcvMailRewardS2CJsonPbFiled.Set("field", "GamerOneKeyRcvMailRewardS2C")
		GamerOneKeyRcvMailRewardS2CJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerOneKeyRcvMailRewardS2CJsonPbFiled)
	}



	GamerNotifyNewFriendS2CObj := new(pb.GamerNotifyNewFriendS2C)
	GamerNotifyNewFriendS2CType := reflect.TypeOf(*GamerNotifyNewFriendS2CObj)
	if GamerNotifyNewFriendS2CType.Kind().String() == "struct" {
		GamerNotifyNewFriendS2CJson, _ := common.Struct2Json(GamerNotifyNewFriendS2CType)
		jsonStructSlice = append(jsonStructSlice, GamerNotifyNewFriendS2CJson)
		GamerNotifyNewFriendS2CJsonPbFiled := simplejson.New()
		GamerNotifyNewFriendS2CJsonPbFiled.Set("field", "GamerNotifyNewFriendS2C")
		GamerNotifyNewFriendS2CJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerNotifyNewFriendS2CJsonPbFiled)
	}



	GamerPVPSyncS2CObj := new(pb.GamerPVPSyncS2C)
	GamerPVPSyncS2CType := reflect.TypeOf(*GamerPVPSyncS2CObj)
	if GamerPVPSyncS2CType.Kind().String() == "struct" {
		GamerPVPSyncS2CJson, _ := common.Struct2Json(GamerPVPSyncS2CType)
		jsonStructSlice = append(jsonStructSlice, GamerPVPSyncS2CJson)
		GamerPVPSyncS2CJsonPbFiled := simplejson.New()
		GamerPVPSyncS2CJsonPbFiled.Set("field", "GamerPVPSyncS2C")
		GamerPVPSyncS2CJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerPVPSyncS2CJsonPbFiled)
	}



	GamerNotifyDelFriendS2CObj := new(pb.GamerNotifyDelFriendS2C)
	GamerNotifyDelFriendS2CType := reflect.TypeOf(*GamerNotifyDelFriendS2CObj)
	if GamerNotifyDelFriendS2CType.Kind().String() == "struct" {
		GamerNotifyDelFriendS2CJson, _ := common.Struct2Json(GamerNotifyDelFriendS2CType)
		jsonStructSlice = append(jsonStructSlice, GamerNotifyDelFriendS2CJson)
		GamerNotifyDelFriendS2CJsonPbFiled := simplejson.New()
		GamerNotifyDelFriendS2CJsonPbFiled.Set("field", "GamerNotifyDelFriendS2C")
		GamerNotifyDelFriendS2CJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerNotifyDelFriendS2CJsonPbFiled)
	}



	GamerCancelMatchS2CObj := new(pb.GamerCancelMatchS2C)
	GamerCancelMatchS2CType := reflect.TypeOf(*GamerCancelMatchS2CObj)
	if GamerCancelMatchS2CType.Kind().String() == "struct" {
		GamerCancelMatchS2CJson, _ := common.Struct2Json(GamerCancelMatchS2CType)
		jsonStructSlice = append(jsonStructSlice, GamerCancelMatchS2CJson)
		GamerCancelMatchS2CJsonPbFiled := simplejson.New()
		GamerCancelMatchS2CJsonPbFiled.Set("field", "GamerCancelMatchS2C")
		GamerCancelMatchS2CJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerCancelMatchS2CJsonPbFiled)
	}



	GamerNotifyMailS2CObj := new(pb.GamerNotifyMailS2C)
	GamerNotifyMailS2CType := reflect.TypeOf(*GamerNotifyMailS2CObj)
	if GamerNotifyMailS2CType.Kind().String() == "struct" {
		GamerNotifyMailS2CJson, _ := common.Struct2Json(GamerNotifyMailS2CType)
		jsonStructSlice = append(jsonStructSlice, GamerNotifyMailS2CJson)
		GamerNotifyMailS2CJsonPbFiled := simplejson.New()
		GamerNotifyMailS2CJsonPbFiled.Set("field", "GamerNotifyMailS2C")
		GamerNotifyMailS2CJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerNotifyMailS2CJsonPbFiled)
	}



	GamerNotifyPVPSyncS2CObj := new(pb.GamerNotifyPVPSyncS2C)
	GamerNotifyPVPSyncS2CType := reflect.TypeOf(*GamerNotifyPVPSyncS2CObj)
	if GamerNotifyPVPSyncS2CType.Kind().String() == "struct" {
		GamerNotifyPVPSyncS2CJson, _ := common.Struct2Json(GamerNotifyPVPSyncS2CType)
		jsonStructSlice = append(jsonStructSlice, GamerNotifyPVPSyncS2CJson)
		GamerNotifyPVPSyncS2CJsonPbFiled := simplejson.New()
		GamerNotifyPVPSyncS2CJsonPbFiled.Set("field", "GamerNotifyPVPSyncS2C")
		GamerNotifyPVPSyncS2CJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerNotifyPVPSyncS2CJsonPbFiled)
	}



	GamerNotifyPVPStateSyncS2CObj := new(pb.GamerNotifyPVPStateSyncS2C)
	GamerNotifyPVPStateSyncS2CType := reflect.TypeOf(*GamerNotifyPVPStateSyncS2CObj)
	if GamerNotifyPVPStateSyncS2CType.Kind().String() == "struct" {
		GamerNotifyPVPStateSyncS2CJson, _ := common.Struct2Json(GamerNotifyPVPStateSyncS2CType)
		jsonStructSlice = append(jsonStructSlice, GamerNotifyPVPStateSyncS2CJson)
		GamerNotifyPVPStateSyncS2CJsonPbFiled := simplejson.New()
		GamerNotifyPVPStateSyncS2CJsonPbFiled.Set("field", "GamerNotifyPVPStateSyncS2C")
		GamerNotifyPVPStateSyncS2CJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerNotifyPVPStateSyncS2CJsonPbFiled)
	}



	GamerChoseHeroS2CObj := new(pb.GamerChoseHeroS2C)
	GamerChoseHeroS2CType := reflect.TypeOf(*GamerChoseHeroS2CObj)
	if GamerChoseHeroS2CType.Kind().String() == "struct" {
		GamerChoseHeroS2CJson, _ := common.Struct2Json(GamerChoseHeroS2CType)
		jsonStructSlice = append(jsonStructSlice, GamerChoseHeroS2CJson)
		GamerChoseHeroS2CJsonPbFiled := simplejson.New()
		GamerChoseHeroS2CJsonPbFiled.Set("field", "GamerChoseHeroS2C")
		GamerChoseHeroS2CJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerChoseHeroS2CJsonPbFiled)
	}



	GamerCheckPVPBattleS2CObj := new(pb.GamerCheckPVPBattleS2C)
	GamerCheckPVPBattleS2CType := reflect.TypeOf(*GamerCheckPVPBattleS2CObj)
	if GamerCheckPVPBattleS2CType.Kind().String() == "struct" {
		GamerCheckPVPBattleS2CJson, _ := common.Struct2Json(GamerCheckPVPBattleS2CType)
		jsonStructSlice = append(jsonStructSlice, GamerCheckPVPBattleS2CJson)
		GamerCheckPVPBattleS2CJsonPbFiled := simplejson.New()
		GamerCheckPVPBattleS2CJsonPbFiled.Set("field", "GamerCheckPVPBattleS2C")
		GamerCheckPVPBattleS2CJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerCheckPVPBattleS2CJsonPbFiled)
	}



	RoomListS2CObj := new(pb.RoomListS2C)
	RoomListS2CType := reflect.TypeOf(*RoomListS2CObj)
	if RoomListS2CType.Kind().String() == "struct" {
		RoomListS2CJson, _ := common.Struct2Json(RoomListS2CType)
		jsonStructSlice = append(jsonStructSlice, RoomListS2CJson)
		RoomListS2CJsonPbFiled := simplejson.New()
		RoomListS2CJsonPbFiled.Set("field", "RoomListS2C")
		RoomListS2CJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, RoomListS2CJsonPbFiled)
	}



	GamerChangeReadyInRoomS2CObj := new(pb.GamerChangeReadyInRoomS2C)
	GamerChangeReadyInRoomS2CType := reflect.TypeOf(*GamerChangeReadyInRoomS2CObj)
	if GamerChangeReadyInRoomS2CType.Kind().String() == "struct" {
		GamerChangeReadyInRoomS2CJson, _ := common.Struct2Json(GamerChangeReadyInRoomS2CType)
		jsonStructSlice = append(jsonStructSlice, GamerChangeReadyInRoomS2CJson)
		GamerChangeReadyInRoomS2CJsonPbFiled := simplejson.New()
		GamerChangeReadyInRoomS2CJsonPbFiled.Set("field", "GamerChangeReadyInRoomS2C")
		GamerChangeReadyInRoomS2CJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerChangeReadyInRoomS2CJsonPbFiled)
	}



	GamerUseItemC2SObj := new(pb.GamerUseItemC2S)
	GamerUseItemC2SType := reflect.TypeOf(*GamerUseItemC2SObj)
	if GamerUseItemC2SType.Kind().String() == "struct" {
		GamerUseItemC2SJson, _ := common.Struct2Json(GamerUseItemC2SType)
		jsonStructSlice = append(jsonStructSlice, GamerUseItemC2SJson)
		GamerUseItemC2SJsonPbFiled := simplejson.New()
		GamerUseItemC2SJsonPbFiled.Set("field", "GamerUseItemC2S")
		GamerUseItemC2SJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerUseItemC2SJsonPbFiled)
	}



	GamerTestChatS2CObj := new(pb.GamerTestChatS2C)
	GamerTestChatS2CType := reflect.TypeOf(*GamerTestChatS2CObj)
	if GamerTestChatS2CType.Kind().String() == "struct" {
		GamerTestChatS2CJson, _ := common.Struct2Json(GamerTestChatS2CType)
		jsonStructSlice = append(jsonStructSlice, GamerTestChatS2CJson)
		GamerTestChatS2CJsonPbFiled := simplejson.New()
		GamerTestChatS2CJsonPbFiled.Set("field", "GamerTestChatS2C")
		GamerTestChatS2CJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerTestChatS2CJsonPbFiled)
	}



	GamerPVPSyncC2SObj := new(pb.GamerPVPSyncC2S)
	GamerPVPSyncC2SType := reflect.TypeOf(*GamerPVPSyncC2SObj)
	if GamerPVPSyncC2SType.Kind().String() == "struct" {
		GamerPVPSyncC2SJson, _ := common.Struct2Json(GamerPVPSyncC2SType)
		jsonStructSlice = append(jsonStructSlice, GamerPVPSyncC2SJson)
		GamerPVPSyncC2SJsonPbFiled := simplejson.New()
		GamerPVPSyncC2SJsonPbFiled.Set("field", "GamerPVPSyncC2S")
		GamerPVPSyncC2SJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerPVPSyncC2SJsonPbFiled)
	}



	GamerAddRoomS2CObj := new(pb.GamerAddRoomS2C)
	GamerAddRoomS2CType := reflect.TypeOf(*GamerAddRoomS2CObj)
	if GamerAddRoomS2CType.Kind().String() == "struct" {
		GamerAddRoomS2CJson, _ := common.Struct2Json(GamerAddRoomS2CType)
		jsonStructSlice = append(jsonStructSlice, GamerAddRoomS2CJson)
		GamerAddRoomS2CJsonPbFiled := simplejson.New()
		GamerAddRoomS2CJsonPbFiled.Set("field", "GamerAddRoomS2C")
		GamerAddRoomS2CJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerAddRoomS2CJsonPbFiled)
	}



	GamerCompoundItemC2SObj := new(pb.GamerCompoundItemC2S)
	GamerCompoundItemC2SType := reflect.TypeOf(*GamerCompoundItemC2SObj)
	if GamerCompoundItemC2SType.Kind().String() == "struct" {
		GamerCompoundItemC2SJson, _ := common.Struct2Json(GamerCompoundItemC2SType)
		jsonStructSlice = append(jsonStructSlice, GamerCompoundItemC2SJson)
		GamerCompoundItemC2SJsonPbFiled := simplejson.New()
		GamerCompoundItemC2SJsonPbFiled.Set("field", "GamerCompoundItemC2S")
		GamerCompoundItemC2SJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerCompoundItemC2SJsonPbFiled)
	}



	GamerCompoundItemS2CObj := new(pb.GamerCompoundItemS2C)
	GamerCompoundItemS2CType := reflect.TypeOf(*GamerCompoundItemS2CObj)
	if GamerCompoundItemS2CType.Kind().String() == "struct" {
		GamerCompoundItemS2CJson, _ := common.Struct2Json(GamerCompoundItemS2CType)
		jsonStructSlice = append(jsonStructSlice, GamerCompoundItemS2CJson)
		GamerCompoundItemS2CJsonPbFiled := simplejson.New()
		GamerCompoundItemS2CJsonPbFiled.Set("field", "GamerCompoundItemS2C")
		GamerCompoundItemS2CJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerCompoundItemS2CJsonPbFiled)
	}



	GamerProcessFriendReqS2CObj := new(pb.GamerProcessFriendReqS2C)
	GamerProcessFriendReqS2CType := reflect.TypeOf(*GamerProcessFriendReqS2CObj)
	if GamerProcessFriendReqS2CType.Kind().String() == "struct" {
		GamerProcessFriendReqS2CJson, _ := common.Struct2Json(GamerProcessFriendReqS2CType)
		jsonStructSlice = append(jsonStructSlice, GamerProcessFriendReqS2CJson)
		GamerProcessFriendReqS2CJsonPbFiled := simplejson.New()
		GamerProcessFriendReqS2CJsonPbFiled.Set("field", "GamerProcessFriendReqS2C")
		GamerProcessFriendReqS2CJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerProcessFriendReqS2CJsonPbFiled)
	}



	GamerGetMailC2SObj := new(pb.GamerGetMailC2S)
	GamerGetMailC2SType := reflect.TypeOf(*GamerGetMailC2SObj)
	if GamerGetMailC2SType.Kind().String() == "struct" {
		GamerGetMailC2SJson, _ := common.Struct2Json(GamerGetMailC2SType)
		jsonStructSlice = append(jsonStructSlice, GamerGetMailC2SJson)
		GamerGetMailC2SJsonPbFiled := simplejson.New()
		GamerGetMailC2SJsonPbFiled.Set("field", "GamerGetMailC2S")
		GamerGetMailC2SJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerGetMailC2SJsonPbFiled)
	}



	GamerNotifyActivitiesS2CObj := new(pb.GamerNotifyActivitiesS2C)
	GamerNotifyActivitiesS2CType := reflect.TypeOf(*GamerNotifyActivitiesS2CObj)
	if GamerNotifyActivitiesS2CType.Kind().String() == "struct" {
		GamerNotifyActivitiesS2CJson, _ := common.Struct2Json(GamerNotifyActivitiesS2CType)
		jsonStructSlice = append(jsonStructSlice, GamerNotifyActivitiesS2CJson)
		GamerNotifyActivitiesS2CJsonPbFiled := simplejson.New()
		GamerNotifyActivitiesS2CJsonPbFiled.Set("field", "GamerNotifyActivitiesS2C")
		GamerNotifyActivitiesS2CJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerNotifyActivitiesS2CJsonPbFiled)
	}



	GamerSubChatChannelC2SObj := new(pb.GamerSubChatChannelC2S)
	GamerSubChatChannelC2SType := reflect.TypeOf(*GamerSubChatChannelC2SObj)
	if GamerSubChatChannelC2SType.Kind().String() == "struct" {
		GamerSubChatChannelC2SJson, _ := common.Struct2Json(GamerSubChatChannelC2SType)
		jsonStructSlice = append(jsonStructSlice, GamerSubChatChannelC2SJson)
		GamerSubChatChannelC2SJsonPbFiled := simplejson.New()
		GamerSubChatChannelC2SJsonPbFiled.Set("field", "GamerSubChatChannelC2S")
		GamerSubChatChannelC2SJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerSubChatChannelC2SJsonPbFiled)
	}



	GamerNotifyNewMailS2CObj := new(pb.GamerNotifyNewMailS2C)
	GamerNotifyNewMailS2CType := reflect.TypeOf(*GamerNotifyNewMailS2CObj)
	if GamerNotifyNewMailS2CType.Kind().String() == "struct" {
		GamerNotifyNewMailS2CJson, _ := common.Struct2Json(GamerNotifyNewMailS2CType)
		jsonStructSlice = append(jsonStructSlice, GamerNotifyNewMailS2CJson)
		GamerNotifyNewMailS2CJsonPbFiled := simplejson.New()
		GamerNotifyNewMailS2CJsonPbFiled.Set("field", "GamerNotifyNewMailS2C")
		GamerNotifyNewMailS2CJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerNotifyNewMailS2CJsonPbFiled)
	}



	GamerNewFriendReqS2CObj := new(pb.GamerNewFriendReqS2C)
	GamerNewFriendReqS2CType := reflect.TypeOf(*GamerNewFriendReqS2CObj)
	if GamerNewFriendReqS2CType.Kind().String() == "struct" {
		GamerNewFriendReqS2CJson, _ := common.Struct2Json(GamerNewFriendReqS2CType)
		jsonStructSlice = append(jsonStructSlice, GamerNewFriendReqS2CJson)
		GamerNewFriendReqS2CJsonPbFiled := simplejson.New()
		GamerNewFriendReqS2CJsonPbFiled.Set("field", "GamerNewFriendReqS2C")
		GamerNewFriendReqS2CJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerNewFriendReqS2CJsonPbFiled)
	}



	GamerPVPStateSyncS2CObj := new(pb.GamerPVPStateSyncS2C)
	GamerPVPStateSyncS2CType := reflect.TypeOf(*GamerPVPStateSyncS2CObj)
	if GamerPVPStateSyncS2CType.Kind().String() == "struct" {
		GamerPVPStateSyncS2CJson, _ := common.Struct2Json(GamerPVPStateSyncS2CType)
		jsonStructSlice = append(jsonStructSlice, GamerPVPStateSyncS2CJson)
		GamerPVPStateSyncS2CJsonPbFiled := simplejson.New()
		GamerPVPStateSyncS2CJsonPbFiled.Set("field", "GamerPVPStateSyncS2C")
		GamerPVPStateSyncS2CJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerPVPStateSyncS2CJsonPbFiled)
	}



	GamerNotifyNewFriendReqS2CObj := new(pb.GamerNotifyNewFriendReqS2C)
	GamerNotifyNewFriendReqS2CType := reflect.TypeOf(*GamerNotifyNewFriendReqS2CObj)
	if GamerNotifyNewFriendReqS2CType.Kind().String() == "struct" {
		GamerNotifyNewFriendReqS2CJson, _ := common.Struct2Json(GamerNotifyNewFriendReqS2CType)
		jsonStructSlice = append(jsonStructSlice, GamerNotifyNewFriendReqS2CJson)
		GamerNotifyNewFriendReqS2CJsonPbFiled := simplejson.New()
		GamerNotifyNewFriendReqS2CJsonPbFiled.Set("field", "GamerNotifyNewFriendReqS2C")
		GamerNotifyNewFriendReqS2CJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerNotifyNewFriendReqS2CJsonPbFiled)
	}



	GamerCancelMatchC2SObj := new(pb.GamerCancelMatchC2S)
	GamerCancelMatchC2SType := reflect.TypeOf(*GamerCancelMatchC2SObj)
	if GamerCancelMatchC2SType.Kind().String() == "struct" {
		GamerCancelMatchC2SJson, _ := common.Struct2Json(GamerCancelMatchC2SType)
		jsonStructSlice = append(jsonStructSlice, GamerCancelMatchC2SJson)
		GamerCancelMatchC2SJsonPbFiled := simplejson.New()
		GamerCancelMatchC2SJsonPbFiled.Set("field", "GamerCancelMatchC2S")
		GamerCancelMatchC2SJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerCancelMatchC2SJsonPbFiled)
	}



	GMC2SObj := new(pb.GMC2S)
	GMC2SType := reflect.TypeOf(*GMC2SObj)
	if GMC2SType.Kind().String() == "struct" {
		GMC2SJson, _ := common.Struct2Json(GMC2SType)
		jsonStructSlice = append(jsonStructSlice, GMC2SJson)
		GMC2SJsonPbFiled := simplejson.New()
		GMC2SJsonPbFiled.Set("field", "GMC2S")
		GMC2SJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GMC2SJsonPbFiled)
	}



	GamerMatchC2SObj := new(pb.GamerMatchC2S)
	GamerMatchC2SType := reflect.TypeOf(*GamerMatchC2SObj)
	if GamerMatchC2SType.Kind().String() == "struct" {
		GamerMatchC2SJson, _ := common.Struct2Json(GamerMatchC2SType)
		jsonStructSlice = append(jsonStructSlice, GamerMatchC2SJson)
		GamerMatchC2SJsonPbFiled := simplejson.New()
		GamerMatchC2SJsonPbFiled.Set("field", "GamerMatchC2S")
		GamerMatchC2SJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerMatchC2SJsonPbFiled)
	}



	GamerUseIconC2SObj := new(pb.GamerUseIconC2S)
	GamerUseIconC2SType := reflect.TypeOf(*GamerUseIconC2SObj)
	if GamerUseIconC2SType.Kind().String() == "struct" {
		GamerUseIconC2SJson, _ := common.Struct2Json(GamerUseIconC2SType)
		jsonStructSlice = append(jsonStructSlice, GamerUseIconC2SJson)
		GamerUseIconC2SJsonPbFiled := simplejson.New()
		GamerUseIconC2SJsonPbFiled.Set("field", "GamerUseIconC2S")
		GamerUseIconC2SJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerUseIconC2SJsonPbFiled)
	}



	GamerDelMailS2CObj := new(pb.GamerDelMailS2C)
	GamerDelMailS2CType := reflect.TypeOf(*GamerDelMailS2CObj)
	if GamerDelMailS2CType.Kind().String() == "struct" {
		GamerDelMailS2CJson, _ := common.Struct2Json(GamerDelMailS2CType)
		jsonStructSlice = append(jsonStructSlice, GamerDelMailS2CJson)
		GamerDelMailS2CJsonPbFiled := simplejson.New()
		GamerDelMailS2CJsonPbFiled.Set("field", "GamerDelMailS2C")
		GamerDelMailS2CJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerDelMailS2CJsonPbFiled)
	}



	GamerChangeMailStateC2SObj := new(pb.GamerChangeMailStateC2S)
	GamerChangeMailStateC2SType := reflect.TypeOf(*GamerChangeMailStateC2SObj)
	if GamerChangeMailStateC2SType.Kind().String() == "struct" {
		GamerChangeMailStateC2SJson, _ := common.Struct2Json(GamerChangeMailStateC2SType)
		jsonStructSlice = append(jsonStructSlice, GamerChangeMailStateC2SJson)
		GamerChangeMailStateC2SJsonPbFiled := simplejson.New()
		GamerChangeMailStateC2SJsonPbFiled.Set("field", "GamerChangeMailStateC2S")
		GamerChangeMailStateC2SJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerChangeMailStateC2SJsonPbFiled)
	}



	GamerBuyItemS2CObj := new(pb.GamerBuyItemS2C)
	GamerBuyItemS2CType := reflect.TypeOf(*GamerBuyItemS2CObj)
	if GamerBuyItemS2CType.Kind().String() == "struct" {
		GamerBuyItemS2CJson, _ := common.Struct2Json(GamerBuyItemS2CType)
		jsonStructSlice = append(jsonStructSlice, GamerBuyItemS2CJson)
		GamerBuyItemS2CJsonPbFiled := simplejson.New()
		GamerBuyItemS2CJsonPbFiled.Set("field", "GamerBuyItemS2C")
		GamerBuyItemS2CJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerBuyItemS2CJsonPbFiled)
	}



	GamerOneKeyRcvMailRewardC2SObj := new(pb.GamerOneKeyRcvMailRewardC2S)
	GamerOneKeyRcvMailRewardC2SType := reflect.TypeOf(*GamerOneKeyRcvMailRewardC2SObj)
	if GamerOneKeyRcvMailRewardC2SType.Kind().String() == "struct" {
		GamerOneKeyRcvMailRewardC2SJson, _ := common.Struct2Json(GamerOneKeyRcvMailRewardC2SType)
		jsonStructSlice = append(jsonStructSlice, GamerOneKeyRcvMailRewardC2SJson)
		GamerOneKeyRcvMailRewardC2SJsonPbFiled := simplejson.New()
		GamerOneKeyRcvMailRewardC2SJsonPbFiled.Set("field", "GamerOneKeyRcvMailRewardC2S")
		GamerOneKeyRcvMailRewardC2SJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerOneKeyRcvMailRewardC2SJsonPbFiled)
	}



	GamerUseIconS2CObj := new(pb.GamerUseIconS2C)
	GamerUseIconS2CType := reflect.TypeOf(*GamerUseIconS2CObj)
	if GamerUseIconS2CType.Kind().String() == "struct" {
		GamerUseIconS2CJson, _ := common.Struct2Json(GamerUseIconS2CType)
		jsonStructSlice = append(jsonStructSlice, GamerUseIconS2CJson)
		GamerUseIconS2CJsonPbFiled := simplejson.New()
		GamerUseIconS2CJsonPbFiled.Set("field", "GamerUseIconS2C")
		GamerUseIconS2CJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerUseIconS2CJsonPbFiled)
	}



	GamerNewHeroC2SObj := new(pb.GamerNewHeroC2S)
	GamerNewHeroC2SType := reflect.TypeOf(*GamerNewHeroC2SObj)
	if GamerNewHeroC2SType.Kind().String() == "struct" {
		GamerNewHeroC2SJson, _ := common.Struct2Json(GamerNewHeroC2SType)
		jsonStructSlice = append(jsonStructSlice, GamerNewHeroC2SJson)
		GamerNewHeroC2SJsonPbFiled := simplejson.New()
		GamerNewHeroC2SJsonPbFiled.Set("field", "GamerNewHeroC2S")
		GamerNewHeroC2SJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerNewHeroC2SJsonPbFiled)
	}



	GamerNewRoomC2SObj := new(pb.GamerNewRoomC2S)
	GamerNewRoomC2SType := reflect.TypeOf(*GamerNewRoomC2SObj)
	if GamerNewRoomC2SType.Kind().String() == "struct" {
		GamerNewRoomC2SJson, _ := common.Struct2Json(GamerNewRoomC2SType)
		jsonStructSlice = append(jsonStructSlice, GamerNewRoomC2SJson)
		GamerNewRoomC2SJsonPbFiled := simplejson.New()
		GamerNewRoomC2SJsonPbFiled.Set("field", "GamerNewRoomC2S")
		GamerNewRoomC2SJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerNewRoomC2SJsonPbFiled)
	}



	GamerGetRewardS2CObj := new(pb.GamerGetRewardS2C)
	GamerGetRewardS2CType := reflect.TypeOf(*GamerGetRewardS2CObj)
	if GamerGetRewardS2CType.Kind().String() == "struct" {
		GamerGetRewardS2CJson, _ := common.Struct2Json(GamerGetRewardS2CType)
		jsonStructSlice = append(jsonStructSlice, GamerGetRewardS2CJson)
		GamerGetRewardS2CJsonPbFiled := simplejson.New()
		GamerGetRewardS2CJsonPbFiled.Set("field", "GamerGetRewardS2C")
		GamerGetRewardS2CJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerGetRewardS2CJsonPbFiled)
	}



	GamerWorldChatC2SObj := new(pb.GamerWorldChatC2S)
	GamerWorldChatC2SType := reflect.TypeOf(*GamerWorldChatC2SObj)
	if GamerWorldChatC2SType.Kind().String() == "struct" {
		GamerWorldChatC2SJson, _ := common.Struct2Json(GamerWorldChatC2SType)
		jsonStructSlice = append(jsonStructSlice, GamerWorldChatC2SJson)
		GamerWorldChatC2SJsonPbFiled := simplejson.New()
		GamerWorldChatC2SJsonPbFiled.Set("field", "GamerWorldChatC2S")
		GamerWorldChatC2SJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerWorldChatC2SJsonPbFiled)
	}



	GamerDelHaveReadMailS2CObj := new(pb.GamerDelHaveReadMailS2C)
	GamerDelHaveReadMailS2CType := reflect.TypeOf(*GamerDelHaveReadMailS2CObj)
	if GamerDelHaveReadMailS2CType.Kind().String() == "struct" {
		GamerDelHaveReadMailS2CJson, _ := common.Struct2Json(GamerDelHaveReadMailS2CType)
		jsonStructSlice = append(jsonStructSlice, GamerDelHaveReadMailS2CJson)
		GamerDelHaveReadMailS2CJsonPbFiled := simplejson.New()
		GamerDelHaveReadMailS2CJsonPbFiled.Set("field", "GamerDelHaveReadMailS2C")
		GamerDelHaveReadMailS2CJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerDelHaveReadMailS2CJsonPbFiled)
	}



	GamerMatchS2CObj := new(pb.GamerMatchS2C)
	GamerMatchS2CType := reflect.TypeOf(*GamerMatchS2CObj)
	if GamerMatchS2CType.Kind().String() == "struct" {
		GamerMatchS2CJson, _ := common.Struct2Json(GamerMatchS2CType)
		jsonStructSlice = append(jsonStructSlice, GamerMatchS2CJson)
		GamerMatchS2CJsonPbFiled := simplejson.New()
		GamerMatchS2CJsonPbFiled.Set("field", "GamerMatchS2C")
		GamerMatchS2CJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerMatchS2CJsonPbFiled)
	}



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



	GamerDelMailC2SObj := new(pb.GamerDelMailC2S)
	GamerDelMailC2SType := reflect.TypeOf(*GamerDelMailC2SObj)
	if GamerDelMailC2SType.Kind().String() == "struct" {
		GamerDelMailC2SJson, _ := common.Struct2Json(GamerDelMailC2SType)
		jsonStructSlice = append(jsonStructSlice, GamerDelMailC2SJson)
		GamerDelMailC2SJsonPbFiled := simplejson.New()
		GamerDelMailC2SJsonPbFiled.Set("field", "GamerDelMailC2S")
		GamerDelMailC2SJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerDelMailC2SJsonPbFiled)
	}



	GamerNotifyItemS2CObj := new(pb.GamerNotifyItemS2C)
	GamerNotifyItemS2CType := reflect.TypeOf(*GamerNotifyItemS2CObj)
	if GamerNotifyItemS2CType.Kind().String() == "struct" {
		GamerNotifyItemS2CJson, _ := common.Struct2Json(GamerNotifyItemS2CType)
		jsonStructSlice = append(jsonStructSlice, GamerNotifyItemS2CJson)
		GamerNotifyItemS2CJsonPbFiled := simplejson.New()
		GamerNotifyItemS2CJsonPbFiled.Set("field", "GamerNotifyItemS2C")
		GamerNotifyItemS2CJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerNotifyItemS2CJsonPbFiled)
	}



	GamerExtractGashaponC2SObj := new(pb.GamerExtractGashaponC2S)
	GamerExtractGashaponC2SType := reflect.TypeOf(*GamerExtractGashaponC2SObj)
	if GamerExtractGashaponC2SType.Kind().String() == "struct" {
		GamerExtractGashaponC2SJson, _ := common.Struct2Json(GamerExtractGashaponC2SType)
		jsonStructSlice = append(jsonStructSlice, GamerExtractGashaponC2SJson)
		GamerExtractGashaponC2SJsonPbFiled := simplejson.New()
		GamerExtractGashaponC2SJsonPbFiled.Set("field", "GamerExtractGashaponC2S")
		GamerExtractGashaponC2SJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerExtractGashaponC2SJsonPbFiled)
	}



	GamerGetBackpackC2SObj := new(pb.GamerGetBackpackC2S)
	GamerGetBackpackC2SType := reflect.TypeOf(*GamerGetBackpackC2SObj)
	if GamerGetBackpackC2SType.Kind().String() == "struct" {
		GamerGetBackpackC2SJson, _ := common.Struct2Json(GamerGetBackpackC2SType)
		jsonStructSlice = append(jsonStructSlice, GamerGetBackpackC2SJson)
		GamerGetBackpackC2SJsonPbFiled := simplejson.New()
		GamerGetBackpackC2SJsonPbFiled.Set("field", "GamerGetBackpackC2S")
		GamerGetBackpackC2SJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerGetBackpackC2SJsonPbFiled)
	}



	GamerPVPStateSyncC2SObj := new(pb.GamerPVPStateSyncC2S)
	GamerPVPStateSyncC2SType := reflect.TypeOf(*GamerPVPStateSyncC2SObj)
	if GamerPVPStateSyncC2SType.Kind().String() == "struct" {
		GamerPVPStateSyncC2SJson, _ := common.Struct2Json(GamerPVPStateSyncC2SType)
		jsonStructSlice = append(jsonStructSlice, GamerPVPStateSyncC2SJson)
		GamerPVPStateSyncC2SJsonPbFiled := simplejson.New()
		GamerPVPStateSyncC2SJsonPbFiled.Set("field", "GamerPVPStateSyncC2S")
		GamerPVPStateSyncC2SJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerPVPStateSyncC2SJsonPbFiled)
	}



	GamerLeaveRoomC2SObj := new(pb.GamerLeaveRoomC2S)
	GamerLeaveRoomC2SType := reflect.TypeOf(*GamerLeaveRoomC2SObj)
	if GamerLeaveRoomC2SType.Kind().String() == "struct" {
		GamerLeaveRoomC2SJson, _ := common.Struct2Json(GamerLeaveRoomC2SType)
		jsonStructSlice = append(jsonStructSlice, GamerLeaveRoomC2SJson)
		GamerLeaveRoomC2SJsonPbFiled := simplejson.New()
		GamerLeaveRoomC2SJsonPbFiled.Set("field", "GamerLeaveRoomC2S")
		GamerLeaveRoomC2SJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerLeaveRoomC2SJsonPbFiled)
	}



	GamerGetMailS2CObj := new(pb.GamerGetMailS2C)
	GamerGetMailS2CType := reflect.TypeOf(*GamerGetMailS2CObj)
	if GamerGetMailS2CType.Kind().String() == "struct" {
		GamerGetMailS2CJson, _ := common.Struct2Json(GamerGetMailS2CType)
		jsonStructSlice = append(jsonStructSlice, GamerGetMailS2CJson)
		GamerGetMailS2CJsonPbFiled := simplejson.New()
		GamerGetMailS2CJsonPbFiled.Set("field", "GamerGetMailS2C")
		GamerGetMailS2CJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerGetMailS2CJsonPbFiled)
	}



	GamerLoginGetDataC2SObj := new(pb.GamerLoginGetDataC2S)
	GamerLoginGetDataC2SType := reflect.TypeOf(*GamerLoginGetDataC2SObj)
	if GamerLoginGetDataC2SType.Kind().String() == "struct" {
		GamerLoginGetDataC2SJson, _ := common.Struct2Json(GamerLoginGetDataC2SType)
		jsonStructSlice = append(jsonStructSlice, GamerLoginGetDataC2SJson)
		GamerLoginGetDataC2SJsonPbFiled := simplejson.New()
		GamerLoginGetDataC2SJsonPbFiled.Set("field", "GamerLoginGetDataC2S")
		GamerLoginGetDataC2SJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerLoginGetDataC2SJsonPbFiled)
	}



	GamerNotifyNewChatS2CObj := new(pb.GamerNotifyNewChatS2C)
	GamerNotifyNewChatS2CType := reflect.TypeOf(*GamerNotifyNewChatS2CObj)
	if GamerNotifyNewChatS2CType.Kind().String() == "struct" {
		GamerNotifyNewChatS2CJson, _ := common.Struct2Json(GamerNotifyNewChatS2CType)
		jsonStructSlice = append(jsonStructSlice, GamerNotifyNewChatS2CJson)
		GamerNotifyNewChatS2CJsonPbFiled := simplejson.New()
		GamerNotifyNewChatS2CJsonPbFiled.Set("field", "GamerNotifyNewChatS2C")
		GamerNotifyNewChatS2CJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerNotifyNewChatS2CJsonPbFiled)
	}



	GamerGetRealTimeRankS2CObj := new(pb.GamerGetRealTimeRankS2C)
	GamerGetRealTimeRankS2CType := reflect.TypeOf(*GamerGetRealTimeRankS2CObj)
	if GamerGetRealTimeRankS2CType.Kind().String() == "struct" {
		GamerGetRealTimeRankS2CJson, _ := common.Struct2Json(GamerGetRealTimeRankS2CType)
		jsonStructSlice = append(jsonStructSlice, GamerGetRealTimeRankS2CJson)
		GamerGetRealTimeRankS2CJsonPbFiled := simplejson.New()
		GamerGetRealTimeRankS2CJsonPbFiled.Set("field", "GamerGetRealTimeRankS2C")
		GamerGetRealTimeRankS2CJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerGetRealTimeRankS2CJsonPbFiled)
	}



	GamerNotifyGashaponInfoS2CObj := new(pb.GamerNotifyGashaponInfoS2C)
	GamerNotifyGashaponInfoS2CType := reflect.TypeOf(*GamerNotifyGashaponInfoS2CObj)
	if GamerNotifyGashaponInfoS2CType.Kind().String() == "struct" {
		GamerNotifyGashaponInfoS2CJson, _ := common.Struct2Json(GamerNotifyGashaponInfoS2CType)
		jsonStructSlice = append(jsonStructSlice, GamerNotifyGashaponInfoS2CJson)
		GamerNotifyGashaponInfoS2CJsonPbFiled := simplejson.New()
		GamerNotifyGashaponInfoS2CJsonPbFiled.Set("field", "GamerNotifyGashaponInfoS2C")
		GamerNotifyGashaponInfoS2CJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerNotifyGashaponInfoS2CJsonPbFiled)
	}



	GamerNotifyDiamonChangeS2CObj := new(pb.GamerNotifyDiamonChangeS2C)
	GamerNotifyDiamonChangeS2CType := reflect.TypeOf(*GamerNotifyDiamonChangeS2CObj)
	if GamerNotifyDiamonChangeS2CType.Kind().String() == "struct" {
		GamerNotifyDiamonChangeS2CJson, _ := common.Struct2Json(GamerNotifyDiamonChangeS2CType)
		jsonStructSlice = append(jsonStructSlice, GamerNotifyDiamonChangeS2CJson)
		GamerNotifyDiamonChangeS2CJsonPbFiled := simplejson.New()
		GamerNotifyDiamonChangeS2CJsonPbFiled.Set("field", "GamerNotifyDiamonChangeS2C")
		GamerNotifyDiamonChangeS2CJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerNotifyDiamonChangeS2CJsonPbFiled)
	}



	GamerNotifyItemChangeS2CObj := new(pb.GamerNotifyItemChangeS2C)
	GamerNotifyItemChangeS2CType := reflect.TypeOf(*GamerNotifyItemChangeS2CObj)
	if GamerNotifyItemChangeS2CType.Kind().String() == "struct" {
		GamerNotifyItemChangeS2CJson, _ := common.Struct2Json(GamerNotifyItemChangeS2CType)
		jsonStructSlice = append(jsonStructSlice, GamerNotifyItemChangeS2CJson)
		GamerNotifyItemChangeS2CJsonPbFiled := simplejson.New()
		GamerNotifyItemChangeS2CJsonPbFiled.Set("field", "GamerNotifyItemChangeS2C")
		GamerNotifyItemChangeS2CJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerNotifyItemChangeS2CJsonPbFiled)
	}



	GamerNotifyLoginOtherS2CObj := new(pb.GamerNotifyLoginOtherS2C)
	GamerNotifyLoginOtherS2CType := reflect.TypeOf(*GamerNotifyLoginOtherS2CObj)
	if GamerNotifyLoginOtherS2CType.Kind().String() == "struct" {
		GamerNotifyLoginOtherS2CJson, _ := common.Struct2Json(GamerNotifyLoginOtherS2CType)
		jsonStructSlice = append(jsonStructSlice, GamerNotifyLoginOtherS2CJson)
		GamerNotifyLoginOtherS2CJsonPbFiled := simplejson.New()
		GamerNotifyLoginOtherS2CJsonPbFiled.Set("field", "GamerNotifyLoginOtherS2C")
		GamerNotifyLoginOtherS2CJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerNotifyLoginOtherS2CJsonPbFiled)
	}



	GamerUseItemS2CObj := new(pb.GamerUseItemS2C)
	GamerUseItemS2CType := reflect.TypeOf(*GamerUseItemS2CObj)
	if GamerUseItemS2CType.Kind().String() == "struct" {
		GamerUseItemS2CJson, _ := common.Struct2Json(GamerUseItemS2CType)
		jsonStructSlice = append(jsonStructSlice, GamerUseItemS2CJson)
		GamerUseItemS2CJsonPbFiled := simplejson.New()
		GamerUseItemS2CJsonPbFiled.Set("field", "GamerUseItemS2C")
		GamerUseItemS2CJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerUseItemS2CJsonPbFiled)
	}



	GetActivityRewardC2SObj := new(pb.GetActivityRewardC2S)
	GetActivityRewardC2SType := reflect.TypeOf(*GetActivityRewardC2SObj)
	if GetActivityRewardC2SType.Kind().String() == "struct" {
		GetActivityRewardC2SJson, _ := common.Struct2Json(GetActivityRewardC2SType)
		jsonStructSlice = append(jsonStructSlice, GetActivityRewardC2SJson)
		GetActivityRewardC2SJsonPbFiled := simplejson.New()
		GetActivityRewardC2SJsonPbFiled.Set("field", "GetActivityRewardC2S")
		GetActivityRewardC2SJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GetActivityRewardC2SJsonPbFiled)
	}



	GamerSplitItemC2SObj := new(pb.GamerSplitItemC2S)
	GamerSplitItemC2SType := reflect.TypeOf(*GamerSplitItemC2SObj)
	if GamerSplitItemC2SType.Kind().String() == "struct" {
		GamerSplitItemC2SJson, _ := common.Struct2Json(GamerSplitItemC2SType)
		jsonStructSlice = append(jsonStructSlice, GamerSplitItemC2SJson)
		GamerSplitItemC2SJsonPbFiled := simplejson.New()
		GamerSplitItemC2SJsonPbFiled.Set("field", "GamerSplitItemC2S")
		GamerSplitItemC2SJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerSplitItemC2SJsonPbFiled)
	}



	GamerLoginS2CObj := new(pb.GamerLoginS2C)
	GamerLoginS2CType := reflect.TypeOf(*GamerLoginS2CObj)
	if GamerLoginS2CType.Kind().String() == "struct" {
		GamerLoginS2CJson, _ := common.Struct2Json(GamerLoginS2CType)
		jsonStructSlice = append(jsonStructSlice, GamerLoginS2CJson)
		GamerLoginS2CJsonPbFiled := simplejson.New()
		GamerLoginS2CJsonPbFiled.Set("field", "GamerLoginS2C")
		GamerLoginS2CJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerLoginS2CJsonPbFiled)
	}



	GamerClubRequestS2CObj := new(pb.GamerClubRequestS2C)
	GamerClubRequestS2CType := reflect.TypeOf(*GamerClubRequestS2CObj)
	if GamerClubRequestS2CType.Kind().String() == "struct" {
		GamerClubRequestS2CJson, _ := common.Struct2Json(GamerClubRequestS2CType)
		jsonStructSlice = append(jsonStructSlice, GamerClubRequestS2CJson)
		GamerClubRequestS2CJsonPbFiled := simplejson.New()
		GamerClubRequestS2CJsonPbFiled.Set("field", "GamerClubRequestS2C")
		GamerClubRequestS2CJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerClubRequestS2CJsonPbFiled)
	}



	GamerChangeMailStateS2CObj := new(pb.GamerChangeMailStateS2C)
	GamerChangeMailStateS2CType := reflect.TypeOf(*GamerChangeMailStateS2CObj)
	if GamerChangeMailStateS2CType.Kind().String() == "struct" {
		GamerChangeMailStateS2CJson, _ := common.Struct2Json(GamerChangeMailStateS2CType)
		jsonStructSlice = append(jsonStructSlice, GamerChangeMailStateS2CJson)
		GamerChangeMailStateS2CJsonPbFiled := simplejson.New()
		GamerChangeMailStateS2CJsonPbFiled.Set("field", "GamerChangeMailStateS2C")
		GamerChangeMailStateS2CJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerChangeMailStateS2CJsonPbFiled)
	}



	GamerGetAllowanceC2SObj := new(pb.GamerGetAllowanceC2S)
	GamerGetAllowanceC2SType := reflect.TypeOf(*GamerGetAllowanceC2SObj)
	if GamerGetAllowanceC2SType.Kind().String() == "struct" {
		GamerGetAllowanceC2SJson, _ := common.Struct2Json(GamerGetAllowanceC2SType)
		jsonStructSlice = append(jsonStructSlice, GamerGetAllowanceC2SJson)
		GamerGetAllowanceC2SJsonPbFiled := simplejson.New()
		GamerGetAllowanceC2SJsonPbFiled.Set("field", "GamerGetAllowanceC2S")
		GamerGetAllowanceC2SJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerGetAllowanceC2SJsonPbFiled)
	}



	GamerDelFriendS2CObj := new(pb.GamerDelFriendS2C)
	GamerDelFriendS2CType := reflect.TypeOf(*GamerDelFriendS2CObj)
	if GamerDelFriendS2CType.Kind().String() == "struct" {
		GamerDelFriendS2CJson, _ := common.Struct2Json(GamerDelFriendS2CType)
		jsonStructSlice = append(jsonStructSlice, GamerDelFriendS2CJson)
		GamerDelFriendS2CJsonPbFiled := simplejson.New()
		GamerDelFriendS2CJsonPbFiled.Set("field", "GamerDelFriendS2C")
		GamerDelFriendS2CJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerDelFriendS2CJsonPbFiled)
	}



	GamerProcessFriendReqC2SObj := new(pb.GamerProcessFriendReqC2S)
	GamerProcessFriendReqC2SType := reflect.TypeOf(*GamerProcessFriendReqC2SObj)
	if GamerProcessFriendReqC2SType.Kind().String() == "struct" {
		GamerProcessFriendReqC2SJson, _ := common.Struct2Json(GamerProcessFriendReqC2SType)
		jsonStructSlice = append(jsonStructSlice, GamerProcessFriendReqC2SJson)
		GamerProcessFriendReqC2SJsonPbFiled := simplejson.New()
		GamerProcessFriendReqC2SJsonPbFiled.Set("field", "GamerProcessFriendReqC2S")
		GamerProcessFriendReqC2SJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerProcessFriendReqC2SJsonPbFiled)
	}



	GamerNotifyAddRoomS2CObj := new(pb.GamerNotifyAddRoomS2C)
	GamerNotifyAddRoomS2CType := reflect.TypeOf(*GamerNotifyAddRoomS2CObj)
	if GamerNotifyAddRoomS2CType.Kind().String() == "struct" {
		GamerNotifyAddRoomS2CJson, _ := common.Struct2Json(GamerNotifyAddRoomS2CType)
		jsonStructSlice = append(jsonStructSlice, GamerNotifyAddRoomS2CJson)
		GamerNotifyAddRoomS2CJsonPbFiled := simplejson.New()
		GamerNotifyAddRoomS2CJsonPbFiled.Set("field", "GamerNotifyAddRoomS2C")
		GamerNotifyAddRoomS2CJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerNotifyAddRoomS2CJsonPbFiled)
	}



	GMS2CObj := new(pb.GMS2C)
	GMS2CType := reflect.TypeOf(*GMS2CObj)
	if GMS2CType.Kind().String() == "struct" {
		GMS2CJson, _ := common.Struct2Json(GMS2CType)
		jsonStructSlice = append(jsonStructSlice, GMS2CJson)
		GMS2CJsonPbFiled := simplejson.New()
		GMS2CJsonPbFiled.Set("field", "GMS2C")
		GMS2CJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GMS2CJsonPbFiled)
	}



	GamerLoginGetDataS2CObj := new(pb.GamerLoginGetDataS2C)
	GamerLoginGetDataS2CType := reflect.TypeOf(*GamerLoginGetDataS2CObj)
	if GamerLoginGetDataS2CType.Kind().String() == "struct" {
		GamerLoginGetDataS2CJson, _ := common.Struct2Json(GamerLoginGetDataS2CType)
		jsonStructSlice = append(jsonStructSlice, GamerLoginGetDataS2CJson)
		GamerLoginGetDataS2CJsonPbFiled := simplejson.New()
		GamerLoginGetDataS2CJsonPbFiled.Set("field", "GamerLoginGetDataS2C")
		GamerLoginGetDataS2CJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerLoginGetDataS2CJsonPbFiled)
	}



	GamerUpgradeHeroS2CObj := new(pb.GamerUpgradeHeroS2C)
	GamerUpgradeHeroS2CType := reflect.TypeOf(*GamerUpgradeHeroS2CObj)
	if GamerUpgradeHeroS2CType.Kind().String() == "struct" {
		GamerUpgradeHeroS2CJson, _ := common.Struct2Json(GamerUpgradeHeroS2CType)
		jsonStructSlice = append(jsonStructSlice, GamerUpgradeHeroS2CJson)
		GamerUpgradeHeroS2CJsonPbFiled := simplejson.New()
		GamerUpgradeHeroS2CJsonPbFiled.Set("field", "GamerUpgradeHeroS2C")
		GamerUpgradeHeroS2CJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerUpgradeHeroS2CJsonPbFiled)
	}



	C2SObj := new(pb.C2S)
	C2SType := reflect.TypeOf(*C2SObj)
	if C2SType.Kind().String() == "struct" {
		C2SJson, _ := common.Struct2Json(C2SType)
		jsonStructSlice = append(jsonStructSlice, C2SJson)
		C2SJsonPbFiled := simplejson.New()
		C2SJsonPbFiled.Set("field", "C2S")
		C2SJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, C2SJsonPbFiled)
	}



	InnerDelGamerFriendObj := new(pb.InnerDelGamerFriend)
	InnerDelGamerFriendType := reflect.TypeOf(*InnerDelGamerFriendObj)
	if InnerDelGamerFriendType.Kind().String() == "struct" {
		InnerDelGamerFriendJson, _ := common.Struct2Json(InnerDelGamerFriendType)
		jsonStructSlice = append(jsonStructSlice, InnerDelGamerFriendJson)
		InnerDelGamerFriendJsonPbFiled := simplejson.New()
		InnerDelGamerFriendJsonPbFiled.Set("field", "InnerDelGamerFriend")
		InnerDelGamerFriendJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, InnerDelGamerFriendJsonPbFiled)
	}



	InnerGamerExitPVPObj := new(pb.InnerGamerExitPVP)
	InnerGamerExitPVPType := reflect.TypeOf(*InnerGamerExitPVPObj)
	if InnerGamerExitPVPType.Kind().String() == "struct" {
		InnerGamerExitPVPJson, _ := common.Struct2Json(InnerGamerExitPVPType)
		jsonStructSlice = append(jsonStructSlice, InnerGamerExitPVPJson)
		InnerGamerExitPVPJsonPbFiled := simplejson.New()
		InnerGamerExitPVPJsonPbFiled.Set("field", "InnerGamerExitPVP")
		InnerGamerExitPVPJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, InnerGamerExitPVPJsonPbFiled)
	}



	InnerGamerLogoutObj := new(pb.InnerGamerLogout)
	InnerGamerLogoutType := reflect.TypeOf(*InnerGamerLogoutObj)
	if InnerGamerLogoutType.Kind().String() == "struct" {
		InnerGamerLogoutJson, _ := common.Struct2Json(InnerGamerLogoutType)
		jsonStructSlice = append(jsonStructSlice, InnerGamerLogoutJson)
		InnerGamerLogoutJsonPbFiled := simplejson.New()
		InnerGamerLogoutJsonPbFiled.Set("field", "InnerGamerLogout")
		InnerGamerLogoutJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, InnerGamerLogoutJsonPbFiled)
	}



	InnerNewGamerFriendObj := new(pb.InnerNewGamerFriend)
	InnerNewGamerFriendType := reflect.TypeOf(*InnerNewGamerFriendObj)
	if InnerNewGamerFriendType.Kind().String() == "struct" {
		InnerNewGamerFriendJson, _ := common.Struct2Json(InnerNewGamerFriendType)
		jsonStructSlice = append(jsonStructSlice, InnerNewGamerFriendJson)
		InnerNewGamerFriendJsonPbFiled := simplejson.New()
		InnerNewGamerFriendJsonPbFiled.Set("field", "InnerNewGamerFriend")
		InnerNewGamerFriendJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, InnerNewGamerFriendJsonPbFiled)
	}



	InnerGetModelDataObj := new(pb.InnerGetModelData)
	InnerGetModelDataType := reflect.TypeOf(*InnerGetModelDataObj)
	if InnerGetModelDataType.Kind().String() == "struct" {
		InnerGetModelDataJson, _ := common.Struct2Json(InnerGetModelDataType)
		jsonStructSlice = append(jsonStructSlice, InnerGetModelDataJson)
		InnerGetModelDataJsonPbFiled := simplejson.New()
		InnerGetModelDataJsonPbFiled.Set("field", "InnerGetModelData")
		InnerGetModelDataJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, InnerGetModelDataJsonPbFiled)
	}



	InnerObj := new(pb.Inner)
	InnerType := reflect.TypeOf(*InnerObj)
	if InnerType.Kind().String() == "struct" {
		InnerJson, _ := common.Struct2Json(InnerType)
		jsonStructSlice = append(jsonStructSlice, InnerJson)
		InnerJsonPbFiled := simplejson.New()
		InnerJsonPbFiled.Set("field", "Inner")
		InnerJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, InnerJsonPbFiled)
	}



	InnerNewGamerChatObj := new(pb.InnerNewGamerChat)
	InnerNewGamerChatType := reflect.TypeOf(*InnerNewGamerChatObj)
	if InnerNewGamerChatType.Kind().String() == "struct" {
		InnerNewGamerChatJson, _ := common.Struct2Json(InnerNewGamerChatType)
		jsonStructSlice = append(jsonStructSlice, InnerNewGamerChatJson)
		InnerNewGamerChatJsonPbFiled := simplejson.New()
		InnerNewGamerChatJsonPbFiled.Set("field", "InnerNewGamerChat")
		InnerNewGamerChatJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, InnerNewGamerChatJsonPbFiled)
	}



	InnerNewPVPResultObj := new(pb.InnerNewPVPResult)
	InnerNewPVPResultType := reflect.TypeOf(*InnerNewPVPResultObj)
	if InnerNewPVPResultType.Kind().String() == "struct" {
		InnerNewPVPResultJson, _ := common.Struct2Json(InnerNewPVPResultType)
		jsonStructSlice = append(jsonStructSlice, InnerNewPVPResultJson)
		InnerNewPVPResultJsonPbFiled := simplejson.New()
		InnerNewPVPResultJsonPbFiled.Set("field", "InnerNewPVPResult")
		InnerNewPVPResultJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, InnerNewPVPResultJsonPbFiled)
	}



	InnerNewGamerFriendReqObj := new(pb.InnerNewGamerFriendReq)
	InnerNewGamerFriendReqType := reflect.TypeOf(*InnerNewGamerFriendReqObj)
	if InnerNewGamerFriendReqType.Kind().String() == "struct" {
		InnerNewGamerFriendReqJson, _ := common.Struct2Json(InnerNewGamerFriendReqType)
		jsonStructSlice = append(jsonStructSlice, InnerNewGamerFriendReqJson)
		InnerNewGamerFriendReqJsonPbFiled := simplejson.New()
		InnerNewGamerFriendReqJsonPbFiled.Set("field", "InnerNewGamerFriendReq")
		InnerNewGamerFriendReqJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, InnerNewGamerFriendReqJsonPbFiled)
	}



	InnerNewGamerMailObj := new(pb.InnerNewGamerMail)
	InnerNewGamerMailType := reflect.TypeOf(*InnerNewGamerMailObj)
	if InnerNewGamerMailType.Kind().String() == "struct" {
		InnerNewGamerMailJson, _ := common.Struct2Json(InnerNewGamerMailType)
		jsonStructSlice = append(jsonStructSlice, InnerNewGamerMailJson)
		InnerNewGamerMailJsonPbFiled := simplejson.New()
		InnerNewGamerMailJsonPbFiled.Set("field", "InnerNewGamerMail")
		InnerNewGamerMailJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, InnerNewGamerMailJsonPbFiled)
	}



	InnerMatchOkObj := new(pb.InnerMatchOk)
	InnerMatchOkType := reflect.TypeOf(*InnerMatchOkObj)
	if InnerMatchOkType.Kind().String() == "struct" {
		InnerMatchOkJson, _ := common.Struct2Json(InnerMatchOkType)
		jsonStructSlice = append(jsonStructSlice, InnerMatchOkJson)
		InnerMatchOkJsonPbFiled := simplejson.New()
		InnerMatchOkJsonPbFiled.Set("field", "InnerMatchOk")
		InnerMatchOkJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, InnerMatchOkJsonPbFiled)
	}



	S2CObj := new(pb.S2C)
	S2CType := reflect.TypeOf(*S2CObj)
	if S2CType.Kind().String() == "struct" {
		S2CJson, _ := common.Struct2Json(S2CType)
		jsonStructSlice = append(jsonStructSlice, S2CJson)
		S2CJsonPbFiled := simplejson.New()
		S2CJsonPbFiled.Set("field", "S2C")
		S2CJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, S2CJsonPbFiled)
	}



	RoomStateObj := new(pb.RoomState)
	RoomStateType := reflect.TypeOf(*RoomStateObj)
	if RoomStateType.Kind().String() == "struct" {
		RoomStateJson, _ := common.Struct2Json(RoomStateType)
		jsonStructSlice = append(jsonStructSlice, RoomStateJson)
		RoomStateJsonPbFiled := simplejson.New()
		RoomStateJsonPbFiled.Set("field", "RoomState")
		RoomStateJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, RoomStateJsonPbFiled)
	}



	RoomGamerStateObj := new(pb.RoomGamerState)
	RoomGamerStateType := reflect.TypeOf(*RoomGamerStateObj)
	if RoomGamerStateType.Kind().String() == "struct" {
		RoomGamerStateJson, _ := common.Struct2Json(RoomGamerStateType)
		jsonStructSlice = append(jsonStructSlice, RoomGamerStateJson)
		RoomGamerStateJsonPbFiled := simplejson.New()
		RoomGamerStateJsonPbFiled.Set("field", "RoomGamerState")
		RoomGamerStateJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, RoomGamerStateJsonPbFiled)
	}



	MailTypeObj := new(pb.MailType)
	MailTypeType := reflect.TypeOf(*MailTypeObj)
	if MailTypeType.Kind().String() == "struct" {
		MailTypeJson, _ := common.Struct2Json(MailTypeType)
		jsonStructSlice = append(jsonStructSlice, MailTypeJson)
		MailTypeJsonPbFiled := simplejson.New()
		MailTypeJsonPbFiled.Set("field", "MailType")
		MailTypeJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, MailTypeJsonPbFiled)
	}



	ChatChannelObj := new(pb.ChatChannel)
	ChatChannelType := reflect.TypeOf(*ChatChannelObj)
	if ChatChannelType.Kind().String() == "struct" {
		ChatChannelJson, _ := common.Struct2Json(ChatChannelType)
		jsonStructSlice = append(jsonStructSlice, ChatChannelJson)
		ChatChannelJsonPbFiled := simplejson.New()
		ChatChannelJsonPbFiled.Set("field", "ChatChannel")
		ChatChannelJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, ChatChannelJsonPbFiled)
	}



	PVPTypeObj := new(pb.PVPType)
	PVPTypeType := reflect.TypeOf(*PVPTypeObj)
	if PVPTypeType.Kind().String() == "struct" {
		PVPTypeJson, _ := common.Struct2Json(PVPTypeType)
		jsonStructSlice = append(jsonStructSlice, PVPTypeJson)
		PVPTypeJsonPbFiled := simplejson.New()
		PVPTypeJsonPbFiled.Set("field", "PVPType")
		PVPTypeJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, PVPTypeJsonPbFiled)
	}



	GashaponTypeObj := new(pb.GashaponType)
	GashaponTypeType := reflect.TypeOf(*GashaponTypeObj)
	if GashaponTypeType.Kind().String() == "struct" {
		GashaponTypeJson, _ := common.Struct2Json(GashaponTypeType)
		jsonStructSlice = append(jsonStructSlice, GashaponTypeJson)
		GashaponTypeJsonPbFiled := simplejson.New()
		GashaponTypeJsonPbFiled.Set("field", "GashaponType")
		GashaponTypeJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GashaponTypeJsonPbFiled)
	}



	ItemTypeObj := new(pb.ItemType)
	ItemTypeType := reflect.TypeOf(*ItemTypeObj)
	if ItemTypeType.Kind().String() == "struct" {
		ItemTypeJson, _ := common.Struct2Json(ItemTypeType)
		jsonStructSlice = append(jsonStructSlice, ItemTypeJson)
		ItemTypeJsonPbFiled := simplejson.New()
		ItemTypeJsonPbFiled.Set("field", "ItemType")
		ItemTypeJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, ItemTypeJsonPbFiled)
	}



	MailStateObj := new(pb.MailState)
	MailStateType := reflect.TypeOf(*MailStateObj)
	if MailStateType.Kind().String() == "struct" {
		MailStateJson, _ := common.Struct2Json(MailStateType)
		jsonStructSlice = append(jsonStructSlice, MailStateJson)
		MailStateJsonPbFiled := simplejson.New()
		MailStateJsonPbFiled.Set("field", "MailState")
		MailStateJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, MailStateJsonPbFiled)
	}



	TaskStateObj := new(pb.TaskState)
	TaskStateType := reflect.TypeOf(*TaskStateObj)
	if TaskStateType.Kind().String() == "struct" {
		TaskStateJson, _ := common.Struct2Json(TaskStateType)
		jsonStructSlice = append(jsonStructSlice, TaskStateJson)
		TaskStateJsonPbFiled := simplejson.New()
		TaskStateJsonPbFiled.Set("field", "TaskState")
		TaskStateJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, TaskStateJsonPbFiled)
	}



	GamerStatusObj := new(pb.GamerStatus)
	GamerStatusType := reflect.TypeOf(*GamerStatusObj)
	if GamerStatusType.Kind().String() == "struct" {
		GamerStatusJson, _ := common.Struct2Json(GamerStatusType)
		jsonStructSlice = append(jsonStructSlice, GamerStatusJson)
		GamerStatusJsonPbFiled := simplejson.New()
		GamerStatusJsonPbFiled.Set("field", "GamerStatus")
		GamerStatusJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerStatusJsonPbFiled)
	}



	ServerInfoObj := new(pb.ServerInfo)
	ServerInfoType := reflect.TypeOf(*ServerInfoObj)
	if ServerInfoType.Kind().String() == "struct" {
		ServerInfoJson, _ := common.Struct2Json(ServerInfoType)
		jsonStructSlice = append(jsonStructSlice, ServerInfoJson)
		ServerInfoJsonPbFiled := simplejson.New()
		ServerInfoJsonPbFiled.Set("field", "ServerInfo")
		ServerInfoJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, ServerInfoJsonPbFiled)
	}



	ServerWarnObj := new(pb.ServerWarn)
	ServerWarnType := reflect.TypeOf(*ServerWarnObj)
	if ServerWarnType.Kind().String() == "struct" {
		ServerWarnJson, _ := common.Struct2Json(ServerWarnType)
		jsonStructSlice = append(jsonStructSlice, ServerWarnJson)
		ServerWarnJsonPbFiled := simplejson.New()
		ServerWarnJsonPbFiled.Set("field", "ServerWarn")
		ServerWarnJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, ServerWarnJsonPbFiled)
	}



	ServerCmdObj := new(pb.ServerCmd)
	ServerCmdType := reflect.TypeOf(*ServerCmdObj)
	if ServerCmdType.Kind().String() == "struct" {
		ServerCmdJson, _ := common.Struct2Json(ServerCmdType)
		jsonStructSlice = append(jsonStructSlice, ServerCmdJson)
		ServerCmdJsonPbFiled := simplejson.New()
		ServerCmdJsonPbFiled.Set("field", "ServerCmd")
		ServerCmdJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, ServerCmdJsonPbFiled)
	}



	ServerHelloObj := new(pb.ServerHello)
	ServerHelloType := reflect.TypeOf(*ServerHelloObj)
	if ServerHelloType.Kind().String() == "struct" {
		ServerHelloJson, _ := common.Struct2Json(ServerHelloType)
		jsonStructSlice = append(jsonStructSlice, ServerHelloJson)
		ServerHelloJsonPbFiled := simplejson.New()
		ServerHelloJsonPbFiled.Set("field", "ServerHello")
		ServerHelloJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, ServerHelloJsonPbFiled)
	}



	GamerNotifyListObj := new(pb.GamerNotifyList)
	GamerNotifyListType := reflect.TypeOf(*GamerNotifyListObj)
	if GamerNotifyListType.Kind().String() == "struct" {
		GamerNotifyListJson, _ := common.Struct2Json(GamerNotifyListType)
		jsonStructSlice = append(jsonStructSlice, GamerNotifyListJson)
		GamerNotifyListJsonPbFiled := simplejson.New()
		GamerNotifyListJsonPbFiled.Set("field", "GamerNotifyList")
		GamerNotifyListJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerNotifyListJsonPbFiled)
	}



	GamerNotifyObj := new(pb.GamerNotify)
	GamerNotifyType := reflect.TypeOf(*GamerNotifyObj)
	if GamerNotifyType.Kind().String() == "struct" {
		GamerNotifyJson, _ := common.Struct2Json(GamerNotifyType)
		jsonStructSlice = append(jsonStructSlice, GamerNotifyJson)
		GamerNotifyJsonPbFiled := simplejson.New()
		GamerNotifyJsonPbFiled.Set("field", "GamerNotify")
		GamerNotifyJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerNotifyJsonPbFiled)
	}



	GamerExitPVPObj := new(pb.GamerExitPVP)
	GamerExitPVPType := reflect.TypeOf(*GamerExitPVPObj)
	if GamerExitPVPType.Kind().String() == "struct" {
		GamerExitPVPJson, _ := common.Struct2Json(GamerExitPVPType)
		jsonStructSlice = append(jsonStructSlice, GamerExitPVPJson)
		GamerExitPVPJsonPbFiled := simplejson.New()
		GamerExitPVPJsonPbFiled.Set("field", "GamerExitPVP")
		GamerExitPVPJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, GamerExitPVPJsonPbFiled)
	}



	RPCObj := new(pb.RPC)
	RPCType := reflect.TypeOf(*RPCObj)
	if RPCType.Kind().String() == "struct" {
		RPCJson, _ := common.Struct2Json(RPCType)
		jsonStructSlice = append(jsonStructSlice, RPCJson)
		RPCJsonPbFiled := simplejson.New()
		RPCJsonPbFiled.Set("field", "RPC")
		RPCJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, RPCJsonPbFiled)
	}



	LeagueDataObj := new(pb.LeagueData)
	LeagueDataType := reflect.TypeOf(*LeagueDataObj)
	if LeagueDataType.Kind().String() == "struct" {
		LeagueDataJson, _ := common.Struct2Json(LeagueDataType)
		jsonStructSlice = append(jsonStructSlice, LeagueDataJson)
		LeagueDataJsonPbFiled := simplejson.New()
		LeagueDataJsonPbFiled.Set("field", "LeagueData")
		LeagueDataJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, LeagueDataJsonPbFiled)
	}



	OnlineLeagueObj := new(pb.OnlineLeague)
	OnlineLeagueType := reflect.TypeOf(*OnlineLeagueObj)
	if OnlineLeagueType.Kind().String() == "struct" {
		OnlineLeagueJson, _ := common.Struct2Json(OnlineLeagueType)
		jsonStructSlice = append(jsonStructSlice, OnlineLeagueJson)
		OnlineLeagueJsonPbFiled := simplejson.New()
		OnlineLeagueJsonPbFiled.Set("field", "OnlineLeague")
		OnlineLeagueJsonPbFiled.Set("type", "struct")
		jsonPbFiledSlice = append(jsonPbFiledSlice, OnlineLeagueJsonPbFiled)
	}


	jsonPb.Set("properties", jsonPbFiledSlice)
	jsonStructSlice = append(jsonStructSlice, jsonPb)
	jsonConfig.Set("config", jsonStructSlice)
	str, _ := jsonConfig.MarshalJSON()
	ioutil.WriteFile("main.json", str, 0777)
}

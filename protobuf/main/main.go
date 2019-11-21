package main

import (
	user_info "StudyGo/protobuf"
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	"os"
)

func main()  {
	jsonFilePath := "/Users/tt/Code/GoCode/StudyGo/protobuf/user.json"
	//pbFilePath := "/Users/tt/Code/GoCode/StudyGo/protobuf/user2.pb"
	//jsonOutFile := "/Users/tt/Code/GoCode/StudyGo/protobuf/output.json"

	userInfo := new(user_info.UserInfo)

	json.Marshal(userInfo)

	file, err := os.Open(jsonFilePath)
	err = jsonpb.Unmarshal(file, userInfo)
	fmt.Println(userInfo.String())

	userInfoUser := new(user_info.UserInfo_User)
	userInfoUser.Age = 25
	userInfoUser.Username = "tt"
	userInfoUser.Graduate = "grade"

	userInfo.UserList = append(userInfo.UserList, userInfoUser)

	marashaler := new(jsonpb.Marshaler)
	out, err := marashaler.MarshalToString(userInfo)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(out)
	}
}
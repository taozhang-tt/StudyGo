package main

import (
	"StudyGo/elasticSearch/api"

	"fmt"
)

func main() {

	var (
		client *api.EsClient
		err    error
		info   string
	)
	//连接服务端
	if client, err = api.NewClient("127.0.0.1", "9200"); err != nil {
		fmt.Println(err)
		return
	}
	//打印服务端信息
	if info, err = client.GetClientInfo(); err != nil {
		fmt.Println(info)
		return
	}

	//检查 & 创建索引
	if exist, err := client.ExistIndex("test_go"); err != nil {
		fmt.Println(err)
		return
	} else if !exist {
		indexConfig := `{"settings":{"number_of_shards":2,"number_of_replicas":1},"mappings":{"properties":{"group_id":{"type":"integer"},"name":{"analyzer":"smartcn","type":"text"},"answer":{"analyzer":"smartcn","type":"text"},"single":{"type":"integer"},"link_table":{"type":"keyword"},"status":{"type":"integer"},"template_id":{"type":"integer"},"created_at":{"type":"date"},"updated_at":{"type":"integer"},"words":{"type":"integer"}}}}`
		if err := client.AddIndex("test_go", indexConfig); err != nil {
			fmt.Println(err)
			return
		}
	}

	//初始化数据
	//doc1 := `{"group_id":1,"name":"如何领取奖励","answer":"公众号领取，应用内签到","single":1,"link_table":"","status":1,"template_id":1,"created_at":"2019-12-11","updated_at":1000,"words":2500}`
	//doc2 := `{"group_id":1,"name":"老师演示了integer和date，那其他的类型聚合本地测试报非法参数异常illegal_argument_exception","answer":"兄嘚，这种情况要看具体问题了，问问题，报了什么错，要贴出来啊，老师演示没问题，基本就是你的问题了","single":1,"link_table":"","status":1,"template_id":1,"created_at":"2019-12-12","updated_at":2000,"words":1500}`
	//doc3 := `{"group_id":1,"name":"elasticsearch中,索引类似于database,类型类似于table;那么 一个索引下可以有多个类型吧?","answer":"好好学习多多思考","single":1,"link_table":"","status":1,"template_id":1,"created_at":"2019-12-12","updated_at":4000,"words":2500}`
	//doc4 := `{"group_id":1,"name":"请问录制视频用的是什么软件","answer":"Mac自带QuickTime Player","single":1,"link_table":"","status":1,"template_id":1,"created_at":"2019-12-13","updated_at":3000,"words":3500}`
	//_ = client.AddDocument("test_go", "_doc", doc1)
	//_ = client.AddDocument("test_go", "_doc", doc2)
	//_ = client.AddDocument("test_go", "_doc", doc3)
	//_ = client.AddDocument("test_go", "_doc", doc4)


	query := `{"query":{"match":{"name":{"query":"类型类似","operator":"and"}}},"highlight":{"fields":{"name":{}}}}`

	client.GetByField("name", "类型类似", "test_go")

	contents, err := client.GetByQuery(query, "test_go")
	if err != nil{
		fmt.Println(err)
		return
	}
	for _, content := range contents{
		fmt.Println(content)
	}
}

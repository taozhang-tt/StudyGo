package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

//etcd 写操作
func main()  {
	var (
		client *clientv3.Client
		err error
		kv clientv3.KV
		putResponse *clientv3.PutResponse
	)
	client, err = clientv3.New(clientv3.Config{
		Endpoints: []string{"122.51.91.106:2379"},
		DialTimeout: 5*time.Second,
	})
	if err != nil {
		fmt.Print(err)
	}

	//用于读写etcd 的键值对
	kv = clientv3.NewKV(client)

	//直接读写
	//if putResponse, err = kv.Put(context.TODO(), "/cron/jobs/job1", "hello world"); err != nil{
	//	fmt.Println(err)
	//} else {
	//	fmt.Println(putResponse.Header.Revision)
	//}

	//带 WithPrevKV 的读写
	if putResponse, err = kv.Put(context.TODO(), "/cron/jobs/job1", "hello tt", clientv3.WithPrevKV()); err != nil{
		fmt.Println(err)
	} else {
		fmt.Println("Revision: ", putResponse.Header.Revision)
		if putResponse.PrevKv != nil {
			fmt.Print("prevKv: ", string(putResponse.PrevKv.Value))
		}
	}
}

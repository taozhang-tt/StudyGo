package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)
//etcd 获取目录下所有kv
func main()  {
	var (
		client *clientv3.Client
		err error
		kv clientv3.KV
		getResponse *clientv3.GetResponse
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

	//带 WithPrevKV 的读写
	if getResponse, err = kv.Get(context.TODO(), "/cron/jobs/", clientv3.WithPrefix()); err != nil{
		fmt.Println(err)
	} else {
		fmt.Println(getResponse.Kvs)
	}
}

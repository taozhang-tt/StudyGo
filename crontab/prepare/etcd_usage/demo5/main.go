package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/mvcc/mvccpb"
	"go.etcd.io/etcd/clientv3"
	"time"
)

//etcd 删除操作
func main()  {
	var (
		client *clientv3.Client
		err error
		kv clientv3.KV
		delResponse *clientv3.DeleteResponse
		keyPair *mvccpb.KeyValue
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

	if delResponse, err = kv.Delete(context.TODO(), "/cron/jobs/job2", clientv3.WithPrevKV()); err != nil{
		fmt.Println(err)
	} else {
		//被删除之前的value
		if delResponse.PrevKvs != nil {
			for _, keyPair = range delResponse.PrevKvs {
				fmt.Println("删除了--- key：", string(keyPair.Key), " value: ", string(keyPair.Value))
			}
		}
	}
}

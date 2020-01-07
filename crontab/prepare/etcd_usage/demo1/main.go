package main

import (
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

func main()  {
	var (
		client *clientv3.Client
		err error
	)
	client, err = clientv3.New(clientv3.Config{
		Endpoints: []string{"122.51.91.106:2379"},
		DialTimeout: 5*time.Second,
	})
	if err != nil {
		fmt.Print(err)
	}
	client = client
}

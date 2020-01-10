package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

//实现kv过期
func main() {
	var (
		client         *clientv3.Client
		err            error
		kv             clientv3.KV
		lease          clientv3.Lease
		leaseGrantResp *clientv3.LeaseGrantResponse
		leaseId        clientv3.LeaseID
		putResp        *clientv3.PutResponse
		getResp        *clientv3.GetResponse
		leaseKeepAliveRespChan <-chan *clientv3.LeaseKeepAliveResponse
		leaseKeepAliveResp *clientv3.LeaseKeepAliveResponse
	)
	client, err = clientv3.New(clientv3.Config{
		Endpoints:   []string{"122.51.91.106:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Print(err)
	}

	//用于读写etcd 的键值对
	kv = clientv3.NewKV(client)

	//申请一个10s的租约
	lease = clientv3.NewLease(client)
	if leaseGrantResp, err = lease.Grant(context.TODO(), 10); err != nil {
		fmt.Println(err)
		return
	}
	leaseId = leaseGrantResp.ID

	//续租，类似发送心跳包，每次续租成功都会有返回
	//if leaseKeepAliveRespChan, err = lease.KeepAlive(context.TODO(), leaseId); err != nil{
	//	fmt.Println(err)
	//	return
	//}

	//5s后取消续约
	ctx, _ := context.WithTimeout(context.TODO(), 5*time.Second)
	if leaseKeepAliveRespChan, err = lease.KeepAlive(ctx, leaseId); err != nil{
		fmt.Println(err)
		return
	}

	//处理续约应答的协程
	go func() {
		for {
			select {
			case leaseKeepAliveResp = <-leaseKeepAliveRespChan:
				if leaseKeepAliveResp == nil {
					fmt.Println("不再续租")
					goto END
				} else {
					fmt.Println("收到自动续租应答：", leaseKeepAliveResp.ID)
				}
			}
		}
		END:
	}()

	//put 一个 kv，让它与租约关联起来，从而实现10s后过期
	if putResp, err = kv.Put(context.TODO(), "/cron/lock/job1", "", clientv3.WithLease(leaseId)); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("写入成功", putResp.Header.Revision)

	for {
		if getResp, err = kv.Get(context.TODO(), "/cron/lock/job1"); err != nil {
			fmt.Println(err)
			return
		}
		if getResp.Count == 0 {
			fmt.Println("kv过期了")
			break
		} else {
			fmt.Println("还没过期", getResp.Kvs)
			time.Sleep(2*time.Second)
		}
	}
}

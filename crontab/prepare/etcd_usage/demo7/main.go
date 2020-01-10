package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/mvcc/mvccpb"
	"go.etcd.io/etcd/clientv3"
	"time"
)

func main()  {
	var (
		client *clientv3.Client
		err error
		kv clientv3.KV
		getResp *clientv3.GetResponse
		watchStartRevision int64
		watcher clientv3.Watcher
		watchRespChan <-chan clientv3.WatchResponse
		watchResp clientv3.WatchResponse
		event *clientv3.Event
	)
	client, err = clientv3.New(clientv3.Config{
		Endpoints: []string{"122.51.91.106:2379"},
		DialTimeout: 5*time.Second,
	})
	if err != nil {
		fmt.Print(err)
		return
	}
	kv = clientv3.NewKV(client)

	//模拟etcd中的kv的变化
	go func() {
		for {
			kv.Put(context.TODO(), "/cron/jobs/job7", "i am job7")

			kv.Delete(context.TODO(), "/cron/jobs/job7")

			time.Sleep(1*time.Second)
		}
	}()

	//先GET到当前到值，并监听后续变化
	if getResp, err = kv.Get(context.TODO(), "/cron/jobs/job7"); err != nil{
		fmt.Println(err)
		return
	}
	//现在key是存在到
	if len(getResp.Kvs) != 0 {
		fmt.Println("当前值", string(getResp.Kvs[0].Value))
	}

	//当前etcd集群事务id，单调递增，从下一个开始监听
	watchStartRevision = getResp.Header.Revision + 1

	//创建一个watcher
	watcher = clientv3.NewWatcher(client)

	//启动监听
	fmt.Println("从该版本向后监听：", watchStartRevision)
	watchRespChan = watcher.Watch(context.TODO(), "/cron/jobs/job7", clientv3.WithRev(watchStartRevision))
	for watchResp = range watchRespChan {
		for _, event = range watchResp.Events {
			switch event.Type {
			case mvccpb.PUT:
				fmt.Println("修改为：", string(event.Kv.Value), "Revision: ", event.Kv.CreateRevision, event.Kv.ModRevision)
			case mvccpb.DELETE:
				fmt.Println("删除了：", string(event.Kv.Value), "Revision: ", event.Kv.ModRevision)
			}
		}
	}
}

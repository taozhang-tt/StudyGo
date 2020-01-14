package main

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
	"fmt"
)

// 任务的执行时间点
type TimePoint struct {
	StartTime int64 `bson:"startTime"`
	EndTime   int64 `bson:"endTime"`
}

// 一条日志
type LogRecord struct {
	JobName   string    `bson:"jobName"`   // 任务名
	Command   string    `bson:"command"`   // shell命令
	Err       string    `bson:"err"`       // 脚本错误
	Content   string    `bson:"content"`   // 脚本输出
	TimePoint TimePoint `bson:"timePoint"` // 执行时间点
}

func main() {
	var (
		client     *mongo.Client
		err        error
		database   *mongo.Database
		collection *mongo.Collection
		record     *LogRecord
		result     *mongo.InsertOneResult
		docId      primitive.ObjectID
	)
	// 1, 建立连接
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	if client, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://122.51.91.106:27017")); err != nil {
		fmt.Println(err)
	}
	// 2, 选择数据库cron
	database = client.Database("cron")

	// 3, 选择表log
	collection = database.Collection("log")

	// 4, 插入记录(bson)
	record = &LogRecord{
		JobName:   "job10",
		Command:   "echo hello",
		Err:       "",
		Content:   "hello",
		TimePoint: TimePoint{StartTime: time.Now().Unix(), EndTime: time.Now().Unix() + 10},
	}

	if result, err = collection.InsertOne(context.TODO(), record); err != nil {
		fmt.Println(err)
		return
	}

	// _id: 默认生成一个全局唯一ID, ObjectID：12字节的二进制
	docId = result.InsertedID.(primitive.ObjectID)
	fmt.Println(docId.Hex())
}

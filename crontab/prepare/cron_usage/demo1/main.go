package main

import (
	"fmt"
	"github.com/gorhill/cronexpr"
	"time"
)

func main()  {
	var (
		err error
		exp *cronexpr.Expression
		now time.Time
		next time.Time
	)

	//分钟，小时，天，月，周
	if exp, err = cronexpr.Parse("*/5 * * * * * *"); err != nil{
		fmt.Println(err)
	}
	now = time.Now()
	fmt.Println("now:", now)
	next = exp.Next(now)
	fmt.Println("next:", next)

	time.AfterFunc(next.Sub(now), func() {
		fmt.Println("被调度了：", next)
	})

	time.Sleep(10 * time.Second)
}

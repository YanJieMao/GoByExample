package main

//基于协程、通道和打点器，支持速率限制。
import (
	"fmt"
	"time"
)

func main() {
	requests := make(chan int, 5)

	for i := 1; i <= 5; i++ {
		requests <- i
	}

	close(requests)

	limiter := time.Tick(200 * time.Millisecond) //任务速率限制的调度器

	for req := range requests {
		<-limiter //通过阻塞接收，限制频率
		fmt.Println("请求", req, time.Now())

	}

	burstyLimiter := make(chan time.Time, 3) //短暂并发，通过缓冲通道

	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	burstyRequest := make(chan int, 5)
	for i := 1; i <= 5; i++ { //传入5个值
		burstyRequest <- i
	}
	close(burstyRequest)
	for rep := range burstyRequest {
		<-burstyLimiter //阻塞请求
		fmt.Println("请求", rep, time.Now())
	}

}

/*
请求 1 2020-12-04 13:35:44.4974996 +0800 CST m=+0.204137501
请求 2 2020-12-04 13:35:44.7062227 +0800 CST m=+0.412860601
请求 3 2020-12-04 13:35:44.8969592 +0800 CST m=+0.603597101
请求 4 2020-12-04 13:35:45.0971244 +0800 CST m=+0.803762301
请求 5 2020-12-04 13:35:45.3007773 +0800 CST m=+1.007415201
请求 1 2020-12-04 13:35:45.3007773 +0800 CST m=+1.007415201
请求 2 2020-12-04 13:35:45.3007773 +0800 CST m=+1.007415201
请求 3 2020-12-04 13:35:45.3007773 +0800 CST m=+1.007415201
请求 4 2020-12-04 13:35:45.5046709 +0800 CST m=+1.211308801
请求 5 2020-12-04 13:35:45.7072585 +0800 CST m=+1.413896401

*/

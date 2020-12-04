package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {

	defer wg.Done() //defer 在return时函数才被调用，在这里return 时，通知 WaitGroup，当前协程的工作已经完成

	fmt.Printf("woker %d 启动\n", id)

	time.Sleep(time.Second)

	fmt.Printf("woker %d 完成\n", id)

}

func main() {
	var wg sync.WaitGroup //用于等待该函数启动的所有协程。

	for i := 1; i <= 5; i++ {
		wg.Add(1) //几个协程几个waitGroup的计数器
		go worker(i, &wg)
	}

	wg.Wait()
}

/*
woker 1 启动
woker 4 启动
woker 3 启动
woker 5 启动
woker 2 启动
woker 5 完成
woker 4 完成
woker 3 完成
woker 1 完成
woker 2 完成

*/

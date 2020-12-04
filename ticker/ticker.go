package main

//定时器执行一次，打点器重复执行
import (
	"fmt"
	"time"
)

func main() {

	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return

			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	time.Sleep(1600 * time.Millisecond) //打点器一旦停止，将不能再从它的通道中接收到值。
	ticker.Stop()
	done <- true
	fmt.Println("Ticker 停止")

}

/*
Tick at 2020-12-04 11:02:26.9667895 +0800 CST m=+0.503678701
Tick at 2020-12-04 11:02:27.4777576 +0800 CST m=+1.014646801
Tick at 2020-12-04 11:02:27.9663214 +0800 CST m=+1.503210601
Ticker 停止
*/

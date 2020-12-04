package main

import (
	"fmt"
	"time"
)

/*
time.NewTimer() 创建，这种类型，timer只会执行一次，
当然，可以在执行完以后通过调用 timer.Reset()
让定时器再次工作，并可以更改时间间隔。
*/

func main() {
	timer1 := time.NewTimer(2 * time.Second)

	<-timer1.C //<-timer1.C 会一直阻塞， 直到定时器的通道 C 明确的发送了定时器失效的值。
	fmt.Println("定时器1")

	timer2 := time.NewTimer(time.Second)

	go func() {
		<-timer2.C
		fmt.Println("定时器2")

	}()

	stop2 := timer2.Stop() //使用定时器可以在定时器触发之前将其取消。

	if stop2 {
		fmt.Println("定时器2停止")
	}

	time.Sleep(2 * time.Second)

}

/*
定时器1
定时器2停止
*/

package main

//用通道来同步协程之间的执行状态。 这里使用阻塞接收的方式，实现了等待另一个协程完成。
import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true //发送通知
}

func main() {

	done := make(chan bool, 1)
	go worker(done)

	msg := <-done //接收通道的通知，如果去除通道同步，woker执行前程序都结束了
	println(msg)

}

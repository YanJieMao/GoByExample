package main

import "fmt"

func ping(pings chan<- string, msg string) { //定义了一个只写通道
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings //只读通道，接收收据
	pongs <- msg   //只写通道，发送数据
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "消息已发送")
	pong(pings, pongs)
	fmt.Println(<-pongs)

}

package main

//通道连接多个协程的管道，两个协程通过通道传值

import "fmt"

func main() {

	messages := make(chan string) //make(chan val-type) 创建一个新的通道

	go func() { messages <- "啦啦啦啦" }() //channel <- 语法 发送 一个新的值到通道中

	msg := <-messages //<-channel 语法从通道中 接收 一个值
	fmt.Println(msg)
	/* 默认发送和接收操作是阻塞的，直到发送方和接收方都就绪。
	这个特性使得可以不使用任何其它的同步操作，
	就可以在程序结尾处等待消息 */
}

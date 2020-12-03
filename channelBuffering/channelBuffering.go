package main

import "fmt"

func main() {
	messages := make(chan string, 2) //一个字符串通道，设置最多允许缓存2个值

	messages <- "啦啦啦"
	messages <- "呀呀呀"

	fmt.Println(<-messages) //因为有缓冲，所以可以发送到通道中，无需并发接收
	fmt.Println(<-messages)

}

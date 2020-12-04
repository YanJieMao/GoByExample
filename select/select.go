package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "一"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "二"
	}()

	for i := 0; i < 2; i++ {
		select { //select的case必须是发送或者接收操作
		case msg1 := <-c1:
			fmt.Println("接收", msg1)

		case msg2 := <-c2:
			fmt.Println("接收", msg2)
		}
	}

}

/*
接收 一
接收 二
*/

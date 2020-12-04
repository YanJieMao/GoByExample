package main

//通过通道和select实现超时操作
import (
	"fmt"
	"time"
)

func main() {

	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "结果 1"

	}()

	select { // select 默认处理第一个已准备好的接收操作，超时1秒，将会执行超时 case。
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("超时 1")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "结果 2"
	}()

	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("超时 2")

	}

}

/*
超时 1
结果 2
*/

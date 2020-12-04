package main

//在select中添加一个default子句实现非阻塞的发送接收

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	select { //非阻塞接收
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	msg := "hi"
	select { //非阻塞发送
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	select { //多路非阻塞选择器
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}

/*
no message received
no message sent
no activity
*/

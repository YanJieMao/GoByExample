package main

import "fmt"

func main() {

	queue := make(chan string, 2)
	queue <- "一"
	queue <- "二"

	close(queue)

	for elem := range queue { //range遍历通道中的值
		fmt.Println(elem)
	}

}

/*
通道是一个引用对象，和 map 类似。
map 在没有任何外部引用时，Go语言程序在运行时（runtime）会自动对内存进行垃圾回收（Garbage Collection, GC）。
类似的，通道也可以被垃圾回收，但是通道也可以被主动关闭。

通道关闭后可以接收数据，但是不可以发送数据，发送数据会panic
从已关闭的通道接收数据时将不会发生阻塞（通道有默认的0值）

*/

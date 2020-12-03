package main

//协程轻量级的
import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	f("执行函数f") //执行一个函数

	go f("1号goroutine") //使用 go f(s) 在一个协程中调用这个函数。并发执行函数

	go func(msg string) { //在匿名函数中启动一个协程
		fmt.Println(msg)
	}("2号goroutine")

	time.Sleep(time.Second)
	fmt.Println("结束")
}

/*
执行函数f : 0
执行函数f : 1
执行函数f : 2
2号goroutine
1号goroutine : 0
1号goroutine : 1
1号goroutine : 2
结束

*/

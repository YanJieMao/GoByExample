package main

import "fmt"

func main() {
	var a = "initial" //声明一个变量并初始化
	fmt.Println(a)

	var b, c int = 1, 2 //一次声明多个变量并初始化
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int //未赋值会初始化为0值
	fmt.Println(e)

	f := "short" //简短声明方式
	fmt.Println(f)

	/*
	   initial
	   1 2
	   true
	   0
	   short

	*/

}

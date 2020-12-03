package main

//闭包：就是能够读取另一个函数内的变量的函数

import "fmt"

func intSeq() func() int { //形成了一个闭包
	i := 0
	return func() int { //匿名函数
		i++
		return i
	}
}

func main() {
	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())
}

/*
1
2
3
1

*/

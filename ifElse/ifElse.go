package main

import "fmt"

func main() {
	if 7%2 == 0 {
		fmt.Println("7是偶数")
	} else {
		fmt.Println("7是奇数")
	}

	if 8%2 == 0 {
		fmt.Println("8被2整除") //可以只要if
	}

	if num := 9; num < 0 { //条件语句前可以有一个声明语句
		fmt.Println(num, "是负数")
	} else if num < 10 {
		fmt.Println(num, "个位数")
	} else {
		fmt.Println(num, "比10大")
	}
	//go没有三目运算符
}

/*
7是奇数
8被2整除
9 个位数

*/

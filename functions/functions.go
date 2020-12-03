package main

import "fmt"

func plus(a int, b int) int {

	return a + b //Go 需要明确的 return，也就是说，它不会自动 return 最后一个表达式的值
}

func plusPlus(a, b, c int) int { //多个连续参数类型相同时，可以只声明最后一个
	return a + b + c
}

func main() {

	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)
}

/*
1+2 = 3
1+2+3 = 6

*/

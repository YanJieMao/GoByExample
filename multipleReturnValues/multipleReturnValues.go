package main

import "fmt"

func vals() (int, int) {
	return 3, 7
}

func main() {

	a, b := vals() //返回的两个值赋值给a,b
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals() //返回的值只要一个7
	fmt.Println(c)

	c, _ = vals() //返回的值只要一个3
	fmt.Println(c)
}

/*
3
7
7
3
*/

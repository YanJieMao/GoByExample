package main

import "fmt"

func main() {
	s := make([]string, 3) //创建不为空的slice需要用make
	fmt.Println("空:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("set:", s)
	fmt.Println("取值:", s[2])

	fmt.Println("长度:", len(s))

	s = append(s, "d") //append返回一个新的slice
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

}

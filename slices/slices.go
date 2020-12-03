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

	c := make([]string, len(s)) //复制数组s给c
	copy(c, s)
	fmt.Println("复制", c)
	//slice 支持通过 slice[low:high] 语法进行“切片”操作
	l := s[2:5] //数组下表（2，5）不包含2，5
	fmt.Println("sl1", l)

	l = s[:5]
	fmt.Println("sl2", l) //0-5不含5

	l = s[2:]
	fmt.Println("sl3", l) //2之后不含2

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	twoD := make([][]int, 3) //slice可以多维，但内部slice长度可以不一致
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}

/*
空: [  ]
set: [a b c]
取值: c
长度: 3
apd: [a b c d e f]
复制 [a b c d e f]
sl1 [c d e]
sl2 [a b c d e]
sl3 [c d e f]
dcl: [g h i]
2d:  [[0] [1 2] [2 3 4]]
*/

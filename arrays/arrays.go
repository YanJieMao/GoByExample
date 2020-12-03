package main

import "fmt"

func main() {

	var a [5]int
	fmt.Println("空数组", a)

	a[4] = 100
	fmt.Println("放入一个", a)
	fmt.Println("取值", a[4])
	fmt.Println("长度:", len(a)) //内置函数获取数组长度

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("输出：", b)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("二维数组 ", twoD)

}

/*
空数组 [0 0 0 0 0]
放入一个 [0 0 0 0 100]
取值 100
长度: 5
输出： [1 2 3 4 5]
二维数组  [[0 1 2] [1 2 3]]
*/

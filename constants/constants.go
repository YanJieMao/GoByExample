package main

import (
	"fmt"
	"math"
)

const s string = "常量" //const用来声明常量

func main() {
	fmt.Println(s)
	const n = 300000000

	const d = 3e20 / n //常数表达式可以执行任意精度的运算
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n)) //数字可以根据需要自动确定类型

	/*常量
	  1e+12
	  1000000000000
	  -0.4395520511837499
	*/

}

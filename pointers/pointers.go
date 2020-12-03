package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("intial:", i)

	zeroval(i)                 //这里传入的是值1，会修改地址为1的值为0
	fmt.Println("zeroval:", i) //变量i的值并没有被修改

	zeroptr(&i) //获取了i的地址，传入i
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)

}

/*
intial: 1
zeroval: 1
zeroptr: 0
pointer: 0xc0000120a0
*/

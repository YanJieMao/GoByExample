package main

import "fmt"

type rect struct { //定义结构体
	width, height int
}

func (r *rect) area() int { // area 是一个拥有 *rect 类型接收器(receiver)的方法。 转换为方法后，我们就可以用面向对象的方法来调用
	return r.width * r.height
}

func (r rect) perim() int { //可以为值类型或者指针类型的接收者定义方法。 这是一个值类型接收者
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}

	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())
	/* 调用方法时，Go 会自动处理值和指针之间的转换。
	想避免在调用方法时产生一个拷贝，或者修改接受结构体的值，
	都可以使用指针来调用方法 */
}

/*
area:  50
perim: 30
area:  50
perim: 30
*/

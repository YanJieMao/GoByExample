package main

import "fmt"

func sum(nums ...int) { //变参函数可以传入任意的参数
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4} //多个值的 slice，想把它们作为参数使用， 这样调用 func(slice...)。
	sum(nums...)
}

package main

// sort 包实现了内建及用户自定义数据类型的排序功能
import (
	"fmt"
	"sort"
)

func main() {

	strs := []string{"c", "a", "b"}
	sort.Strings(strs) //原地排序，直接改变给定的切片
	fmt.Println("字符串:", strs)

	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("整数: ", ints)

	s := sort.IntsAreSorted(ints) //检查一个切片是否有序
	fmt.Println("是否有序 ", s)
}

/*
字符串: [a b c]
整数:  [2 4 7]
是否有序  true
*/

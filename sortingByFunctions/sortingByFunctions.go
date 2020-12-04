package main

import (
	"fmt"
	"sort"
)

type byLength []string //创建一个类型

func (s byLength) Len() int { //实现了 sort.Interface 接口的 Len、Less 和 Swap 方法
	return len(s)
}
func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func main() {
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(byLength(fruits)) //转换为bylenth类型然后排序
	fmt.Println(fruits)
}

/*
[kiwi peach banana]
*/

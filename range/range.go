package main

import "fmt"

func main() {

	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums { //range 在数组和 slice 中提供对每项的索引和值的访问。
		sum += num
	}

	fmt.Println("sum", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs { //range 也可以只遍历 map 的键。
		fmt.Println("key:", k)
	}

	for i, c := range "go" { //第一个返回值是字符的起始字节位置，然后第二个是字符本身。
		fmt.Println(i, c)
	}

}

/*
sum 9
index 1
a -> apple
b -> banana
key: a
key: b
0 103
1 111

*/

package main

import "fmt"

func main() {
	m := make(map[string]int) //创建一个空 map，需要使用内建函数 make：make(map[key-type]val-type)

	m["k1"] = 7 //典型的 make[key] = val 语法来设置键值对。
	m["k2"] = 13

	fmt.Println("map", m)

	v1 := m["k1"] //使用 name[key] 来获取一个键的值。

	fmt.Println("v1:", v1)

	fmt.Println("len:", len(m)) //len显示键值对数量

	delete(m, "k2")
	fmt.Println("map", m) //内建函数 delete 可以从一个 map 中移除键值对。

	_, prs := m["k2"]
	fmt.Println("prs:", prs) //当从一个 map 中取值时，还有可以选择是否接收的第二个返回值，该值表明了 map 中是否存在这个键。

	n := map[string]int{"foo": 1, "bar": 2} //声明初始化一个map
	fmt.Println("map:", n)

	/*
	   map map[k1:7 k2:13]
	   v1: 7
	   len: 2
	   map map[k1:7]
	   prs: false
	   map: map[bar:2 foo:1]

	*/

}

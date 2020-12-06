package main

import (
	"crypto/sha1" //多个 crypto/* 包中实现了一系列散列函数
	"fmt"
)

func main() {
	s := "sha1 this string"

	h := sha1.New()

	h.Write([]byte(s)) //如果是一个字符串， 需要使用 []byte(s) 将其强制转换成字节数组。

	bs := h.Sum(nil) //Sum 得到最终的散列值的字符切片。

	fmt.Println(s)
	fmt.Printf("%x\n", bs) //散列结果格式化为 16 进制字符串。
}

/*
sha1 this string
cf23df2207d99a74fbe169e3eba035e633b65d94
*/

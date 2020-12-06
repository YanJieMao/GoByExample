package main

//从字符串中解析数字
import (
	"fmt"
	"strconv" // strconv 包提供了数字解析能力。
)

func main() {

	f, _ := strconv.ParseFloat("1.234", 64) //使用 ParseFloat，这里的 64 表示解析的数的位数
	fmt.Println(f)

	i, _ := strconv.ParseInt("123", 0, 64) //ParseInt 解析整型数时， 例子中的参数 0 表示自动推断字符串所表示的数字的进制。 64 表示返回的整型数是以 64 位存储的。
	fmt.Println(i)

	d, _ := strconv.ParseInt("0x1c8", 0, 64) //ParseInt 会自动识别出字符串是十六进制数。
	fmt.Println(d)

	u, _ := strconv.ParseUint("789", 0, 64) //ParseUint 也是可用的。
	fmt.Println(u)

	k, _ := strconv.Atoi("135") //Atoi 是一个基础的 10 进制整型数转换函数。
	fmt.Println(k)

	_, e := strconv.Atoi("wat")
	fmt.Println(e)
}

/*
1.234
123
456
789
135
strconv.Atoi: parsing "wat": invalid syntax


*/

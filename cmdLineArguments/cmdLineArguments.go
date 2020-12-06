package main

//命令行参数
import (
	"fmt"
	"os"
)

func main() {

	argsWithProg := os.Args        //os.Args 提供原始命令行参数访问功能。
	argsWithoutProg := os.Args[1:] //切片中的第一个参数是该程序的路径， 而 os.Args[1:]保存了程序全部的参数。

	arg := os.Args[3] //可以使用标准的下标方式取得单个参数的值。

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
}

/*
$ go build command-line-arguments.go
$ ./command-line-arguments a b c d
[./command-line-arguments a b c d]
[a b c d]
c

*/

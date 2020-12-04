package main

import (
	"fmt"
	"os"
)

func main() {

	f := createFile("C:/Users/Troila/Desktop/defer.txt")
	defer closeFile(f) //在函数返回时才会被调用，确保文件被关闭
	writeFile(f)
}

func createFile(p string) *os.File { //新建文件
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) { //文件写入数据
	fmt.Println("writing")
	fmt.Fprintln(f, "data")

}

func closeFile(f *os.File) { //文件关闭
	fmt.Println("closing")
	err := f.Close()

	if err != nil { //错误检查
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

/*
creating
writing
closing
运行结果会在桌面创建一个defer.txt文件，里面被写入data
*/

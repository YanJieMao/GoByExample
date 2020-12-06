package main

import (
	"fmt"
	"path/filepath" //filepath 包为 文件路径 ，提供了方便的跨操作系统的解析和构建函数
	"strings"
)

func main() {

	p := filepath.Join("dir1", "dir2", "filename")
	fmt.Println("p:", p)

	fmt.Println(filepath.Join("dir1//", "filename"))
	fmt.Println(filepath.Join("dir1/../dir1", "filename"))

	fmt.Println("Dir(p):", filepath.Dir(p))
	fmt.Println("Base(p):", filepath.Base(p))

	fmt.Println(filepath.IsAbs("dir/file"))
	fmt.Println(filepath.IsAbs("/dir/file"))

	filename := "config.json"

	ext := filepath.Ext(filename) //某些文件名包含了扩展名（文件类型）。 我们可以用 Ext 将扩展名分割出来。
	fmt.Println(ext)

	fmt.Println(strings.TrimSuffix(filename, ext)) //获取文件名清除扩展名后的值，请使用 strings.TrmSuffix。

	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)

	rel, err = filepath.Rel("a/b", "a/c/t/file") //Rel 寻找 basepath 与 targpath 之间的相对路径。 如果相对路径不存在，则返回错误。
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)
}

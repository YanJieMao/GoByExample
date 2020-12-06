package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	err := os.Mkdir("subdir", 0755) //在当前工作目录下，创建一个子目录。
	check(err)

	defer os.RemoveAll("subdir") //使用 defer 删除这个目录。 os.RemoveAll 会删除整个目录（类似于 rm -rf）

	createEmptyFile := func(name string) { //个用于创建临时文件的帮助函数。
		d := []byte("")
		check(ioutil.WriteFile(name, d, 0644))
	}

	createEmptyFile("subdir/file1")

	err = os.MkdirAll("subdir/parent/child", 0755) //创建一个有层级的目录，使用 MkdirAll 函数，并包含其父目录。 这个类似于命令 mkdir -p。
	check(err)

	createEmptyFile("subdir/parent/file2")
	createEmptyFile("subdir/parent/file3")
	createEmptyFile("subdir/parent/child/file4")

	c, err := ioutil.ReadDir("subdir/parent") //ReadDir 列出目录的内容，返回一个 os.FileInfo 类型的切片对象。
	check(err)

	fmt.Println("Listing subdir/parent")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	err = os.Chdir("subdir/parent/child") //Chdir 可以修改当前工作目录，类似于 cd
	check(err)

	c, err = ioutil.ReadDir(".")
	check(err)

	fmt.Println("Listing subdir/parent/child")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	err = os.Chdir("../../..")
	check(err)

	fmt.Println("Visiting subdir")
	err = filepath.Walk("subdir", visit) //Walk 接受一个路径和回调函数，用于处理访问到的每个目录和文件。
}

func visit(p string, info os.FileInfo, err error) error { //filepath.Walk 遍历访问到每一个目录和文件后，都会调用 visit。
	if err != nil {
		return err
	}
	fmt.Println(" ", p, info.IsDir())
	return nil
}

/*
Listing subdir/parent
  child true
  file2 false
  file3 false
Listing subdir/parent/child
  file4 false
Visiting subdir
  subdir true
  subdir\file1 false
  subdir\parent true
  subdir\parent\child true
  subdir\parent\child\file4 false
  subdir\parent\file2 false
  subdir\parent\file3 false


*/

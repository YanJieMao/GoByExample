package main

//临时文件
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

	f, err := ioutil.TempFile("", "sample") //创建临时文件最简单的方法是调用 ioutil.TempFile 函数
	check(err)

	fmt.Println("Temp file name:", f.Name()) //文件名以 ioutil.TempFile 函数的第二个参数作为前缀， 剩余的部分会自动生成

	defer os.Remove(f.Name())

	_, err = f.Write([]byte{1, 2, 3, 4})
	check(err)

	dname, err := ioutil.TempDir("", "sampledir") //ioutil.TempDir 的参数与 TempFile 相同， 但是它返回的是一个 目录名 ，而不是一个打开的文件。
	fmt.Println("Temp dir name:", dname)

	defer os.RemoveAll(dname)

	fname := filepath.Join(dname, "file1") //通过拼接临时目录和临时文件合成完整的临时文件路径，并写入数据。
	err = ioutil.WriteFile(fname, []byte{1, 2}, 0666)
	check(err)
}

/*
Temp file name: C:\Users\Hao\AppData\Local\Temp\sample859665995
Temp dir name: C:\Users\Hao\AppData\Local\Temp\sampledir314593070

*/

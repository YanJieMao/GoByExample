package main

import "os"

func main() {
	panic("a problem")

	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err) ////当函数返回我们不知道如何处理（或不想处理）的错误值时，中止操作。
	}

}

/*
panic: a problem

goroutine 1 [running]:
main.main()
	c:/Users/Troila/go/base/GoByExample/panic/panic.go:6 +0x45
exit status 2

*/

package main

import (
	"fmt"
	"os"
)

func main() {

	defer fmt.Println("!") //os.Exit 时 defer 将不会 被执行， 所以这里的 fmt.Println 将永远不会被调用。

	os.Exit(3) //Go 不使用在 main 中返回一个整数来指明退出状态。 如果你想以非零状态退出，那么你就要使用 os.Exit。
}

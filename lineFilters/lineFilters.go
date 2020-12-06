package main

//行过滤器
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin) //用带缓冲的 scanner 包装无缓冲的 os.Stdin

	for scanner.Scan() {

		ucl := strings.ToUpper(scanner.Text()) //输出转换为大写后的行。

		fmt.Println(ucl)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

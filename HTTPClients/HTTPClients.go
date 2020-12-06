package main

import (
	"bufio"
	"fmt"
	"net/http" //Go 标准库的 net/http 包为 HTTP 客户端和服务端提供了出色的支持。
)

func main() {

	resp, err := http.Get("http://baidu.com") //向服务端发送一个 HTTP GET 请求。
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status) //打印 HTTP response 状态.

	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

/*
Response status: 200 OK
<html>
<meta http-equiv="refresh" content="0;url=http://www.baidu.com/">
</html>
*/

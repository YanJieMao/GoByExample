package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {

	s := "postgres://user:pass@host.com:5432/path?k=v#f" //解析这个 URL 示例，它包含了一个 scheme、 认证信息、主机名、端口、路径、查询参数以及查询片段

	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	fmt.Println(u.Scheme) //直接访问 scheme。

	fmt.Println(u.User)
	fmt.Println(u.User.Username()) //User 包含了所有的认证信息， 这里调用 Username 和 Password 来获取单独的值。
	p, _ := u.User.Password()
	fmt.Println(p)

	fmt.Println(u.Host)
	host, port, _ := net.SplitHostPort(u.Host) //Host 包含主机名和端口号。使用 SplitHostPort 提取它们。
	fmt.Println(host)
	fmt.Println(port)

	fmt.Println(u.Path)
	fmt.Println(u.Fragment) //提取路径和 # 后面的查询片段信息。

	fmt.Println(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery) // k=v 这种格式的查询参数，可以使用 RawQuery 函数。
	fmt.Println(m)
	fmt.Println(m["k"][0])
}

/*
postgres
user:pass
user
pass
host.com:5432
host.com
5432
/path
f
k=v
map[k:[v]]
v

*/

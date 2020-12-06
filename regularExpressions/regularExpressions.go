package main

//正则表达式.参考 regexp 包文档

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {

	match, _ := regexp.MatchString("p([a-z]+)ch", "peach") //检测字符串是否符合一个表达式
	fmt.Println(match)

	r, _ := regexp.Compile("p([a-z]+)ch") //通过 Compile 得到一个优化过的 Regexp 结构体。

	fmt.Println(r.MatchString("peach"))

	fmt.Println(r.FindString("peach punch")) //查找首次匹配的字符串， 但是它的返回值是，匹配开始和结束位置的索引，而不是匹配的内容

	fmt.Println(r.FindStringIndex("peach punch"))

	fmt.Println(r.FindStringSubmatch("peach punch"))

	fmt.Println(r.FindStringSubmatchIndex("peach punch"))

	fmt.Println(r.FindAllString("peach punch pinch", -1))

	fmt.Println(r.FindAllStringSubmatchIndex(
		"peach punch pinch", -1))

	fmt.Println(r.FindAllString("peach punch pinch", 2))

	fmt.Println(r.Match([]byte("peach")))

	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println(r)

	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}

/*
true
true
peach
[0 5]
[peach ea]
[0 5 1 3]
[peach punch pinch]
[[0 5 1 3] [6 11 7 9] [12 17 13 15]]
[peach punch]
true
p([a-z]+)ch
a <fruit>
a PEACH


*/

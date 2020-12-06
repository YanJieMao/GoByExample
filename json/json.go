package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type response1 struct {
	Page   int
	Fruits []string
}

type response2 struct {
	Page   int      `json:"page"`   //只有 可导出 的字段才会被 JSON 编码/解码。必须以大写字母开头的字段才是可导出的。
	Fruits []string `json:"fruits"` //给结构字段声明标签来自定义编码的 JSON 数据的键名。
}

func main() {

	bolB, _ := json.Marshal(true) //基本数据类型到 JSON 字符串的编码过程。
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	slcD := []string{"apple", "peach", "pear"} //一些切片和 map 编码成 JSON 数组和对象的例子
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	res1D := &response1{ //JSON 包可以自动的编码你的自定义类型。 编码的输出只包含可导出的字段，并且默认使用字段名作为 JSON 数据的键名。
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	res2D := &response2{ //给结构字段声明标签来自定义编码的 JSON 数据的键名。
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	byt := []byte(`{"num":6.13,"strs":["a","b"]}`) //将 JSON 数据解码为 Go 值的过程

	var dat map[string]interface{} //提供一个 JSON 包可以存放解码数据的变量

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	num := dat["num"].(float64) //访问嵌套的值需要一系列的转化
	fmt.Println(num)

	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)

	/*
		将 JSON 解码为自定义数据类型。 这样做的好处是，
		可以为我们的程序增加附加的类型安全性，
		并在访问解码后的数据时不需要类型断言。
	*/

	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)

	/*
		我们总是使用 byte和 string 作为数据和 JSON 表示形式之间的中介。
		当然，我们也可以像 os.Stdout 一样直接将 JSON 编码流传输到 os.Writer 甚至 HTTP 响应体。
	*/
}

/*
true
1
2.34
"gopher"
["apple","peach","pear"]
{"apple":5,"lettuce":7}
{"Page":1,"Fruits":["apple","peach","pear"]}
{"page":1,"fruits":["apple","peach","pear"]}
map[num:6.13 strs:[a b]]
6.13
a
{1 [apple peach]}
apple
{"apple":5,"lettuce":7}
*/

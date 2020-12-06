package main

//encoding.xml 包为 XML 和 类-XML 格式提供了内建支持。

import (
	"encoding/xml"
	"fmt"
)

type Plant struct { //Plant 结构将被映射到 XML
	XMLName xml.Name `xml:"plant"` //XMLName 字段名称规定了该结构的 XML 元素名称； id,attrr 表示 Id 字段是一个 XML 属性，而不是嵌套元素。
	Id      int      `xml:"id,attr"`
	Name    string   `xml:"name"`
	Origin  []string `xml:"origin"`
}

func (p Plant) String() string {
	return fmt.Sprintf("Plant id=%v, name=%v, origin=%v",
		p.Id, p.Name, p.Origin)
}

func main() {
	coffee := &Plant{Id: 27, Name: "Coffee"}
	coffee.Origin = []string{"Ethiopia", "Brazil"}

	out, _ := xml.MarshalIndent(coffee, " ", "  ") // 使用 MarshalIndent 生成可读性更好的输出结果。
	fmt.Println(string(out))

	fmt.Println(xml.Header + string(out)) //明确的为输出结果添加一个通用的 XML 头部信息。

	var p Plant
	if err := xml.Unmarshal(out, &p); err != nil { //Unmarshal 将 XML 格式字节流解析到 Plant 结构。
		panic(err)
	}
	fmt.Println(p)

	tomato := &Plant{Id: 81, Name: "Tomato"}
	tomato.Origin = []string{"Mexico", "California"}

	type Nesting struct {
		XMLName xml.Name `xml:"nesting"`
		Plants  []*Plant `xml:"parent>child>plant"` //parent>child>plant 字段标签告诉编码器，将 Plants 中的元素嵌套在 <parent><child> 里面。
	}

	nesting := &Nesting{}
	nesting.Plants = []*Plant{coffee, tomato}

	out, _ = xml.MarshalIndent(nesting, " ", "  ")
	fmt.Println(string(out))
}

/*
<plant id="27">
   <name>Coffee</name>
   <origin>Ethiopia</origin>
   <origin>Brazil</origin>
 </plant>
<?xml version="1.0" encoding="UTF-8"?>
 <plant id="27">
   <name>Coffee</name>
   <origin>Ethiopia</origin>
   <origin>Brazil</origin>
 </plant>
Plant id=27, name=Coffee, origin=[Ethiopia Brazil]
 <nesting>
   <parent>
     <child>
       <plant id="27">
         <name>Coffee</name>
         <origin>Ethiopia</origin>
         <origin>Brazil</origin>
       </plant>
       <plant id="81">
         <name>Tomato</name>
         <origin>Mexico</origin>
         <origin>California</origin>
       </plant>
     </child>
   </parent>
 </nesting>


*/

package main

import (
	b64 "encoding/base64" //引入了 encoding/base64 包， 并使用别名 b64 代替默认的 base64。
	"fmt"
)

func main() {

	data := "abc123!?$*&()'-=@~"

	sEnc := b64.StdEncoding.EncodeToString([]byte(data)) //同时支持标准 base64 以及 URL 兼容 base64。 这是使用标准编码器进行编码的方法。 编码器需要一个 []byte，因此我们将 string 转换为该类型。
	fmt.Println(sEnc)

	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec))
	fmt.Println()

	uEnc := b64.URLEncoding.EncodeToString([]byte(data)) //URL base64 格式进行编解码。
	fmt.Println(uEnc)
	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))

	/*
		标准 base64 编码和 URL base64 编码的 编码字符串存在稍许不同（后缀为 + 和 -）， 但是两者都可以正确解码为原始字符串。
	*/
}

/*
YWJjMTIzIT8kKiYoKSctPUB+
abc123!?$*&()'-=@~

YWJjMTIzIT8kKiYoKSctPUB-
abc123!?$*&()'-=@~

*/

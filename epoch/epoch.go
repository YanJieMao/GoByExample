package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now() //分别使用 time.Now 的 Unix 和 UnixNano， 来获取从 Unix 纪元起，到现在经过的秒数和纳秒数。
	secs := now.Unix()
	nanos := now.UnixNano()
	fmt.Println(now)

	millis := nanos / 1000000 //UnixMillis 是不存在的，所以要得到毫秒数的话， 你需要手动的从纳秒转化一下。
	fmt.Println(secs)
	fmt.Println(millis)
	fmt.Println(nanos)

	fmt.Println(time.Unix(secs, 0)) //Unix 纪元起的整数秒或者纳秒转化到相应的时间。
	fmt.Println(time.Unix(0, nanos))
}

/*
2020-12-06 18:39:07.6409069 +0800 CST m=+0.001952401
1607251147
1607251147640
1607251147640906900
2020-12-06 18:39:07 +0800 CST
2020-12-06 18:39:07.6409069 +0800 CST
*/

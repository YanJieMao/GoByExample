package main

//有关时间的例子
import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println //输出

	now := time.Now()
	p(now)

	then := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p(then)

	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())

	p(then.Weekday())

	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	diff := now.Sub(then)
	p(diff)

	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	p(then.Add(diff))
	p(then.Add(-diff))
}

/*
2020-12-04 15:31:53.2294186 +0800 CST m=+0.002992001
2009-11-17 20:34:58.651387237 +0000 UTC
2009
November
17
20
34
58
651387237
UTC
Tuesday
true
false
false
96826h56m54.578031363s
96826.9484938976
5.809616909633856e+06
3.4857701457803136e+08
348577014578031363
2020-12-04 07:31:53.2294186 +0000 UTC
1998-11-01 09:38:04.073355874 +0000 UTC

*/

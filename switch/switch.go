package main

import (
	"fmt"
	"time"
)

func main() {

	i := 2
	fmt.Print(i, "写做")
	switch i {
	case 1:
		fmt.Println("一")
	case 2:
		fmt.Println("二")

	case 3:
		fmt.Println("三")

	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("周末啦")
	default:
		fmt.Println("打工人的工作日")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("上午")
	default:
		fmt.Println("下午")

	}

	whatAmI := func(i interface{}) { //类型开关 (type switch) 比较类型而非值。
		switch t := i.(type) {
		case bool:
			fmt.Println("我是bool")
		case int:
			fmt.Println("我是int")

		default:
			fmt.Printf("不知道啥类型%T\n", t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}

/*
2写做二
打工人的工作日
上午
我是bool
我是int
不知道啥类型string
*/

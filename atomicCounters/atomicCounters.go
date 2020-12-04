package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

//原子计数器

func main() {

	var ops uint64 //无符号整型变量

	var wg sync.WaitGroup

	for i := 0; i < 50; i++ { //启动50个协程
		wg.Add(1)

		go func() {
			for c := 0; c < 1000; c++ {

				atomic.AddUint64(&ops, 1) //AddUint64 来让计数器自动增加
			}
			wg.Done()
		}()

		//fmt.Println(atomic.LoadUint64(&ops))  //用它来读取原子更新时的数据，协程并发执行，这个值并不是有序增加的
	}

	wg.Wait()
	fmt.Println("ops:", ops)

}

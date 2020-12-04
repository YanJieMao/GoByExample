package main

//互斥锁，在协程间安全地访问数据
import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func main() {

	var state = make(map[int]int)

	var mutex = &sync.Mutex{} //同步对state的访问

	var readOps uint64
	var writeOps uint64

	for r := 0; r < 100; r++ { //100个协程，重复读取state
		go func() {
			total := 0
			for {

				key := rand.Intn(5)
				mutex.Lock() //加锁
				total += state[key]
				mutex.Unlock() //解锁
				atomic.AddUint64(&readOps, 1)

				time.Sleep(time.Millisecond)
			}
		}()
	}

	for w := 0; w < 10; w++ { //10个协程，持续写入
		go func() {
			for {
				key := rand.Intn(5)
				val := rand.Intn(100)
				mutex.Lock()
				state[key] = val
				mutex.Unlock()
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)

	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)

	mutex.Lock()
	fmt.Println("state:", state)
	mutex.Unlock()
}

/*
readOps: 64692
writeOps: 6470
state: map[0:92 1:12 2:95 3:16 4:40]

*/

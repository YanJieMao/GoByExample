package main

import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("接收job", j)
			} else {
				fmt.Println("接收所有job")
				done <- true
				return

			}

		}

	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("发送job", j)
	}

	close(jobs) //关闭jobs
	fmt.Println("发送所有job")

	<-done //通道同步等待任务结束

}

/*
发送job 1
接收job 1
接收job 2
发送job 2
发送job 3
发送所有job
接收job 3
接收所有job
*/

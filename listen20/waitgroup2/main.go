package main

import (
	"fmt"
	"sync"
	"time"
)

func process(i int, wg *sync.WaitGroup) {
	fmt.Println("started Goroutine ", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("Goroutine %d ended\n", i)
	wg.Done()  //释放一个计数器
}
func main() {
	no := 3
	var wg sync.WaitGroup  //创建一个waitgroup
	fmt.Println("wait return")
	for i := 0; i < no; i++ {
		wg.Add(1) //添加WaitGroup计数器
		go process(i, &wg)
	}
	wg.Wait()  //阻塞，知道WaitGroup计数器等于0
	fmt.Println("All go routines finished executing")
}

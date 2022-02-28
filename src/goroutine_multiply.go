package main

import (
	"fmt"
	"sync"
)

//使用sync.WaitGroup同步线程池来实现goroutine线程的同步
var wg sync.WaitGroup

func hello2(i int) {
	//线程做完结束就登记-1
	defer wg.Done()
	fmt.Println("你好 Go线程!——", i)
}

func main() {
	for i := 0; i < 10; i++ {
		//启动一个线程就登记+1
		wg.Add(1)
		go hello2(i)
	}
	//等待所有的登记线程都结束
	wg.Wait()
}

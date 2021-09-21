package main

import (
	"fmt"
	"runtime"
	"time"
)

func a() {
	for i := 1; i < 10; i++ {
		fmt.Println("A:", i)
	}
}

func b() {
	for i := 1; i < 10; i++ {
		fmt.Println("B:", i)
	}
}

func main() {
	//用GOMAXPROCS参数来确定需要使用多少个OS线程来同时执行Go代码，默认值是机器上的CPU核心数
	runtime.GOMAXPROCS(1)
	//两个任务只有一个逻辑核心，此时是做完一个任务再做另一个任务, 这里先做B再做A
	go a()
	go b()
	time.Sleep(time.Second)
}

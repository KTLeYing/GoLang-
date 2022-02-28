package main

import (
	"fmt"
	"runtime"
	"time"
)

func a1() {
	for i := 1; i < 10; i++ {
		fmt.Println("A:", i)
	}
}

func b1() {
	for i := 1; i < 10; i++ {
		fmt.Println("B:", i)
	}
}

func c() {
	for i := 0; i < 10; i++ {
		fmt.Println("C", i)
	}
}

func d() {
	for i := 0; i < 10; i++ {
		fmt.Println("D", i)
	}
}

func main() {
	//用GOMAXPROCS参数来确定需要使用多少个OS线程来同时执行Go代码，默认值是机器上的CPU核心数
	runtime.GOMAXPROCS(4)
	//将逻辑核心数设为2，此时两个任务并行执行，代码如下。, 这里B和A同时进行
	go a1()
	go b1()
	go c()
	go d()
	time.Sleep(time.Second)
}

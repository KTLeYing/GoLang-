package main

import "fmt"

func main() {
	ch := make(chan int, 5)
	for i := 4; i >= 0; i-- {
		ch <- i
		//ret := <- ch
		//fmt.Println(ret)
	}
	//channel通道先进先出
	for i := 0; i < 5; i++ {
		//一个个输出ch
		ret := <-ch
		fmt.Println(ret)
	}

	fmt.Println("发送成功")
}

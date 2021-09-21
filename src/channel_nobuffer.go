package main

import "fmt"

func recv(c chan int) {
	ret := <-c
	fmt.Println("接收成功", ret)
}

func main() {
	ch := make(chan int)
	//开启一个线程来接收数据
	go recv(ch)
	//发送数据
	ch <- 11
	fmt.Println("发送成功")
}

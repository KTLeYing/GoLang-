package main

import "fmt"

func hello1() {
	fmt.Println("第 1 步...")
}

func main() {
	//启动一个go线程去执行hello()函数
	go hello1()
	fmt.Println("第 2 步...")
}

//这一次的执行结果只打印了 第 2 步...，并没有打印 第 1 步...。为什么呢？
//【在程序启动时，Go程序就会为main()函数创建一个默认的goroutine】
//当main()函数返回的时候该goroutine就结束了，所有在main()函数中启动的goroutine会一同结束，main函数所在的goroutine就像是权利的游戏中的夜王，其他的goroutine都是异鬼，夜王一死它转化的那些异鬼也就全部GG了。

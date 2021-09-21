package main

import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	//开启go线程将0-100的数发送到ch1中
	go func() {
		for i := 0; i < 100; i++ {
			ch1 <- i
		}
		close(ch1)
	}()
	//开启一个go线程从ch1中接收值，并将该值的平方发送到ch2中
	go func() {
		for {
			//通道关闭后再取值 ok = false //判断通道通道关闭方式1
			i, ok := <-ch1
			if !ok {
				break
			}
			//将该数的平方发送到ch2通道
			ch2 <- i * i
		}
		close(ch2)
	}()
	//在go线程中从ch2中接收值打印   //判断通道关闭方式2
	for i := range ch2 {
		//通道关闭后会退出for range循环来获取值
		fmt.Println(i)
	}
}

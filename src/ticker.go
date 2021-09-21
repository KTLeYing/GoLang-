package main

import (
	"fmt"
	"time"
)

//时间到了，多次执行
func main() {
	// 1. 获取ticker对象
	ticker := time.NewTimer(5 * time.Second)
	i := 0
	//子协程
	go func() {
		for {
			//<- ticker.C
			i++
			fmt.Println(<-ticker.C)
			if i == 5 {
				//停止
				ticker.Stop()
			}
		}
	}()
	for {
	}
}

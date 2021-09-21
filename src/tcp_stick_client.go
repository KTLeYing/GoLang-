package main

import (
	"fmt"
	"net"
)

func main() {
	//建立连接
	conn, err := net.Dial("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("建立连接失败, 错误：", err)
		return
	}
	//延迟关闭连接
	defer conn.Close()
	for i := 0; i < 20; i++ {
		msg := "你好，服务器，我来找你聊天了!!!"
		conn.Write([]byte(msg))
	}
}

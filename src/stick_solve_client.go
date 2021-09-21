package main

import (
	"fmt"
	"net"
)

//粘包解决——客户端
func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("建立连接失败， 错误：", err)
		return
	}
	//延迟关闭连接
	defer conn.Close()
	//发送消息给服务器
	for i := 0; i < 20; i++ {
		msg := "你好，服务器，我来找你聊天了!!!"
		data, err := Encode(msg)
		if err != nil {
			fmt.Println("消息编码失败， 错误：", err)
			return
		}
		//发送一条消息
		conn.Write(data)
	}
}

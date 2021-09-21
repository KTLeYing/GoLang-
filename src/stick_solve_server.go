package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		//解码读取到的消息
		msg, err := Decode(reader)
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Println("解码消息失败， 错误：", err)
			return
		}
		fmt.Println("收到客户端的数据：", msg)
	}
}

//粘包解决——服务器
func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("建立连接失败， 错误：", err)
		return
	}
	//延迟关闭连接
	defer listen.Close()
	//接收客户端消息
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("接收消息客户端失败, 错误：", err)
			continue
		}
		//开启一个go协程处理消息
		go process(conn)
	}
}

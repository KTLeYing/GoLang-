package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

//处理过程
func process2(conn net.Conn) {
	//关闭连接
	defer conn.Close()
	readre := bufio.NewReader(conn)
	var buf [1024]byte
	for {
		//读取客户端数据
		n, err := readre.Read(buf[:])
		//如果是读到消息末尾
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("读取客户端数据失败， 错误：", err)
			break
		}
		recvStr := string(buf[:n])
		fmt.Println("收到客户端的数据：", recvStr)
	}
}

//服务器
func main() {
	//监听服务器并建立连接
	listen, err := net.Listen("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("建立连接失败，错误：", err)
		return
	}
	//关闭连接(延迟)
	defer listen.Close()
	//监听接收数据
	for {
		//每次接收到客户端的数据
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("接收数据失败， 错误：", err)
			continue
		}
		//开启一个线程来处理消息数据
		go process2(conn)
	}
}

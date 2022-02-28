package main

import (
	"bufio"
	"fmt"
	"net"
)

// 处理函数
func process1(conn net.Conn) {
	defer conn.Close()
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		//读取数据
		n, err := reader.Read(buf[:])
		if err != nil {
			fmt.Println("read from client fail,", err)
			break
		}
		receiveStr := string(buf[:n])
		fmt.Println("收到客户端的数据：", receiveStr)
		//服务器向客户端发送数据
		conn.Write([]byte(receiveStr))
	}
}

func main() {
	listen, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		fmt.Println("listen fail, ", err)
		return
	}
	for {
		//建立连接
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept fail,", err)
			continue
		}
		//启动一个goroutine来启动线程
		go process1(conn)
	}
}

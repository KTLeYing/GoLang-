package main

import (
	"fmt"
	"net"
)

//UDP客户端
func main() {
	//建立连接
	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 8888,
	})
	if err != nil {
		fmt.Println("连接服务器失败， err:", err)
		return
	}
	//延迟关闭连接
	defer socket.Close()
	sendData := []byte("你好， 服务器， 我给你发消息来了！！！")
	//发送数据给服务器
	_, err = socket.Write(sendData)
	if err != nil {
		fmt.Println("发送数据失败， err:", err)
		return
	}
	data := make([]byte, 4096)
	//监听和接收服务器的消息（即刚刚客户端自己发送的数据）
	n, remoteAddr, err := socket.ReadFromUDP(data)
	if err != nil {
		fmt.Println("接收数据失败， err:", err)
		return
	}
	fmt.Printf("【客户端接收到的数据】data:%v \t addr:%v \t count:%v\n", string(data[:n]), remoteAddr, n)
}

package main

import (
	"fmt"
	"net"
)

//UPD服务端
func main() {
	//建立连接
	listen, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 8888,
	})
	if err != nil {
		fmt.Println("监听失败, err: ", err)
		return
	}
	//关闭连接（延迟）
	defer listen.Close()
	for {
		var data [1024]byte
		//接收客户端数据
		n, addr, err := listen.ReadFromUDP(data[:])
		if err != nil {
			fmt.Println("接收数据失败, err: ", err)
			continue
		}
		fmt.Printf("【服务器接收到的数据】data:%v \t addr:%v \t cout:%v\n", string(data[:n]), addr, n)
		//发送数据给客户端
		sendData := []byte("你好，客户端， 我已经成功收到你的消息!!!")
		_, err = listen.WriteToUDP(sendData, addr)
		if err != nil {
			fmt.Println("发送数据失败，err:", err)
			continue
		}
	}

}

package main

import (
	"fmt"
	"net/http"
)

func main() {
	//res, _ := http.Get("http://localhost:8000/go")
	res, _ := http.Get("http://www.baidu.com")
	fmt.Println(res)
	defer res.Body.Close()
	//200 OK
	fmt.Println(res.Status)
	fmt.Println(res.Header)

	buf := make([]byte, 1024)
	for {
		//接收服务器消息
		n, err := res.Body.Read(buf)
		if err != nil {
			fmt.Println(err)
			return
		} else {
			fmt.Println("读取完毕...")
			result := string(buf[:n])
			fmt.Println(result)
			break
		}
	}

}

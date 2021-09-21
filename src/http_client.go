package main

import (
	"fmt"
	"net/http"
)

func main() {
	//单独写回调函数
	http.HandleFunc("/go", myHandler)
	//监听服务器和建立连接
	http.ListenAndServe("127.0.0.1:8000", nil)
}

//handler回调函数
func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.RemoteAddr, "连接成功")
	//请求方式：GET POST  DELETE  PUT  UPDATE
	fmt.Println("url:", r.URL.Path)
	fmt.Println("header: ", r.Header)
	fmt.Println("body: ", r.Body)
	//回复
	w.Write([]byte("www.5lmh.com"))
}

package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	fmt.Println("client start...")
	time.Sleep(1 * time.Second)
	//1、连接远程服务器，得到连接
	conn, err := net.Dial("tcp", "127.0.0.1:8999")
	if err != nil {
		fmt.Println("client start err,exit!")
		return
	}
	//2、调用write方法，写数据
	for {
		_, err := conn.Write([]byte("hello world"))
		if err != nil {
			return
		}
		buf := make([]byte, 512)
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read err ", err)
			return
		}
		fmt.Printf("server call back = %s,cnt=%d\n", buf, cnt)
		time.Sleep(1*time.Second)
	}
}

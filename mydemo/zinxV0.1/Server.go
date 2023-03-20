package main

import "zinx/znet"

func main() {
	//1、创建server
	s:=znet.NewServer("zinxV0.1")
	s.Start()
	//2、启动server
	s.Serve()
}

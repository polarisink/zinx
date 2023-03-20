package main

import (
	"fmt"
	"zinx/ziface"
	"zinx/znet"
)

type PingRouter struct {
	znet.BaseRouter
}

func (r *PingRouter) PreHandle(request ziface.IRequest) {
	fmt.Println("call router preHandle")
	_, err := request.GetConn().GetTcpConnection().Write([]byte("before ping...\n"))
	if err != nil {
		fmt.Println("call back before ping error")
		return
	}
}
func (r *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("call router handle")
	_, err := request.GetConn().GetTcpConnection().Write([]byte("ping... ping...\n"))
	if err != nil {
		fmt.Println("call back ping... ping... error")
		return
	}
}
func (r *PingRouter) PostHandle(request ziface.IRequest) {
	fmt.Println("call router postHandle")
	_, err := request.GetConn().GetTcpConnection().Write([]byte("before ping...\n"))
	if err != nil {
		fmt.Println("call back after ping error")
		return
	}
}

func main() {
	//1、创建server
	s := znet.NewServer("zinxV0.3")
	//添加router
	s.AddRouter(&PingRouter{})
	//启动server
	s.Serve()
}

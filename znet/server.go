package znet

import (
	"fmt"
	"net"
	"zinx/ziface"
)

type Server struct {
	Name, IpVersion, Ip string
	Port                int
	Router              ziface.IRouter
}

func (s *Server) AddRouter(router ziface.IRouter) {
	s.Router = router
	fmt.Println("add router success!")
}

func (s *Server) Start() {
	//1、获取一个tcp的addr
	fmt.Printf("[Start] Server Listenner at Ip: %s, Port: %d , is starting\n", s.Ip, s.Port)
	addr, err := net.ResolveTCPAddr(s.IpVersion, fmt.Sprintf("%s:%d", s.Ip, s.Port))
	if err != nil {
		fmt.Println("resolve tcp addr error: ", err)
		return
	}
	//2、监听服务器地址
	listen, err := net.ListenTCP(s.IpVersion, addr)
	if err != nil {
		fmt.Println("listen ", s.IpVersion, " error ", err)
		return
	}
	fmt.Println("start zinx server success,", s.Name, " listening ")
	var cid uint32
	cid = 0

	//3、阻塞等待服务器链接，处理客户端连接业务
	for {
		conn, err := listen.AcceptTCP()
		if err != nil {
			fmt.Println("Accept err", err)
			continue
		}

		//将处理新连接的业务方法和conn进行绑定得到连接模块
		dealConn := NewConnection(conn, cid, s.Router)
		cid++
		go dealConn.Start()
	}
}

func (s *Server) Stop() {
	//TODO 关闭服务器资源
}

func (s *Server) Serve() {
	s.Start()

	//TODO 其他初始化操作
	select {}
}

func NewServer(name string) ziface.IServer {
	s := &Server{
		Name:      name,
		IpVersion: "tcp4",
		Ip:        "0.0.0.0",
		Port:      8080,
		Router:    nil,
	}
	return s
}

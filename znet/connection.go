package znet

import (
	"fmt"
	"net"
	"zinx/ziface"
)

// @author lqs
// @date 2023/3/15 10:06 PM

type Connection struct {
	//当前连接的tcp套接字
	Conn *net.TCPConn
	//连接id
	ConnId uint32
	//是否关闭
	isClosed bool
	//告知当前连接一家退出停止channel
	ExitChan chan bool
	//该连接处理的方法router
	Router ziface.IRouter
}

func (c *Connection) Start() {
	fmt.Println("conn start... connId = ", c.ConnId)
	//TODO 启动当前连接写数据的业务,MaxConn
	go c.StartReader()
}

func (c *Connection) Stop() {
	fmt.Println("conn stop,,,connId = ", c.ConnId)

	if c.isClosed {
		return
	}
	c.isClosed = true
	c.Conn.Close()
	close(c.ExitChan)
}

func (c *Connection) GetTcpConnection() *net.TCPConn {
	return c.Conn
}

func (c *Connection) GetConnId() uint32 {
	return c.ConnId
}

func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

func (c *Connection) Send(data []byte) error {
	//TODO implement me
	panic("implement me")
}

func (c *Connection) StartReader() {
	fmt.Println("reader goroutine is running...")
	defer fmt.Println("connId = ", c.ConnId, " reader is exit, remote addr is ", c.RemoteAddr().String())
	defer c.Stop()
	for {
		//buf := make([]byte, utils.GlobalObject.MaxPackageSize)
		//_, err := c.Conn.Read(buf)
		//if err != nil {
		//	continue
		//}

		//创建拆包解包对象
		dp := NewDataPack()
		//读取客户端msg head二进制流8个字节

		//拆包，得到msgId和msgDataLen放在msg消息中

		//根据dataLen，再次读取data，放在msg的data中
		req := Request{
			conn: c,
			data: buf,
		}
		//执行注册的路由方法
		go func(request ziface.IRequest) {
			c.Router.PreHandle(request)
			c.Router.Handle(request)
			c.Router.PostHandle(request)
		}(&req)
	}
}

// NewConnection 新建连接
func NewConnection(conn *net.TCPConn, connId uint32, router ziface.IRouter) *Connection {
	return &Connection{
		Conn:     conn,
		ConnId:   connId,
		Router:   router,
		isClosed: false,
		ExitChan: make(chan bool, 1),
	}
}

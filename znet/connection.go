package znet

import (
	"net"
	"zinx/ziface"
)

/**
 * @author lqs
 * @date 2023/3/15 10:06 PM
 */

type Connection struct {
	//当前连接的tcp套接字
	Conn *net.TCPConn
	//连接id
	ConnId uint32
	//是否关闭
	isClosed bool
	//连接的业务方法
	handleApi ziface.HandleFunc
	//告知当前连接一家退出停止channel
	ExitChan chan bool
}

func (c *Connection) Start() {
	//TODO implement me
	panic("implement me")
}

func (c *Connection) Stop() {
	//TODO implement me
	panic("implement me")
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

// NewConnection 新建连接
func NewConnection(conn *net.TCPConn, connId uint32, callBackApi ziface.HandleFunc) *Connection {
	return &Connection{
		Conn:      conn,
		ConnId:    connId,
		handleApi: callBackApi,
		isClosed:  false,
		ExitChan:  make(chan bool, 1),
	}
}

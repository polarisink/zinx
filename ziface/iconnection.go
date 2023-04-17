package ziface

import "net"

type IConnection interface {
	// Start 启动连接，让当前连接准备开始工作
	Start()
	// Stop 停止连接，结束当前工作
	Stop()
	// GetTcpConnection 获取当前点链接的绑定socket conn
	GetTcpConnection() *net.TCPConn
	// GetConnId 获取当前连接模块的连接id
	GetConnId() uint32
	// RemoteAddr 获取远程客户端的tcp状态和ip
	RemoteAddr() net.Addr
	// Send 发送数据给远程客户端
	Send(data []byte) error
	// SetProperty 设置连接属性
	SetProperty(string, interface{})
	// GetProperty 获取属性
	GetProperty(string) (interface{}, error)
	RemoveProperty(string)
}

// HandleFunc 定义一个处理业务连接的方法
type HandleFunc func(*net.TCPConn, []byte, int) error

package znet

import (
	"errors"
	"fmt"
	"io"
	"net"
	"sync"
	"zinx/utils"
	"zinx/ziface"
)

// @author lqs
// @date 2023/3/15 10:06 PM

type Connection struct {
	TcpServer ziface.IServer
	//当前连接的tcp套接字
	Conn *net.TCPConn
	//连接id
	ConnId uint32
	//是否关闭
	isClosed bool
	//告知当前连接一家退出停止channel
	ExitChan chan bool
	//无缓冲通道，用于读写goroutine之间的消息通信
	msgChan chan []byte
	//消息的管理msgId和对于的业务处理api
	MsgHandler ziface.IMsgHandle
	//连接树形集合
	property     map[string]interface{}
	propertyLock sync.RWMutex
}

func (c *Connection) SetProperty(key string, value interface{}) {
	c.propertyLock.Lock()
	defer c.propertyLock.Unlock()
	c.property[key] = value
}

func (c *Connection) GetProperty(s string) (interface{}, error) {
	c.propertyLock.RLock()
	defer c.propertyLock.RUnlock()
	if value, ok := c.property[s]; ok {
		return value, nil
	}
	return nil, errors.New("no property found")
}

func (c *Connection) RemoveProperty(s string) {
	c.propertyLock.Lock()
	defer c.propertyLock.Unlock()
	delete(c.property, s)
}

func (c *Connection) Start() {
	fmt.Println("conn start... connId = ", c.ConnId)
	//TODO 启动当前连接写数据的业务,MaxConn
	go c.StartReader()
	go c.StartWriter()
	//执行hook函数
	c.TcpServer.CallOnConnStart(c)
}

func (c *Connection) Stop() {
	fmt.Println("conn stop,,,connId = ", c.ConnId)

	if c.isClosed {
		return
	}
	c.isClosed = true
	c.Conn.Close()
	c.ExitChan <- true
	c.TcpServer.Stop()
	c.TcpServer.CallOnConnStop(c)
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
		headData := make([]byte, dp.GetHeadLen())
		if _, err := io.ReadFull(c.GetTcpConnection(), headData); err != nil {
			fmt.Println("read msg head error", err)
			break
		}
		msg, err := dp.Unpack(headData)
		if err != nil {
			fmt.Println("unpack error: ", err)
			break
		}
		var data []byte
		if msg.GetMsgLen() > 0 {
			data = make([]byte, msg.GetMsgLen())
			if _, err := io.ReadFull(c.GetTcpConnection(), data); err != nil {
				fmt.Println("read msg data error: ", err)
				break
			}
		}
		msg.SetData(data)
		//读取客户端msg head二进制流8个字节

		//拆包，得到msgId和msgDataLen放在msg消息中
		//根据dataLen，再次读取data，放在msg的data中
		req := Request{
			conn: c,
			msg:  msg,
		}
		//执行注册的路由方法
		//已经开启工作池，直接发送
		if utils.GlobalObject.WorkerPoolSize > 0 {
			c.MsgHandler.SendMsgToTaskQueue(&req)
		} else {
			c.MsgHandler.DoMsgHandler(&req)
		}
	}
}

// SendMsg 提供一个sendMsg方法，将要发送给客户的数据，先进行封包，再发送
func (c *Connection) SendMsg(msgId uint32, data []byte) error {
	if c.isClosed {
		return fmt.Errorf("connection closed when send msg")
	}
	dp := NewDataPack()
	binaryMsg, err := dp.Pack(NewMsgPackage(msgId, data))
	if err != nil {
		fmt.Println("pack err msg id = ", msgId)
		return errors.New("pack err msg")
	}
	if _, err := c.Conn.Write(binaryMsg); err != nil {
		fmt.Println("write msg id, ", msgId, " error: ", err)
		return errors.New("conn write error")
	}
	if err != nil {
		return err
	}
	//将data进行封包
	return nil
}

// StartWriter 写消息goroutine,专门发送给客户端消息的模块
func (c *Connection) StartWriter() {
	fmt.Println("writer goroutine is running")
	defer fmt.Println(c.RemoteAddr().String(), "conn writer exit!")
	for {
		select {
		case data := <-c.msgChan:
			if _, err := c.Conn.Write(data); err != nil {
				fmt.Println("write data err: ", err)
				return
			}
		case <-c.ExitChan:
			//reader退出，writer也退出
			return
		}
	}
}

// NewConnection 新建连接
func NewConnection(server ziface.IServer, conn *net.TCPConn, connId uint32, router ziface.IMsgHandle) *Connection {
	c := &Connection{
		TcpServer:  server,
		Conn:       conn,
		ConnId:     connId,
		MsgHandler: router,
		isClosed:   false,
		ExitChan:   make(chan bool, 1),
		msgChan:    make(chan []byte),
	}
	c.TcpServer.GetConnMgr().Add(c)
	return c
}

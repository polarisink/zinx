package znet

import (
	"fmt"
	"io"
	"net"
	"testing"
)

// @author lqs
// @date 2023/3/20 8:21 PM

// 测试dataPack拆包和分包
func TestDataPack(t *testing.T) {

	//模拟创建服务器
	listener, err := net.Listen("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("server listen err: ", err)
		return
	}

	//创建go函数负责从客户端处理业务
	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				continue
			}
			go func(conn net.Conn) {
				dp := NewDataPack()
				for {
					headData := make([]byte, dp.GetHeadLen())
					_, err := io.ReadFull(conn, headData)
					if err != nil {
						fmt.Println("read head error")
						return
					}
					msgHead, err := dp.Unpack(headData)
					if err != nil {
						fmt.Println("server unpack err ", err)
						return
					}
					if msgHead.GetMsgLen() > 0 {
						msg := msgHead.(*Message)
						msg.Data = make([]byte, msg.GetMsgLen())
						_, err := io.ReadFull(conn, msg.Data)
						if err != nil {
							return
						}
						//完整的一个消息已经读取完毕
						fmt.Println("---> receive msgId: ", msg.Id, " , dataLen = ", msg.DataLen, " data = ", msg.Data)
					}
				}
			}(conn)
		}
	}()

	//模拟客户端
	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("client dial err ", err)
		return
	}

	//创建一个分包对象dp
	dp := NewDataPack()
	//模拟粘包，两个msg一起发
	msg1 := &Message{
		Id:      1,
		DataLen: 5,
		Data:    []byte{'h', 'e', 'l', 'l', 'o'},
	}
	msg2 := &Message{
		Id:      2,
		DataLen: 5,
		Data:    []byte{'w', 'o', 'r', 'l', 'd'},
	}
	sendData1, err := dp.Pack(msg1)
	if err != nil {
		fmt.Println("send msg1 error", err)
		return
	}
	sendData2, err := dp.Pack(msg2)
	if err != nil {
		fmt.Println("send msg2 error", err)
		return
	}
	sendData1 = append(sendData1, sendData2...)
	conn.Write(sendData1)

	//客户端阻塞
	select {}
}

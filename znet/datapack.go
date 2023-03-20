package znet

import (
	bytes "bytes"
	"encoding/binary"
	"fmt"
	"zinx/utils"
	"zinx/ziface"
)

// @author lqs
// @date 2023/3/16 10:06 PM

type DataPack struct {
}

func NewDataPack() *DataPack {
	return &DataPack{}
}

func (d *DataPack) GetHeadLen() uint32 {
	//dataLen + id
	return 8
}

func (d *DataPack) Pack(message ziface.IMessage) ([]byte, error) {
	//创建一个缓冲
	buf := bytes.NewBuffer([]byte{})
	//将dataLen写进buf
	if err := writeData(buf, message.GetDataLen()); err != nil {
		return nil, err
	}
	//将msgId写进buf
	if err := writeData(buf, message.GetMsgId()); err != nil {
		return nil, err
	}
	//将data写入buf
	if err := writeData(buf, message.GetData()); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (d *DataPack) Unpack(bytesData []byte) (ziface.IMessage, error) {
	buf := bytes.NewReader(bytesData)
	msg := &Message{}
	if err := binary.Read(buf, binary.LittleEndian, &msg.DataLen); err != nil {
		return nil, err
	}
	if err := binary.Read(buf, binary.LittleEndian, &msg.Id); err != nil {
		return nil, err
	}
	//先对长度判断
	if utils.GlobalObject.MaxPackageSize > 0 && msg.DataLen > utils.GlobalObject.MaxPackageSize {
		return nil, fmt.Errorf("too large message data received")
	}
	/*if err := binary.Read(buf, binary.LittleEndian, &msg.Data); err != nil {
		return nil, err
	}*/
	return msg, nil
}

func writeData(buf *bytes.Buffer, any interface{}) error {
	err := binary.Write(buf, binary.LittleEndian, any)
	if err != nil {
		return err
	}
	return nil
}

package znet

import (
	"bytes"
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
	bytes.NewBuffer([]byte{})
	//TODO implement me
	panic("implement me")
}

func (d *DataPack) Unpack(bytes []byte) (ziface.IMessage, error) {
	//TODO implement me
	panic("implement me")
}

package znet

import "zinx/ziface"

// @author lqs
// @date 2023/3/16 8:42 PM

type Request struct {
	conn ziface.IConnection
	data []byte
}

func (r *Request) GetConn() ziface.IConnection {
	return r.conn
}

func (r *Request) GetData() []byte {
	return r.data
}

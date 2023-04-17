package ziface

// @author lqs
// @date 2023/4/16 4:27 PM

type IConnManager interface {
	Add(IConnection)
	Remove(IConnection)
	Get(connId uint32) (IConnection, error)
	Len() int
	ClearConn()
}

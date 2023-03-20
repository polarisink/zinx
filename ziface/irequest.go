package ziface

// @author lqs
// @date 2023/3/16 8:40 PM

type IRequest interface {
	GetConn() IConnection

	GetData() []byte
}

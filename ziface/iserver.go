package ziface

// IServer 服务器接口
type IServer interface {
	// Start 启动服务器
	Start()
	// Stop 停止服务器
	Stop()
	// Serve 运行服务器
	Serve()
	// AddRouter 路由功能
	AddRouter(uint32, IRouter)

	GetConnMgr() IConnManager

	SetOnConnStart(func(connection IConnection))

	SetOnConnStop(func(connection IConnection))

	CallOnConnStart(connection IConnection)

	CallOnConnStop(connection IConnection)
}

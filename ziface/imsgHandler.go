package ziface

// @author lqs
// @date 2023/3/28 10:39 PM

type IMsgHandle interface {
	// DoMsgHandler 调度对于的router处理方法
	DoMsgHandler(IRequest)

	// AddRouter 为消息添加router
	AddRouter(uint32, IRouter)

	// StartWorkerPool 启动worker工作池子
	StartWorkerPool()

	SendMsgToTaskQueue(IRequest)
}

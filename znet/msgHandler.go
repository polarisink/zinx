package znet

import (
	"fmt"
	"zinx/utils"
	"zinx/ziface"
)

// @author lqs
// @date 2023/3/29 6:51 PM

type MsgHandle struct {
	Apis map[uint32]ziface.IRouter
	//负责worker取任务的消息队列
	TaskQueue      []chan ziface.IRequest
	WorkerPoolSize uint32
}

func NewMsgHandle() *MsgHandle {
	return &MsgHandle{
		Apis:           make(map[uint32]ziface.IRouter, 100),
		WorkerPoolSize: utils.GlobalObject.WorkerPoolSize,
		TaskQueue:      make([]chan ziface.IRequest, utils.GlobalObject.MaxWorkerTaskLen),
	}
}

func (m *MsgHandle) DoMsgHandler(request ziface.IRequest) {
	//从request中找到msgId
	handler, ok := m.Apis[request.GetMsgId()]
	if !ok {
		fmt.Println("api msgId = ", request.GetMsgId(), " is not found! need register handler")
	}
	//TODO npe
	handler.PreHandle(request)
	handler.Handle(request)
	handler.PostHandle(request)
	//根据msgId调度对于router
}

func (m *MsgHandle) AddRouter(msgId uint32, router ziface.IRouter) {
	if _, ok := m.Apis[msgId]; ok {
		panic("repeat api, msgId = " /*+ strconv.Itoa(msgId)*/)
	}
	m.Apis[msgId] = router
	fmt.Println("add api msgId = ", msgId, " success!")
}

func (m *MsgHandle) StartWorkerPool() {
	//根据workerPoolSize，每个worker用一个go承载
	for i := 0; i < int(m.WorkerPoolSize); i++ {
		//一个worker被启动
		//当前worker对应channel消息队列 开辟空间
		m.TaskQueue[i] = make(chan ziface.IRequest, utils.GlobalObject.MaxWorkerTaskLen)
		//启动当前worker
		go m.StartOneWorker(i, m.TaskQueue[i])
	}
}

func (m *MsgHandle) StartOneWorker(workerId int, taskQueue chan ziface.IRequest) {
	fmt.Println("worker id = ", workerId, " is started...")
	for {
		select {
		case request := <-taskQueue:
			m.DoMsgHandler(request)

		}
	}
}

func (m *MsgHandle) SendMsgToTaskQueue(request ziface.IRequest) {
	//平均分配
	connId := request.GetConn().GetConnId()
	workerId := connId % m.WorkerPoolSize
	fmt.Println("add connId = ", connId, " request msg id = ", request.GetMsgId(), " to worker id = ", workerId)
	m.TaskQueue[workerId] <- request
}

package znet

import (
	"zinx/ziface"
)

// @author lqs
// @date 2023/3/16 8:57 PM

// BaseRouter 定义他的初衷是嵌入BaseRouter基类，按需重写
type BaseRouter struct {
}

// PreHandle Router全部继承BaseRouter,不需要实现PreHandle和PostHandle
func (b *BaseRouter) PreHandle(request ziface.IRequest) {
}

func (b *BaseRouter) Handle(request ziface.IRequest) {
}

func (b *BaseRouter) PostHandle(request ziface.IRequest) {
}

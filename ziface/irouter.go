package ziface

// @author lqs
// @date 2023/3/16 8:49 PM

// IRouter 路由接口
type IRouter interface {
	// PreHandle 处理之前
	PreHandle(request IRequest)

	// Handle  处理
	Handle(request IRequest)

	// PostHandle 处理之后
	PostHandle(request IRequest)
}

package utils

import (
	"encoding/json"
	"os"
	"zinx/ziface"
)

// @author lqs
// @date 2023/3/16 10:05 PM

// GlobalObj 全局参数 🐶
type GlobalObj struct {
	//Server
	TcpServer ziface.IServer //全局server对象
	Host      string         //当前服务器监听的ip
	TcpPort   int            //服务器主机监听的端口
	Name      string         //服务器名字

	//Zinx
	Version          string //版本号
	MaxConn          int    //最大连接数
	MaxPackageSize   uint32 //数据包最大值
	WorkerPoolSize   uint32 //当前业务工作worker池的goroutine数量
	MaxWorkerTaskLen uint32 //zinx允许用户最多的worker数量
}

// GlobalObject 定义全局对外GlobalObj
var GlobalObject *GlobalObj

// 初始化
func init() {
	//默认值
	GlobalObject = &GlobalObj{
		Name:             "ZinxServerApp",
		Version:          "V0.4",
		TcpPort:          8080,
		Host:             "0.0.0.0",
		MaxConn:          1000,
		MaxPackageSize:   4096,
		WorkerPoolSize:   10,
		MaxWorkerTaskLen: 1024,
	}
	//从json加载
	//GlobalObject.Reload()
}

func (g *GlobalObj) Reload() {

	data, err := os.ReadFile("./utils/conf/zinx.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &GlobalObject)
	if err != nil {
		panic(err)
	}
}

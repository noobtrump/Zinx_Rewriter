package ziface

//定义服务器接口
type IServer interface {
	//启动服务器
	Start()
	//停止服务器
	Stop()
	//开启业务服务
	Serve()
	//路由功能：给当前服务注册一个路由业务方法，供客户端连接处理
	AddRouter(msgId uint32, router IRouter)
}

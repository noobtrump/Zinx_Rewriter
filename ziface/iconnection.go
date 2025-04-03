package ziface

import "net"

// IConnection 定义连接接口
// 该接口抽象了网络连接的基本操作
type IConnection interface {
	// Start 启动连接，开始读写数据
	Start()

	// Stop 停止连接，释放资源
	Stop()

	// GetTCPConnection 获取底层的TCP连接对象
	// 返回: *net.TCPConn TCP连接指针
	GetTCPConnection() *net.TCPConn

	// GetConnID 获取当前连接的唯一ID
	// 返回: uint32 连接ID
	GetConnID() uint32

	// RemoteAddr 获取远程客户端地址信息
	// 返回: net.Addr 远程地址
	RemoteAddr() net.Addr
	//直接将Message数据发送给远程的TCP客户端
	SendMsg(msgId uint32, data []byte) error
	//直接将Message数据发送给远程的TCP客户端(有缓冲)
	SendBuffMsg(msgId uint32, data []byte) error //添加带缓冲发送消息接口

}

// HandFunc 定义处理函数的类型
// 参数:
//   *net.TCPConn: TCP连接对象
//   []byte: 接收到的数据
//   int: 数据长度
// 返回: error 处理过程中产生的错误
type HandFunc func(*net.TCPConn, []byte, int) error

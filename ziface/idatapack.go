package ziface

/*
	封包数据和拆包数据
	直接面向TCP连接中的数据流，为传输数据添加头部信息，处理TCP粘包
*/

type IDataPack interface {
	//获取包头长度
	GetHeadLen() uint32
	//封包方法
	Pack(msg IMessage) ([]byte, error)
	//拆包方法
	Unpack([]byte) (IMessage, error)
}

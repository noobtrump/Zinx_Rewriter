package znet

import (
	"fmt"
	"net"
)

//IServer 接口的实现
type Server struct {
	// 定义服务器名称，IP链接方式，绑定IP地址，及端口
	Name      string
	IPVersion string
	IP        string
	Port      int
}

//============== 实现 ziface.IServer 里的全部接口方法 ========

func (s *Server) Start() {
	fmt.Printf("[START] Server listenner at IP: %s, Port %d, is starting\n", s.IP, s.Port)


	go func ()  {
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil{
			fmt.Println("resolve tcp addr err:", err)
			return
		}

		listenner, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listen", s.IPVersion, "err", err)
			return
		}

		fmt.Println("start Zinx server", s.Name, "success, now listenning...")

		for {
			
		}
		
	}

	
}

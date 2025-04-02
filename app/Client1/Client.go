package main

import (
	"Zinx_Rewriter/znet"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

/*
定义一个全局的客户端工厂，动态选择实现：
*/
var clients = map[string]func(){
	"client0": func() {
		fmt.Println("Client Test ... start")
		//3秒之后发起测试请求，给服务端开启服务的机会
		time.Sleep(3 * time.Second)

		conn, err := net.Dial("tcp", "127.0.0.1:7777")
		if err != nil {
			fmt.Println("client start err, exit!")
			return
		}

		for {
			//发封包message消息
			dp := znet.NewDataPack()
			msg, _ := dp.Pack(znet.NewMsgPackage(0, []byte("Zinx V0.6 Client0 Test Message")))
			_, err := conn.Write(msg)
			if err != nil {
				fmt.Println("write error err ", err)
				return
			}

			//先读出流中的head部分
			headData := make([]byte, dp.GetHeadLen())
			_, err = io.ReadFull(conn, headData) //ReadFull 会把msg填充满为止
			if err != nil {
				fmt.Println("read head error")
				break
			}
			//将headData字节流 拆包到msg中
			msgHead, err := dp.Unpack(headData)
			if err != nil {
				fmt.Println("server unpack err:", err)
				return
			}

			if msgHead.GetDataLen() > 0 {
				//msg 是有data数据的，需要再次读取data数据
				msg := msgHead.(*znet.Message)
				msg.Data = make([]byte, msg.GetDataLen())

				//根据dataLen从io中读取字节流
				_, err := io.ReadFull(conn, msg.Data)
				if err != nil {
					fmt.Println("server unpack data err:", err)
					return
				}

				fmt.Println("==> Recv Msg: ID=", msg.Id, ", len=", msg.DataLen, ", data=", string(msg.Data))
			}

			time.Sleep(1 * time.Second)
		}
	},
	"client1": func() {
		fmt.Println("Client Test ... start")
		//3秒之后发起测试请求，给服务端开启服务的机会
		time.Sleep(3 * time.Second)

		conn, err := net.Dial("tcp", "127.0.0.1:7777")
		if err != nil {
			fmt.Println("client start err, exit!")
			return
		}

		for {
			//发封包message消息
			dp := znet.NewDataPack()
			msg, _ := dp.Pack(znet.NewMsgPackage(1, []byte("Zinx V0.6 Client1 Test Message")))
			_, err := conn.Write(msg)
			if err != nil {
				fmt.Println("write error err ", err)
				return
			}

			//先读出流中的head部分
			headData := make([]byte, dp.GetHeadLen())
			_, err = io.ReadFull(conn, headData) //ReadFull 会把msg填充满为止
			if err != nil {
				fmt.Println("read head error")
				break
			}
			//将headData字节流 拆包到msg中
			msgHead, err := dp.Unpack(headData)
			if err != nil {
				fmt.Println("server unpack err:", err)
				return
			}

			if msgHead.GetDataLen() > 0 {
				//msg 是有data数据的，需要再次读取data数据
				msg := msgHead.(*znet.Message)
				msg.Data = make([]byte, msg.GetDataLen())

				//根据dataLen从io中读取字节流
				_, err := io.ReadFull(conn, msg.Data)
				if err != nil {
					fmt.Println("server unpack data err:", err)
					return
				}

				fmt.Println("==> Recv Msg: ID=", msg.Id, ", len=", msg.DataLen, ", data=", string(msg.Data))
			}

			time.Sleep(1 * time.Second)
		}
	},
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: go run main.go [client0|client1]")
	}
	if fn, ok := clients[os.Args[1]]; ok {
		fn()
	} else {
		log.Fatal("Invalid client type")
	}
}

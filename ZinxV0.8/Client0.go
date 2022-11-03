package main

import (
	"fmt"
	"github.com/lemuzhi/zinx/znet"
	"io"
	"net"
	"time"
)

//模拟客户端

func main() {

	fmt.Println("client0 start ......")

	time.Sleep(1 * time.Second)

	//1.链接远程服务器，得到一个conn
	//conn, err := net.Dial("tcp4", "192.168.131.128:8999")
	conn, err := net.Dial("tcp4", "127.0.0.1:8999")
	if err != nil {
		fmt.Println("client start err, exit")
		return
	}

	for {
		//发送封包的message消息
		dp := znet.NewDataPack()

		binaryMsg, err := dp.Pack(znet.NewMsgPackage(0, []byte("Zinx client0 Test Message")))
		if err != nil {
			fmt.Println("pack error ", err)
			return
		}

		if _, err := conn.Write(binaryMsg); err != nil {
			fmt.Println("write error ", err)
			return
		}

		//服务器就应该给我们回复一个message数据，MsgID:1 Ping..ping..ping

		// 1、先读取流中的head部分，得到ID河dataLen

		binaryHead := make([]byte, dp.GetHeadLen())
		if _, err = io.ReadFull(conn, binaryHead); err != nil {
			fmt.Println("read head error ", err)
			break
		}

		//将二进制的head拆包到msg结构体中
		msgHead, err := dp.Unpack(binaryHead)
		if err != nil {
			fmt.Println("unpack head error ", err)
			break
		}

		if msgHead.GetMsgLen() > 0 {
			//msg是有数据的
			//2、再根据DataLen进行第二次读取，将data读出来
			msg := msgHead.(*znet.Message)
			msg.Data = make([]byte, msg.GetMsgLen())

			if _, err = io.ReadFull(conn, msg.Data); err != nil {
				fmt.Println("read msg data error, ", err)
				return
			}

			fmt.Println("recv server msg : ID= ", msg.Id, ", len = ", msg.DataLen, ", data = ", string(msg.Data))

		}

		time.Sleep(1 * time.Second)
	}
}

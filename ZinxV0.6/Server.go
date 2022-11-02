package main

import (
	"fmt"
	"github.com/lemuzhi/zinx/ziface"
	"github.com/lemuzhi/zinx/znet"
)

/*
	基于zinx开发的服务端应用程序
*/

// ping test 自定义路由
type PingRouter struct {
	znet.BaseRouter
}

// Test Handle
func (p *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call PingRouter Handle...")
	//先读取客户端的数据，再回写ping...ping...ping
	fmt.Println("recv from client: msgID = ", request.GetMsgID(), ", data = ", string(request.GetData()))

	err := request.GetConnection().SendMsg(200, []byte("ping...ping...ping"))
	if err != nil {
		fmt.Println(err)
	}
}

// hello Zinx 自定义路由
type HelloZinxRouter struct {
	znet.BaseRouter
}

// Test Handle
func (h *HelloZinxRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call HelloZinxRouter Handle...")
	//先读取客户端的数据，再回写ping...ping...ping
	fmt.Println("recv from client: msgID = ", request.GetMsgID(), ", data = ", string(request.GetData()))

	err := request.GetConnection().SendMsg(201, []byte("Hello Welcome to zinx"))
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	//1.使用zinx的api，创建一个server句柄
	s := znet.NewServer("[zinx V0.6]")

	//2.给当前zinx框架添加一个自定义的router
	s.AddRouter(0, &PingRouter{})
	s.AddRouter(1, &HelloZinxRouter{})

	//3.启动server
	s.Serve()
}

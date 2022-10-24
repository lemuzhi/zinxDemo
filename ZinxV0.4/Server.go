package main

import (
	"fmt"
	"github.com/lemuzhi/zinx/ziface"
	"github.com/lemuzhi/zinx/znet"
)

/*
	基于zinx开发的服务端应用程序
*/

//ping test 自定义路由
type PingRouter struct {
	znet.BaseRouter
}

// Test PreHandle
func (p *PingRouter) PreHandle(request ziface.IRequest)  {
	fmt.Println("Call Router PreHandle...")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("before ping...\n"))
	if err != nil {
		fmt.Println("call back before ping error")
	}
}

// Test Handle
func (p *PingRouter) Handle(request ziface.IRequest)  {
	fmt.Println("Call Router Handle...")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("ping... ping... ping...\n"))
	if err != nil {
		fmt.Println("call back ping... ping... ping... ping error")
	}
}

// Test PostHandle
func (p *PingRouter) PostHandle(request ziface.IRequest)  {
	fmt.Println("Call Router PostHandle...")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("after ping...\n"))
	if err != nil {
		fmt.Println("call back after ping... ping error")
	}
}

func main() {
	//1.使用zinx的api，创建一个server句柄
	s := znet.NewServer("[zinx V0.3]")

	//2.给当前zinx框架添加一个自定义的router
	s.AddRouter(&PingRouter{})

	//3.启动server
	s.Serve()
}



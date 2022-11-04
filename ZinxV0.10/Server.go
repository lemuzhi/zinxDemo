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

// DoConnectionBegin 创建链接之后执行钩子函数
func DoConnectionBegin(conn ziface.IConnection) {
	fmt.Println("==> DoConnectionBegin is Called... ")
	if err := conn.SendMsg(202, []byte("DoConnection BEGIN")); err != nil {
		fmt.Println(err)
	}

	//给当前连接设置一些属性
	fmt.Println("Set conn property ... ")
	conn.SetProperty("Name", "零度")
	conn.SetProperty("Email", "1163648924@qq.com")
	conn.SetProperty("GitHub", "https://www.github.com/lemuzhi")
}

// DoConnectionLost 链接断开之前需要处理的函数
func DoConnectionLost(conn ziface.IConnection) {
	fmt.Println("==> DoConnectionLost is Called... ")
	fmt.Println("conn ID = ", conn.GetConnID(), " is Lost...")

	//获取连接属性
	if name, err := conn.GetProperty("Name"); err == nil {
		fmt.Println("Name = ", name)
	}
	if name, err := conn.GetProperty("Email"); err == nil {
		fmt.Println("Email = ", name)
	}
	if name, err := conn.GetProperty("GitHub"); err == nil {
		fmt.Println("GitHub = ", name)
	}
}

func main() {
	//1.使用zinx的api，创建一个server句柄
	s := znet.NewServer("[zinx V0.8]")

	//2、链接的Hook钩子函数
	s.SetOnConnStart(DoConnectionBegin)
	s.SetOnConnStop(DoConnectionLost)

	//3	.给当前zinx框架添加一个自定义的router
	s.AddRouter(0, &PingRouter{})
	s.AddRouter(1, &HelloZinxRouter{})

	//4.启动server
	s.Serve()
}

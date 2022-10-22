package main

import "github.com/lemuzhi/zinx/znet"

//基于zinx开发的服务端应用程序

func main() {
	//使用zinx的api，创建一个server句柄
	s := znet.NewServer("[zinx V0.1]")
	s.Serve()
}



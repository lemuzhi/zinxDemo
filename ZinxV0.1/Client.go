package main

import (
	"fmt"
	"net"
	"time"
)

//模拟客户端

func main() {

	fmt.Println("client start ......")

	time.Sleep(1 * time.Second)

	//1.链接远程服务器，得到一个conn
	conn, err := net.Dial("tcp4", "192.168.131.128:8999")
	if err != nil {
		fmt.Println("client start err, exit")
		return
	}

	for  {
		//2.链接调用write，写数据
		var data string
		_, err = fmt.Scan(&data)
		if err != nil {
			fmt.Println("read input err: ", err)
			continue
		}
		_, err = conn.Write([]byte(data))
		if err != nil {
			fmt.Println("write data err: ", err)
			continue
		}

		buf := make([]byte, 512)
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read buf err: ", err)
			continue
		}
		fmt.Printf("server call back:%s, cnt = %d\n", buf[:], cnt)
	}
}

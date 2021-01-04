package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main()  {

	conn,err := net.Dial("tcp","127.0.0.1:8888")
	if err != nil {
		fmt.Println("client dial err=",err)
		return
	}

	//功能1 客户端发送单行数据，然后退出
	reader := bufio.NewReader(os.Stdin)  //os.stdin代表标准输入【终端】

	for {
		//从终端读取一行输入，发送给服务端
		line,err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("readerString err=",err)
		}

		line = strings.Trim(line," \r\n")
		if line == "exit" {
			fmt.Println("客户端退出了")
			break
		}

		//再将line发送给服务器
		_,err = conn.Write([]byte(line+"\n"))
		if err != nil {
			fmt.Println("conn.Write err",err)
		}
	}
}






















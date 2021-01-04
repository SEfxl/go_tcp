package main

import (
	"fmt"
	"io"
	"net"
)

func process(conn net.Conn)  {
	//循环的接受客户端发送的数据
	defer conn.Close()  //一定要关闭

	for{
		//创建一个新的切片
		buf := make([]byte,1024)
		//等待客户端通过conn发送信息,如果客户端链接没有数据,协程将一直阻塞在这里
		//fmt.Println("服务端:"+conn.LocalAddr().String()+" 在等待客户端:"+conn.RemoteAddr().String()+"的输入")
		n,err:=conn.Read(buf) //从conn中读取数据

		if err == io.EOF {
			fmt.Println("客户端退出",err)
			return
		}

		//显示客户端发送的内容到服务器的终端
		fmt.Print(string(buf[:n]))
	}

}

func main()  {

	fmt.Println("服务端开始监听....")
	//tcp表示使用的是tcp协议
	//0.0.0.0:8888表示在本地监听8888端口
	listen,err := net.Listen("tcp","0.0.0.0:8888")

	if err != nil {
		fmt.Println("listen err = ",err)
		return
	}

	defer listen.Close() //延时关闭
	//循环等待客户端来链接
	for {
		fmt.Println("等待客户端来链接。。。。")
		conn,err := listen.Accept()
		if err != nil {
			fmt.Println("Accept() err = ",err)
		} else {
			fmt.Printf("Accept() conn =%v, 客户端IP=%v\n",conn,conn.RemoteAddr().String())
		}

		//准备一个协程,为客户端服务
		go process(conn)
	}



}

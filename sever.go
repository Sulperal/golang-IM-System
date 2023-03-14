package main

import (
	"fmt"
	"net"
)

// 定义类型
type Server struct {
	Ip   string
	Port int
}

// 创建sever
// 构造器
func NewServer(ip string, port int) *Server {
	//创建一个对象
	server := &Server{
		Ip:   ip,
		Port: port,
	}
	//返回对象的地址，
	return server
}

func (this *Server) Handler(conn net.Conn) {
	//...当前链接的任务
	fmt.Println("connect success")
}

// 启动服务器
func (this *Server) Start() {
	//socket listen
	//Sprintf拼接字符串
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	if err != nil {
		fmt.Printf("net.Listen err: %v\n", err)
		return
	}
	//close listen socket
	defer listener.Close()

	for {
		//accept
		//返回一个读写的套接字
		conn, err2 := listener.Accept()
		if err2 != nil {
			fmt.Printf("listener accept err2: %v\n", err2)
			return
		}
		//handle
		//开一个协程去处理业务
		go this.Handler(conn)
	}

}

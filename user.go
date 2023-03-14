package main

import "net"

type User struct {
	Name string
	Addr string
	C    chan string
	conn net.Conn
}

//创建用户
//通过建立的链接创建User
func NewUser(conn net.Conn) *User {

	userAddr := conn.RemoteAddr().String()

	user := &User{
		Name: userAddr,
		Addr: userAddr,
		C:    make(chan string),
		conn: conn,
	}
	//创建的同时启动监听当前user channel消息的goroutine
	go user.ListenMessage()
	return user

}

//监听当前的User channel, 一旦通道中有消息，马上将其发送给对端客户端
func (this *User) ListenMessage() {
	for {
		msg := <-this.C

		this.conn.Write([]byte(msg + "\r\n"))
	}
}

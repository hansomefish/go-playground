package main

import (
	"fmt"
	"net"
)

type Server struct {
	Ip string
	Port int
	Message chan string
}

// get server instance
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip: ip,
		Port: port,
	}

	return server
}

func (server *Server) BroadCast(user *User, msg string) {
	//sendMsg := "[" + user.Address + "]" + ": " + msg
	//for server.
}

// handle data
func handleData(conn net.Conn) {
	//fmt.Println("connect successfully!")
	user := NewUser(conn)
	fmt.Println(user.Address, "已上线")


}

func (server *Server) Start() {
	// listen connection
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", server.Ip, server.Port))
	if err != nil {
		fmt.Println("listen socket err", err)
		return
	}
	defer listener.Close()
	// accept connection
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept conn err", err)
			continue
		}
		// handle data in goroutine
		go handleData(conn)
	}
}
package main

import "net"

type User struct {
	Address string
	c chan string
	conn net.Conn
}

func NewUser(conn net.Conn) *User {
	addr := conn.RemoteAddr().String()
	user := &User{
		Address: addr,
		c:       make(chan string),
		conn:    conn,
	}

	return user
}

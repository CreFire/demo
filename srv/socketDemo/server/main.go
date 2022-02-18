package main

import (
	"net"
	"sync"
)

type nameHandler struct {
	eventsLock   sync.RWMutex
	Conn         net.Conn
	onConnect    func(conn net.Conn) error
	onDisconnect func(conn net.Conn, msg string)
	onError      func(conn net.Conn, err error)
}

func main() {
	conn, err := net.
		net.Listen(":8000")
}

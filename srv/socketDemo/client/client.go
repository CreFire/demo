package main

import "net"

// 定义客户端对象，以实现对客户端连接的封装
type client struct {
	// 客户端连接对象
	conn net.Conn

	// 接收到的消息内容
	content []byte
}

// 新建客户端对象
// conn：连接对象
// 返回值：客户端对象的指针
func newClient(_conn net.Conn) *client {
	return &client{
		conn:    _conn,
		content: make([]byte, 0, 1024),
	}
}

// 追加内容
// content：新的内容
// 返回值：无
func (clientObj *client) appendContent(content []byte) {
	clientObj.content = append(clientObj.content, content...)
}

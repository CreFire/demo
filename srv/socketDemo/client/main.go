package main

import (
	"bufio"
	"fmt"
	"github.com/sirupsen/logrus"
	"net"
	"os"
)

func main() {
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.InfoLevel)

	tcpAddr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:8000")

	if err != nil {
		fmt.Println(os.Stderr, "Fatal error: ", err)
		os.Exit(1)
	}

	//建立服务器连接
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Println(conn.RemoteAddr().String(), os.Stderr, "Fatal error:", err)
		os.Exit(1)
	}
	sender(conn)
}

func sender(conn *net.TCPConn) {

	for {
		inStr := reader()
		msgBack, err := conn.Write([]byte(inStr)) //给服务器发信息
		if err != nil {
			fmt.Println(conn.RemoteAddr().String(), "服务器反馈")
			os.Exit(1)
		}
		buffer := make([]byte, 1024)
		msg, err := conn.Read(buffer) //接受服务器信息

		fmt.Println(conn.RemoteAddr().String(), "服务器反馈：", string(buffer[:msg]),string(buffer[:msg]))
		fmt.Println(msgBack, "；实际发送了", len(inStr),inStr)

		conn.Write([]byte("ok")) //在告诉服务器，它的反馈收到了。
	}
}

func reader() string {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	return input.Text()
}
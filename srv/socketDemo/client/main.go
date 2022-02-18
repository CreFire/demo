package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"net"
	"os"
	"time"
)

func main() {
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.InfoLevel)
	// 连接指定的端口
	msg := ""
	conn, err := net.DialTimeout("tcp", "127.0.0.1:8000", 2*time.Second)
	if err != nil {
		msg = fmt.Sprintf("Dial Error: %s", err)
	} else {
		msg = fmt.Sprintf("Connect to the server. (local address: %s)", conn.LocalAddr())
	}
	var clientObj  = newClient(conn)
	defer func() {
		conn.Close()
		clientObj = nil
	}()

	// 死循环，不断地读取数据，解析数据，发送数据
	for {
		// 先读取数据，每次读取1024个字节
		readBytes := make([]byte, 1024)

		// Read方法会阻塞，所以不用考虑异步的方式
		n, err := conn.Read(readBytes)
		if err != nil {
			logrus.Errorf(fmt.Sprintf("读取消息错误：%s，本次读取的字节数为：%d", err, n))
			os.Exit(1)
		}

		// 将读取到的数据追加到已获得的数据的末尾
		clientObj.appendContent(readBytes[:n])

		// 已经包含有效的数据，处理该数据
		handleClient()
	}
}
}
package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"log"
	"net"
	"os"
	"sync"
)

const (
	msgLength = 10
)

type nameHandler struct {
	eventsLock   sync.RWMutex
	Conn         net.Conn
	onConnect    func(conn net.Conn) error
	onDisconnect func(conn net.Conn, msg string)
	onError      func(conn net.Conn, err error)
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		logrus.Errorf("err:%v",err)
	}
	for {
		conn, err := listener.Accept() //开启监听
		if err != nil {
			fmt.Println("Accept is err!: ", err)
			continue
		}
		//发生了连接
		fmt.Println("tcp connect success:", conn.RemoteAddr().String())
		//go handleConnection(conn)
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	buffer := make([]byte, 1024)
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			logrus.Errorf("conn close Error:%v",err)
		}
	}(conn)
	for {
		//接受客户端消息
		msg, err := conn.Read(buffer)

		if err != nil {
			//接受错误
			fmt.Println("connection err!:", err)
			return
		}
		//接受正确
		fmt.Print(conn.RemoteAddr().String())
		fmt.Println("receive data: ", string(buffer[:msg]))

		//反馈给客户端
		bufferReturn := fmt.Sprintf("我收到了 %v \n",string(buffer[:msg]))
		msgW, errW := conn.Write([]byte(bufferReturn))

		//确认客户端没有收到回执
		if errW != nil {
			fmt.Print(conn.RemoteAddr().String(), msgW)
			fmt.Println("没有收到回执")
			return
		}

		//确认客户端收到回执
		msg, err = conn.Read(buffer)
		fmt.Println(conn.RemoteAddr().String(), "客户端收到回执", string(buffer[:msg]), "客户收到了", msgW, "；实际发送了", len(bufferReturn))
	}

}


func init(){
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.DebugLevel)
	viper.AddConfigPath(".")    // 设置配置文件和可执行二进制文件在用一个目录
	viper.SetConfigFile("../config.yaml")
	err := viper.ReadInConfig() // 根据以上配置读取加载配置文件
	if err != nil {
		log.Fatal(err)// 读取配置文件失败致命错误
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		//viper配置发生变化了 执行响应的操作
		logrus.Errorf("Config file changed:%v \n", e.Name)
	})
}

package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	// tcp server端

	// 1. 本地端口启动服务
	lis, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// 2. 等别人来连接
	for {
		conn, err := lis.Accept()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		go processConn(conn)
	}

}

func processConn(conn net.Conn) {
	// 3. 与客户端通信
	defer conn.Close()
	tmp := make([]byte, 128, 128)
	reader := bufio.NewReader(os.Stdin)
	for {
		n, err := conn.Read(tmp)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(string(tmp[:n]))

		fmt.Println("请回复: ")
		msg, _ := reader.ReadString('\n')
		msg = strings.TrimSpace(msg)
		fmt.Scanln(&msg)
		if msg == "exit" {
			break
		}
		fmt.Println(msg)
	}

}

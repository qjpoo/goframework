package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	// tcp  client

	// 1. 与server连立连接
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// 2. 发送数据
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("请说话: ")
		msg, _ := reader.ReadString('\n')
		msg = strings.TrimSpace(msg)
		fmt.Scanln(&msg)
		if msg == "exit" {
			break
		}
		conn.Write([]byte(msg))
	}

	// 3. 关闭
	conn.Close()

}

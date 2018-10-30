package main

import (
	"bytes"
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port ", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)
	for {
		sendData(conn)
	}
}

func sendData(conn net.Conn) {
	input := ""
	// 等待用户输入
	fmt.Scanf("%s", &input)
	// 发送给服务端
	_, err := conn.Write([]byte(input))
	checkError(err)
	var chunk []byte
	for {
		// 读出服务端的返回值
		data := make([]byte, 10)
		_, err := conn.Read(data)
		checkError(err)
		chunk = append(chunk, data...)
		index := bytes.Index(data, []byte("|"))
		if index == -1 {
			continue
		}
		splitData := bytes.Split(chunk, []byte("|"))
		fmt.Println("res:", string(splitData[0]))
		break
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

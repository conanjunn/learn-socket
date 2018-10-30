package main

import (
	"bytes"
	"fmt"
	"net"
)

func main() {
	// 创建地址对象
	tcpAddr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:8888")
	if err != nil {
		fmt.Printf("resolveTCPAddr error: %v\n", err)
		return
	}
	// 创建tcp服务器
	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		fmt.Printf("listenTCP error: %v\n", err)
		return
	}

	fmt.Printf("server start\n")

	for {
		// 等待客户端连接 用for来不断获取连接(客户端可能有多个)，防止进程自动退出
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("accept error: %v\n", err)
		}
		fmt.Printf("accept success\n")
		// 处理客户端连接
		go handlerClient(conn)
	}
}

func handlerClient(conn net.Conn) {
	var chunk []byte
	for {
		// 一次读取128字节的数据
		data := make([]byte, 10)
		// 读取客户端传来得数据，用for来不断读取
		dataLen, err := conn.Read(data)
		if err != nil {
			fmt.Printf("read err: %v\n", err)
			break
		}
		fmt.Printf("read wait, len=%d\n", dataLen)
		index := bytes.Index(data, []byte("|"))
		if index == -1 {
			chunk = append(chunk, data...)
			continue
		}
		splitData := bytes.Split(data, []byte("|"))
		wholeData := append(chunk, splitData[0]...)
		chunk = splitData[1]

		str := string(wholeData)
		if str == "close" {
			fmt.Printf("conn will close\n")
			conn.Close()
			fmt.Printf("conn closed\n")
			break
		}
		go func() {
			res := fmt.Sprintf("%s%s", str, "|")
			fmt.Printf("%s\n", str)
			// 回应客户端数据
			conn.Write([]byte(res))
		}()
	}
}

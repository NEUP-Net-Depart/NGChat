package server

import (
	"net"
	"fmt"
	"log"
)

// 处理连接。
func handle_connection(conn net.Conn) {
	// TODO: 实现连接的处理

	// NOTE: 下面的代码仅供测试使用！
	fmt.Fprintln(conn, "Hello connection world!")
	log.Printf("已成功接受地址为 %s 的传入连接。", conn.RemoteAddr().String())
	conn.Close()
}
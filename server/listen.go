package server

import (
	"net"
	"NGChat/config"
	"log"
)

// 开始监听。
// 注意：失败将会产生 panic。
func start_listening() {
	// 提示信息常量
	const listen_error = "无法启动监听于地址 %s。"
	const listen_success = "成功！服务器正工作于地址 %s ..."
	const accept_error = "无法接受连接？"

	// 尝试监听
	listener, err := net.Listen("tcp", config.ServerCfg.ListenAddr)
	if err != nil {
		log.Panicf(listen_error, config.ServerCfg.ListenAddr)
	}

	// 成功后输出信息
	log.Printf(listen_success, config.ServerCfg.ListenAddr)

	// 处理连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln(accept_error)
			continue
		}
		go handle_connection(conn)
	}
}
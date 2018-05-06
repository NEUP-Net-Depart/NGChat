package server

import (
	"net"
	"config"
	"log"
)

// 开始监听。
// 注意：失败将会产生 panic。
func startListening() {
	// 提示信息常量
	const msgListenError = "无法启动监听于地址 %s。"
	const msgListenSuccess = "成功！服务器正工作于地址 %s ..."
	const msgAcceptError = "无法接受连接？"

	// 尝试监听
	listener, err := net.Listen("tcp", config.ServerCfg.ListenAddr)
	if err != nil {
		log.Panicf(msgListenError, config.ServerCfg.ListenAddr)
	}

	// 成功后输出信息
	log.Printf(msgListenSuccess, config.ServerCfg.ListenAddr)

	// 处理连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln(msgAcceptError)
			continue
		}
		go handleConnection(conn)
	}
}
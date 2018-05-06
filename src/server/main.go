package server

import (
	"runtime"
	"config"
	"log"
)


// 服务器初始化函数。
// 载入配置文件并进行必须的操作。
func Init() {
	log.Println("正在初始化服务器...")
	defer log.Println("初始化完成.")

	// 读入配置文件
	config.LoadConfig()
	// 设置 GOMAXPROCS
	runtime.GOMAXPROCS(config.ServerCfg.GoMaxProcs)

	// TODO: 添加其他初始化操作
}


// 启动服务器。
func Start() {
	log.Println("启动服务器中...")

	// 尝试开始监听
	startListening()
}
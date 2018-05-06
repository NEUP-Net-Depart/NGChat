package config

import (
	"github.com/BurntSushi/toml"
	"log"
)

// 配置文件路径。
const PathFileConfig = "config.toml"

// 所有的配置对象将作为全局变量存储在此处，可使用 LoadConfig() 方法以从配置文件中加载数据。
var ServerCfg = ServerConfig{}



// 从配置文件中载入配置到内存中的配置对象。
// 注意：若载入失败，将会触发 panic 而终止程序运行。
func LoadConfig() {
	// 错误信息常量
	const msgPanic = "从配置文件 %s 中读取配置 %s 时发生严重错误。"

	// ServerConfig 的读取
	_, err := toml.DecodeFile(PathFileConfig, &ServerCfg)
	if err != nil {
		log.Panicf(msgPanic, PathFileConfig, "ServerConfig")
	}

	// TODO: 添加更多配置的读取
}


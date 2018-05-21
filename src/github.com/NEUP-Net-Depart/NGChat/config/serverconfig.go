package config

// 服务器配置结构体。
type ServerConfig struct {
	ListenAddr	string	`toml:"listen_addr"`	// 监听地址。
	GoMaxProcs	int		`toml:"go_maxprocs"`	// GOMAXPROCS。
}

